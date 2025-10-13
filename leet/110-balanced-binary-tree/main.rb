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
# @return {Boolean}
def is_balanced(root)
  back(root, 0) >= 0
end

def back(node, c)
  return c unless node
  l = back(node.left, c + 1)
  r = back(node.right, c + 1)
  return -1 if l == -1 || r == -1 || (l - r).abs > 1

  [l, r].max
end