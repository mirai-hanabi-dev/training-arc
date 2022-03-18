class Solution:
    def firstMissingPositive(self, nums):
        n = len(nums)
        nums = [0 if v < 0 or v > n else v for v in nums]

        for idx in range(n):
            if nums[idx] == 0:
                continue
            if nums[idx] == idx + 1:
                continue
            while True:
                pos = nums[idx] - 1
                nums[idx], nums[pos] = nums[pos], nums[idx]
                if nums[idx] == 0:
                    break
                if nums[idx] == idx + 1:
                    break
                if nums[idx] == nums[nums[idx] - 1]:
                    break

        for idx in range(n):
            if nums[idx] != idx + 1:
                return idx + 1

        return n + 1
