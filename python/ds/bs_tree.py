from __future__ import annotations
from typing import Any


class Node:
    left: Node | None
    right: Node | None
    data: Any

    def __init__(self, data, left=None, right=None) -> None:
        self.data = data
        self.left = None
        self.right = None


class BSTree:
    root: Node | None
    count: int

    def __init__(self):
        self.root = None
        self.count = 0

    def is_empty(self) -> bool:
        return self.get_root() is None

    def size(self) -> int:
        return self.count

    def height(self) -> int:
        return 0

    def insert(self, data) -> None:
        new_node = Node(data)
        if not self.root:
            self.root = new_node
            return
        current = self.root
        while True:
            if data < current.data:
                if current.left:
                    current = current.left
                else:
                    current.left = new_node
                    break
            else:
                if current.right:
                    current = current.right
                else:
                    current.right = new_node
                    break

    def contains(self, data) -> bool:
        if self.is_empty():
            return False
        curr = self.root
        while curr:
            if curr.data == data:
                return True
            curr = curr.left if data < curr.data else curr.right
        return False

    def remove(self, data) -> None:
        # find the node
        curr = self.root
        parent = curr
        while curr and curr.data != data:
            parent = curr
            curr = curr.left if data < curr.data else curr.right

        if curr is None:
            return

        # remove leaf node
        if not curr.left and not curr.right:
            if curr == self.root:
                self.root = None
            elif parent.left == curr:
                parent.left = None
            else:
                parent.right = None
        # remove node w/ 1 child
        elif not curr.left or not curr.right:
            child = curr.left if curr.left else curr.right
            if curr == self.root:
                self.root = child
            elif parent.left == curr:
                parent.left = child
            else:
                parent.right = child
        # remove node w/ 2 children
        else:
            succ = curr.right
            succ_parent = curr

            while succ.left:
                succ_parent = succ
                succ = succ.left

            curr.val = succ.val

            if succ_parent.left == succ:
                succ_parent.left = succ.right
            else:
                succ_parent.right = succ.right

    def get_root(self) -> Node | None:
        return self.root


def in_order_traversal(root: Node | None, path: list = []) -> list:
    if not root:
        return path

    in_order_traversal(root.left, path)
    path.append(root.data)
    in_order_traversal(root.right, path)
    return path


def pre_order_traversal(root: Node | None, path: list = []) -> list:
    if not root:
        return path

    path.append(root.data)
    pre_order_traversal(root.left)
    pre_order_traversal(root.right)

    return path


def post_order_traversal(root: Node | None, path: list = []) -> list:
    if not root:
        return path

    post_order_traversal(root.left, path)
    post_order_traversal(root.right, path)
    path.append(root.data)

    return path


def level_order_traversal(root: Node | None, path: list = []) -> list:
    queue = []
    if not root:
        return path

    queue = [root]
    while queue:
        curr = queue.pop(0)
        path.append(curr.data)
        if curr.left:
            queue.append(curr.left)
        if curr.right:
            queue.append(curr.right)
    return path
