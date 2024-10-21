from typing import Optional
import unittest
from dataclasses import dataclass

from diameter_of_binary_tree import Solution, TreeNode


class TestSolution(unittest.TestCase):
    def to_tn(self, lst: list[int | None]):
        if not lst:
            return None
        nodes = [TreeNode(n) if n else None for n in lst]
        kids = nodes[::-1]
        root = kids.pop()
        for node in nodes:
            if node:
                if kids:
                    node.left = kids.pop()
                if kids:
                    node.right = kids.pop()
        return root

    def cases(self):
        @dataclass
        class Case:
            given: list[int | None]
            expected: int

        return [
            Case(given=[1, 2], expected=1),
            Case(given=[1, 2, 3], expected=2),
            Case(given=[1, 2, 3, 4, 5], expected=3),
            Case(given=[1, 2, 3, 4, 5, 6], expected=4),
        ]

    def test_diameter_of_binary_tree(self):
        for case in self.cases():
            sol = Solution()
            given = self.to_tn(case.given)
            actual = sol.diameterOfBinaryTree(given)
            assert actual == case.expected
