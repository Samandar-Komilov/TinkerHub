package main

import "testing"

func ExampleMain() {
	main()
	// Output: Hello World!
}

func TestGreet_English(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello World!",
		},
		"Russian": {
			lang: "ru",
			want: "Привет, мир!",
		},
		"Uzbek": {
			lang: "uz",
			want: "Assalom dunyo!",
		},
		"Latin": {
			lang: "lt",
			want: `unsupported language: "lt"`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
