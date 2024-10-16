### DBマイグレーション

- 一番簡単なコマンド 最新までまとめてup
docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://root:rootpassword@tcp(localhost:33060)/go_todo_local" up


- ホストマシンにあるマイグレーションファイルをコンテナにマウントし、コンテナ内のマイグレーションファイルがあるパスを指定、dbに接続して最後にupコマンド。upの後にマイグレーションファイルの番号を指定することで、その番号までのマイグレーションを実行する。


- 一番簡単なコマンド 1つdown
docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://root:rootpassword@tcp(localhost:33060)/go_todo_local" down 1
