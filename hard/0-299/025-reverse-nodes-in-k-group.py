# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverse(self, head, tail):
        dummyHead = head

        nextNode = None
        prevNode = None
        while True:
            nextNode = head.next
            head.next = prevNode
            if head == tail:
                break

            prevNode = head
            head = nextNode

        return tail, dummyHead

    def reverseKGroup(self, head, k):
        dummy = prev = ListNode()
        dummy.next = left = right = head
        while True:
            count = 1
            if not left:
                return dummy.next

            while right.next and count < k:
                count += 1
                right = right.next

            if count == k:
                nextNode = right.next
                newLeft, newRight = self.reverse(left, right)

                prev.next = newLeft
                newRight.next = nextNode

                left = right = nextNode
                prev = newRight
            else:
                return dummy.next


node1 = ListNode(1)
node2 = ListNode(2)
node3 = ListNode(3)
node4 = ListNode(4)
node5 = ListNode(5)
node1.next = node2
node2.next = node3
node3.next = node4
node4.next = node5

newNode = Solution().reverseKGroup(node1, 1)

while newNode:
    print(newNode.val)
    newNode = newNode.next
