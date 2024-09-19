from typing import Optional
import unittest
from dataclasses import dataclass

from merge_two_sorted_lists import ListNode, Solution


def areIdentical(head1, head2):
    while head1 is not None and head2 is not None:
        if head1.val != head2.val:
            return False

        # Move to next nodes in both lists
        head1 = head1.next
        head2 = head2.next

    # If linked lists are identical,
    # both 'head1' and 'head2' must be None
    return head1 is None and head2 is None


def test_merge_two_list_empty():
    s = Solution()
    a = ListNode(1)
    b = ListNode(2)

    c = ListNode(1)
    c.next = ListNode(2)
    actual = s.mergeTwoLists(a, b)
    assert actual is not None
    assert actual.val == c.val
    assert actual.next and c.next
    assert actual.next.val == c.next.val


def test_merge_two_list():
    s = Solution()
    a = ListNode(0).add(2)
    b = ListNode(1).add(3)
    actual = s.mergeTwoLists(a, b)

    expected = ListNode(0).add(1)
    # expected.next = ListNode(1)
    assert actual is not None
    assert actual.val == expected.val
    assert actual.next and expected.next
    assert actual.next.val == expected.next.val
    # assert areIdentical(actual, expected)


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            givenA: Optional[ListNode]
            givenB: Optional[ListNode]
            expected: Optional[ListNode]

        return [
            Case(givenA=None, givenB=ListNode(), expected=ListNode()),
            Case(givenA=ListNode(), givenB=ListNode(), expected=ListNode().add()),
            Case(givenA=ListNode(2), givenB=ListNode(1), expected=ListNode(1).add(2)),
            Case(
                givenA=ListNode(2).add(4),
                givenB=ListNode(1).add(3),
                expected=ListNode(1).addall([2, 3, 4]),
            ),
        ]

    def test_can_construct(self):
        for case in self.cases():
            sol = Solution()
            assert case.expected == sol.mergeTwoLists(case.givenA, case.givenB)
