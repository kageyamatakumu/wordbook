# Node.jsのベースイメージを取得
FROM node:latest

# ターゲットディレクトリを作成
RUN mkdir -p /app

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# package.json ファイルをコピー
COPY package.json .

# 依存関係をインストール
RUN npm install

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN npm run build

# ポートを公開
EXPOSE 3000

# コンテナ起動時に実行されるコマンド
CMD ["npm", "start"]
