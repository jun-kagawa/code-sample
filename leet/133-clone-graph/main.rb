# Definition for a Node.
class Node
    attr_accessor :val, :neighbors
    def initialize(val = 0, neighbors = nil)
		  @val = val
		  neighbors = [] if neighbors.nil?
        @neighbors = neighbors
    end
end

# @param {Node} node
# @return {Node}
def cloneGraph(node)
  return nil if node.nil?
  root = Node.new node.val, node.neighbors
  return root if root.neighbors.size == 0
  h = {}
  h[root.val] = root
  back(root, h)
  root
end

def back(parent, h)
  new_neighbors = []
  parent.neighbors.each do |node|
    unless h.include? node.val
        new_node = Node.new node.val, node.neighbors
        h[new_node.val] = new_node
        back(new_node, h)
    end
    new_neighbors << h[node.val]
  end
  parent.neighbors = new_neighbors
end

