require "minitest/autorun"

def final_value_after_operations(operations)
  operations.map {|op| op.include?("++") ? 1 : -1}.sum()
end

class TestFinalValueAfterOperations < Minitest::Test
  def test_final_value_after_operations
    assert_equal 1, final_value_after_operations(["--X","X++","X++"])
    assert_equal 3, final_value_after_operations(["++X","X++","X++"])
    assert_equal 0, final_value_after_operations(["X++","++X","--X","X--"])
  end
end
