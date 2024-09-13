import unittest
from dataclasses import dataclass

from valid_anagram import Solution



class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            givenA: str
            givenB: str
            expected: bool

        return [
            Case(givenA="anagram", givenB="nagaram", expected=True),
            Case(givenA="pins", givenB="spin", expected=True),
            Case(givenA="rat", givenB="car", expected=False),
            Case(givenA="dormitory", givenB="dirtyroom", expected=True),
            Case(givenA="astronomer", givenB="moonstarer", expected=True),
            Case(givenA="harrypotterandthegobletoffire", givenB="frontedelaboratetrophyfreight", expected=True),
            Case(givenA="harrypotterandthegobletoffire", givenB="harrypotterandthegobletsoffire", expected=False),
        ]

    def test_is_palindrome(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.isAnagram(case.givenA, case.givenB)

