package main
import . "nc_tools"
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

//https://www.nowcoder.com/share/jump/2615771881734778311231

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param head ListNode类 
 * @return ListNode类
*/
func ReverseList( head *ListNode ) *ListNode {
    var prev *ListNode
    curr:=head
    for curr!=nil{
        nextTemp:=curr.Next
        curr.Next=prev // 把当前节点插到反转链表的头部
        prev=curr // 更新反转链表头节点
        curr=nextTemp // 更新原链表的头结点
    }
    return prev
}
