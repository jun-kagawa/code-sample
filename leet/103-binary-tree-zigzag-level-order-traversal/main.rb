
class TreeNode
  attr_accessor :val, :left, :right
  def initialize(val = 0, left = ni, right = nil)
    @val = val
    @left = left
    @right= right
  end
end

def zigzag_level_order(root)
  m = Hash.new { |h, k| h[k] = []}
  traverse(root, 0, m)
  m.map do |k, v|
    k.even? ? v : v.reverse
  end
end

def traverse(node, level, m)
  return unless node
  m[level] << node.val

  traverse(node.left, level + 1, m)
  traverse(node.right, level + 1, m)
end

def zigzag_level_order(root)
  return [] unless root

  result = []
  queue = [root]
  level = 0

  until queue.empty?
    vals = queue.map(&:val)
    vals.reverse! if level.odd?
    result << vals

    queue = queue.flat_map { |node| [node.left, node.right].compact }
    level += 1
  end

  result
end