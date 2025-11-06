require "minitest/autorun"

def get_sneaky_numbers(nums)
  nums.each_with_object(Hash.new(0)) {|n, acc| acc[n] += 1}.filter {|k, v| v > 1}.keys
end

class TestGetSneakyNumbers < Minitest::Test
  def test_get_sneaky_numbers
    assert_equal [0, 1], get_sneaky_numbers([0,1,1,0]).sort
    assert_equal [2, 3], get_sneaky_numbers([0,3,2,1,3,2]).sort
    assert_equal [4, 5], get_sneaky_numbers([7,1,5,4,3,4,6,0,9,5,8,2]).sort
  end
end