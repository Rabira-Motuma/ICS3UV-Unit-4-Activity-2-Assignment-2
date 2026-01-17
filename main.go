/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2026-01-5
 * Fileoverview: This program prints a pattern of numbers in a right-angled triangle.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	// variables
	var StartString string
	var interestString string
	var needString string
	var balance int
	var years int
	var start int
	var interest int
	var need int

	reader := bufio.NewReader(os.Stdin)

	// prompts
	fmt.Printf("Enter the starting bank account amount: ")
	StartString, _ = reader.ReadString('\n')
	StartString = strings.TrimSpace(StartString)
	start, _ = strconv.Atoi(StartString)

	fmt.Printf("Enter the yearly interest rate (as a percentage): ")
	interestString, _ = reader.ReadString('\n')
	interestString = strings.TrimSpace(interestString)
	interest, _ = strconv.Atoi(interestString)

	fmt.Printf("Enter the amount needed for post-secondary education: ")
	needString, _ = reader.ReadString('\n')
	needString = strings.TrimSpace(needString)
	need, _ = strconv.Atoi(needString)

	// for calculations

	balance = start

	for balance < need {
  balance += balance * interest;
  years++;
}

  fmt.Printf("It will take %d years for the starting bank account to reach the required amount for post-secondary education.", years)

	fmt.Println("\nDone.")
}
