# Log

Related issue #2

## Nginx

http://articles.slicehost.com/2010/8/27/reading-nginx-web-logs

Combined log format is used by modern servers

````
127.0.0.1 - - [01/Dec/2016:22:41:58 -0800] "GET / HTTP/1.1" 200 396 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.100 Safari/537.36"
127.0.0.1 - - [01/Dec/2016:22:41:58 -0800] "GET /favicon.ico HTTP/1.1" 404 209 "http://localhost/" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.100 Safari/537.36"
127.0.0.1 - - [01/Dec/2016:22:42:53 -0800] "GET /sb.html HTTP/1.1" 404 209 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.100 Safari/537.36"
````

The second entry above is "-", which is what gets logged when there's nothing to put in that part of the log. In this case, the entry would represent the name of a remote log, if one were being used. You'll pretty much always see "-" here.

The third entry above is another "-". That slot contains the username the web client was authorized under, if any. If you enabled password protection for a file or directory, then the username the visitor used to log in would be recorded here.

`ip - - [date] "Method url HTTP Version" status response-size "refer-url" user-agent`
