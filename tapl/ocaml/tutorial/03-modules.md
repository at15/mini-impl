# Modules

## Basic usage

- each file automatically defines a module

amodule.ml

````
let hello () = print_endline "Hello"
````

bmodule.ml

````
Amodule.hello ()
````

`make 03-modules`

## Interfaces and signatures

Only expose part of the module to user by using a `mli` file

amodule.mli

````
val hello : unit -> unit
(** Displays a greeting message. *)
````

`.mli` files must be compiled just before the matching `.ml` files. They are compiled using `ocamlc`, even if .ml files are compiled to native code using `ocamlopt`

````
ocamlc -c amodule.mli
ocamlopt -c amodule.ml
````

## Abstract types

````
type date = { day : int;  month : int;  year : int }
````

- omit, (not in mli)
- copy and paste
- abstract, only give name `type date`
- fields made read-only `type data = private { ... }`

- [ ] TODO: can't figure out the first two typing,,,,

````
type date
val create : ?days:int -> ?months:int -> ?years:int -> unit -> date
val sub : date -> date -> date
val years : date -> float
````

- [ ] TODO: how does ocamel's type inference know which one to use, it seems to be using the closest one

````
type date {day: int; month: int; year: int}
type date2 {day: int; month: int}
let years d = d.day
val years : date2 -> int = <fun>  
let years (d: date) = d.day;;
val years : date -> int = <fun>  
````

## Submodules

use `module ... end` to define a submodule in a file, which is  also a module

````
module Hello = struct
  let message = "Hello"
  let hello () = print_endline message
end
let goodbye () = print_endline "Goodbye"
let hello_goodbye () =
  Hello.hello ();
  goodbye ()
````

````
let () =
  Example.Hello.hello ();
  Example.goodbye ()
````
