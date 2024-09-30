from typing import Optional
import unittest
from dataclasses import dataclass

from reverse_linked_list import Solution
from utils import ListNode


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: Optional[ListNode]
            expected: Optional[ListNode]

        return [
            Case(given=ListNode(1), expected=ListNode(1)),
            Case(given=ListNode(1, ListNode(2)), expected=ListNode(2, ListNode(1))),
            Case(
                given=ListNode(1, ListNode(2, ListNode(3))),
                expected=ListNode(3, ListNode(2, ListNode(1))),
            ),
        ]

    def test_reverse_linked_list(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.reverseList(case.given)
            assert actual == case.expected
