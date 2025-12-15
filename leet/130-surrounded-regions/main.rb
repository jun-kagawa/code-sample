require 'minitest/autorun'

def solve(board)
  memo = Set.new
  nomemo = Set.new
  m = board.size
  n = board[0].size
  m.times do |i|
    n.times do |j|
      next if nomemo.include?([i, j])

      next unless board[i][j] == 'O'

      if check(board, i, j, memo)
        memo.to_a.each do |k, l|
          board[k][l] = 'X'
        end
      else
        nomemo.merge(memo)
      end
      memo = Set.new
    end
  end
end

def check(board, i, j, memo)
  return false if i - 1 == -1 || j - 1 == -1
  return false if i + 1 >= board.size || j + 1 >= board[0].size

  memo.add?([i, j])
  index = []
  index << [i, j + 1] if board[i][j + 1] == 'O'
  index << [i, j - 1] if board[i][j - 1] == 'O'
  index << [i + 1, j] if board[i + 1][j] == 'O'
  index << [i - 1, j] if board[i - 1][j] == 'O'
  index.map do |z|
    next if memo.include?(z)

    check(board, z[0], z[1], memo)
  end.filter { |k| !k.nil? }.all?
end

class TestSolve < Minitest::Test
  def test_solve
    board = [%w[X X X X], %w[X O O X], %w[X X O X], %w[X O X X]]
    solve(board)
    assert_equal [%w[X X X X], %w[X X X X], %w[X X X X], %w[X O X X]], board
  end

  def test_solve_1
    board = [%w[X O X], %w[O X O], %w[X O X]]
    solve(board)
    assert_equal [%w[X O X], %w[O X O], %w[X O X]], board
  end
end

# XOX
# OXO
# XOX
