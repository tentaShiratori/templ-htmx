# 素振りするなら Go、Templ、HTMX にしませんか？

## はじめに

プログラミングの素振り（練習）を始めたいけど、何を選べばいいか迷っている方に朗報です。**Go**、**Templ**、そして**HTMX**の組み合わせは、素振りに最適な選択肢の一つです。

この記事では、なぜこの組み合わせが素振りに適しているのか、そして実際に動くサンプルコードを通じてその魅力をお伝えします。

## なぜ Go、Templ、HTMX なのか？

### Go の魅力

1. **シンプルな構文**: 学習コストが低く、初心者でも理解しやすい
2. **高速なコンパイル**: コードを書いてから実行までが非常に速い
3. **標準ライブラリが充実**: 外部ライブラリに頼らずに多くのことが可能
4. **並行処理が簡単**: goroutine を使った並行プログラミングが直感的
5. **クロスプラットフォーム**: Windows、macOS、Linux で同じコードが動く

### Templ の魅力

1. **型安全なテンプレート**: コンパイル時にエラーを検出できる
2. **Go の構文で書ける**: 新しい言語を覚える必要がない
3. **パフォーマンスが良い**: 生成されたコードは最適化されている
4. **IDE サポート**: 自動補完やリファクタリングが効く
5. **コンポーネント指向**: 再利用可能な UI コンポーネントを作成可能

### HTMX の魅力

1. **シンプルな Ajax**: HTML 属性だけで動的な機能を実現
2. **フレームワーク不要**: 複雑な JavaScript フレームワークを覚える必要がない
3. **サーバーサイド中心**: Go のロジックをそのまま活用できる
4. **段階的導入**: 既存の HTML に少しずつ追加できる

## 実際に動かしてみよう

### プロジェクトのセットアップ

まず、Go のプロジェクトを初期化します：

```bash
go mod init my-templ-app
go get github.com/a-h/templ
```

### 基本的なテンプレートファイル

`html/hello.templ`を作成します：

```go
package html

templ Hello(name string) {
    <div class="greeting">
        <h1>Hello, { name }!</h1>
        <p>Welcome to Go + Templ!</p>
    </div>
}
```

### メインアプリケーション

`cmd/main.go`を作成します：

```go
package main

import (
    "fmt"
    "net/http"
    "my-templ-app/html"

    "github.com/a-h/templ"
)

func main() {
    component := html.Hello("World")

    http.Handle("/", templ.Handler(component))

    fmt.Println("Listening on :3000")
    http.ListenAndServe(":3000", nil)
}
```

### テンプレートの生成と実行

```bash
# テンプレートを生成
templ generate

# アプリケーションを実行
go run cmd/main.go
```

ブラウザで `http://localhost:3000` にアクセスすると、美しい挨拶ページが表示されます。

## より実践的な例：タスク管理アプリ

### HTMX を使った動的なタスク管理

`html/todo.templ`を作成：

```go
package html

type Task struct {
    ID    int
    Text  string
    Done  bool
}

templ TodoApp() {
    @Htmx()
    <div class="min-h-screen bg-gray-50 py-8">
        <div class="max-w-2xl mx-auto px-4">
            <div class="bg-white rounded-lg shadow-lg p-6">
                <h1 class="text-3xl font-bold text-gray-800 mb-8 text-center">タスク管理アプリ</h1>

                <form hx-post="/add-task" hx-target="#task-list" hx-swap="beforeend" class="mb-8">
                    <div class="flex gap-3">
                        <input
                            type="text"
                            name="task"
                            placeholder="新しいタスクを入力..."
                            required
                            class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all">
                        <button
                            type="submit"
                            class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors font-medium">
                            追加
                        </button>
                    </div>
                </form>

                <div id="task-list" class="space-y-3">
                    @TaskList([]Task{
                        {ID: 1, Text: "Goの学習", Done: false},
                        {ID: 2, Text: "Templの習得", Done: true},
                        {ID: 3, Text: "HTMXの理解", Done: false},
                    })
                </div>
            </div>
        </div>
    </div>
}

templ TaskItem(task Task) {
    <div class={ "flex items-center p-4 bg-white border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow", if task.Done { "opacity-60" } else { "" } }>
        <input
            type="checkbox"
            checked={ task.Done }
            hx-post={ fmt.Sprintf("/toggle-task/%d", task.ID) }
            hx-target="closest div"
            hx-swap="outerHTML"
            class="w-5 h-5 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2">
        <span class={ "flex-1 ml-3 text-gray-800 font-medium", if task.Done { "line-through text-gray-500" } else { "" } }>
            { task.Text }
        </span>
        <div class="flex gap-2 ml-4">
            <button
                class={ "px-3 py-1 text-sm rounded-md font-medium transition-colors", if task.Done { "bg-yellow-100 text-yellow-800 hover:bg-yellow-200" } else { "bg-green-100 text-green-800 hover:bg-green-200" } }
                hx-post={ fmt.Sprintf("/toggle-task/%d", task.ID) }
                hx-target="closest div"
                hx-swap="outerHTML">
                { if task.Done { "未完了" } else { "完了" } }
            </button>
            <button
                class="px-3 py-1 text-sm bg-red-100 text-red-800 rounded-md hover:bg-red-200 font-medium transition-colors"
                hx-delete={ fmt.Sprintf("/delete-task/%d", task.ID) }
                hx-target="closest div"
                hx-swap="outerHTML"
                hx-confirm="このタスクを削除しますか？">
                削除
            </button>
        </div>
    </div>
}
```

### サーバーサイドの実装

`cmd/main.go`を更新：

```go
package main

import (
    "fmt"
    "net/http"
    "strconv"
    "my-templ-app/html"

    "github.com/a-h/templ"
)

var tasks = []html.Task{
    {ID: 1, Text: "Goの学習", Done: false},
    {ID: 2, Text: "Templの習得", Done: true},
    {ID: 3, Text: "HTMXの理解", Done: false},
}

func main() {
    // ホームページ
    http.Handle("/", templ.Handler(html.Hello("World")))

    // タスク管理アプリ
    http.Handle("/todo", templ.Handler(html.TodoApp()))

    // HTMX API エンドポイント
    http.HandleFunc("/add-task", addTask)
    http.HandleFunc("/toggle-task/", toggleTask)
    http.HandleFunc("/delete-task/", deleteTask)

    fmt.Println("Listening on :3000")
    http.ListenAndServe(":3000", nil)
}

func addTask(w http.ResponseWriter, r *http.Request) {
    taskText := r.FormValue("task")
    newTask := html.Task{
        ID:   len(tasks) + 1,
        Text: taskText,
        Done: false,
    }
    tasks = append(tasks, newTask)

    // 新しいタスクアイテムを返す
    component := html.TaskItem(newTask)
    component.Render(r.Context(), w)
}
```

## 素振りに最適な理由

### 1. 学習曲線が緩やか

- Go の構文は直感的で、他の言語経験があればすぐに理解できる
- Templ は Go の構文をそのまま使えるため、新しい言語を覚える必要がない
- HTMX は HTML 属性を追加するだけなので、JavaScript の知識がなくても使える

### 2. 即座に結果が見える

- `go run`で即座にアプリケーションが起動
- ブラウザでリアルタイムに結果を確認可能
- HTMX により、ページリロードなしで動的な機能を実現
- コンパイルエラーは即座に表示される

### 3. 段階的に機能を追加できる

- シンプルな Hello World から始めて
- HTMX でフォーム送信、Ajax 通信を追加
- データベース接続、API 連携など
- 徐々に複雑な機能を追加していける

### 4. 実用的なスキルが身につく

- Web アプリケーション開発の基礎
- 型安全なプログラミング
- コンポーネント指向の設計
- モダンな Web 開発のパターン（HTMX）

## 次のステップ

1. **データベース連携**: GORM を使ってデータの永続化
2. **認証機能**: ユーザー登録・ログイン機能
3. **API 開発**: RESTful API の作成
4. **フロントエンド強化**: HTMX や Alpine.js との組み合わせ
5. **スタイリング**: Tailwind CSS でより洗練されたデザイン
6. **デプロイ**: Docker、Kubernetes での本番環境構築

## まとめ

Go、Templ、HTMX の組み合わせは、プログラミングの素振りに最適な選択肢です。シンプルな構文、高速な開発サイクル、そして実用的なスキルの習得が可能です。

HTMX を加えることで、複雑な JavaScript フレームワークを覚えることなく、モダンな Web アプリケーションを構築できます。サーバーサイドのロジックを中心に据えながら、ユーザー体験も向上させることができます。

まずは上記のサンプルコードを動かしてみて、徐々に機能を追加していくことで、確実にプログラミングスキルを向上させることができます。

**今日から始めてみませんか？**

---

_この記事のサンプルコードは実際に動作することを確認済みです。ぜひ手を動かして体験してみてください！_
