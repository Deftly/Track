/*
 * Overview:
 * The rest and spread operators in JavaScript are both represented by three
 * dots(...) but they serve different purproses depending on the context. The
 * rest operator collects multiple elements into an array, while the spread
 * operator expands an iterable into individual elements.
 *
 * Importance:
 * - Flexibility: These operators provide a concise way to work with arrays and
 *   objects.
 * - Immutability: They help in creating new arrays and objects without modifying
 *   the original ones, which is crucial in React for state management.
 * - Function arguments: The rest operator simplifies working with variable
 *   numbers of function arguments.
 */

// Rest Operator:
const ex1 = () => {
  function sum(...numbers) {
    return numbers.reduce((total, num) => total + num, 0);
  }
  console.log(sum(1, 2, 3, 4)); // 10
};

// Spread Operator:
const ex2 = () => {
  // Arrays
  const arr1 = [1, 2, 3];
  const arr2 = [...arr1, 4, 5];

  console.log(arr2); // [1, 2, 3, 4, 5]

  // Objects
  const obj1 = { x: 1, y: 2 };
  const obj2 = { ...obj1, z: 3 };

  console.log(obj2); // {x: 1, y: 2, z: 3}

  // Copying arrays and objects:
  const originalArray = [1, 2, 3];
  const copiedArray = [...originalArray];

  const originalObject = { a: 1, b: 2 };
  const copiedObject = { ...originalObject };

  // Spreading elements into a function call
  const numbers = [1, 2, 3];
  console.log(Math.max(...numbers));
};

// Combining Rest and Spread:
const ex3 = () => {
  const [first, ...rest] = [1, 2, 3, 4, 5];
  const newArray = [0, ...rest];

  console.log(first); // 1
  console.log(rest); // [2, 3, 4, 5]
  console.log(newArray); // [0, 2, 3, 4, 5]
};
