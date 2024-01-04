package concurrency

import (
	"maps"
	"testing"
	"time"
)

const (
	google = "http://google.com"
	blog   = "http://blog.gypsydave5.com"
	waat   = "waat://furhurterwe.geds"
)

func mockWebSiteChecker(url string) bool {
	return url != waat
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		google, blog, waat,
	}

	want := UrlResponses{
		google: true,
		blog:   true,
		waat:   false,
	}

	got := CheckWebsites(mockWebSiteChecker, websites)

	if !maps.Equal(want, got) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}
