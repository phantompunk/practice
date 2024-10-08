from typing import Optional
import unittest
from dataclasses import dataclass

from maximum_depth_of_binary_tree import Solution, TreeNode


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: list[int|None]
            expected: int

        return [
            Case(given=[3,9,20,None,None,15,7], expected=3),
            Case(given=[1,None,2], expected=2),
        ]
    
    def toTreeNode(self, lst:list) -> Optional[TreeNode]:
        if not lst:
            return None

        root = TreeNode(lst[0])
        queue = [root]

        i = 1
        while i < len(lst):
            node = queue.pop(0)

            if i < len(lst) and lst[i] is not None:
                node.left = TreeNode(lst[i])
                queue.append(node.left)
            i += 1

            if i < len(lst) and lst[i] is not None:
                node.right = TreeNode(lst[i])
                queue.append(node.right)
            i += 1

        return root

    def test_maximum_depth_of_binary_tree(self):
        for case in self.cases():
            sol = Solution()
            given = self.toTreeNode(case.given)
            actual = sol.maxDepth(given) 
            assert actual == case.expected  
