require "minitest/autorun"
require_relative "main"

class TestMinDeletions < Minitest::Test
  def test_min_deletions
    assert_equal 0, min_deletions("aab")
    assert_equal 2, min_deletions("aaabbbcc")
    assert_equal 2, min_deletions("ceabaacb")
  end
end
