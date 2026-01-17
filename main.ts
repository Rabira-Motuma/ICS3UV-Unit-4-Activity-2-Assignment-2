/**
* @author Rabira Motuma
* @version 1.0.0
* @date 2026-01-13
* @fileoverview This program calculates banking and tuition info for a student.
*/

// variables
let StartString: string;
let interestString: string;
let needString: string;
let balance: number;
let years: number = 0;
let start: number;
let interest: number;
let need: number;

// prompts
StartString = prompt("Enter the starting bank account amount:") || "0";
start = parseFloat(StartString);

interestString = prompt("Enter the yearly interest rate (as a percentage):") || "0";
interest = parseFloat(interestString) / 100;

needString = prompt("Enter the amount needed for post-secondary education:") || "0";
need = parseFloat(needString);

// calculations

balance = start;

while (balance < need) {
  balance += balance * interest;
  years++;
}

// output
console.log(`It will take ${years} years for the starting bank account to reach the required amount for post-secondary education.`);

console.log("\nDone.");
