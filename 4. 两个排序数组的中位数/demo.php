<?php
/**
 * @param array $nums1
 * @param array $nums2
 *
 * @return float
 */
function findMedianSortedArrays(array $nums1, array $nums2)
{
    $nums   = array_merge($nums1, $nums2);
    $length = count($nums);
    sort($nums);
    
    return is_int($length / 2)
        ? ($nums[$length / 2] + $nums[$length / 2 - 1]) / 2
        : $nums[floor($length / 2)];
}

echo findMedianSortedArrays([1, 3], [2]) . PHP_EOL;
echo findMedianSortedArrays([1, 2], [3, 4]) . PHP_EOL;
echo findMedianSortedArrays([1, 3, 5, 100], [2]) . PHP_EOL;
