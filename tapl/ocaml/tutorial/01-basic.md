# The Basics

http://ocaml.org/learn/tutorials/basics.html

````
utop # let average a b = (a +. b) /. 2.0;;
val average : float -> float -> float = <fun>  
```

- statically typed
- type inference
  - no function overload
- the last expression in a function is the returned result

Basic types

````
int 31-bit, 1 bit is used for GC
char no support for Unicode or UTF-8
unit ()
````

Need explicit cast

````
utop # float_of_int 1;;
- : float = 1.   
````

Ordinary functions and recursive functions, `rec` keyword is needed for correct lookup

````
utop # let rec range a b =
if a > b then []
else a :: range (a+1) b;;
val range : int -> int -> int list = <fun>   
utop # range 1 5;;
- : int list = [1; 2; 3; 4; 5]   
````

`a :: range (a+1) b` is list construction I suppose

Types of function

- [ ] currying later chapter

Polymorphic function

- `a'` means any type, like the `Top` in TAPL?

````
utop # let give_me_three x = 3;;
val give_me_three : 'a -> int = <fun>
````

Type inference
