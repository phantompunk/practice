import unittest
from dataclasses import dataclass
from best_time_to_buy_and_sell_stock import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[int]
            expected: int

        return [
            Case(given=[], expected=0),
            Case(given=[1, 1, 1], expected=0),
            Case(given=[1, 7, 3], expected=6),
            Case(given=[7, 1, 5, 3, 6, 4], expected=5),
            Case(given=[7, 6, 4, 3, 1], expected=0),
            Case(given=[1, 6, 2, 3, 7], expected=6),
            Case(given=[7, 6, 7, 8, 3], expected=2),
            Case(given=[3,5,6,7,2,9], expected=7),
        ]

    def test_max_profit(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.maxProfit(case.given)
