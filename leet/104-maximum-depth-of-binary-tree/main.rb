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
def max_depth(root)
  back(root, 0)
end

def back(node, level)
  return level unless node
  [back(node.left, level + 1),
  back(node.right, level + 1)].max
end

def max_depth(root)
  return 0 unless root
  1 + [max_depth(root.left), max_depth(root.right)].max
end
