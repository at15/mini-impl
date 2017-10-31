# Real World OCaml

https://realworldocaml.org/v1/en/html/a-guided-tour.html

- seems more structured than the tutorial on the website

variant type

The `|` character separates the different cases of the variant (the first | is optional), and each case has a capitalized tag, like Circle, Rect or Segment, to distinguish that case from the others

````
# type scene_element =
    | Circle  of circle_desc
    | Rect    of rect_desc
    | Segment of segment_desc
  ;;
type scene_element =
    Circle of circle_desc
  | Rect of rect_desc
  | Segment of segment_desc
````

ref type

The ref type comes predefined in the standard library, but there's nothing really special about it. It's just a record type with a single mutable field called contents, you can re implement it

````
# type 'a ref = { mutable contents : 'a }

  let ref x = { contents = x }
  let (!) r = r.contents
  let (:=) r x = r.contents <- x
  ;;
type 'a ref = { mutable contents : 'a; }
val ref : 'a -> 'a ref = <fun>
val ( ! ) : 'a ref -> 'a = <fun>
val ( := ) : 'a ref -> 'a -> unit = <fun>
````

https://realworldocaml.org/v1/en/html/variables-and-functions.html

https://stackoverflow.com/questions/27631215/how-should-i-do-clear-screen-on-ocaml-toplevel  Ctrl+I is better

````
let clear () = Sys.command("clear");;
````

nameless function

````
(fun x -> x + 1);;
````

named function

````
let plusone = (fun x -> x + 1);;
let plusone x = x + 1;; (* the syntax sugar *)
````

> Functions and let bindings have a lot to do with each other. In some sense, you can think of the parameter of a function as a variable being bound to the value passed by the caller. Indeed, the following two expressions are nearly equivalen

multiargument, curry

````
# let abs_diff x y = abs (x - y);;
val abs_diff : int -> int -> int = <fun>
````

a multiargument function is actually a function that returns a function

````
# let abs_diff =
    (fun x -> (fun y -> abs (x - y)));;
val abs_diff : int -> int -> int = <fun>
````

The practice of applying some of the arguments of a curried function to get a new function is called partial application.

Recursive Functions

`let rec`

Infix operator is also functions

````
(+) 1 2
- : int = 3
````

Labeled Arguments

````
# let ratio ~num ~denom = float num /. float denom;;
val ratio : num:int -> denom:int -> float = <fun>
# ratio ~num:3 ~denom:10;;
- : float = 0.3
# ratio ~denom:10 ~num:3;;
- : float = 0.3
````

FIXED: but how to have the example in TAPL (bool -> bool) -> bool -> bool, need explicit, or use mli https://stackoverflow.com/questions/6005176/ocaml-explicit-type-signatures

````
utop # let f (f2: bool -> bool) = f2;;
val f : (bool -> bool) -> bool -> bool = <fun>
````
