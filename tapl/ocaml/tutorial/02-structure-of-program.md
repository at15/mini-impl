# The Structure of OCaml Programs

http://ocaml.org/learn/tutorials/structure_of_ocaml_programs.html

## Local "variables" (really local expressions)

`let name = expression in` used to define a named local expression,
and name can be used later

````
utop # let average a b =
  let sum = a +. b in
  sum /. 2.0;;
val average : float -> float -> float = <fun>   
utop # let average a b =
  let sum = a +. b in
  sum /. 2.0;;
val average : float -> float -> float = <fun>   
````

````
# let f a b =
    (a +. b) +. (a +. b) ** 2.;;
val f : float -> float -> float = <fun>
# let f a b =
    let x = a +. b in
    x +. x ** 2.;;
val f : float -> float -> float = <fun>
````

The second one is faster if the compiler don't perform optimization

## Global "variables" (really global expressions)

- [ ] TODO: read_whole_file is undefined

````
let html =
  let content = read_whole_file file in
  GHtml.html_from_string content
  ;;
````

- There is no space allocated to "store" the "html pointer". Nor is it possible to assign anything to html

## Let-bindings

Any use of `let ...`, whether at the top level (globally) or within a function, is often called a let-binding

- [ ] TODO: is it more like constant?
- 11.5 in TAPL P124

## References: real variables

 References are very similar to pointers in C/C++. In Java, all variables which store objects are really references (pointers) to the objects. In Perl, references are references - the same thing as in OCaml

- you have to explicit dereference it to use the value in store

````
utop # let my_ref = ref 0;;
val my_ref : int ref = {contents = 0}     
utop # let my_ref = ref 0;;
val my_ref : int ref = {contents = 0}   
utop # my_ref;;
- : int ref = {contents = 10086}   
utop # my_ref + 1;;
Error: This expression has type int ref                                                                                                              but an expression was expected of type int   
utop # !my_ref + 1;;
- : int = 10087   
````  

## Nested functions

- lexical scoping

````
# let read_whole_channel chan =
    let buf = Buffer.create 4096 in
    let rec loop () =
      let newline = input_line chan in
      Buffer.add_string buf newline;
      Buffer.add_char buf '\n';
      loop ()
    in
    try
      loop ()
    with
      End_of_file -> Buffer.contents buf;;
val read_whole_channel : in_channel -> string = <fun>
````

## Modules and `open`

## The sequence operator `;`

- can be seen as `unit -> 'b -> 'b`, ignore the value for first one and return the second one
  - as we see in TAPL

**in OCaml, nearly everything is an expression**, including control flow like `if/then/else`

OCaml's definition of what is an expression is just a little wider than C's.
In fact, C has the concept of "statements" — but all of C's statements are just expressions in OCaml of type unit (combined with the ; operator)

All functions in OCaml can be expressed as

````
let name [parameters] = expression ;;
````

## The disappearance of `;;`

- `;;` is mostly used in the toplevel and tutorials to mark the end of an OCaml phrase and send this phrase to the toplevel in order to evaluate it.
- discouraged and not needed in real code
  - `let () = expression ()` or
  - `let _ = expression ()`

## Putting it all together: some real code

- `?foo` optional argument, `~foo` named argument
- `foo#bar` method invocation
