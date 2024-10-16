from typing import Optional
import unittest
from dataclasses import dataclass

from invert_binary_tree import Solution, TreeNode


class TestSolution(unittest.TestCase):
    def to_l(self, root: Optional[TreeNode]) -> list[int]:
        if not root:
            return []
        values = []
        curr = root
        queue = [root]
        while queue:
            curr = queue.pop(0)
            values.append(curr.val)
            if curr.left:
                queue.append(curr.left)
            if curr.right:
                queue.append(curr.right)
        return values

    def to_tn(self, nodes: list) -> Optional[TreeNode]:
        if not nodes:
            return None
        root = TreeNode(nodes[0])
        queue = [root]
        index = 1
        while queue and index < len(nodes):
            node = queue.pop(0)
            if index < len(nodes):
                node.left = TreeNode(nodes[index])
                queue.append(node.left)
                index += 1
            if index < len(nodes):
                node.right = TreeNode(nodes[index])
                queue.append(node.right)
                index += 1
        return root

    def cases(self):
        @dataclass
        class Case:
            given: list[int | None]
            expected: list[int | None]

        return [
            Case(given=[], expected=[]),
            Case(given=[2, 1, 3], expected=[2, 3, 1]),
            Case(given=[4, 2, 7, 1, 3, 6, 9], expected=[4, 7, 2, 9, 6, 3, 1]),
        ]

    def test_invert_binary_tree(self):
        for case in self.cases():
            sol = Solution()
            given = self.to_tn(case.given)
            actual = sol.invertTree(given)
            assert self.to_l(actual)== case.expected
