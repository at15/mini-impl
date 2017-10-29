# OCaml

The original implementation from https://www.cis.upenn.edu/~bcpierce/tapl/checkers/

- [Tutorials notes](tutorial)

## Installation

````
sudo apt-get install ocaml opam camelp4
opam init
opam install merlin # auto complete
opam install utop # the REPL
````

````
Quick setup for VIM
-------------------
Append this to your .vimrc to add merlin to vim's runtime-path:
  let g:opamshare = substitute(system('opam config var share'),'\n$','','''')
  execute "set rtp+=" . g:opamshare . "/merlin/vim"

Also run the following line in vim to index the documentation:
  :execute "helptags " . g:opamshare . "/merlin/vim/doc"
````

## TAPL

### Arith

Actually it's pretty close to CMPS203 assignment https://github.com/at15/CMPS203/tree/master/hw/hw1

### info

`support.ml` defines support functions, `support.mli` define the interface (signature)

`info` is a variant type, which could be a tuple (string, int, int) labeled as FI, or UNKNOWN:q:""

````
module Error = struct
type info = FI of string * int * int | UNKNOWN
end

type info = Error.info
````

and in

````
let printInfo =
  (* In the text of the book, file positions in error messages are replaced
     with the string "Error:" *)
  function
    FI(f,l,c) ->
      print_string f;
      print_string ":";
      print_int l; print_string ".";
      print_int c; print_string ":"
  | UNKNOWN ->
      print_string "<Unknown file and line>: "
````

### term

`term` is variant type defined in `syntax.ml`

- see http://ocaml.org/learn/tutorials/data_types_and_matching.html, it's like C struct with union

````ocaml
type term =
    TmTrue of info
  | TmFalse of info
  | TmIf of info * term * term * term
  | TmZero of info
  | TmSucc of info * term
  | TmPred of info * term
  | TmIsZero of info * term

(* Extracting file info *)

let tmInfo t = match t with
    TmTrue(fi) -> fi
  | TmFalse(fi) -> fi
  | TmIf(fi,_,_,_) -> fi
  | TmZero(fi) -> fi
  | TmSucc(fi,_) -> fi
  | TmPred(fi,_) -> fi
  | TmIsZero(fi,_) -> fi
````    

### eval

evaluation is defined in `core.ml`
