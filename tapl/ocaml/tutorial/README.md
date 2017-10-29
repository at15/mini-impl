# OCamel Tutorials

- [Basics](01-basic.md)
- [Structure of Program](02-structure-of-program.md)
  - let-binding
  - reference
  - nested function
    - lexical scoping ?= static scoping
  - `;` operator, **everything is expression**, C's statement is expression of type `unit`
- [Modules](03-modules.md)


## Book

Real World OCaml

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
