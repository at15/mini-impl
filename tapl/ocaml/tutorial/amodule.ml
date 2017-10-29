(* FIXME: this is not working .... *)
module Hello : Hello_type = struct
  let message = "Hello"
  let hello () = print_endline message
end
