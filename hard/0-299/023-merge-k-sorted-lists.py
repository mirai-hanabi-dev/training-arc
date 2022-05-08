from dataclasses import dataclass, field
from queue import PriorityQueue
from typing import Any


@dataclass(order=True)
class PrioritizedItem:
    priority: int
    item: Any = field(compare=False)


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeKLists(self, lists):
        q = PriorityQueue()
        head = ListNode()
        current = head

        for headNode in lists:
            if headNode:
                q.put(PrioritizedItem(
                    priority=headNode.val,
                    item=headNode.next))

        if q.empty():
            return None

        while not q.empty():
            item = q.get()
            val = item.priority
            nextNode = item.item

            current.val = val
            if nextNode:
                q.put(PrioritizedItem(
                    priority=nextNode.val,
                    item=nextNode.next))

            if not q.empty():
                current.next = ListNode()
                current = current.next

        return head
