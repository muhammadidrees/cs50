package cs50

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

/**
 * Prompts user for a line of text from standard input and returns
 * it as a string (char *), sans trailing line ending. Supports
 * CR (\r), LF (\n), and CRLF (\r\n) as line endings. If user
 * inputs only a line ending, prompts user again. Returns ""
 * upon error or no input whatsoever (i.e., just EOF).
 */
func GetString(format string, a ...interface{}) string {

	// initialize a reader
	reader := bufio.NewReader(os.Stdin)

	// initialize an input string
	input := ""

	// keep promting the user until he gives input i.e. some character and not just
	// press the enter key giving new-line checking for CR (Mac OS), LF (Linux), 
	// and CRLF (Windows) new line charactersa
	for input == "" || input == "\n" || input == "\r" || input == "\r\n"{

		// display prompt to user 
		fmt.Fprintf(os.Stdout, format, a...)

		// read till new line i.e. till the user hits the enter key
		x, e := reader.ReadString('\n')

		// on error return ""
		if e != nil{
			return "";
		}

		// store the typed input [x] in variable input
		input = x

	}

	// remove the escape character i.e. \r\n in case of windows from 
	// and \r or \n in MAC or Linux input string that is placed by
	// hitting the enter key
	if (input[len(input) - 2] == '\r') {
		return input[:len(input) - 2]
	} else {
		return input[:len(input) - 1]
	}	
}

/**
 * Prompts user for a line of text from standard input and returns the
 * equivalent char; if text is not a single char, user is prompted
 * to retry. If line can't be read, returns "".
 */
 func GetChar(format string, a ...interface{}) byte {
	 
	// initialize the input variable
	input := ""

	// keep prompting user until he inputs a single character
	for input == "" || len(input) > 1 {
		// get input string from command line
		input = GetString(format, a...)
	}

	// return input character
	return input[0]
 }

/**
 * Prompts user for a line of text from standard input and returns the
 * equivalent int; if text does not represent an int in [-2^31, 2^31 - 1)
 * or would cause underflow or overflow, user is prompted to retry.
 */
func GetInt(format string, a ...interface{}) int {

	// keep prompting user until he inputs a single character
	for {

		// get input string from command line
		input := GetString(format, a...)

		// convert string to int
		i, err := strconv.Atoi(input)

		// on successful conversion return integer value
		if err == nil {
			return i
		}
	}
}

/**
 * Prompts user for a line of text from standard input and returns the
 * equivalent float as precisely as possible; if text does not represent
 * a float or if value would cause underflow or overflow, user is prompted
 * to retry.
 */
func GetFloat(format string, a ...interface{}) float64 {

	// keep prompting user until he inputs a single character
	for {

		// get input string from command line
		input := GetString(format, a...)

		// convert string to float
		i, err := strconv.ParseFloat(input, 64)

		// on successful conversion return float value
		if err == nil {
			return i
		}
	}
}