require 'minitest/autorun'

def generate(num_rows)
  result = [[1]]
  return result if num_rows == 1
  result << [1,1]
  return result if num_rows == 2
  (num_rows - 2).times do |i|
    arr = [1]
    (result.last.size - 1).times do |j|
      arr << result.last[j] + result.last[j+1]
    end
    arr << 1
    result << arr
  end
  result
end

class TestGenerate < Minitest::Test
  def test_generate
    assert_equal [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]], generate(5)
    assert_equal [[1]], generate(1)
    assert_equal [[1], [1,1]], generate(2)
    assert_equal [[1], [1,1], [1,2,1]], generate(3)
  end
end