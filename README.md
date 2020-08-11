# Set-Cookie_sample
golang EchoでSet-Cookieのサンプルを作ってみる。

## 参考にした記事一覧
- [Go言語フレームワークechoのcookieでつまずいた話。](https://www.pnkts.net/2018/04/25/golang-echo-cookie/)
- [Cookieの基本](https://qiita.com/yassh/items/2088c6d026fb6b66806e)
- [Echo公式ドキュメント](https://echo.labstack.com/)

## 起動手順
``` bash
$ git clone git@github.com:Takumaron/Set-Cookie_sample.git

$ cd Set-Cookie_sample

$ make run
```

## 検証
### WriteCookie
[:input]に指定したPathParamの値をset-cookieヘッダーに含めてレスポンスを返してくれる。[:input]で指定した値がブラウザのCookieに格納される。

http://localhost:1323/write/:input

### ReadCookie
現状ブラウザに保存されているCookieの中から、inputというKeyに紐づくValueをサーバ側で取得し、取得した値をレスポンスとして返却してくれる。
もし、ブラウザのCookieにデータが存在しない場合はサーバから`input: default value`というset-cookieヘッダー付きでレスポンスが返却される。

http://localhost:1323/read
