# !/bin/bash
# MySQLサーバーが起動するまで待機する
until mysqladmin ping -h mysql -P 3306 --silent;
do
    echo "MySQLの起動を待っています"
    sleep 2;
done
echo "APIを起動します"
exec go run main.go;