package calc

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("result %d expected %d given, %v", result, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slicesEqual(result, expected) {
		t.Errorf("result %v expected %v", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, result, expected []int) {
		t.Helper()
		if !slicesEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	}

	t.Run("make the sums of tails of some slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, result, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, result, expected)
	})
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
