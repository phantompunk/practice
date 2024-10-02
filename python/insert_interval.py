from typing import List


class Solution:
    def insert(
        self, intervals: List[List[int]], newInterval: List[int]
    ) -> List[List[int]]:
        idx = 0
        length = len(intervals)
        newIntervals: List[List[int]] = []

        while idx < length and intervals[idx][1] < newInterval[0]:
            newIntervals.append(intervals[idx])
            idx += 1

        # merge
        while idx < length and intervals[idx][0] <= newInterval[1]:
            newInterval[0] = min(newInterval[0], intervals[idx][0])
            newInterval[1] = max(newInterval[1], intervals[idx][1])
            idx += 1
        newIntervals.append(newInterval)

        # add rest
        while idx < length:
            newIntervals.append(intervals[idx])
            idx += 1

        return newIntervals

        # newStart = newInterval[0]
        # newEnd = newInterval[1]

        # index = 0
        # curr = intervals[index]
        # while curr and index < len(intervals) and curr[1] < newStart:
        #     newIntervals.append(curr)
        #     index+=1
        #     curr = intervals[index]

        # overlap = False
        # addInterval = False
        # for idx, interval in enumerate(intervals):
        #     start = interval[0]
        #     end = interval[1]
        #
        #     if not addInterval and end >= newStart and newEnd >= start:
        #         overlap = True
        #     elif not addInterval and end < newStart:
        #         print("add curr", interval)
        #         newIntervals.append(interval)
        #     elif not addInterval and newEnd < start:
        #         print("add new", newInterval)
        #         newIntervals.append(newInterval)
        #         newIntervals.append(interval)
        #         addInterval = True
        #     else:
        #         print("add rest", interval)
        #         newIntervals.append(interval)
        #
        # if not addInterval and not overlap:
        #     print("add interval", newInterval)
        #     newIntervals.append(newInterval)
        #
        # index = 0
        # final = []
        # while not addInterval and index + 1 < len(newIntervals):
        #     start = newIntervals[index][0]
        #     end = newIntervals[index][1]
        #     if not addInterval and end >= newStart and newEnd >= start:
        #         print("insert")
        #         final.append([start, newEnd])
        #         addInterval = True
        #     else:
        #         final.append(newIntervals[index])
        #     index += 1
        # # for idx, interval in enumerate(intervals[len(newIntervals):]):
        # #     if newEnd < interval[0]:
        # #         newIntervals.append(newInterval)
        #
        # return final
