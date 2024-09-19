from __future__ import annotations


class Node:
    next: Node | None
    prev: Node | None
    data: any

    def __init__(self, data) -> None:
        self.data = data
        self.prev = None
        self.next = None


class LinkedList:
    head: Node | None
    tail: Node | None
    n: int

    def __init__(self) -> None:
        self.head = None
        self.tail = None
        self.n = 0

    def add(self, data):
        """Append element to the end of the list"""
        return self.add_last(data)
        # node = Node(data)
        # self.n += 1
        # if not self.tail:
        #     self.head = self.tail = node
        #     return
        #
        # node.prev = self.tail
        # self.tail.next = node
        # self.tail = node

    def add_at(self, data, index: int):
        """Insert element to the specified position in the list"""
        if index < 0 or index > self.n:
            return None
        elif index == self.n:
            return self.add(data)
        elif index == 0:
            return self.add_first(data)

        self.n += 1
        curr = self.head
        while curr and index > 0:
            curr = curr.next
            index -= 1

        node = Node(data)
        node.prev = curr.prev
        node.next = curr
        curr.prev = node

        if node.prev:
            node.prev.next = node

    def add_first(self, data):
        """Insert element to the beginning of the list"""
        node = Node(data)
        self.n += 1
        if not self.head:
            self.head = self.tail = node
            return

        node.next = self.head
        self.head.prev = node
        self.head = node

    def add_last(self, data):
        """Append element to the end of the list"""
        node = Node(data)
        self.n += 1
        if not self.tail:
            self.head = self.tail = node
            return

        node.prev = self.tail
        self.tail.next = node
        self.tail = node

    def clear(self):
        """Remove all elements from the list"""
        self.head = None
        self.tail = None
        self.n = 0

    def index_of(self, value) -> int:
        """Return the index of the first occurence of the element or -1 if DNE"""
        index = 0
        curr = self.head
        while curr:
            if curr.data == value:
                return index
            curr = curr.next
            index += 1
        return -1

    def peek(self):
        """Retrieves but does not remove the first element from the list or None if empty"""
        if self.head:
            return self.head.data
        return

    def peek_first(self):
        """Retrieves but does not remove the first element from the list or None if empty"""
        return self.peek()

    def peek_last(self):
        """Retrieves but does not remove the last element from the list"""
        if self.tail:
            return self.tail.data
        return

    def remove(self):
        """Retrieves and remove the first element from the list"""
        curr = self.head
        if curr is None:
            return

        self.n-=1
        if curr.next:
            curr.next.prev = None
        self.head = curr.next
        curr.next = None
        return curr.data

    def remove_at(self, index: int):
        """Retrieves and remove the element at the specified index from the list"""
        if index < 0 or index > self.n:
            return None

        curr = self.head
        while curr and index > 0:
            curr = curr.next
            index -= 1

        if curr is None:
            return None

        self.n-=1
        if curr.prev:
            curr.prev.next = curr.next
        if curr.prev:
            pass

        # return curr.data

    def remove_first(self):
        """Retrieves and remove the first element from the list"""
        self.n-=1
        curr = self.head
        if curr is None:
            return None
        if curr.next:
            curr.next.prev = None
        self.head = curr.next
        curr.next = None

        return curr.data

    def remove_last(self):
        """Retrieves and remove the last element from the list"""
        self.n-=1
        curr = self.tail
        if curr is None:
            return None
        if curr.prev:
            curr.prev.next = None
        self.tail = curr.prev
        return curr.data

    def size(self) -> int:
        return self.n
