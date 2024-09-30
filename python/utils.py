# Definition for singly-linked list.
from __future__ import annotations
from typing import Optional


class ListNode:
    val: int
    next: Optional[ListNode]

    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __eq__(self, other):
        if not isinstance(other, ListNode):
            return False
        current_self = self
        current_other = other
        while current_self and current_other:
            if current_self.val != current_other.val:
                return False
            current_self = current_self.next
            current_other = current_other.next
        return current_self is None and current_other is None

    def add(self, data=0, next=None) -> ListNode:
        self.next = ListNode(data, next)
        return self

    def addall(self, data: list[int]) -> ListNode:
        curr = self
        while curr.next:
            curr = curr.next
        for item in data:
            curr.next = ListNode(item)
            curr = curr.next
        return self
