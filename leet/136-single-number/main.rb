require 'minitest/autorun'

def single_number(nums)
  return nums[0] if nums.size == 1

  h = Hash.new {|hash, k| hash[k] = 0 }
  nums.map do |n|
    h[n] += 1
  end
  h.map do |k, v|
    return k if v == 1
  end
end

class TestSingleNumber < Minitest::Test
  def test_single_number
    assert_equal 1, single_number([2, 2, 1])
    assert_equal 4, single_number([4, 1, 2, 1, 2])
    assert_equal 1, single_number([1])
  end
end
