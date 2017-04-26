class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        result = []
        for i in range(len(nums) - 2):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            r = self.__twoSum(nums[i + 1:], nums[i], 0 - nums[i])
            if len(r) > 0:
                result.extend(r)
        return result

    def __twoSum(self, nums, num, target):
        i, j = 0, len(nums) - 1
        r = []
        while i < j:
            if nums[i] + nums[j] == target:
                r.append([num, nums[i], nums[j]])
                i += 1
                j -= 1
                while i < j and nums[i] == nums[i - 1]:
                    i += 1
                while i < j and nums[j] == nums[j + 1]:
                    j -= 1
            elif nums[i] + nums[j] < target:
                i += 1
                while i < j and nums[i] == nums[i - 1]:
                    i += 1
            else:
                j -= 1
                while i < j and nums[j] == nums[j + 1]:
                    j -= 1
        return r


if __name__ == "__main__":
    s = Solution()
    print(s.threeSum([-1, 0, 1, 2, -1, -4]))
