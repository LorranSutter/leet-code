class Solution:
    def lexSmallest(self, s: str) -> str:
        smallest = s

        for i in range(len(s)):
            s_revers_front = s[:i][::-1] + s[i:]
            if s_revers_front < smallest:
                smallest = s_revers_front
            s_revers_back = s[:-i] + s[-i:][::-1]
            if s_revers_back < smallest:
                smallest = s_revers_back

        return smallest


s = Solution()
print(s.lexSmallest("dcab") == "acdb")
print(s.lexSmallest("abba") == "aabb")
print(s.lexSmallest("zxy") == "xzy")
print(s.lexSmallest("lk") == "kl")
