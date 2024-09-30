# Definition for singly-linked list.
from typing import Optional
from utils import ListNode


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        """Reverse the list
        Given the following linked list:
        A --> B --> C
        Return:
        C --> B --> A
        """
        curr = head
        prev = None
        while curr:
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next
        return prev
