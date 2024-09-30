from dataclasses import dataclass
from typing import Optional


# Definition for singly-linked list.
class ListNode:
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


class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        explore = head
        pointer = head

        length = 0
        while explore:
            explore = explore.next
            length += 1

        print(pointer)
        mid = length // 2
        while mid > 0 and pointer:
            pointer = pointer.next
            mid -= 1

        return pointer
