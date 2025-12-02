require 'minitest/autorun'

def get_row(row_index)
  return [1] if row_index == 0
  result = [1,1]
  return result if row_index == 1
  (row_index - 1).times do |i|
    arr = [1]
    (result.size - 1).times do |j|
      arr << result[j] + result[j+1]
    end
    arr << 1
    result = arr
  end
  result
end

class TestGetRow < Minitest::Test
  def test_get_row
    assert_equal [1,3,3,1], get_row(3)
    assert_equal [1], get_row(0)
    assert_equal [1,1], get_row(1)
    assert_equal [1,4,6,4,1], get_row(4)
  end
end