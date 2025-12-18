require 'minitest/autorun'


# @param {String} s
# @return {String[][]}
def partition(s)
  splited = s.split('')
  rr = []
  splited.size.times do |i|
    r = []
    splited.size.times do |j|
      s = splited[j,i+1].join('')
      r << s if s == s.reverse
      break if s == splited.join('')
    end
    rr << r if r.size > 0
  end
  rr
end

def partition(s)
  return [[]] if s == ""
  ans = []
  s.size.times do |i|
      if s[..i] == s[..i].reverse
        partition(s[i+1..]).map do |suf|
          ans << ([s[..i]] + suf)
        end
      end
    end
  ans
end

class TestPartition < Minitest::Test
  def test_partition
    assert_equal [['a', 'a', 'b'], ['aa', 'b']], partition('aab')
    assert_equal [['a']], partition('a')
    assert_equal [["c","d","d"],["c","dd"]], partition('cdd')
  end
end

