require 'minitest/autorun'

require 'set'

class DSU
  attr_reader :parent, :size

  def initialize(n)
    @parent = (0..n).to_a
    @size = Array.new(n + 1, 1)
  end

  def find(x)
    @parent[x] == x ? x : @parent[x] = find(@parent[x])
  end

  def unite(a, b)
    a = find(a)
    b = find(b)
    return if a == b
    a, b = b, a if @size[a] < @size[b]
    @parent[b] = a
    @size[a] += @size[b]
  end
end

def process_queries(c, connections, queries)
  dsu = DSU.new(c)
  connections.each { |a, b| dsu.unite(a, b) }
  p dsu.parent

  # 各成分にヒープ（最小値取り出し用）
  heaps = Hash.new { |h, k| h[k] = [] }
  (1..c).each do |i|
    r = dsu.find(i)
    heaps[r] << i
  end
  heaps.each_value(&:sort!) # Ruby ではsort済配列をheap代わりに使う

  offline = Array.new(c + 1, false)
  ans = []

  queries.each do |t, x|
    if t == 2
      offline[x] = true
    else
      if !offline[x]
        ans << x
      else
        r = dsu.find(x)
        heap = heaps[r]
        # 既にオフラインのものを削除
        heap.shift while !heap.empty? && offline[heap[0]]
        ans << (heap.empty? ? -1 : heap[0])
      end
    end
  end

  ans
end


class TestProcessQueries < Minitest::Test
  def test_process_queries
    assert_equal [3,2,3], process_queries(5, [[1,2],[2,3],[3,4],[4,5]], [[1,3],[2,1],[1,1],[2,2],[1,2]])
  end
end