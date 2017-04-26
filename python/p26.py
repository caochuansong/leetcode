class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums:
            return 0
        currentIndex = 0
        for i in range(1, len(nums)):
            if nums[i] != nums[currentIndex]:
                nums[currentIndex + 1] = nums[i]
                currentIndex += 1
        return currentIndex + 1
