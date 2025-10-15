# @param {String} s
# @return {Integer}
def min_deletions(s)
  m = Hash.new(0)
  s.each_char { |s| m[s] += 1}
  count_arr = m.sort_by {|k, v| -v}.map {|pair| pair.last}

  count = 0
  i = 1
  while i < count_arr.size
    until count_arr[i] < count_arr[i-1] || count_arr[i] == 0
      count_arr[i] -= 1
      count += 1
    end
    i += 1
  end
  count
end