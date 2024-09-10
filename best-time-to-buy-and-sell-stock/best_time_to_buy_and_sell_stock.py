class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        profit = 0
        for i, x in enumerate(prices):
            for y in prices[i+1:]:
                profit = max(y-x, profit)
        return profit
