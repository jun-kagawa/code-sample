require 'minitest/autorun'

def longest_consecutive(nums)
  hash = Set.new nums
  longest = 0
  hash.to_a.each do |num|
    if hash.include?(num-1)
      next
    end
    streak = 1

    while hash.include?(num+1)
      streak += 1
      num += 1
    end
    longest = [longest, streak].max
  end
  longest
end

class TestLongestConsecutive < Minitest::Test
  def test_longest_consecutive
     assert_equal 4, longest_consecutive([100,4,200,1,3,2])
     assert_equal 9, longest_consecutive([0,3,7,2,5,8,4,6,0,1])
     assert_equal 3, longest_consecutive([1,0,1,2])
     assert_equal 2, longest_consecutive([0,-1])
     assert_equal 3, longest_consecutive([0,-1,1])
     assert_equal 5, longest_consecutive([4,0,-4,-2,2,5,2,0,-8,-8,-8,-8,-1,7,4,5,5,-4,6,6,-3])
    assert_equal 4, longest_consecutive([9,1,-3,2,4,8,3,-1,6,-2,-4,7])
    assert_equal 6, longest_consecutive([1,-8,7,-2,-4,-4,6,3,-4,0,-7,-1,5,1,-9,-3])
  end
end
