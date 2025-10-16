require "minitest/autorun"
require_relative "main"

class TestMinimizeTheDifference < Minitest::Test
  def test_minimize_the_difference
    # assert_equal 0, minimize_the_difference([[1,2,3],[4,5,6],[7,8,9]], 13)
    # assert_equal 94, minimize_the_difference([[1],[2],[3]], 100)
    assert_equal 1, minimize_the_difference([[1,2,9,8,7]], 6)
  end
end
