require "minitest/autorun"

def has_same_digits(s)
  while s.size > 2
    s = (0..s.size-2).map do |n|
      ((s[n].to_i + s[n+1].to_i) % 10).to_s
    end.join
  end
  s[0] == s[1]
end

class TestHasSameDigits < Minitest::Test
  def test_has_same_digits_1
    assert_equal true, has_same_digits("3902")
  end
  def test_has_same_digits_2
    assert_equal false, has_same_digits("34789")
  end
end