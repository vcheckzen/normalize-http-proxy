# Normalize HTTP Proxy

A http proxy server chaining a upstream which needs authentication headers.

```txt
local -> [np] -> upstream -> destination
```

## Usage

- Normal use

```shell
➜  ./np baidu
Normalize Proxy: listens at 127.0.0.1:8888.
Http upstream is http://cloudnproxy.baidu.com:443, extra header(s) is(are):
X-T5-Auth: ZjQxNDIh
Https upstream is http://cloudnproxy.baidu.com:443, extra header(s) is(are):
X-T5-Auth: ZjQxNDIh
Verbose output is off.
```

```shell
➜  curl -x localhost:8888 https://baidu.com
<html>
<head><title>302 Found</title></head>
<body bgcolor="white">
<center><h1>302 Found</h1></center>
<hr><center>bfe/1.0.8.18</center>
</body>
</html>
```

- [Chaining with vanila vmess](chain/vanila-vmess.json)
- [Chaining with vless + ws](chain/vless-ws-tls.json)

## Credit

- [Pre-proxy](https://github.com/v2ray/v2ray-core/issues/1736)
- [Pre-proxy](https://guide.v2fly.org/app/parent.html#%E5%9F%BA%E6%9C%AC%E9%85%8D%E7%BD%AE-v2ray-4-21-0)
- [Proxy chain](https://oi.0w0.io/2020/06/28/v2ray%E9%85%8D%E7%BD%AE%E5%89%8D%E7%BD%AE%E4%BB%A3%E7%90%86-%E4%BB%A3%E7%90%86%E9%93%BE-%E9%93%BE%E5%BC%8F%E4%BB%A3%E7%90%86%E8%BD%AC%E5%8F%91/)
- [Baidu http proxy](https://yunz.blog4j.top/articles/2021/07/30/1627615492050.html)
- [Dedicated traffic](https://rainchan.win/2021/07/30/%E5%A6%82%E4%BD%95%E4%BC%98%E9%9B%85%E4%BD%BF%E7%94%A8%E5%AE%9A%E5%90%91%E6%B5%81%E9%87%8F/)
- [Reasonable use of tricky operations on the Internet](https://xlmy.net/2020/11/07/%E7%90%86%E5%88%A9%E7%94%A8%E7%BD%91%E7%BB%9C%E4%B8%8A%E7%9A%84%E9%AA%9A%E6%93%8D%E4%BD%9C/)
- [Release two traffic-free interfaces](http://bbs.liuxingw.com/t/15074.html)
- [Douyu/KingCard traffic-free dynamic interface](http://www.saoml.com/index.php/2020/02/21/8.html)