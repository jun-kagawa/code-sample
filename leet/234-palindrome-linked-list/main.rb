require 'minitest/autorun'

# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end
# @param {ListNode} head
# @return {Boolean}
def is_palindrome(head)
  arr = []
  until head.nil?
    arr << head.val
    head = head.next
  end
  arr == arr.reverse
end

class TestIsPalindrome < Minitest::Test
  def test_is_palindrome
    assert_equal true, is_palindrome
  end
end