require 'minitest/autorun'

def is_palindrome(s)
  m = Set.new ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
    'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
  ]
  str = s.upcase.split("").map do |ss|
    m.include?(ss) ? ss : nil
  end.filter {|ss| !ss.nil?}.join
  str == str.reverse
end

class TestIsPalindrome < Minitest::Test
  def test_is_palindrome
    assert_equal true, is_palindrome('A man, a plan, a canal: Panama')
    assert_equal false, is_palindrome('race a car')
    assert_equal true, is_palindrome(' ')
    assert_equal true, is_palindrome('')
    assert_equal false, is_palindrome('0P')
  end
end

