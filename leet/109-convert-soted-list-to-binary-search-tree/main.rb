# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end
# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val = 0, left = nil, right = nil)
        @val = val
        @left = left
        @right = right
    end
end
# @param {ListNode} head
# @return {TreeNode}
def sorted_list_to_bst(head)
  return unless head
  return TreeNode.new(head.val) unless head.next

  prev, slow, fast = head, head.next, head.next.next

  while fast&.next
    prev, slow, fast = slow, slow.next, fast.next.next
  end

  prev.next = nil
  TreeNode.new(slow.val).tap do |root|
    root.left = sorted_list_to_bst(head)
    root.right = sorted_list_to_bst(slow.next)
  end
end