package signal

import (
	"testing"
)

func TestSample(t *testing.T) {
	const n = 256

	a := make(Discrete, n)
	b := make(Discrete, n)

	a.Sample(SineFunc, 1/n, 0)
	b.Sample(a.At, 1/n, 0)

	for i := range a {
		if a[i] != b[i] {
			t.Fail()
		}
	}

	if t.Failed() {
		for i := range a {
			t.Logf("%v: a[%.4f] b[%.4f]\n", i, a[i], b[i])
		}
	}
}

var resf float64

func BenchmarkDiscreteInterp(b *testing.B) {
	sig := Sine()
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		resf = sig.Interp(float64(n) / float64(len(sig)))
	}
}

func BenchmarkDiscreteAt(b *testing.B) {
	sig := Sine()
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		resf = sig.At(float64(n) / float64(len(sig)))
	}
}

func BenchmarkDiscreteIndex(b *testing.B) {
	sig := Sine()
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		resf = sig.Index(n)
	}
}
