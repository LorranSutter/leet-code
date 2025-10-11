class Solution:
    def scoreBalance(self, s: str) -> bool:
        s = [ord(letter) - 96 for letter in s]
        l = len(s)

        sum_init, sum_end = s[0], s[-1]
        i, j = 0, l - 1
        while True:
            if j - i == 1:
                return sum_init == sum_end
            if sum_init > sum_end:
                j -= 1
                sum_end += s[j]
            else:
                i += 1
                sum_init += s[i]


s = Solution()
print(s.scoreBalance("adcb") == True)
print(s.scoreBalance("bace") == True)
print(s.scoreBalance("baece") == True)
