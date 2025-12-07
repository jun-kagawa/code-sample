require 'minitest/autorun'

def max_profit(prices)
  i = 0
  j = 1
  m = 0
  while j < prices.size
    buy = prices[i]
    current = prices[j]
    if buy < current
      profit = current - buy
      m = [m, profit].max
    else
      i = j
    end
    j += 1
  end
  m
end


class TestMaxProfit < Minitest::Test
  def test_max_profit
    assert_equal 5, max_profit([7,1,5,3,6,4])
    assert_equal 0, max_profit([7,6,4,3,1])
  end
end