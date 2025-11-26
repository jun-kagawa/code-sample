class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
        @val = val
        @left = left
        @right = right
    end
end

def path_sum(root, target_sum)
  return [] if root.nil?
  r = []
  recur(root, target_sum, [], r)
  r
end

def recur(node, target_sum, current, results)
  current << node.val
  if node.left.nil? && node.right.nil?
    if target_sum == current.sum
      results << current.dup
    end
    current.pop
    return
  end

  if !node.left.nil?
    recur(node.left, target_sum, current, results)
  end

  if !node.right.nil?
    recur(node.right, target_sum, current, results)
  end
  current.pop
end
