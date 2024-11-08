package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two strings",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with string and number",
			struct {
				Name string
				Age  int
			}{"Chris", 27},
			[]string{"Chris"},
		},
		{
			"struct with nested fields",
			Person{"Chris",
				Profile{"London", 27}},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{"Chris",
				Profile{"London", 27}},
			[]string{"Chris", "London"},
		},
		{
			"handle slices",
			[]Profile{
				Profile{"London", 27},
				Profile{"Berlin", 27},
			},
			[]string{"London", "Berlin"},
		},
		{
			"handle arrays",
			[2]Profile{
				Profile{"London", 27},
				Profile{"Hamburg", 27},
			},
			[]string{"London", "Hamburg"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("handle maps", func(t *testing.T) {
		var got []string
		testMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		walk(testMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("handle channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{"Rome", 1}
			aChannel <- Profile{"Madrid", 2}
			close(aChannel)
		}()

		var got []string
		want := []string{"Rome", "Madrid"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("handle functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{"Berlin", 33}, Profile{"Katowice", 33}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected haystack %v to contain %v but it didn´t", haystack, needle)
	}
}
