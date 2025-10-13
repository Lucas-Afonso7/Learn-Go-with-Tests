/*
Problem: Diagonal Difference
Given a square matrix, calculate the absolute difference between the sums of its diagonals.
Example:
1 2 3
4 5 6
9 8 9

The left-to-right diagonal = 1 + 5 + 9 = 15
The right-to-left diagonal = 3 + 5 + 9 = 17
Their absolute difference is |15 - 17| = 2

Function Description
Complete the function with the following parameter:

arr: a 2D array of integers

Return

The absolute difference in sums along the diagonals

Input Format
The first line contains a single integer n, the number of rows and columns in the square matrix.
Each of the next n lines describes a row and consists of space-separated integers.
Constraints

1 ≤ n ≤ 100

Sample Input
3
11 2 4
4 5 6
10 8 -12
Sample Output
15
Explanation
The primary diagonal is:
11
   5
     -12
Sum across the primary diagonal: 11 + 5 + (-12) = 4
The secondary diagonal is:
     4
   5
10
Sum across the secondary diagonal: 4 + 5 + 10 = 19
Difference: |4 - 19| = 15
Note: |x| is the absolute value of x
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
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
    // Write your code here
    primarySum := int32(0)
    secondarySum := int32(0)
    
    for i := 0; i < len(arr); i++ {
        primarySum += arr[i][i]
        secondarySum += arr[i][len(arr)-1-i]
    }
    diference := primarySum - secondarySum
    if diference < 0 {
        diference = -diference
    }
    return diference
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var arr [][]int32
    for i := 0; i < int(n); i++ {
        arrRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(n) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := diagonalDifference(arr)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
