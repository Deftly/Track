# Dynamic Dispatch

## What is Dynamic Dispatch?
Dynamic dispatch, also known as dynamic method dispatch or runtime polymorphism, is a feature in object-oriented programming languages where the method to be invoked is determined at runtime based on the actual type of the object, rather than the declared type of the variable holding the object reference.

In languages with dynamic dispatch, when a method is called on an object, the runtime system looks up the actual type of the object and invokes the appropriate implementation of the method based on that type. This allows for polymorphic behavior, where different objects can respond to the same method call in different ways.

## Go and Dynamic Dispatch
Go's approach to polymorphism and method invocation is different from language like C++, Java, or C# that heavily rely on dynamic dispatch.

In Go, methods are defined on types, not on objects. When a method is called on a value of a certain type, the method corresponding to that type is invoked. Go uses a static dispatch mechanism, where the method to be called is determined at compile-time based on the type of the value.

Go achieves polymorphic behavior through interfaces. An interface defines a set of methods, and any type that implements all the methods of an interface is said to implement that interface. When a variable of an interface type holds a value of a concrete type, the method calls on that variable are dispatched to the appropriate implementation based on the underlying concrete type.

However, the dispatch in Go is still static in nature. The compiler generates the necessary code to perform the method dispatch based on type information available at compile-time. This doesn't involve the same kind of runtime lookup and dynamic dispatch as in languages with virtual methods.

## Method Dispatch with Interfaces in Go
In Go, when you declare a variable of an interface type, the variable can hold any value that implements the methods defined by the interface. The actual type of the value stored in the interface variable is known as the concrete type.

When you call a method on an interface variable, the compiler needs to generate the code to dispatch the method call to the appropriate implementation based on the concrete type of the value. This is where the static dispatch comes into play.

The compiler uses type information available at compile-time to generate the necessary code for method dispatch. Here's how the process works:
1. Type Assertion: When you assign a value to an interface variable, the compiler performs a type assertion to ensure that the value implements the methods required by the interface. If the value doesn't implement all the necessary methods, a compile-time error is generated.
2. Type Information: The compiler maintains type information for each value stored in an interface variable. This type information includes the concrete type of the value and the memory layout of that type.
3. Method Table (itable): For each concrete type that implements an interface, the compiler generates a method table, also known as an itable. The method table is a data structure that maps the methods of the interface to the corresponding implementations of the concrete type.
4. Interface Value: An interface value in Go consists of two components: a type pointer and a value pointer. The type pointer points to the type information of the concrete value, which includes the method table. The value pointer points to the actual value stored in the interface.
5. Method Dispatch: When a method is called on an interface variable, the compiler generates code to perform the following steps:
  - Retrieve the type information from the interface value's type pointer.
  - Look up the corresponding method in the method table based on the method signature.
  - Call the method implementation associated with the concrete type.

The compiler generates this dispatch code based on the type information available at compile-time. It doesn't involve any runtime type information or dynamic lookup.

