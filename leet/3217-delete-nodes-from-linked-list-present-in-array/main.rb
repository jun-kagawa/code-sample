# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end
# @param {Integer[]} nums
# @param {ListNode} head
# @return {ListNode}
def modified_list(nums, head)
  s = nums.to_set
  dummy = ListNode.new(0, head)
  pre = dummy

  while pre.next
    if s.include?(pre.next.val)
      pre.next = pre.next.next
    else
      pre = pre.next
    end
  end
  dummy.next
end
