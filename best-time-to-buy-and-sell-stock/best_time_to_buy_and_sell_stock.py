class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        profit = 0
        lowest = prices[0]
        for x in prices:
            diff = x - lowest
            profit = max(diff, profit)
            lowest = min(x, lowest)
        return profit
