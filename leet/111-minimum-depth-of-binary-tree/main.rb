# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
        @val = val
        @left = left
        @right = right
    end
end
# @param {TreeNode} root
# @return {Integer}
def min_depth(root)
  back(root, 1)
end

def back(node, c)
  return -1 unless node
  return c if node.left.nil? && node.right.nil?
  l = back(node.left, c + 1)
  r = back(node.right, c + 1)

  return r if l == -1
  return l if r == -1
  [l, r].min
end