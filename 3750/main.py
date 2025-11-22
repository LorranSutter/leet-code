class Solution:
    def minimumFlips(self, n: int) -> int:
        s = str(bin(n))[2:]
        flip_count = 0
        for i in range(len(s)//2):
            if s[i] != s[-i-1]:
                flip_count += 1
        return flip_count*2

s = Solution()
print(s.minimumFlips(7) == 0)
print(s.minimumFlips(10) == 4)
print(s.minimumFlips(11) == 2)
print(s.minimumFlips(17) == 0)
print(s.minimumFlips(19) == 2)