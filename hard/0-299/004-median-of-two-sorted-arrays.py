# Break two arrays into 4 smaller arrays:
# leftNums1 & rightNums1 & leftNums2 & rightNums2
# where leftNums1 and leftNums2 contain the first half of the merged array.

class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        m = len(nums1)
        n = len(nums2)
        if m > n:
            return self.findMedianSortedArrays(nums2, nums1)

        left = 0
        right = m
        while left <= right:
            mid1 = left + (right - left) // 2
            mid2 = (m + n + 1) // 2 - mid1

            maxLeftNums1 = -float('Inf') if mid1 == 0 else nums1[mid1 - 1]
            minRightNums1 = float('Inf') if mid1 == m else nums1[mid1]

            maxLeftNums2 = -float('Inf') if mid2 == 0 else nums2[mid2 - 1]
            minRightNums2 = float('Inf') if mid2 == n else nums2[mid2]

            if maxLeftNums1 <= minRightNums2 and maxLeftNums2 <= minRightNums1:
                length = m + n
                if length % 2 == 1:
                    return float(max(maxLeftNums1, maxLeftNums2))
                else:
                    return (max(maxLeftNums1, maxLeftNums2)
                            + min(minRightNums1, minRightNums2)) / 2
            elif maxLeftNums1 > minRightNums2:
                right = mid1 - 1
            else:
                left = mid1 + 1


nums1 = [1, 3]
nums2 = [2]
print(Solution().findMedianSortedArrays(nums1, nums2))

nums1 = [1, 2]
nums2 = [3, 4]
print(Solution().findMedianSortedArrays(nums1, nums2))

nums1 = [2]
nums2 = []
print(Solution().findMedianSortedArrays(nums1, nums2))
