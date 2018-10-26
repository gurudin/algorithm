var twoSum = function(nums, target) {
  for (let i = 0; i < nums.length; i++) {
    for (let j = (i+1); j < nums.length; j++) {
      if (nums[i] + nums[j] == target) {
        return [i, j];
      }
    }
  }
}

var twoSumB = function(nums, target) {
  var tmp = Array;
  nums.forEach((value, key) => {
    tmp[value] = key;
  });

  for (let i = 0; i < nums.length; i++) {
    if (typeof tmp[target - nums[i]] != 'undefined' && i != tmp[target - nums[i]]) {
      return [i, tmp[target - nums[i]]];
    }
  }
}

let nums = [1, 2, 3, 4, 5, 6];
let target = 4;
console.log(twoSum(nums, target));
console.log(twoSumB(nums, target));
