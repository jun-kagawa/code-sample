require "minitest/autorun"
require "prime"

def diagonal_prime(nums)
  arr = []
  nums.each_with_index do |_, i|
    j = nums.size - i - 1
    arr << nums[i][i] if nums[i][i]
    arr << nums[i][j] if nums[i][j]
  end
  arr.uniq.sort.reverse.each do |n|
    if Prime::prime?(n)
        return n
    else
        next
    end
  end
  0
end

class TestDiagonalPrime < Minitest::Test
  def test_diagonal_prime
    assert_equal 11, diagonal_prime([[1,2,3],[5,6,7],[9,10,11]])
    assert_equal 17, diagonal_prime([[1,2,3],[5,17,7],[9,11,10]])
  end
end
