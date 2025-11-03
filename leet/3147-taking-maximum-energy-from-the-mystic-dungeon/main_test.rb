require "minitest/autorun"

def maximum_energy(energy, k)
  energy = energy.reverse
  energy.each_with_index do |n, i|
    if i - k >= 0
      val = energy[i-k]
      energy[i] += val
    end
  end
  energy.max
end

class TestMaximumEnergy < Minitest::Test
  def test_maximum_energy
    assert_equal 3, maximum_energy([5,2,-10,-5,1], 3)
    assert_equal(-1, maximum_energy([-2,-3,-1], 2))
  end
end