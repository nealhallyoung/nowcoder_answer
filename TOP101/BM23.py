# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#
# 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
#
# 
# @param root TreeNode类 
# @return int整型一维数组
# BM23 二叉树的前序遍历
# https://www.nowcoder.com/share/jump/2615771881735305368186
class Solution:
    def preorder(self, list: List[int], node: TreeNode):
        # 遇到空节点则返回
        if node == None:
            return
        # 先遍历根节点
        list.append(node.val)
        # 再去左子树
        self.preorder(list, node.left)
        # 最后去右子树
        self.preorder(list, node.right)

    def preorderTraversal(self, root: TreeNode) -> List[int]:
        # 添加遍历结果的数组
        list = []
        # 递归前序遍历
        self.preorder(list, root)
        return list
