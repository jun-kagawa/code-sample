
# Definition for a Node.
class Node
    attr_accessor :val, :left, :right, :next
    def initialize(val)
        @val = val
        @left, @right, @next = nil, nil, nil
    end
end

# @param {Node} root
# @return {Node}
def connect(root)
  return root if root.nil?
  q = [root]
  while q.size != 0
    nodes = []
    before = nil
    q.each do |node|
      node.next = before
      before = node
      nodes << node.right unless node.right.nil?
      nodes << node.left unless node.left.nil?
    end
    q = nodes
  end
  return root
end