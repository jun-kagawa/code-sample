require "minitest/autorun"

def unique_paths(grid)
  m = grid.size
  n = grid[0].size
  prev = [[1,1]]
  (1...n).each do |i|
    if grid[0][i-1] == 1
      prev << [0,0]
      next
    end
    prev << prev[i-1]
  end

  p "prev", prev
  cur = []
  (1...m).each do |i|
    if grid[i-1][0] == 1
      cur << [0,0]
    else
      cur << prev[0]
    end

    (1...n).each do |j|
      if grid[i][j-1] == 1 && j - 1 != 0
        cur << prev[j].first
      elsif grid[i-1][j] == 1 && prev[j] != 0
        cur << cur[j-1]
      else
        cur << prev[j] + cur[j-1]
      end
    end
    p "cur", cur
    prev = cur.dup
    cur = []
  end
  prev.last
end

class TestUniquePaths < Minitest::Test
  # def test_unique_paths_1
  #   assert_equal 5, unique_paths([[0,1,0],[0,0,1],[1,0,0]])
  # end
  # def test_unique_paths_2
  #   assert_equal 2, unique_paths([[0,0], [0,0]])
  # end
  # def test_unique_paths_3
  #   assert_equal 1, unique_paths([[0,1,1],[1,1,0]])
  # end
  # def test_unique_paths_4
  #   assert_equal 0, unique_paths([[0,1,0,0],[0,0,1,0]])
  # end
  # def test_unique_paths_5
  #   assert_equal 2, unique_paths([[0,0,0],[0,1,0]])
  # end
  # def test_unique_paths_6
  #   assert_equal 2, unique_paths([[0,0],[0,1],[0,0]])
  # end
  # def test_unique_paths_7
  #   assert_equal 3, unique_paths([[0,0,1,1,0],[0,0,0,0,0]])
  # end
  # def test_unique_paths_8
  #   assert_equal 2, unique_paths([[0,0],[1,0],[0,0]])
  # end
  def test_unique_paths_9
    assert_equal 2, unique_paths([[0,0,0],[1,1,1],[0,1,0]])
  end
end