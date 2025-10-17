require "minitest/autorun"

def find_smallest_integer(nums, value)
  freq = Array.new(value, 0)
  nums.each {|n| freq[n % value] += 1}

  mex = 0
  loop do
    r = mex % value
    return mex if freq[r] == 0 
    freq[r] -= 1
    mex += 1
  end
end

class TestFindSmallestInteger < Minitest::Test
  def test_find_smallest_integer
    assert_equal 4, find_smallest_integer([1,-10,7,13,6,8], 5)
    assert_equal 2, find_smallest_integer([1,-10,7,13,6,8], 7)
  end
end
