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
	var balance float64
	var years float64
	var start float64
	var interest float64
	var need float64

	reader := bufio.NewReader(os.Stdin)

	// prompts
	fmt.Printf("Enter the starting bank account amount: ")
	StartString, _ = reader.ReadString('\n')
	StartString = strings.TrimSpace(StartString)
	start, _ = strconv.ParseFloat(StartString, 64)

	fmt.Printf("Enter the yearly interest rate (as a percentage): ")
	interestString, _ = reader.ReadString('\n')
	interestString = strings.TrimSpace(interestString)
	interest, _ = strconv.ParseFloat(interestString, 64)
	interest = interest/100

	fmt.Printf("Enter the amount needed for post-secondary education: ")
	needString, _ = reader.ReadString('\n')
	needString = strings.TrimSpace(needString)
	need, _ = strconv.ParseFloat(needString, 64)

	// for calculations

	balance = start

	for balance < need {
  balance += balance * interest;
  years++;
}

  fmt.Printf("It will take %2.f years for the starting bank account to reach the required amount for post-secondary education.\n", years)

	fmt.Println("\nDone.")
}
