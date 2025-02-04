import unittest

from valid_anagram import Solution


class TestSolution(unittest.TestCase):
    def test_empty_strings(self):
        self.assertTrue(Solution().isAnagram("", ""))

    def test_basic_anagram(self):
        self.assertTrue(Solution().isAnagram("anagram", "nagaram"))

    def test_not_anagram(self):
        self.assertFalse(Solution().isAnagram("rat", "car"))

    def test_different_lengths(self):
        self.assertFalse(Solution().isAnagram("abc", "ab"))

    def test_case_sensitive(self):  # Anagrams are usually case-sensitive
        self.assertFalse(Solution().isAnagram("Anagram", "nagaram"))

    def test_punctuation(self):
        self.assertFalse(Solution().isAnagram("hello!", "olleh"))  # Punctuation matters

    def test_long_strings(self):
        s = "a" * 100000  # Long string for performance testing
        t = "a" * 100000
        self.assertTrue(Solution().isAnagram(s, t))

    def test_long_strings_not_anagram(self):
        s = "a" * 100000
        t = "a" * 99999 + "b"
        self.assertFalse(Solution().isAnagram(s, t))

    def test_unicode_characters(self):
        self.assertTrue(
            Solution().isAnagram("你好", "好你")
        )  # Example with Chinese characters

    def test_string_with_numbers(self):
        self.assertFalse(Solution().isAnagram("123", "12"))

    def test_string_with_special_characters(self):
        self.assertTrue(Solution().isAnagram("!@#$", "$#@!"))

    def test_one_string_empty(self):
        self.assertFalse(Solution().isAnagram("abc", ""))

    def test_one_string_long(self):
        self.assertFalse(Solution().isAnagram("a", "aaaa"))

    def test_identical_strings(self):
        self.assertTrue(Solution().isAnagram("apple", "apple"))

    def test_same_length_diff(self):
        self.assertFalse(Solution().isAnagram("aacc", "ccac"))
