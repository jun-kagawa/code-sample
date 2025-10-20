require "minitest/autorun"

def find_lex_smallest_string(s, a, b)
  h = {}
  q = [s]
  small = s

  until q.empty?
    cur = q.pop
    next if h[cur]

    h[cur] = true
    small = cur if cur.to_i < small.to_i

    add = add_operation(cur, a)
    rota = rotate(cur, b)

    q << add unless h[add]
    q << rota unless h[rota]
  end
  small
end

def add_operation(s, a)
  s.chars.each_with_index.map do |c, i|
    i.odd? ? ((c.to_i + a) % 10).to_s : c
  end.join
end

def rotate(s, b)
  s[s.size-b...s.size] + s[0...s.size-b]
end

class TestFindLexSmallestString < Minitest::Test
  def test_find_lex_smallest_string
    assert_equal "2050", find_lex_smallest_string("5525", 9, 2)
    assert_equal "24", find_lex_smallest_string("74", 5, 1)
    assert_equal "0011", find_lex_smallest_string("0011", 4, 2)
  end
end