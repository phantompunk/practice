import unittest
from dataclasses import dataclass

from maximum_subarray import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[int]
            expected: int

        return [
            Case(given=[-2,1,-3,4,-1,2,1,-5,4], expected=6),
            Case(given=[4], expected=4),
            Case(given=[5,4,-1,7,8], expected=23),
        ]

    def test_maximum_subarray(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.maxSubArray(case.given)  
            assert actual == case.expected  
