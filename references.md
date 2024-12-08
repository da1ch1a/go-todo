## html/template

- GoのWebフレームワーク Echoの "Template Rendering" を利用してHTMLページを作成する  
https://yhidetoshi.hatenablog.com/entry/2022/03/28/213000

## echo

- echo Templates  
https://echo.labstack.com/docs/templates

## migration

- golang-migrate/migrate  
https://github.com/golang-migrate/migrate?tab=readme-ov-file#docker-usage
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#usage

## sqlboiler

- MySQLのDATEとDATETIMEはGoで出力するとデフォルトは[]byteになるので、これをtime.Timeに変換するためにDB接続時にパラメータにparseTime=trueを指定する  
https://qiita.com/stern327/items/cb325b44a0335deea402

## zerolog, sqldblogger

- zerolog  
- 高速、JSON形式で出力できる、ログレベルごとに異なる出力先を設定可能  
https://github.com/rs/zerolog

- sqldblogger  
- A logger for Go SQL database driver without modify existing *sql.DB stdlib usage.
https://github.com/simukti/sqldb-logger

- init関数  
- mainパッケージでない場合はimportするだけで呼び出されるため初期化に使われる。mainパッケージで書くとmain関数より先に実行される。
https://qiita.com/tenntenn/items/7c70e3451ac783999b4f

## docker

- docker-composeでコンテナが起動するまで待つ  
- alpineではbashではなくshなので注意  
https://github.com/Eficode/wait-for
https://docs.docker.jp/compose/startup-order.html
