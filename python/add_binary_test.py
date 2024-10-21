import unittest
from dataclasses import dataclass

from add_binary import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            givenA: str
            givenB: str
            expected: str

        return [
            Case(givenA="11", givenB="1", expected="100"),
        ]

    def test_add_binary(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.addBinary(case.givenA, case.givenB)  
            assert actual == case.expected  
