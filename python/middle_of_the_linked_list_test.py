import unittest
from dataclasses import dataclass

from middle_of_the_linked_list import Solution
from utils import ListNode


class TestSolution(unittest.TestCase):
    def to_ll(self, values):
        if not values:
            return None
        head = ListNode(values[0])
        current = head
        for value in values[1:]:
            current.next = ListNode(value)
            current = current.next
        return head

    def cases(self):
        @dataclass
        class Case:
            given: list[int]
            expected: list[int]

        return [
            Case(given=[1], expected=[1]),
            Case(given=[1, 2, 3], expected=[2, 3]),
            Case(given=[1, 2, 3, 4, 5], expected=[3, 4, 5]),
            Case(given=[1, 2, 3, 4, 5, 6], expected=[4, 5, 6]),
        ]

    def test_middle_of_the_linked_list(self):
        for case in self.cases():
            sol = Solution()
            given = self.to_ll(case.given)
            actual = sol.middleNode(given)
            expected = self.to_ll(case.expected)
            assert actual == expected
