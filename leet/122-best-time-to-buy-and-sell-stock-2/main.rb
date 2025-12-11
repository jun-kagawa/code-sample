require 'minitest/autorun'

def max_profit(prices)
  buy = 0
  sum = 0
  size = prices.size - 1
  while buy < size
    if prices[buy] < prices[buy+1]
      sum += prices[buy+1] - prices[buy]
    end
    buy += 1
  end
  sum
end

def max_profit(prices)
  prices.map.with_index do |price, i|
    next 0 if prices[i+1].nil?
    price < prices[i+1] ? prices[i+1] - price : 0
  end.sum
end

class TestMaxProfit < Minitest::Test
  def test_max_profit
    assert_equal 7, max_profit([7,1,5,3,6,4])
    assert_equal 4, max_profit([1,2,3,4,5])
    assert_equal 0, max_profit([7,6,4,3,1])
  end
end

