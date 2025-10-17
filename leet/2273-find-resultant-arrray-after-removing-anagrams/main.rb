require "minitest/autorun"

def remove_anagrams(words)
  words.reduce([]) do |arr, word|
    s = word.split("").sort.join
    if arr[-1].nil? || arr.last.last != s
      arr << [word, s]
    end
    arr
  end.map { |arr| arr[0]}
end

class TestRemoveAnagrams < Minitest::Test
  def test_remove_anagrams
    assert_equal ["abba", "cd"], remove_anagrams(["abba","baba","bbaa","cd","cd"])
    assert_equal ["a","b","c","d","e"], remove_anagrams(["a","b","c","d","e"])
  end
end