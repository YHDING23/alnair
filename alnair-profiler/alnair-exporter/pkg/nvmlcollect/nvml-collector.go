package nvmlcollect

import (
	"fmt"
	"log"
	"time"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/prometheus/client_golang/prometheus"
)

// nvmlCollector has two metrics to report from nvml server
type nvmlCollector struct {
	alnairGpuUtil  *prometheus.Desc
	alnairGmemUtil *prometheus.Desc
}

//return data struct from nvml.DeviceGetProcessUtilization
type GPUStat struct {
	Pid       uint32
	SmUtil    uint32
	MemUtil   uint32
	Uuid      string
	TimeStamp uint64 // timestamp
}

//NewCollector return a prometheus.Collector, must implement Describe and Collect function
func NewCollector(nodeName string) prometheus.Collector {
	return &nvmlCollector{
		alnairGpuUtil: prometheus.NewDesc(
			//name of the metrics
			"alnair_gpu_util",
			//help text
			"GPU utilization at process level",
			//the metric's variable label dimensions.
			[]string{"pid", "gpu_uuid"},
			//the metric's constant label dimensions.
			prometheus.Labels{"nodeName": nodeName},
		),
		alnairGmemUtil: prometheus.NewDesc("alnair_gpu_mem_used",
			"Used GPU memory at process level", []string{"pid", "gpu_uuid"}, prometheus.Labels{"nodeName": nodeName},
		),
	}
}

func (collector *nvmlCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.alnairGpuUtil
	ch <- collector.alnairGmemUtil
}

func (collector *nvmlCollector) Collect(ch chan<- prometheus.Metric) {
	//v2 := rand.Intn(1000)
	metrics := GetGpuStat()

	//ch <- prometheus.MustNewConstMetric(collector.alnairGmemUsed, prometheus.CounterValue, float64(v2))
	//it is OK to have an empty metrics when no processes, nothing reported under this metrics
	for _, g := range metrics {
		ch <- prometheus.MustNewConstMetric(collector.alnairGpuUtil, prometheus.GaugeValue, float64(g.SmUtil), fmt.Sprint(g.Pid), g.Uuid)
		ch <- prometheus.MustNewConstMetric(collector.alnairGmemUtil, prometheus.GaugeValue, float64(g.MemUtil), fmt.Sprint(g.Pid), g.Uuid)
	}
	//possibility add timestamp
	//https://gist.github.com/fl64/a86b3d375deb947fa11099dd374660da

}

func GetGpuStat() []GPUStat {
	var stats []GPUStat
	//log.Printf("nvml collector got pulled.") //verify if collector function got called at the pulling frequency set up at the Prometheus Config
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		log.Printf("Unable to initialize NVML: %v", nvml.ErrorString(ret))
		return stats
	}
	defer func() {
		ret := nvml.Shutdown()
		if ret != nvml.SUCCESS {
			log.Printf("Unable to shutdown NVML: %v", nvml.ErrorString(ret))
		}
	}()

	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		log.Printf("Unable to get device count: %v", nvml.ErrorString(ret))
		return stats
	}
	for i := 0; i < count; i++ {
		device, ret := nvml.DeviceGetHandleByIndex(i)
		if ret != nvml.SUCCESS {
			log.Printf("Unable to get device at index %d: %v", i, nvml.ErrorString(ret))
			continue
		}

		uuid, ret := device.GetUUID()
		if ret != nvml.SUCCESS {
			log.Printf("Unable to get uuid of device at index %d: %v", i, nvml.ErrorString(ret))
			continue
		}

		processInfos, ret := device.GetComputeRunningProcesses()
		if ret != nvml.SUCCESS {
			log.Printf("Unable to get process info for device %s: %v", uuid, nvml.ErrorString(ret))
			continue
		}
		if len(processInfos) > 0 {
			log.Printf("Found %d processes on device %s\n", len(processInfos), uuid)
			for pi, processInfo := range processInfos {
				log.Printf("\t[%2d] ProcessInfo: %+v\n", pi, processInfo)
			}

			var lastSeenTimeStamp uint64
			lastSeenTimeStamp = uint64(time.Now().UnixMicro())
			lastSeenTimeStamp = lastSeenTimeStamp - 1e6
			processUtilizationSample, ret := nvml.DeviceGetProcessUtilization(device, lastSeenTimeStamp)
			if ret != nvml.SUCCESS {
				log.Printf("Unable to get device %s's process utilization: %v", uuid, nvml.ErrorString(ret))
				continue
			}
			log.Printf("last seen timestamp %d, %v\n", lastSeenTimeStamp, processUtilizationSample)
			for _, sample := range processUtilizationSample {
				if sample.Pid == 0 { // 0 means empty return
					continue
				}
				stats = append(stats, GPUStat{Pid: sample.Pid,
					MemUtil:   sample.MemUtil,
					SmUtil:    sample.SmUtil,
					Uuid:      uuid,
					TimeStamp: sample.TimeStamp,
				})
			}
			if len(stats) == 0 {
				log.Printf("process level utilization is empty, GPU cards may not support function nvml.DeviceGetProcessUtilization")
			} else {
				log.Printf("Pulled process level GPU utils %v", stats)
			}
		}
		//sample output
		//Found 1 processes on device GPU-af943c68-b15d-be46-901e-187129cdc536
		//[ 0] ProcessInfo: {Pid:24142 UsedGpuMemory:12512657408 GpuInstanceId:4294967295 ComputeInstanceId:4294967295}
		//last seen timestamp 1656912044933412, [{19043 1656912048930155 43 24 0 0}] //sm util, mem util, dec util, enc util
	}
	return stats
}
