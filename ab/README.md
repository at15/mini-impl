# AB

- [Void HTTP Server](server)
- [Load](load.go)

## Usage 

- `cd server && go run main.go` to start a http server on `localhost:8000`, everything is hard coded, https code is commented out.
- `cd ab && go run main.go` to make request to `localhost:8000`, everything is hard coded right now


## Key points 

- go lang currency 
- net/http/httptrace
- aggregate data

## TODO

- [ ] the QPS in `hey` means QPS per worker? 
- [x] timeout in both client and server https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/

## NOTE

- need to change the file descriptor. `ulimit -n` show current one `ulimit -n 4096` to set limit in current shell. 
- the max file descriptor can be found using `cat /proc/sys/fs/file-max`
- `invalid memory address or nil pointer dereference`, response can be nil when err is not nil 
- need to call `res.Body.Close()` so both server and client and release resource to avoid file descriptor running out.
- MUST make shallow copy of request, otherwise `res.Body.Close()` would cause all the reuqest being cancelled
- `connect: cannot assign requested address` https://github.com/golang/go/issues/16012 seems need to set maxidle connection per host
run `ab/ab/main` serveral times and this problem will appear
    - [OS does not release port right after the connection is closed](http://grokbase.com/t/gg/golang-nuts/156b4z7w57/go-nuts-dial-tcp-127-0-0-1-3333-cant-assign-requested-address-when-calling-http-get-too-quickly)
    - running hey several times does not have this problem 
    - need to specify transport for client in order to reuse connection (TODO: why not do it by default)

## Ref 

- https://github.com/rakyll/hey