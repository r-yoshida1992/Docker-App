# Dockerイメージの指定
FROM ubuntu:16.04

# ホスト側からDockerコンテナ内にコピーする
COPY helloworld.sh /usr/local/bin

# Dockerコンテナ実行前に行われるコマンドを定義
CMD ["helloworld.sh"]