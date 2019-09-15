package calc

import "testing"

func TestSum(t *testing.T) {
	result := Sum(10, 20)

	if result != 30 {
		t.Errorf("TestSum is failed. expect:%d, actual:%d", 30, result)
	}

	t.Logf("result is a %d", result)
}
