from typing import List
import unittest
from dataclasses import dataclass

from flood_fill import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            image: List[List[int]]
            sr: int
            sc: int
            color: int
            expected: List[List[int]]

        return [
            Case(
                image=[[1, 1, 1], [1, 1, 0], [1, 0, 1]],
                sr=1,
                sc=1,
                color=2,
                expected=[[2, 2, 2], [2, 2, 0], [2, 0, 1]],
            ),
            Case(
                image=[[1, 1, 1, 1], [1, 1, 1, 0], [1, 0, 1, 1]],
                sr=1,
                sc=1,
                color=2,
                expected=[[2, 2, 2, 2], [2, 2, 2, 0], [2, 0, 2, 2]],
            ),
            Case(
                image=[[0, 0, 0], [0, 0, 0]],
                sr=0,
                sc=0,
                color=0,
                expected=[[0, 0, 0], [0, 0, 0]],
            ),
            Case(
                image=[[0,0,0],[0,1,0]],
                sr=0,
                sc=0,
                color=2,
                expected=[[2,2,2],[2,1,2]],
            ),
            Case(
                image=[[0,0,1],[1,0,1]],
                sr=0,
                sc=0,
                color=2,
                expected=[[2,2,1],[1,2,1]],
            ),
        ]

    def test_flood_fill(self):
        for case in self.cases():
            sol = Solution()
            actual = sol.floodFill(case.image, case.sr, case.sc, case.color)
            assert actual == case.expected
