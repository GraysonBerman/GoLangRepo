## Excersise on Termination

Show that the following programs terminate:

```
while m =/= n do
  if m > n then m := m — n else n := n — m
  ```
  
  and 
  
 ```
 while m =/= n  do
  if m > n then m : = m — n
  else begin h :=m; m :=n; n := h endterminate.go
  ```
  
  Implementation in Go:
  
  ```go
  package main
import "fmt"
  
  func terminate1(m, n int) {
    for m != n {
      if m > n {
        m -= n
      }else {
       n -= m
      }
    }
    fmt.Println("Done!")
  }

  func terminate2(m, n int) {
  	for m != n {
      if m > n {
        m -= n
      }else {
       h := m
       m = n
       n = h
      }
    }
    fmt.Println("Done!")
  }

  func main() {
   terminate1(5,7)
   terminate2(5,7)
  }
  ```
  
  The program above will output:
  ```
  Done!
  Done!
  ```
  ### Assumptions
  In order for this program to work, we need to assume that the inputs are positive. The way for a function to terminate is to have a easure function
  that is decreasing with every step. In this case, the programs are decreasing towards 0 so if the inputs are not positive, then the function
  will continue to increase in the negative direction and never terminate.
