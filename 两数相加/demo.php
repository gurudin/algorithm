<?php
class ListNode
{
    public $val = null;
    public $next = null;

    public function __construct(int $data = null)
    {
        $this->val = $data;
    }
}

class LinkedList
{
    public $firstNode = null;

    public function insert(int $data = null)
    {
        $newNode = new ListNode($data);
        
        if ($this->firstNode === null) {
            $this->firstNode = $newNode;
        } else {
            $currentNode = $this->firstNode;
            while ($currentNode->next !== null) {
                $currentNode = $currentNode->next;
            }

            $currentNode->next = $newNode;
        }
    }

    public function show()
    {
        $currentNode = $this->firstNode;
        while ($currentNode !== null) {
            echo $currentNode->val . "\n";
            $currentNode = $currentNode->next;
        }
    }
}

function addTwoNumbers(ListNode $l1, ListNode $l2)
{
    $currentL1 = $l1;
    $currentL2 = $l2;

    $tmp1 = $tmp2 = '';
    while ($currentL1 !== null && $currentL2 !== null) {
        $tmp1 .= $currentL1->val;
        $tmp2 .= $currentL2->val;

        $currentL1 = $currentL1->next;
        $currentL2 = $currentL2->next;
    }

    $sum = $tmp1 + $tmp2;
    $linked_res = new LinkedList();
    for ($i=strlen($sum)-1; $i>=0; $i--) {
        $linked_res->insert((int)substr($sum, $i, 1));
    }

    return $linked_res->firstNode;
}

$linked_list1 = new LinkedList();
$linked_list1->insert(2);
$linked_list1->insert(4);
$linked_list1->insert(3);

$linked_list2 = new LinkedList();
$linked_list2->insert(5);
$linked_list2->insert(6);
$linked_list2->insert(4);

print_r(
    addTwoNumbers(
        $linked_list1->firstNode,
        $linked_list2->firstNode
    )
);
