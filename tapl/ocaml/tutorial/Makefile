.PHONY: 03-modules
03-modules:
	ocamlc -c amodule.mli
	ocamlopt -c amodule.ml
	ocamlopt -c bmodule.ml
	ocamlopt -o hello amodule.cmx bmodule.cmx

.PHONY: clean
clean:
	rm -f hello *.o *.cmx *.cmi
