import unittest
from dataclasses import dataclass

from first_bad_version import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: int
            expected: int

        return [
            Case(given=1, expected=1),
            Case(given=5, expected=4),
            Case(given=50, expected=45),
            Case(given=2474, expected=1274),
        ]

    def test_first_bad_version(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.firstBadVersion(case.given)
            assert actual == case.expected
