class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        假设数组是A，每次左边缘为l，右边缘为r，还有中间位置是m。在每次迭代中，分三种情况：
        （1）如果target==A[m]，那么m就是我们要的结果，直接返回；
        （2）如果A[m]<A[r]，那么说明从m到r一定是有序的（没有受到rotate的影响），那么我们只需要判断target是不是在m到r之间，如果是则把左边缘移到m+1，否则就target在另一半，即把右边缘移到m-1。
        （3）如果A[m]>=A[r]，那么说明从l到m一定是有序的，同样只需要判断target是否在这个范围内，相应的移动边缘即可。
        """
        l, h = 0, len(nums) - 1
        while l <= h:
            m = l + ((h - l) >> 1)
            if nums[m] == target:
                return m

            if nums[m] < nums[h]:
                if nums[m] < target <= nums[h]:
                    l = m + 1
                else:
                    h = m - 1
            else:
                if nums[l] <= target < nums[m]:
                    h = m - 1
                else:
                    l = m + 1
        return -1


if __name__ == "__main__":
    s = Solution()
    print(s.search([4, 5, 6, 7, 0, 1, 2], 5))