from typing import Callable, Optional
from PIL import Image
from lib.AlnairJobDataset import *


class ImageNetDataset(AlnairJobDataset):
    def __init__(self, keys,
                transform: Optional[Callable] = None,
                target_transform: Optional[Callable] = None):
        super().__init__(keys)
        self.keys = list(self.data.keys())
        classes, class_to_idx = self.find_classes(self.root)
        self.classes = classes
        self.class_to_idx = class_to_idx

        self.transform = transform
        self.target_transform = target_transform
    
    def find_classes(self):
        classes = set([x.split('/')[2] for x in self.keys])
        if len(classes) == 0:
            raise FileNotFoundError(f"Couldn't find any class.")
        class_to_idx = {cls_name: i for i, cls_name in enumerate(classes)}
        return classes, class_to_idx

    def __preprocess__(self):
        samples = []
        targets = []
        for target_class in sorted(self.class_to_idx.keys()):
            cls_keys = list(filter(lambda x: target_class in x.split('/'), self.keys))
            for key in sorted(cls_keys):
                samples.append(self.data[key])
                targets.append(target_class)
        return samples, targets
    
    def __getitem__(self, index: int):
        img, target = self.data[index], self.targets[index]
        img = Image.frombytes(img)    
        if self.transform is not None:
            img = self.transform(img)
        if self.target_transform is not None:
            target = self.target_transform(target)
        return img, target

    def __len__(self) -> int:
        return len(self.data)