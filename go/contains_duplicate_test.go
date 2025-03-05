import unittest
from contains_duplicate import Solution


class TestSolution(unittest.TestCase):
  def test_empty_strings(self):
    self.assertTrue(Solution())
