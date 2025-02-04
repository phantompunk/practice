import unittest

from contains_duplicate import Solution


class TestContainsDuplicate(unittest.TestCase):
    def test_empty_array(self):
        self.assertFalse(Solution().containsDuplicate([]))

    def test_single_element_array(self):
        self.assertFalse(Solution().containsDuplicate([1]))

    def test_no_duplicates(self):
        self.assertFalse(Solution().containsDuplicate([1, 2, 3, 4, 5]))

    def test_one_duplicate(self):
        self.assertTrue(Solution().containsDuplicate([1, 2, 3, 1]))

    def test_multiple_duplicates(self):
        self.assertTrue(Solution().containsDuplicate([1, 2, 3, 1, 2, 4]))

    def test_all_duplicates(self):
        self.assertTrue(Solution().containsDuplicate([1, 1, 1, 1, 1]))

    def test_negative_numbers(self):
        self.assertTrue(Solution().containsDuplicate([-1, -2, -3, -1]))

    def test_mixed_positive_negative(self):
        self.assertTrue(Solution().containsDuplicate([-1, 2, -3, 2]))

    def test_large_numbers(self):
        self.assertFalse(
            Solution().containsDuplicate([10**9, 2 * 10**9, 3 * 10**9])
        )  # Check for potential overflow

    def test_large_array_no_duplicates(self):
        self.assertFalse(
            Solution().containsDuplicate(list(range(10000)))
        )  # Performance test

    def test_large_array_with_duplicate(self):
        nums = list(range(9999))
        nums.append(9998)  # Introduce a duplicate
        self.assertTrue(Solution().containsDuplicate(nums))

    def test_zero_present(self):
        self.assertTrue(Solution().containsDuplicate([0, 1, 2, 0]))

    def test_duplicate_at_beginning(self):
        self.assertTrue(Solution().containsDuplicate([1, 1, 2, 3]))

    def test_duplicate_at_end(self):
        self.assertTrue(Solution().containsDuplicate([1, 2, 3, 1]))

    def test_duplicate_in_middle(self):
        self.assertTrue(Solution().containsDuplicate([1, 2, 1, 3]))
