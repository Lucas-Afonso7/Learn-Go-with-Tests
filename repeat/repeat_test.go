package integers

import "testing"

func TestRepeat(t *testing.T) {
	repetitions := Repeat("a")
	expected := "aaaaa"

	if repetitions != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repetitions )
	}
}

func BenchmarkRepetir(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}