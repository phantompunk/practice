# The isBadVersion API is already defined for you.
def isBadVersion(v: int) -> bool:
    print("Version", v)
    if v >= 1274:
        return True
    if 50 >= v and v >= 45:
        return True
    if v == 5 or v == 4:
        return True
    if v == 1:
        return True
    return False


class Solution:
    def firstBadVersion(self, n: int) -> int:
        lo, hi = 1, n
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if isBadVersion(mid):
                hi = mid 
            else:
                lo = mid+1
        return lo
