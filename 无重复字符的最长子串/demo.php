<?php
function lengthOfLongestSubstring(string $s)
{
    $tmp_str = '';
    $max_len = 0;
    // $max_str = '';

    for ($i = 0; $i < strlen($s); $i++) {
        $tmp_len = strpos($s, $s{$i}, $i + 1) === false ? strlen($s) : strpos($s, $s{$i}, $i + 1);
        $tmp_str = substr($s, $i, $tmp_len - $i);
        
        for ($j = 1; $j < strlen($tmp_str); $j++) {
            if (strpos($tmp_str, $tmp_str{$j}, $j + 1) !== false) {
                $tmp_str = substr($tmp_str, 0, strpos($tmp_str, $tmp_str{$j}, $j + 1));
            }
        }

        // $max_str = strlen($tmp_str) > $max_len
        //     ? $tmp_str
        //     : $max_str;

        $max_len = strlen($tmp_str) > $max_len
            ? strlen($tmp_str)
            : $max_len;
    }

    // echo $max_str . PHP_EOL;
    return $max_len;
}

echo lengthOfLongestSubstring('abcbcabcbb') . PHP_EOL;
echo lengthOfLongestSubstring('bbbbb') . PHP_EOL;
echo lengthOfLongestSubstring('pwwkew') . PHP_EOL;
