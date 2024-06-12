# Compilation and Linking

## Compilation Stage
The compilation stage is the process of converting the source code into object code, which is a lower-level representation that can be executed by the computer.
1. Preprocessing: The preprocessor reads the source code and handles directives such as `#include`, `#define`, and conditional compilation statements. It expands macros, includes header files, and performs any necessary text substitutions.
2. Lexical Analysis: The lexical analyzer, also known as the scanner or tokenizer, breaks down the preprocessed code into a sequence of tokens. It identifies and categorizes the individual elements of the code, such as keywords, identifiers, constants, and operators.
3. Syntax Analysis: The syntax analyzer, also called the parser, takes the sequence of tokens and checks if they conform to the grammar rules of the C language. It constructs an abstract syntax tree (AST) that represents the structure of the code.
4. Semantic Analysis: The semantic analyzer performs various checks on the AST to ensure the code is semantically correct. It verifies type consistency, checks for undeclared variables, and performs type conversions as needed.
5. Intermediate Code Generation: The compiler generates an intermediate representation (IR) of the code, such as three-address code or quadruples. This IR is a lower-level representation that is closer to machine code but still independent of the target machine architecture.
6. Optimization: The optimization phase applies various techniques to improve the efficiency and performance of the generated code. Common optimizations include constant folding, dead code elimination, loop unrolling, and code reordering.
7. Code Generation: The final step of the compilation stage is to generate the object code from the optimized intermediate representation. The code generator translates the IR into the target machine language, such as assembly code or machine instructions.

## Linking Stage:
The linking stage is the process of combining the compiled object files, along with any necessary libraries, into an executable program.
1. Symbol Resolution: The linker resolves references to external symbols, such as functions and global variables, by matching them with their definitions in other object files or libraries.
2. Relocation: The linker adjusts the memory addresses of the code and data to create a cohesive executable. It ensures that all references to memory locations are properly resolved.
3. Library Linking: If the program uses external libraries, the linker searches for the required library files and includes them in the final executable. It resolves the references to library functions and variables.
4. Symbol Table Generation: The linker generates a symbol table that contains information about the symbols used in the program, such as function names and global variable names. This table is used for debugging and runtime symbol resolution.
5. Executable Generation: Finally, the linker combines all the object files, resolved symbols, and libraries into a single executable file that can be run on the target machine.
