require 'minitest/autorun'

# @param {Integer[][]} triangle
# @return {Integer}
def minimum_total(triangle)
  result = triangle[0]
  triangle[1..].each do |arr|
    result = arr.map.with_index do |v, i|
      if i == 0
        v + result[i]
      else
        v + [result[i-1], result[i]].filter{|v| !v.nil?}.min
      end
    end
  end
  result.min
end

class TestMinimumTotal < Minitest::Test
  def test_minimum_total
    assert_equal 11, minimum_total([[2],[3,4],[6,5,7],[4,1,8,3]])
    assert_equal 5, minimum_total([[2],[3,4]])
    assert_equal -10, minimum_total([[-10]])
    assert_equal -1, minimum_total([[-1],[2,3],[1,-1,-3]])
    assert_equal -1, minimum_total([[-1],[3,2],[-3,1,-1]])
  end
end

# -1
# 3,2
# -3,1,-1

# -1
# 2,1
# -1,2,1