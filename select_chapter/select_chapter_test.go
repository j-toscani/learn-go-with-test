package select_chapter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createDelayedTestServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader((http.StatusOK))
	}))
}

func TestRacer(t *testing.T) {
	t.Run("returns an error if a server doens`t repsond within 10 seconds", func(t *testing.T) {
		slowServer := createDelayedTestServer(20 * time.Millisecond)
		fastServer := createDelayedTestServer(19 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		_, err := ConfigurableRacer(slowUrl, fastUrl, time.Millisecond*10)

		if err == nil {
			t.Errorf("expected an error but did not get one")
		}
	})
	t.Run("returns the first server that responds", func(t *testing.T) {
		slowServer := createDelayedTestServer(3 * time.Millisecond)
		fastServer := createDelayedTestServer(2 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Errorf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
