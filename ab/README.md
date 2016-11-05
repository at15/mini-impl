# AB

- [Void HTTP Server](server)

## Key points 

- go lang currency 
- net/http/httptrace
- aggregate data

## TODO

- [ ] the QPS in `hey` means QPS per worker? 

## NOTE

- need to change the file descriptor. `ulimit -n` show current one `ulimit -n 4096` to set limit in current shell. 
- the max file descriptor can be found using `cat /proc/sys/fs/file-max`
- `invalid memory address or nil pointer dereference`

## Ref 

- https://github.com/rakyll/hey