/**
 * 链表模型
 */
function ListNode(val) {
  this.val = val;
  this.next = null;
}

/**
 * 链表操作类
 */
function List() {
  this.List = null;
  this.insert = function(val) {
    var newNode = new ListNode(val);

    if (this.List == null) {
      this.List = newNode;
    } else {
      currentNode = this.List;
      while (currentNode.next !== null) {
        currentNode = currentNode.next;
      }
      
      currentNode.next = newNode;
    }
  };
}

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
  var sum = [];
  var cnt = 1;

  var currentL1 = l1;
  sum.push(currentL1.val);
  while (currentL1.next !== null) {
    currentL1 = currentL1.next;
    sum.push(currentL1.val);
  }

  var currentL2 = l2;
  sum[0] += currentL2.val;
  while (currentL2.next !== null) {
    currentL2 = currentL2.next;
    
    if (typeof sum[cnt] == 'undefined') {
      sum.push(currentL2.val);
    } else {
      sum[cnt] += currentL2.val;
    }

    cnt++;
  }

  for (let i = 0; i < sum.length; i++) {
    if (sum[i] >= 10) {
      if (typeof sum[i + 1] == 'undefined') {
        sum.push(1);
      } else {
        sum[i + 1] ++;
      }
      sum[i] -= 10;
    }
  }

  var l = new List();
  sum.forEach(val =>{
    l.insert(val);
  });

  return l.List;
};

function main() {
  // var data = [
  //   [2, 4, 3],
  //   [5, 6, 4]
  // ];

  // var data = [
  //   [1, 8],
  //   [0]
  // ];

  // var data = [
  //   [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
  //   [5, 6, 4]
  // ];

  var data = [
    [0],
    [7, 3]
  ];

  var l1 = new List();
  var l2 = new List();
  data[0].forEach(val =>{
    l1.insert(val);
  });
  data[1].forEach(val =>{
    l2.insert(val);
  });

  console.log(addTwoNumbers(l1.List, l2.List));
}

main();
