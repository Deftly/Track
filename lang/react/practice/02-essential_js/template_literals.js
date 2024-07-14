/*
 * Overview:
 * Template literals, introduced in ES6, are a way to create strings in
 * JavaScript that allows for more flexible and readbale string formatting.
 * They use backticks(`) instead of single or double quotes and support
 * multi-line strings and string interpolation.
 *
 * Importance:
 * - Readability
 * - Expressions in strings: They allow embedding expressions directly in strings.
 * - Multi-line strings
 */

// Basic Syntax:
const text = "World";
const greeting = `Hello, ${text}!`;

console.log(greeting); // Hello, World!

// Multi-line Strings:
const multiLine = `
  This is a
  multi-line
  string.
`;

console.log(multiLine);
// Output:
//  This is a
//  multi-line
//  string.

// Expression Interpolation:
const a = 5;
const b = 10;
console.log(`Fifteen is ${a + b} and not ${2 * a + b}.`);

// Nesting Templates
const person = { name: "John", age: 30 };
const message = `${person.name} is ${person.age > 18 ? `an adult` : `a minor`}.`;
console.log(message); // John is an adult.

// Tagged Templates:
function highlight(strings, ...values) {
  return strings.reduce(
    (acc, str, i) =>
      `${acc}${str}<span class="highlight">${values[i] || ""}</span>`,
    "",
  );
}

const name1 = "Alice";
const age = 28;
const highlightedText = highlight`${name1} is ${age} years old.`;

console.log(highlightedText);

// Raw Strings:
console.log(`Line 1\nLine 2`);
// Line 1
// Line 2

console.log(String.raw`Line 1\nLine 2`);
// Line 1\nLine 2
