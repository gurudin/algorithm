/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function (s) {
  var len = s.length;
  for (let i = 0; i < s.length; i++) {
    let longs = '';
    let tmpChar = s[i];

    if (s.substring(i+1, len).indexOf(s[i]) > -1) {
      
    }
  }
  
};

longestPalindrome('babad');
