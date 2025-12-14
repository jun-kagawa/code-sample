require 'minitest/autorun'

# Definition for a binary tree node.
class TreeNode
  attr_accessor :val, :left, :right
  def initialize(val = 0, left = nil, right = nil)
    @val = val
    @left = left
    @right = right
  end
end

def sum_numbers(root)
  return if root.nil?
  if root.left.nil? && root.right.nil?
    return root.val
  end
  l = back(root.left, root.val.to_s)
  r = back(root.right, root.val.to_s)
  l = l.nil? ? 0 : l
  r = r.nil? ? 0 : r
  l + r
end

def back(node, str)
  if node.nil?
    return
  end
  str += node.val.to_s
  if node.left.nil? && node.right.nil?
    return str.to_i
  end
  l = back(node.left, str)
  r = back(node.right, str)
  str = str[...-1]
  l = l.nil? ? 0 : l
  r = r.nil? ? 0 : r
  l + r
end

# @param {TreeNode} root
# @return {Integer}
def sum_numbers(root)
  dfs(root, 0)
end

def dfs(node, current_sum)

  return 0 if node.nil?



  current_sum = current_sum * 10 + node.val



  return current_sum if node.left.nil? && node.right.nil?



  dfs(node.left, current_sum) + dfs(node.right, current_sum)

end



class TestSumNumbers < Minitest::Test

  def test_sum_numbers

    # root = [1,2,3]

    root = TreeNode.new(1, TreeNode.new(2), TreeNode.new(3))

    assert_equal 25, sum_numbers(root)



    # root = [4,9,0,5,1]

    root2 = TreeNode.new(4,

      TreeNode.new(9, TreeNode.new(5), TreeNode.new(1)),

      TreeNode.new(0)

    )

    assert_equal 1026, sum_numbers(root2)

  end

end
