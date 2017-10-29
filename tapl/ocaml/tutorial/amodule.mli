(* FIXME: this is not working as written in the example http://ocaml.org/learn/tutorials/modules.html .... *)
module Hello : sig
  val hello : unit -> unit
  (** Displays a greeting message *)
end
