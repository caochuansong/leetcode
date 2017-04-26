class Solution(object):
    def threeSumClosest(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        nums.sort()
        result = nums[0] + nums[1] + nums[2]
        minDiff = abs(result - target)
        for i in range(len(nums) - 2):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            begin = i + 1
            end = len(nums) - 1
            while begin < end:
                s = nums[i] + nums[begin] + nums[end]
                if s == target:
                    return s
                else:
                    diff = abs(target - s)
                    if diff < minDiff:
                        minDiff = diff
                        result = s
                    if s < target:
                        begin += 1
                        while begin < end and nums[begin] == nums[begin - 1]:
                            begin += 1
                    else:
                        end -= 1
                        while begin < end and nums[end] == nums[end + 1]:
                            end -= 1
        return result
