from ds.queue import Queue


def test_init_queue():
    queue = Queue()
    assert queue.size() == 0
    assert queue.poll() is None
    assert queue.peek() is None


def test_poll_peek():
    queue = Queue()
    queue.offer("A")
    queue.offer("B")
    assert queue.size() == 2
    assert queue.peek() == "A"
    assert queue.poll() == "A"
    assert queue.peek() == "B"
    assert queue.size() == 1

def test_remove_from_queue():
    q = Queue()
    q.offer("A")
    q.offer("B")
    q.offer("C")
    assert q.poll() == "A"
    assert q.poll() == "B"
    assert q.peek() == "C"
