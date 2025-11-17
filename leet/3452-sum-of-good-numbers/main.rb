require "minitest/autorun"

def sum_of_good_numbers(nums, k)
  nums.map.with_index do |n, i|
    min = i - k < 0 ? i + k : i - k
    next -1 if nums[min].nil? && nums[i+k].nil?
    if nums[min].nil?
      nums[i+k] < n ? n : -1
    elsif nums[i+k].nil?
      nums[min] < n ? n : -1
    else
      (nums[min] < n && nums[i+k] < n) ? n : -1
    end
  end.filter { |n| n != -1}.sum()
end

def sum_of_good_numbers(nums, k)
  nums.each_with_index.map do |n, i|
    left = i - k >= 0 ? nums[i - k] : nil
    right = nums[i+k]
    next unless [left, right].compact.all? { |x| n > x}
    n
  end.compact.sum
end

class TestSumOfGoodNumbers < Minitest::Test
  def test_sum_of_good_numbers
    assert_equal 12, sum_of_good_numbers([1,3,2,1,5,4], 2)
    assert_equal 2, sum_of_good_numbers([2,1], 1)
  end
end