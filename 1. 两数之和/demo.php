<?php
function towSum(array $nums, int $target)
{
    for ($i = 0; $i < count($nums); $i++) {
        for ($j = ($i+1); $j < count($nums); $j++) {
            if ($nums[$i] + $nums[$j] == $target) {
                return [$i, $j];
            }
        }
    }
}

function towSumB(array $nums, int $target)
{
    $count = count($nums);
    $tmp = [];
    
    for ($i = 0; $i < $count; $i++) {
        $tmp[$nums[$i]] = $i;
    }

    for ($j = 0; $j < $count; $j++) {
        if (isset($tmp[$target - $nums[$j]]) && $j != $tmp[$target - $nums[$j]]) {
            return [$j, $tmp[$target - $nums[$j]]];
        }
    }
}

function towSumC(array $nums, int $target)
{
    $hash = [];

    for ($i=0; $i<count($nums); $i++) {
        $tmp_num = $nums[$i];

        if (isset($hash[$tmp_num])) {
            return [$hash[$tmp_num], $i];
        } else {
            $hash[$target - $tmp_num] = $i;
        }
    }
}

$nums = [1, 2, 3, 4, 5, 6];
$target = 11;

print_r(towSum($nums, $target));
print_r(towSumB($nums, $target));
print_r(towSumC($nums, $target));
