class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        如果我们不寻找那个元素先，而是直接相等的时候也向一个方向继续夹逼，如果向右夹逼，最后就会停在右边界，而向左夹逼则会停在左边界，如此用停下来的两个边界就可以知道结果了，只需要两次二分查找
        """
        left, right = -1, -1

        ll, lr = 0, len(nums) - 1
        while ll <= lr:
            m = (ll + lr) // 2
            if nums[m] < target:
                ll = m + 1
            else:
                lr = m - 1

        rl, rr = 0, len(nums) - 1
        while rl <= rr:
            m = (rl + rr) // 2
            if nums[m] <= target:
                rl = m + 1
            else:
                rr = m - 1

        if ll <= rr:
            left = ll
            right = rr

        return [left, right]
