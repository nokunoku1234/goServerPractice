# サーバー起動方法
1. パッケージの依存関係に修正がある場合は `go mod tidy` を実行
2. `docker compose build` を実行
3. `docker compose up -d` を実行

# データモデルの作成
1. `go run -mod=mod entgo.io/ent/cmd/ent new X` を実行
2. `go generate ./ent` を実行（generate.goの設定でこれで生成可能になっている）

# DBへのアクセス方法
- `docker compose exec ${SERVICE_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}` を実行
  - `${SERVICE_NAME}` はdocker-compose.ymlに記載

# SQLクエリの処理速度などを計測
- `EXPLAIN (ANALYZE, BUFFERS, VERBOSE) 計測したいSQL文`
- 内容次第でIndexの作成などを検討
- width: 高ければ転送するデータ量が重く、ネットワークやメモリ負荷が高い → 改善の余地あり
  - Columnの設定を見直して不要な情報を取得しないようにする