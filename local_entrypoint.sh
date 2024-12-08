#!/bin/sh

mysql_host="go-todo-mysql"
mysql_user="user"
mysql_port="3306"
mysql_password="password"

echo "go-todo-mysqlへの接続を試みています: $mysql_host:$mysql_port"
if ! mysql -h "$mysql_host" -P "$mysql_port" -u "$mysql_user" -p"$mysql_password" -e 'SELECT 1'; then
    echo "go-todo-mysqlへの接続に失敗しました"
    exit 1
fi

echo "go-todo-mysqlへの接続に成功しました"

# airを起動
exec air -c .air.toml
