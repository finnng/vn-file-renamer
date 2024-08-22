# File Renamer

## Why This Exists

This tool was created to automate the process of renaming files in a specific format. It's particularly useful for scenarios where you have a series of numbered files (e.g., "1.png", "2.png", etc.) and you want to append descriptive names to them while maintaining their original order.

## How It Works

1. The program reads all files in the current directory.
2. It filters for files that match the pattern of a number followed by a file extension (e.g., "1.png", "2.txt").
3. It reads a list of new names from a file called "name.txt".
4. The new names are processed to remove diacritics, convert to lowercase, and replace spaces with dashes.
5. The program then renames each file by appending the processed new name to the original number, maintaining the original file extension.

For example, if you have:
- Files: "1.png", "2.png", "3.png"
- Contents of name.txt: "First Image", "Second Image", "Third Image"

The result will be:
"1-first-image.png", "2-second-image.png", "3-third-image.png"

## Examples

The program handles various cases of input text. Here are some examples of how input from `name.txt` is processed:

1. Basic English text:
   Input: "Hello World"
   Output: "hello-world"

2. Vietnamese text with diacritics:
   Input: "Xin chào thế giới"
   Output: "xin-chao-the-gioi"

3. Mixed case text:
   Input: "This is a TEST"
   Output: "this-is-a-test"

4. Vietnamese text with mixed case:
   Input: "Đây là MỘT câu TIẾNG Việt"
   Output: "day-la-mot-cau-tieng-viet"

5. Text with spaces and punctuation:
   Input: "Spaces   and  !@#$%^&*()  Punctuation"
   Output: "spaces-and-punctuation"

6. Text with non-letter prefix:
   Input: "1. Hello World"
   Output: "hello-world"

7. Vietnamese text with non-letter prefix:
   Input: "2. Xin chào"
   Output: "xin-chao"

8. Text with numbers and special characters:
   Input: "H3llo W0rld! Spec!al Ch@rs"
   Output: "h3llo-w0rld-spec-al-ch-rs"

9. Preserving case for non-Vietnamese characters:
   Input: "IBM and VIỆT NAM"
   Output: "ibm-and-viet-nam"

These examples demonstrate how the program:
- Converts text to lowercase (except for non-Vietnamese uppercase characters)
- Removes diacritics from Vietnamese characters
- Replaces spaces and non-alphanumeric characters with dashes
- Removes non-letter prefixes
- Handles mixed Vietnamese and English text

When using this program, ensure your `name.txt` file contains the desired names, one per line. The program will process each line and use the result to rename the corresponding numbered file in the directory.

## Steps to Run This Code Locally

1. Ensure you have Go installed on your system.
2. Clone this repository or download the `main.go` file.
3. Create a `name.txt` file in the same directory as `main.go`. Each line in this file should contain a name for the corresponding numbered file.
4. Place the files you want to rename in the same directory.
5. Open a terminal and navigate to the directory containing `main.go`.
6. Run the program with the command:

```
go run main.go
```

## How to Build for Target OS

To build the program for a specific operating system, you can use Go's cross-compilation feature. Here are examples for common target systems:

### For Windows (from any OS):

```
GOOS=windows GOARCH=amd64 go build -o file_renamer.exe main.go
```

### For macOS (from any OS):
```
GOOS=darwin GOARCH=amd64 go build -o file_renamer main.go
```

### For Linux (from any OS):
```
GOOS=linux GOARCH=amd64 go build -o file_renamer main.go
```

Replace `amd64` with `386` for 32-bit systems if needed.

After building, transfer the executable to the target system and run it in the directory containing the files to be renamed and the `name.txt` file.

Note: Ensure that you have the necessary permissions to rename files in the directory where you run this program.
