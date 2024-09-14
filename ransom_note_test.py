import unittest
from dataclasses import dataclass

from ransom_note import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            givenA: str
            givenB: str
            expected: bool

        return [
            Case(givenA="abc", givenB="cba", expected=True),
            Case(givenA="abcc", givenB="abc", expected=False),
            Case(givenA="abccd", givenB="dabc", expected=False),
            Case(givenA="abccd", givenB="dabcc", expected=True),
        ]

    def test_can_construct(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.canConstruct(case.givenA, case.givenB)
