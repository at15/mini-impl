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
