from collections.abc import Iterable


class ArrayList:
    def __init__(self, capacity: int = 2):
        self.cap = capacity
        self.data = [0] * self.cap
        self.n = 0

    def __len__(self) -> int:
        return self.size()

    def add(self, element, index: int = 0) -> None:
        if self.size() == self.capacity():
            self._resize()

        self.data[self.n] = element
        self.n += 1

    def add_all(self, collection: Iterable) -> bool:
        for x in collection:
            self.add(x)
        return True

    def remove(self, obj):
        item = None
        for i in range(self.size()):
            if self.data[i] == obj:
                item = obj
                self.remove_at(i)
        return item

    def remove_at(self, index: int) -> bool:
        if self.n < index < 0:
            raise IndexError()

        for i in range(index, self.n - 1):
            self.data[i] = self.data[i + 1]
        self.n -= 1
        return True

    def get(self, index: int):
        if self.n < index <= 0:
            return IndexError
        return self.data[index]

    def set(self, index, element):
        if self.n < index <= 0:
            return IndexError
        self.data[index] = element

    def clear(self):
        self.data = [0] * self.capacity()
        self.n = 0

    def _resize(self):
        self.cap = 2 * self.cap
        arr = [0] * self.cap
        for i in range(self.size()):
            arr[i] = self.data[i]
        self.data = arr

    def size(self) -> int:
        return self.n

    def capacity(self) -> int:
        return self.cap
