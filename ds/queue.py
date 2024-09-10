class Node:
    def __init__(self, value, next=None) -> None:
        self.value = value
        self.next = next


class Queue:
    def __init__(self) -> None:
        self.head = None
        self.tail = None
        self.n = 0

    def offer(self, value):
        """Add an element to the end of the queue"""
        node = Node(value)
        self.n += 1
        if self.tail is None:
            self.head = self.tail = node

        self.tail.next = node
        self.tail = node

    def peek(self):
        """Look at the head element of the queue without removing it"""
        if self.head is None:
            return None
        return self.head.value

    def poll(self):
        """Remove the head element and return the objects value"""
        if self.head is None:
            return None
        tmp = self.head
        self.head = self.head.next
        self.n -= 1
        return tmp.value

    def size(self):
        return self.n
