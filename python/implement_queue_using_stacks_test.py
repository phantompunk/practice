import unittest
from dataclasses import dataclass

from implement_queue_using_stacks import MyQueue


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: int | None
            method: str
            expected: int | bool | None

        return [
            Case(given=None, method="empty", expected=True),
            Case(given=1, method="push", expected=None),
            Case(given=2, method="push", expected=None),
            Case(given=None, method="peek", expected=1),
            Case(given=None, method="pop", expected=1),
            Case(given=None, method="empty", expected=False),
            Case(given=None, method="pop", expected=2),
            Case(given=None, method="peek", expected=None),
            Case(given=None, method="pop", expected=None),
        ]

    def test_my_queue(self):
        sol = MyQueue()
        for case in self.cases():
            match case.method:
                case "push":
                    actual = sol.push(case.given)
                case "pop":
                    actual = sol.pop()
                case "empty":
                    actual = sol.empty()
                case "peek":
                    actual = sol.peek()

            assert case.expected == actual
