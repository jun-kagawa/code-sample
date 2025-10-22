require "minitest/autorun"

def find_the_prefix_common_array(a, b)
  Array.new(a.size, nil).each_with_index.map do |_, i|
    (a[0..i] & b[0..i]).size
  end
end

class TestFindThePrefixCommonArray < Minitest::Test
  def test_find_the_prefix_common_array
    assert_equal [0,2,3,4], find_the_prefix_common_array([1,3,2,4], [3,1,2,4])
    assert_equal [0, 1, 3], find_the_prefix_common_array([2,3,1], [3,1,2])
    assert_equal [1,2,3], find_the_prefix_common_array([1,2,3], [1,2,3])
  end
end