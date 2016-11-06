# AB

- [Void HTTP Server](server)

## Key points 

- go lang currency 
- net/http/httptrace
- aggregate data

## TODO

- [ ] the QPS in `hey` means QPS per worker? 
- [ ] timeout in both client and server https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/

## NOTE

- need to change the file descriptor. `ulimit -n` show current one `ulimit -n 4096` to set limit in current shell. 
- the max file descriptor can be found using `cat /proc/sys/fs/file-max`
- `invalid memory address or nil pointer dereference`, response can be nil when err is not nil 
- need to call `res.Body.Close()` so both server and client and release resource to avoid file descriptor running out.
- MUST make shallow copy of request, otherwise `res.Body.Close()` would cause all the reuqest being cancelled

## Ref 

- https://github.com/rakyll/hey