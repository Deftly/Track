/*
 * Overview:
 * Destructuring is a JavaScript feature that allows you to extract values from
 * objects or arrays and assign them to variables in a more concise and readable
 * way. It provides a shorthand syntax for unpacking values from data structures
 * into distinct variables.
 *
 * Importance:
 * - Cleaner code: Destructuring makes your code more concise and easier to read.
 * - Flexible assignment: Allows extraction of multiple values in a single line of code.
 * - Default values: You can assign default values when destructuring, this is
 * - particularly useful in React for handling props.
 */

const ex1 = () => {
  const person = { name: "John", age: 30, city: "New York" };
  const { name, age } = person;

  console.log(name); // John
  console.log(age); // 30

  const colors = ["red", "green", "blue"];

  const [first, second] = colors;
  console.log(first); // red
  console.log(second); // green
};

// Nested Destructuring:
const ex2 = () => {
  const user = {
    id: 42,
    displayName: "jdoe",
    fullName: {
      firstName: "John",
      lastName: "Doe",
    },
  };

  const {
    id,
    displayName,
    fullName: { firstName, lastName },
  } = user;

  console.log(id); // 42
  console.log(displayName); // jdoe
  console.log(firstName); // John
  console.log(lastName); // Doe
};

// Default Values:
const ex3 = () => {
  const { name = "Guest", age = 18 } = {};
  console.log(name); // Guest
  console.log(age); // 18
};

// Rest Pattern:
const ex4 = () => {
  const { id, ...rest } = { id: 1, name: "John", age: 30 };
  console.log(id); // 1
  console.log(rest); // {name: "John", age: 30}
};
