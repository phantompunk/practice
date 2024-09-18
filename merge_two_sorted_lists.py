from __future__ import annotations
from dataclasses import dataclass
from typing import Optional


@dataclass
class ListNode:
    val: int
    next: Optional[ListNode]
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
    def add(self, data=0) -> ListNode:
        if self.next is None:
            self.next = ListNode(data)
        return self
    def addall(self, data: list[int]) -> ListNode:
        curr = self
        while curr.next:
            curr = curr.next
        for item in data:
            curr.next = ListNode(item) 
            curr = curr.next
        return self

class Solution:
    def mergeTwoLists(
        self, list1: Optional[ListNode], list2: Optional[ListNode]
    ) -> Optional[ListNode]:
        merged = ListNode()
        curr = merged
        while list1 and list2:
            if list1.val < list2.val:
                curr.next = list1
                list1 = list1.next
            else:
                curr.next = list2
                list2 = list2.next

            curr = curr.next

        if list1:
            curr.next = list1
        else:
            curr.next = list2

        return merged.next
