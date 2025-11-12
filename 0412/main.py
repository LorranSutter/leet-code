from typing import List


class Solution:
    def fizzBuzz(self, n: int) -> List[str]:
        answer = ["" for _ in range(n)]
        countThree = 0
        countFive = 0

        for i in range(n):
            countThree += 1
            countFive += 1

            if countThree == 3 and countFive == 5:
                answer[i] = "FizzBuzz"
                countThree = 0
                countFive = 0
                continue
            elif countThree == 3:
                answer[i] = "Fizz"
                countThree = 0
                continue
            elif countFive == 5:
                answer[i] = "Buzz"
                countFive = 0
                continue
            else:
                answer[i] = str(i+1)
            
        return answer