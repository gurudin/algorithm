/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
    var nums = nums1.concat(nums2);
    var length = nums.length;
    nums = nums.sort(function(a, b) {
        return a - b;
    });

    return Number.isInteger(length / 2)
        ? (nums[length / 2] + nums[length / 2 - 1]) / 2
        : nums[Math.floor(length / 2)];
};

console.log(findMedianSortedArrays([1, 3], [2]));
console.log(findMedianSortedArrays([1, 2], [3, 4]));
console.log(findMedianSortedArrays([1, 3, 5, 100], [2]));
