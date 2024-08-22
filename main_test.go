package main

import (
	"os"
	"testing"
)

func TestTrimNonLetterPrefix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1. Hello World", "Hello World"},
		{"2. Xin chào", "Xin chào"},
		{"...Test", "Test"},
		{"123ABC", "ABC"},
		{"!@#$%^&*()Việt Nam", "Việt Nam"},
	}

	for _, test := range tests {
		result := trimNonLetterPrefix(test.input)
		if result != test.expected {
			t.Errorf("trimNonLetterPrefix(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestRemoveDiacritics(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Xin chào thế giới", "Xin chao the gioi"},
		{"Đây là một câu tiếng Việt", "Day la mot cau tieng Viet"},
		{"Hello World", "Hello World"},
		{"Café au lait", "Cafe au lait"},
	}

	for _, test := range tests {
		result := removeDiacritics(test.input)
		if result != test.expected {
			t.Errorf("removeDiacritics(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestToDashCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello-world"},
		{"Xin chào thế giới", "xin-chao-the-gioi"},
		{"This is a TEST", "this-is-a-test"},
		{"Đây là MỘT câu TIẾNG Việt", "day-la-mot-cau-tieng-viet"},
		{"Spaces   and  !@#$%^&*()  Punctuation", "spaces-and-punctuation"},
	}

	for _, test := range tests {
		result := toDashCase(test.input)
		if result != test.expected {
			t.Errorf("toDashCase(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestGetFormatedFileName(t *testing.T) {
	nameFile := "./name.txt"
	bakFile := "./name.txt.bak"

	// Create a temporary file
	content := []byte("1. Hello World\n2. Xin chào thế giới\n3. This is a TEST\n4. Đây là MỘT câu TIẾNG Việt")
	tmpfile, err := os.CreateTemp("", nameFile[2:])
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Temporarily replace the name.txt path
	os.Rename(nameFile, bakFile)
	os.Rename(tmpfile.Name(), nameFile)
	defer func() {
		os.Remove(nameFile)
		os.Rename(bakFile, nameFile)
	}()

	expected := []string{
		"hello-world",
		"xin-chao-the-gioi",
		"this-is-a-test",
		"day-la-mot-cau-tieng-viet",
	}

	result := getFormatedFileName()

	if len(result) != len(expected) {
		t.Fatalf("getFormatedFileName() returned %d items, expected %d", len(result), len(expected))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("getFormatedFileName()[%d] = %q, expected %q", i, v, expected[i])
		}
	}
}

func TestGetNumberedFileName(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1.png", 1},
		{"10.jpg", 10},
		{"100.txt", 100},
	}

	for _, test := range tests {
		result := getNumberedFileName(test.input)
		if result != test.expected {
			t.Errorf("getNumberedFileName(%q) = %d, expected %d", test.input, result, test.expected)
		}
	}
}
