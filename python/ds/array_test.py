from ds.array import ArrayList


def test_array_size_and_capacity():
    arr = ArrayList()
    assert arr.size() == 0
    assert arr.capacity() == 2

    arr = ArrayList(10)
    assert arr.size() == 0
    assert arr.capacity() == 10


def test_add_to_array():
    arr = ArrayList()
    arr.add("a")
    arr.add("b")
    arr.add("c")
    assert arr.size() == 3
    assert arr.capacity() == 4


def test_add_all_to_array():
    arr = ArrayList()
    items = [1, 2, 3, 4]
    arr.add_all(items)
    assert arr.size() == 4


def test_remove_from_array():
    arr = ArrayList()
    items = [1, 2, 3, 4]
    arr.add_all(items)

    arr.remove(3)
    assert arr.size() == 3

    arr.remove_at(0)
    assert arr.size() == 2


def test_get_from_array():
    arr = ArrayList()
    items = [1, 2, 3, 4]
    arr.add_all(items)

    assert arr.get(0) == 1


def test_set_item_in_array():
    arr = ArrayList()
    items = [1, 2, 3, 4]
    arr.add_all(items)
    arr.set(0, 9)
    arr.set(3, 9)
    assert arr.get(0) == 9
    assert arr.get(3) == 9
    assert arr.size() == 4


def test_clear_array():
    arr = ArrayList()
    items = [1, 2, 3, 4]
    arr.add_all(items)
    arr.clear()

    assert arr.size() == 0
