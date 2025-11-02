require "minitest/autorun"

def total_money(n)
  pre = 0
  sum = 0
  n.times do |i|
    pre += 1
    sum += pre
    if ((i+1) % 7 == 0)
      pre -= 6
    end
  end
  sum
end

class TestTotalMoney < Minitest::Test
  def test_total_money
    assert_equal 10, total_money(4)
    assert_equal 37, total_money(10)
    assert_equal 96, total_money(20)
  end
end