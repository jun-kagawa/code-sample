
def minimize_the_difference(mat, target)
  base = 0
  magic = mat.map do |row|
    sorted = row.sort.uniq
    base += sorted.first
    sorted.drop(1).map { |v| v - sorted.first}
  end
  min_diff = (target - base).abs
  return min_diff if base >= target

  h = { base => true}

  magic.each do |row|
    more = []
    row.each do | delta |
      h.keys.each do |k|
        v = k + delta
        next if v > 870
        more << v unless h[v]
        return 0 if v == target
      end
    end
    more.each {|mv| h[mv] = true}
  end

  h.keys.reduce(min_diff) do |acc, v|
    [acc, (target - v).abs].min
  end
end