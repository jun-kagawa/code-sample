class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
        @val = val
        @left = left
        @right = right
    end
end

def has_path_sum(root, target_sum)
  if root.nil?
    return false
  end
  recur(root, 0, target_sum)   
end

def recur(node, current, target_sum)
  current += node.val
  if node.left.nil? && node.right.nil?
    p current
    return current == target_sum
  end

  if !node.left.nil?
    r1 = recur(node.left, current, target_sum)
    if r1
      return true
    end
  end

  if !node.right.nil?
    recur(node.right, current, target_sum)
  else
    r1
  end
end
