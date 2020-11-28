# Dockerイメージの指定
FROM golang:1.9

# イメージビルド時にDockerコンテナ内で実行するコマンドを定義
RUN mkdir /echo

# ホスト側からDockerコンテナ内にコピーする
COPY main.go /echo

# Dockerコンテナ実行前に行われるコマンドを定義
CMD ["go", "run", "/echo/main.go"]