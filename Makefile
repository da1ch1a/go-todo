## compose-up: ローカル用のdocker-compose起動
compose-up:
	docker compose up -d

## compose-down: ローカル用のdocker-composeの破棄
compose-down:
	docker compose down

# マイグレーションを最新までまとめてupする
migrate-db:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://root:rootpassword@tcp(localhost:33060)/go_todo_local" up

# マイグレーションを1つdownする
migrate-down-db:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "mysql://root:rootpassword@tcp(localhost:33060)/go_todo_local" down 1

# マイグレーションファイルを作成する
migrate-new:
	docker run -v ./migrations:/migrations --network host migrate/migrate create -ext sql -dir /migrations/ -seq file_name
