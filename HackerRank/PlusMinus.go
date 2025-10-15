/*
Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with 6 places after the decimal.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to 10^-4 are acceptable.

Example

There are 6 elements: two positive, two negative, and one zero. Their ratios are 0.5, 0.333333, and 0.166667. Results are printed as:

0.500000
0.333333
0.166667

Function Description

Complete the plusMinus function below.

func plusMinus(arr []int32)

Parameters

arr – an array of integers

Print

Print the ratios of positive, negative, and zero values in the array. Each value should be printed on a separate line with 6 digits after the decimal.
The function should not return a value.

Input Format

The first line contains an integer, n, the size of the array.
The second line contains n space-separated integers that describe arr.

Constraints

0 < n ≤ 100

-100 ≤ arr[i] ≤ 100

Sample Input
6
-4 3 -9 0 4 1

Sample Output
0.500000
0.333333
0.166667

Explanation

There are 3 positive numbers, 2 negative numbers, and 1 zero in the array.
The proportions are:

Positive: 3/6 = 0.500000

Negative: 2/6 = 0.333333

Zero: 1/6 = 0.166667
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
	var positive, negative, zero int32
	length := int32(len(arr))
	for _, value := range arr {
		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("%.6f\n", float64(positive)/float64(length))
	fmt.Printf("%.6f\n", float64(negative)/float64(length))
	fmt.Printf("%.6f\n", float64(zero)/float64(length))
	}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
