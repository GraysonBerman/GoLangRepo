# Syntax
During the first few weeks of class we started talking about the following concepts:
* Syntax
* Semantics
* Soundness
* Completeness

Syntax is the definition of language and semantics are the rules we give to that language. 
Soundness verifys that our rules agree the semantics given. Completeness requires that we have defines all the rules 
needed to complete any computation in our language. 

### The Syntax of Go
The syntax of Go is simple and easy to comprehend having previous experience with other programming languages.
There are no semi-colons to separate statements, instead using white-space (similar to Python). The main method is run in 
the `func main()` method and you need to import specific packages to use certain features.

Here is a basic Hello, World program in Go:

```go
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```
