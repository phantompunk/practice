from __future__ import annotations


class Node:
    def __init__(self, value, next=None) -> None:
        self.val = value
        self.next = next


class Stack:
    def __init__(self) -> None:
        self.head = None
        self.size = 0

    def empty(self) -> bool:
        """Checks if stack is empty"""
        return self.size == 0 and self.head is None

    def peek(self):
        """Look at the top element without removing it"""
        if self.head is None:
            return None
        return self.head.val

    def pop(self):
        """Remove the top element and return that objects value"""
        self.size = max(0, self.size - 1)
        if self.head is None:
            return None

        if self.size == 0:
            tmp = self.head
            self.head = None
            return tmp.val

        node = self.head
        self.head = node.next
        return node.val

    def push(self, value):
        """Push an element to the top of the stack"""
        node = Node(value)
        self.size += 1

        node.next = self.head
        self.head = node

    def search(self, value):
        """Return the index of the value."""
        index = 1
        curr = self.head
        while curr:
            print(curr.val)
            if curr.val == value:
                return index
            curr = curr.next
            index += 1
        return -1
