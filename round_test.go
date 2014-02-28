package round

import "testing"

func TestCompilation(t *testing.T) {
	result := Compilation(10000)
	if len(result) != 9999 {
		t.Fatal("Wrong number of pairings")
	}
	if len(result[0]) != 10000 {
		t.Fatal("Wrong number of related results")
	}
}

func TestCompilationDetailed(t *testing.T) {
	result := Compilation(6)
	if len(result) != 5 {
		t.Fatal("Wrong number of pairings")
	}
	if len(result[0]) != 6 {
		t.Fatal("Wrong number of related results")
	}
}

func BenchmarkCompilation(b *testing.B) {
	Compilation(5000)
}
func BenchmarkLargeCompilation(b *testing.B) {
	Compilation(50000)
}

func TestPrintMatchups(t *testing.T) {
	PrintMatchups([]string{"abba", "bards", "cats", "ducks", "elephants"})
}
