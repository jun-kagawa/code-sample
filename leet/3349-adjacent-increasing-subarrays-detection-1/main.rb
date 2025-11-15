require "minitest/autorun"

def has_increasing_subarrays(nums, k)
  nums.each_with_index do |n, i|
    a = nums[i...i+k]
    next unless is_sorted(a) || a.size != k
    b = nums[i+k...i+k+k]
    next if b.nil? || b.size != k
    return true if is_sorted(b)
  end
  false
end

def is_sorted(nums)
  return false if nums.nil?
  nums.each_with_index do |n, i|
    next if i == 0
    return false if nums[i-1] >= n
  end
end

class TestHasIncreasingSubArrays < Minitest::Test
  def test_has_increasing_subarrays
    assert_equal true, has_increasing_subarrays([2,5,7,8,9,2,3,4,3,1], 3)
    assert_equal false, has_increasing_subarrays([1,2,3,4,4,4,4,5,6,7], 5)
    assert_equal false, has_increasing_subarrays([-3,-19,-8,-16], 2)
  end
end
