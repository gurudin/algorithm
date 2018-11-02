/**
 * @param {string}
 * @return {number}
 */
var lengthOfLongestSubstringA = function (s) {
  var tmpStr;
  var maxLen = 0;
  
  for (let i = 0; i < s.length; i++) {
    var tmpLen = s.indexOf(s[i], i + 1) > -1 ? s.indexOf(s[i], i + 1) : s.length;
    var tmpStr = s.substring(i, tmpLen);

    for (let j = 1; j < tmpStr.length; j++) {
      if (tmpStr.indexOf(tmpStr[j], j + 1) > -1) {
        tmpStr = tmpStr.substring(0, tmpStr.indexOf(tmpStr[j], j + 1));
      }
    }

    maxLen = tmpStr.length > maxLen
      ? tmpStr.length
      : maxLen;
  }

  return maxLen;
};

/**
 * @param {string}
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
  var maxLen = tmpLen = start = 0;
  var map = [];

  for (let i = 0; i < s.length; i++) {
    if (map[s[i]] !== undefined && start <= map[s[i]]) {
      start = map[s[i]];
      tmpLen = i - map[s[i]];
    } else {
      tmpLen++;
    }

    map[s[i]] = i;

    maxLen = maxLen < tmpLen ? tmpLen : maxLen;
  }

  return maxLen;
};

console.log(lengthOfLongestSubstring("abba"));
console.log(lengthOfLongestSubstring("aab"));
console.log(lengthOfLongestSubstring('abcabcbb'));
console.log(lengthOfLongestSubstring('bbbbb'));
console.log(lengthOfLongestSubstring('pwwkew'));

