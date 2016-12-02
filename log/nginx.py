# parse nginx log
import re

# parse combined log format
def parse_nginx_log(log_str):
    exp = '((?:\d{1,3}\.){3}\d{1,3})\s-\s-\s\[(.*)\]\s"(\S*)\s(\S*)\sHTTP/(\S*)"\s(\d*)\s(\d*)\s"(\S*)"\s"(.*)"'
    r = re.compile(exp)
    return r.findall(log_str)


success_str = '127.0.0.1 - - [01/Dec/2016:22:41:58 -0800] "GET / HTTP/1.1" 200 396 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.100 Safari/537.36"'
print(parse_nginx_log(success_str))

success_str_with_refer ='127.0.0.1 - - [01/Dec/2016:22:41:58 -0800] "GET /favicon.ico HTTP/1.1" 404 209 "http://localhost/" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.100 Safari/537.36"'
print(parse_nginx_log(success_str_with_refer))

notfound_str = '127.0.0.1 - - [01/Dec/2016:22:42:53 -0800] "GET /sb.html HTTP/1.1" 404 209 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.100 Safari/537.36"'
print(parse_nginx_log(notfound_str))
