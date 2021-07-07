# String Rewriting Exercise

#### Rewrite Rules
```go
ab -> ba
ba -> ab
aa -> 
b -> 
```

Strings to Rewrite :


* abba:
```
-> baba
-> baab
-> bb
-> b
-> b
```

> the last step is not correct as there is no rule b->b


* bababa:
```
-> abbaba
-> abbbaa
-> abbb
-> abb
-> ab
-> ba
...
```

> you show me that there is an infinite reduction sequence and this is correct
> can you also reduce to a normal form?

In order for an ARS to be considered terminating, there needs to be a defined measure function. 

> this is not correct: the existence of a measure function is a sufficient condition for termination, not a necessary one

The ARS above does not contain a measure function 

> I dont understand what it means to say "does not contain a measure function"

and therefore does not terminate. Because `ab -> ba` and `ba -> ab`, the string potentially remians the same length and stuck in a cycle of rewriting.

> correct

> what about the other questions: How many equivalence classes does ⟷∗ have? Can you describe them in a nice way? What are the normal forms? Can you change the rules so that the ARS becomes terminating without changing its equivalence classes? Which measure function proves termination of your modified system? Write down a question or two about strings that can be answered using the ARS. Think about whether this amounts to giving a semantics to the ARS.
