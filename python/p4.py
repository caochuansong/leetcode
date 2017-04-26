class Solution:
    # @return a float
    # 转化为求第k最小数
    def findMedianSortedArrays(self, a, b):
        m = len(a)
        n = len(b)
        k = (m + n) // 2
        return self.__find_kth_element(a, m, b, n, k + 1) if (m + n) & 0x01 \
            else (self.__find_kth_element(a, m, b, n, k) + self.__find_kth_element(a, m, b, n, k + 1)) / 2.0

    def __find_kth_element(self, a, m, b, n, k):
        if m > n:
            return self.__find_kth_element(b, n, a, m, k)
        if m == 0:
            return b[k - 1]
        if k == 1:
            return min(a[0], b[0])

        pa = min(k // 2, m)
        pb = k - pa
        if a[pa - 1] < b[pb - 1]:
            return self.__find_kth_element(a[pa:m], m - pa, b, n, k - pa)
        elif a[pa - 1] > b[pb - 1]:
            return self.__find_kth_element(a, m, b[pb:n], n - pb, k - pb)
        else:
            return a[pa - 1]

