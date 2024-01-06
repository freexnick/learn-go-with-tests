package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name    string
	Profile Profile
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	var contains bool

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"User"},
			[]string{"User"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"User", "City1"},
			[]string{"User", "City1"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"User", 1},
			[]string{"User"},
		},
		{
			"nested fields",
			Person{
				"User",
				Profile{1, "City1"},
			},
			[]string{"User", "City1"},
		},
		{
			"pointers to things",
			&Person{
				"User",
				Profile{1, "City1"},
			},
			[]string{"User", "City1"},
		},
		{
			"slices",
			[]Profile{
				{1, "City1"},
				{2, "City2"},
			},
			[]string{"City1", "City2"},
		},
		{
			"arrays",
			[2]Profile{
				{1, "City1"},
				{2, "City2"},
			},
			[]string{"City1", "City2"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"User": "Test",
			"City": "City1",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Test")
		assertContains(t, got, "City1")
	})

	t.Run("with channels", func(t *testing.T) {
		var got []string
		want := []string{"City1", "City2"}
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{1, "City1"}
			aChannel <- Profile{2, "City2"}
			close(aChannel)
		}()

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		var got []string
		want := []string{"City1", "City2"}

		aFunction := func() (Profile, Profile) {
			return Profile{1, "City1"}, Profile{2, "City2"}
		}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
