require 'minitest/autorun'

def smallest_number(n)
  (0...(n.to_s(2).length)).map {|n| '1'}.sum('').to_i(2)
end

class TestSmallestNumber < Minitest::Test
  def test_smallest_number
    assert_equal 7, smallest_number(5)
    assert_equal 15, smallest_number(10)
    assert_equal 3, smallest_number(3)
  end
end