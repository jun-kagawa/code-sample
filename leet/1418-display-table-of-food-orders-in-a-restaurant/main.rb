require "minitest/autorun"

def display_table(orders)
  table = Hash.new {|h, k| h[k] = Hash.new(0)}
  cols = []

  orders.each do |_, table_num, food|
    table[table_num][food] += 1
    cols << food
  end
  cols = cols.uniq.sort
  headers = ["Table", *cols]
  rows = table.sort_by {|k, _| k.to_i}.map do |table_num, foods|
    [table_num, *cols.map {|col| foods[col].to_s}]
  end
  [headers, *rows]
end

class TestDisplayTable < Minitest::Test
  def test_display_table
    assert_equal [["Table","Beef Burrito","Ceviche","Fried Chicken","Water"],["3","0","2","1","0"],["5","0","1","0","1"],["10","1","0","0","0"]], display_table([["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]])
    assert_equal [["Table","Canadian Waffles","Fried Chicken"],["1","2","0"],["12","0","3"]], display_table([["James","12","Fried Chicken"],["Ratesh","12","Fried Chicken"],["Amadeus","12","Fried Chicken"],["Adam","1","Canadian Waffles"],["Brianna","1","Canadian Waffles"]])
    assert_equal [["Table","Bean Burrito","Beef Burrito","Soda"],["2","1","1","1"]], display_table([["Laura","2","Bean Burrito"],["Jhon","2","Beef Burrito"],["Melissa","2","Soda"]])
  end
end