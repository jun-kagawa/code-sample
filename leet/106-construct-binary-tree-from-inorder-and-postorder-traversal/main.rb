# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
        @val = val
        @left = left
        @right = right
    end
end
# @param {Integer[]} inorder
# @param {Integer[]} postorder
# @return {TreeNode}
def build_tree(inorder, postorder)
  m = inorder.each_with_index.to_h
  i = postorder.size - 1

  helper = lambda do |j, k|
    return nil if j > k || i < 0

    val = postorder[i]
    i -= 1
    idx = m[val]
    node = TreeNode.new(val)
    node.right = helper.call(idx + 1, k)
    node.left = helper.call(j, idx - 1)
    node
  end
  helper.call(0, postorder.size - 1)
end