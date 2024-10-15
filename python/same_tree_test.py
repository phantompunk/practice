from typing import Optional
import unittest
from dataclasses import dataclass

from same_tree import Solution, TreeNode


class TestSolution(unittest.TestCase):
    def to_tn(self, nodes: list) -> Optional[TreeNode]:
        if not nodes:
            return None
        root = TreeNode(nodes[0])
        queue = [root]
        index = 1
        while queue and index < len(nodes):
            node = queue.pop(0)
            if index+1 < len(nodes):
                node.left = TreeNode(nodes[index+1])
                queue.append(node.left)
                index+=1
            if index+1 < len(nodes):
                node.right = TreeNode(nodes[index+1])
                queue.append(node.right)
                index+=1
        return root

    def cases(self):
        @dataclass
        class Case:
            given: list[int | None]
            target: list[int | None]
            expected: bool

        return [
            Case(given=[0], target=[], expected=False),
            Case(given=[], target=[], expected=True),
            Case(given=[1], target=[], expected=False),
            Case(given=[1, 2, 3], target=[1, 2, 3], expected=True),
            Case(given=[1, 2, None, 4], target=[1, 2, None, 4], expected=True),
            Case(given=[1, 2, None, 4,5], target=[1, 2, 3], expected=False),
            Case(given=[1, 2, 3, 4], target=[1, 2, 3], expected=False),
            Case(given=[1, 2, 3, 4], target=[1, 2, 3], expected=False),
        ]

    def test_same_tree(self):
        for case in self.cases():
            sol = Solution()
            given = self.to_tn(case.given)
            target = self.to_tn(case.target)
            actual = sol.isSameTree(given, target)
            assert actual == case.expected
