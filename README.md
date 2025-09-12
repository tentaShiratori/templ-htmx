# Go + Templ + HTMX + Tailwind CSS アプリケーション

## セットアップ

### 1. 依存関係のインストール

```bash
# pnpmでNode.jsの依存関係をインストール
make install

# または直接実行
pnpm install
```

### 2. Tailwind CSS のビルド

```bash
# 本番用ビルド
make css

# 開発用（ファイル監視）
make css-watch
```

### 3. Go アプリケーションの実行

```bash
# テンプレート生成
make gen

# アプリケーション実行
make run
```

## 開発ワークフロー

### 開発モード

```bash
# ターミナル1: Tailwind CSSの監視
make css-watch

# ターミナル2: Goアプリケーションの実行
make run
```

### 本番ビルド

```bash
# すべてをビルド
make build
```

## 利用可能なコマンド

- `make install` - Node.js の依存関係をインストール
- `make css` - Tailwind CSS を本番用にビルド
- `make css-watch` - Tailwind CSS を監視モードでビルド
- `make gen` - Templ テンプレートを生成
- `make fmt` - Go と Templ ファイルをフォーマット
- `make build` - アプリケーションをビルド
- `make run` - アプリケーションを実行
- `make dev` - 開発モード（CSS 監視）

## プロジェクト構造

```
.
├── cmd/                 # Goアプリケーション
├── html/               # Templテンプレート
├── src/                # Tailwind CSS入力ファイル
├── static/             # 静的ファイル（CSS出力先）
├── package.json        # Node.js依存関係
├── tailwind.config.js  # Tailwind設定
└── makefile           # ビルドコマンド
```

## アクセス

- ホームページ: http://localhost:3000
- カウンター: http://localhost:3000/counter
- タスク管理: http://localhost:3000/todo
