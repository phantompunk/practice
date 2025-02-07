import unittest

from group_anagrams import Solution


class TestGroupAnagrams(unittest.TestCase):
    def test_empty_list(self):
        self.assertEqual(Solution().groupAnagrams([]), [])

    def test_single_word(self):
        self.assertEqual(Solution().groupAnagrams(["cat"]), [["cat"]])

    def test_basic_anagrams(self):
        expected = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
        actual = Solution().groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
        # Sort the inner lists to handle any order within a group
        for i in range(len(expected)):
            expected[i].sort()
        for i in range(len(actual)):
            actual[i].sort()
        self.assertEqual(actual, expected)

    def test_with_empty_string(self):
        self.assertEqual(
            Solution().groupAnagrams([""]), [[""]]
        )  # Empty string is its own anagram

    def test_long_strings(self):
        self.assertEqual(
            Solution().groupAnagrams(["aaaaaaaaaa", "aaaaaaaaab"]),
            [["aaaaaaaaaa"], ["aaaaaaaaab"]],
        )

    def test_with_unicode(self):
        self.assertEqual(Solution().groupAnagrams(["你好", "好你"]), [["你好", "好你"]])

    def test_with_special_characters(self):
        self.assertEqual(Solution().groupAnagrams(["!@#", "#!@"]), [["!@#", "#!@"]])

    def test_with_numbers(self):
        self.assertEqual(Solution().groupAnagrams(["123", "321"]), [["123", "321"]])

    def test_mixed_characters(self):
        self.assertEqual(Solution().groupAnagrams(["abc1", "1cba"]), [["abc1", "1cba"]])

    # def test_with_duplicates(self):
    #     expected = [
    #         ["eat", "tea", "ate"],
    #         ["tan", "nat"],
    #         ["bat"],
    #         ["eat", "tea", "ate"],
    #     ]  # Duplicate anagram groups
    #     actual = Solution().groupAnagrams(
    #         ["eat", "tea", "tan", "ate", "nat", "bat", "eat", "tea", "ate"]
    #     )
    #     for i in range(len(expected)):
    #         expected[i].sort()
    #     for i in range(len(actual)):
    #         actual[i].sort()
    #     self.assertEqual(actual, expected)
    #
    # def test_large_input(self):
    #     strs = [
    #         "".join(sorted(str(i))) for i in range(1000)
    #     ]  # Generate a large list of anagrams
    #     self.assertEqual(
    #         len(Solution().groupAnagrams(strs)), 1000
    #     )  # Should be 1000 groups
    #
    # def test_with_mixed_case(self):
    #     expected = [["Eat", "tea", "ate"], ["tan", "nat"], ["bat"]]  # Case sensitive
    #     actual = Solution().groupAnagrams(["Eat", "tea", "tan", "ate", "nat", "bat"])
    #     for i in range(len(expected)):
    #         expected[i].sort()
    #     for i in range(len(actual)):
    #         actual[i].sort()
    #     self.assertEqual(actual, expected)
