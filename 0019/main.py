import sys
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))

from utils.utils import run_tests, ListNode, make_list


class Solution:
    def removeNthFromEnd(self, head: ListNode | None, n: int) -> ListNode | None:
        num_elements = 0
        node = head
        while node != None:
            node = node.next
            num_elements += 1

        if num_elements == n:
            return head.next

        node = head
        for _ in range(num_elements - n - 1):
            node = node.next

        node.next = node.next.next
        return head


list1, list1result = make_list([1, 2, 3, 4, 5]), make_list([1, 2, 3, 5])
list2, list2result = make_list([1]), make_list([])
list3, list3result = make_list([1, 2]), make_list([1])
run_tests(
    Solution().removeNthFromEnd,
    [
        {"input": [list1, 2], "expected": list1result},
        {"input": [list2, 1], "expected": list2result},
        {"input": [list3, 1], "expected": list3result},
    ],
)
