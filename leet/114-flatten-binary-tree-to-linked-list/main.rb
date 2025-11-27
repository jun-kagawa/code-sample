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
# @return {Void} Do not return anything, modify root in-place instead.
def flatten(root)
  return nil if root.nil?
  left = recur(root.left) if root.left
  right = recur(root.right) if root.right
  root.right = left
  root.left = nil
  if left.nil?
    root.right = right
  else
    until left.right.nil?
    left = left.right
    end
  left.right = right
  end
end

def recur(node)
  dummy = TreeNode.new
  back(node, dummy)

  return dummy.right
end

def back(node, result)
  result.right = TreeNode.new(node.val)
  result = result.right
  back(node.left, result) unless node.left.nil?
  until result.right.nil?
    result = result.right
  end
  back(node.right, result) unless node.right.nil?
end