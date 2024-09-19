from typing import Optional
import unittest
from dataclasses import dataclass

from linked_list_cycle import ListNode, Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: Optional[ListNode]
            expected: bool

        a = ListNode(0)
        b = ListNode(0)
        return [
            Case(given=None, expected=False),
            Case(given=a.add(1), expected=False),
            Case(given=b.add(3, b), expected=True),
            Case(given=ListNode(2).add(3).add(2), expected=False),
        ]

    def test_linked_list_cycle(self):
        for case in self.cases():
            sol = Solution()
            print(case.given, case.expected)
            actual = sol.hasCycle(case.given)
            assert actual == case.expected
