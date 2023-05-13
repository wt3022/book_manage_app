# 環境構築
.envファイルを作成し、dokcerコンテナを起動します。

ターミナルで以下のコマンドを実行します
```bash
cp project.env .env

docker compose build

docker compose up -d
```

VSCodeを使用している場合はコンテナに接続し、コマンドラインを使用して開発する場合は、以下のコマンドを実行してコンテナに接続してください。
```bash
docker compose exec go_app bash
```


# コンテナ内
コンテナ内に接続後、appsディレクトリで以下のコマンドを入力してください。<br>
(server.envにはサーバーの起動情報が記載されています。)
```bash
cp server.env .env
```