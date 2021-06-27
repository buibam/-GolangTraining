package main

import (
	"testing"
)

func BenchmarkMovieRead(b *testing.B) {
	readJSON("movie.json", func(data map[string]interface{}) bool {
		return data["year"].(float64) >= 2010
	})
}

func BenchmarkMovieReadToken(b *testing.B) {
	readJSONToken("movie.json", func(data map[string]interface{}) bool {
		return data["year"].(float64) >= 2010
	})
}

func BenchmarkQRead(b *testing.B) {
	readJSON("question.json", func(data map[string]interface{}) bool {
		return data["show_number"].(string) == "4680"
	})
}

func BenchmarkQReadToken(b *testing.B) {
	readJSONToken("question.json", func(data map[string]interface{}) bool {
		return data["show_number"].(string) == "4680"
	})
}