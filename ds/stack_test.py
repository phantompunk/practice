from ds.stack import Stack


def test_init_stack():
    stack = Stack()
    assert stack.empty()
    assert stack.pop() is None
    assert stack.search("A") == -1


def test_push_pop_peek():
    stack = Stack()
    stack.push("A")
    assert stack.peek() == "A"
    assert stack.pop() == "A"
    assert stack.peek() is None


def test_search():
    stack = Stack()
    stack.push("A")
    stack.push("B")
    stack.push("C")
    assert stack.search("A") == 3
    assert stack.search("D") == -1
