class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        // 新建一个链表 用来存储相加后的和
        ListNode* head = new ListNode(0);
        ListNode* cur = head;

        // 进位，即使l1,l2指向null，进位存在也得增加节点
        int plus = 0;
        // 什么时候结束操作？l1，l2，进位都为空的时候
        while(l1||l2||plus){
            int sum = (l1?l1->val:0) + (l2?l2->val:0) + plus;
            plus = sum/10;
            cur->next = new ListNode(sum%10);
            cur = cur->next;

            l1 = l1?l1->next:l1;
            l2 = l2?l2->next:l2;
        }
        return head->next;
    }
};