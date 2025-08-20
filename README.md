# Go-sample-projects-
sample projects 

**Functions** in Go are central to the language, serving as blocks of code that perform specific tasks, taking inputs, processing them, and generating outputs.
 A function is declared using the func keyword, followed by the function name, parameters in parentheses, and the return type.
 For example, a function that adds two integers can be declared as func plus(a int, b int) int { return a + b }.
** Methods:** A method in Go is declared like a function but includes a receiver parameter before the function name. The receiver specifies the type the method is associated with.
 func (receiver ReceiverType) MethodName(params) returnType {
    // Method implementation
}
A type's method set is the collection of methods defined for it. The method set differs based on whether the type is used as a value or a pointer:

Value type (T): Includes methods with value receivers only.
Pointer type (*T): Includes methods with both value and pointer receivers.
