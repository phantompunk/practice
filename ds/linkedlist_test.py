from ds.linkedlist import LinkedList


def test_init_linkedlist():
    ll = LinkedList()
    ll.clear()
    assert ll
    assert ll.size() == 0
    assert ll.peek() == None
    assert ll.remove() == None
    assert ll.size() == 0


def test_linkedlist():
    ll = LinkedList()

    ll.add("a")
    ll.add_first("b")
    ll.add_last("c")

    assert ll.size() == 3
    assert ll.peek_first() == "b"
    assert ll.peek_last() == "c"

    ll.add_at("d", 1)
    ll.add_last("e")
    assert ll.size() == 5
    assert ll.peek_last() == "e"
    assert ll.index_of("d") == 1


def test_remove_from_linkedlist():
    ll = LinkedList()

    ll.add("a")
    ll.add_first("b")
    ll.add_last("c")

    assert ll.remove() == "b"
    assert ll.remove_first() == "a"
    assert ll.remove_last() == "c"
    assert ll.size() == 0
