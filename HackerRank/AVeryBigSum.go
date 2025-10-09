/*
In this challenge, you need to calculate and print the sum of elements in an array, considering that some integers may be very large.

Function Description

Complete the aVeryBigSum function.
It must return the sum of all array elements.

aVeryBigSum has the following parameter(s):

int ar[n]: an array of integers

Return

long: the sum of all array elements

Input Format

The first line of the input consists of an integer n.
The next line contains n space-separated integers contained in the array.

Output Format

Return the integer sum of the elements in the array.

Constraints
1 ≤ n ≤ 10
0 ≤ ar[i] ≤ 10^10

Sample Input
5
1000000001 1000000002 1000000003 1000000004 1000000005

Sample Output
5000000015

Note

The range of the 32-bit integer is (−2^31) to (2^31−1).
When we add several integer values, the resulting sum might exceed this range.
You might need to use long int (C/C++/Java) or an equivalent type to store such sums.
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
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
    // Write your code here
    var sum int64 = 0
    for _, num := range ar {
        sum += num
    }
    return sum
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    arCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ar []int64

    for i := 0; i < int(arCount); i++ {
        arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        ar = append(ar, arItem)
    }

    result := aVeryBigSum(ar)

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
