require 'minitest/autorun'

def number_of_stable_arrays(zero, one, limit)
  result = []
  nums = [*Array.new(zero, 0), *Array.new(one, 1)]
  used = Array.new(nums.size, false)

  dfs = ->(path) do
    if path.size == nums.size
      result << path.dup
      return
    end

    nums.each_with_index do |n, i|
      next if used[i]
      next if path.last(limit).count(n) >= limit
      path << n
      used[i] = true
      dfs.call(path)
      used[i] = false
      path.pop
    end
  end

  dfs.call([])
  result.uniq.size
end


class TestNumberOfStableArrays < Minitest::Test
  def test_number_of_stable_arrays
    assert_equal 2, number_of_stable_arrays(1, 1, 2)
    assert_equal 1, number_of_stable_arrays(1, 2, 1)
    assert_equal 14, number_of_stable_arrays(3, 3, 2)
  end
end