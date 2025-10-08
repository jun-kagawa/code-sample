# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
        @val = val
        @left = left
        @right = right
    end
end
# @param {Integer[]} preorder
# @param {Integer[]} inorder
# @return {TreeNode}
def build_tree(preorder, inorder)
  m = inorder.each_with_index.to_h
  i = 0

  helper = lambda do |j, k|
    return nil if j > k
    val = preorder[i]
    i += 1
    node = TreeNode.new(val)
    idx = m[val]
    node.left = helper.call(j, idx - 1)
    node.right = helper.call(idx + 1, k)
    node
  end
  helper.call(0, inorder.size - 1)
end