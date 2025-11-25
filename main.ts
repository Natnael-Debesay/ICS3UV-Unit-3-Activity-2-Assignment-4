/**
 * @author Natnael Debesay
 * @version 1.0.0
 * @date 2025-11-24
 * @fileoverview This program calculates the interest value at the end of ten years.
 */

import * as readlineSync from "readline-sync";
const YEARS = 10;

// input
const initial = Number(readlineSync.question("Enter initial value: "));
const rate = Number(readlineSync.question("Enter interest rate (e.g., 0.07): "));

function calcInterest(previousValue: number): number {
  return Number ((previousValue* rate).toFixed(2));
}

function calcNewValue(previousValue: number): number {
  return Number((previousValue + calcInterest(previousValue)).toFixed(2));
}

// Year 1
const interest1 = calcInterest(initial);
const value1 = calcNewValue(initial);

// Year 2
const interest2 = calcInterest(value1);
const value2 = calcNewValue(value1);

// Year 3
const interest3 = calcInterest(value2);
const value3 = calcNewValue(value2);

// Year 4
const interest4 = calcInterest(value3);
const value4 = calcNewValue(value3);

// Year 5
const interest5 = calcInterest(value4);
const value5 = calcNewValue(value4);

// Year 6
const interest6 = calcInterest(value5);
const value6 = calcNewValue(value5);

// Year 7
const interest7 = calcInterest(value6);
const value7 = calcNewValue(value6);

// Year 8
const interest8 = calcInterest(value7);
const value8 = calcNewValue(value7);

// Year 9
const interest9 = calcInterest(value8);
const value9 = calcNewValue(value8);

// Year 10
const interest10 = calcInterest(value9);
const value10 = calcNewValue(value9);

// Output
console.log("\nYear | Initial Value | Interest Gained | Current Value");
console.log("--------------------------------------------------------");

console.log(`1    | $${initial.toFixed(2)}    | $${interest1.toFixed(2)}    | $${value1.toFixed(2)}`);
console.log(`2    | $${initial.toFixed(2)}    | $${interest2.toFixed(2)}    | $${value2.toFixed(2)}`);
console.log(`3    | $${initial.toFixed(2)}    | $${interest3.toFixed(2)}    | $${value3.toFixed(2)}`);
console.log(`4    | $${initial.toFixed(2)}    | $${interest4.toFixed(2)}    | $${value4.toFixed(2)}`);
console.log(`5    | $${initial.toFixed(2)}    | $${interest5.toFixed(2)}    | $${value5.toFixed(2)}`);
console.log(`6    | $${initial.toFixed(2)}    | $${interest6.toFixed(2)}    | $${value6.toFixed(2)}`);
console.log(`7    | $${initial.toFixed(2)}    | $${interest7.toFixed(2)}    | $${value7.toFixed(2)}`);
console.log(`8    | $${initial.toFixed(2)}    | $${interest8.toFixed(2)}    | $${value8.toFixed(2)}`);
console.log(`9    | $${initial.toFixed(2)}    | $${interest9.toFixed(2)}    | $${value9.toFixed(2)}`);
console.log(`10    | $${initial.toFixed(2)}    | $${interest10.toFixed(2)}    | $${value10.toFixed(2)}`);

console.log("\nDone.")
