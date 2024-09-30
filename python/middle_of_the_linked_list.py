from typing import Optional
from utils import ListNode


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
