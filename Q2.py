class Solution:
    def maxSumOfSquares(self, num: int, sum: int) -> str:
        if sum > num:
            return ""
        
        res = 0
        current_sum = sum
        partial_sum = 0
        count_digits = num
        while True:
            if count_digits == 0:
                break
            for i in range(9,0,-1):
                if current_sum // i > 0:
                    current_sum -= i
                    partial_sum += i
                    res += i * 10**count_digits
                    count_digits -= 1
                    break

        # 2 3

        # 3 0

        # 3 20

        # 9 9 2

        return res


s = Solution()
print(s.maxSumOfSquares(2,3) == "30")
print(s.maxSumOfSquares(2,17) == "98")
print(s.maxSumOfSquares(1,10) == "")