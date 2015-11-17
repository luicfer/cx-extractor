package extractor

import (
	"io/ioutil"
	"testing"
)

func Benchmark_Extractor(b *testing.B) {
	b.StopTimer()
	f, _ := ioutil.ReadFile("./test.html")
	html := string(f)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Extractor(html)
	}
}
