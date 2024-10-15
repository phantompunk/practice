from __future__ import with_statement

from ds.bs_tree import BSTree, in_order_traversal, level_order_traversal, post_order_traversal, pre_order_traversal


def test_init_bst():
    bst = BSTree()
    assert bst.size() == 0
    assert bst.height() == 0
    assert bst.get_root() is None
    assert bst.is_empty()


def test_add_bst():
    bst = BSTree()
    bst.insert(10)
    bst.insert(20)
    assert bst.get_root().data == 10
    assert bst.contains(20)
    assert not bst.is_empty()


def test_bst_contains():
    bst = BSTree()
    bst.insert(10)
    assert bst.contains(10)
    assert not bst.contains(11)


def test_bst_remove():
    bst = BSTree()
    bst.insert(10)
    bst.remove(10)
    assert not bst.contains(10)
    assert bst.is_empty()


def test_bst_inorder():
    bst = BSTree()
    bst.insert(10)
    bst.insert(5)
    bst.insert(15)
    bst.insert(11)
    assert in_order_traversal(bst.root) == [5, 10, 11, 15]


def test_bst_preorder():
    bst = BSTree()
    bst.insert(10)
    bst.insert(5)
    bst.insert(15)
    assert pre_order_traversal(bst.root) == [10, 5, 15]


def test_bst_postorder():
    bst = BSTree()
    bst.insert(10)
    bst.insert(5)
    bst.insert(15)
    assert post_order_traversal(bst.root) == [5, 10, 15]
    
def test_bst_levelorder():
    bst = BSTree()
    bst.insert(10)
    bst.insert(5)
    bst.insert(15)
    assert level_order_traversal(bst.root) == [10, 5, 15]






















