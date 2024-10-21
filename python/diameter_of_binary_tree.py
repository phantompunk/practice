# Definition for a binary tree node.
from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        diameter = 0
        if root is None:
            return diameter

        queue = []
        if root.left:
            diameter += 1
            queue.append(root.left)
        if root.right:
            diameter += 1
            queue.append(root.right)

        while queue:
            curr = queue.pop(0)
            print(curr.left)
            if curr.left or curr.right:
                diameter += 1
                print("dia", diameter)
            if curr.left:
                queue.append(curr.left)
            if curr.right:
                queue.append(curr.right)

        return diameter
