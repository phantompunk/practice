from __future__ import annotations, with_statement

from dataclasses import dataclass
from typing import Optional


# Definition for singly-linked list.
@dataclass
class ListNode(object):
    val: int
    next: Optional[ListNode]

    def __init__(self, x, next=None):
        self.val = x
        self.next = next

    def add(self, data=0, next=None) -> ListNode:
        self.next = ListNode(data, next)
        return self


class Solution(object):
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        """
        :type head: ListNode
        :rtype: bool
        """
        slow = fast = head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next

            if slow == fast:
                return True
        return False
