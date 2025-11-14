require "minitest/autorun"

def max_frequency(nums, k, num_operations)
  cnt = Hash.new(0)
  diff = Hash.new(0)

  nums.each do |n|
    cnt[n] += 1
    diff[n] += 0
    diff[n-k] += 1
    diff[n+k+1] -= 1
  end

  running = 0
  ans = 0
  diff.keys.sort.each do |k|
    running += diff[k]
    free = cnt[k]
    cur = [running, free + num_operations].min
    ans = [ans, cur].max
  end
  ans
end

class TestMaxFrequency < Minitest::Test
  def test_max_frequency
    assert_equal 2, max_frequency([1,4,5], 1, 2)
    assert_equal 2, max_frequency([5,11,20,20], 5, 1)
    assert_equal 2, max_frequency([88, 53], 27, 2)
    assert_equal 1, max_frequency([1, 90], 76, 1)
  end
end

