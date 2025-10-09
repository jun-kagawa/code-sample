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
# @return {Integer[][]}
def level_order_bottom(root)
  r = []
  return r if root.nil?

  queue = [root]
  until queue.empty?
    vals = queue.map(&:val)
    r.insert(0, vals)

    queue = queue.flat_map { |node| [node.left, node.right].compact }
  end
  r
end