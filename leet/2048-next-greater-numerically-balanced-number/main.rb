require "minitest/autorun"

def next_beautiful_number(n)
  return 1224444 if n >= 700000
  loop do
    n += 1
    return n if is_balanced(n)
  end
end

def is_balanced(n)
  h = Hash.new(0)
  n.to_s.chars.each do |s|
    h[s.to_i] += 1
  end
  h.all? do |k, v|
    k == v
  end
end

class TestNextBeautifulNumber < Minitest::Test
  def test_next_beautiful_number
    assert_equal 22, next_beautiful_number(1)
    assert_equal 1333, next_beautiful_number(1000)
    assert_equal 3133, next_beautiful_number(3000)
    assert_equal 1224444, next_beautiful_number(780001)
  end
end