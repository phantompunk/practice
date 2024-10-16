# Definition for a binary tree node.
from __future__ import annotations
from typing import Optional


class TreeNode:
    left: TreeNode | None
    right: TreeNode | None

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root:
            return None

        queue = [root]
        while queue:
            curr = queue.pop(0)
            if curr.left:
                queue.append(curr.left)
            if curr.right:
                queue.append(curr.right)

            tmp = curr.left
            curr.left = curr.right
            curr.right = tmp
        return root
