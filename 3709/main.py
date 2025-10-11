import bisect


class ExamTracker:

    def __init__(self):
        self.times = [0]
        self.scores = [0]

    def record(self, time: int, score: int) -> None:
        # TODO Have to use prefix sum for the scores
        self.times.append(time)
        self.scores.append(score + self.scores[-1])

    def totalScore(self, startTime: int, endTime: int) -> int:
        startTime_idx = bisect.bisect_left(self.times, startTime)
        endTime_idx = bisect.bisect_right(self.times, endTime)

        # total_score = 0
        # for i in range(startTime_idx , endTime_idx + 1):
        #     if startTime <= self.times[i] <= endTime:
        #         total_score += self.scores[i]

        return self.scores[endTime_idx - 1] - self.scores[startTime_idx - 1]

    def search(self, time: int) -> int:
        left, mid = 0, 0
        right = len(self.times) - 1

        while left <= right:
            mid = left + (right - left) // 2

            if self.times[mid] == time:
                return mid
            if self.times[mid] < time:
                left = mid + 1
            else:
                right = mid - 1

        return mid


# Your ExamTracker object will be instantiated and called as such:
obj = ExamTracker()
obj.record(1, 98)
print(obj.totalScore(1, 1) == 98)
obj.record(5, 99)
print(obj.totalScore(1, 3) == 98)
print(obj.totalScore(1, 5) == 197)
print(obj.totalScore(3, 4) == 0)
print(obj.totalScore(2, 5) == 99)
