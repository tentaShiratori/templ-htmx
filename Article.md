# 素振りするなら Go、Templ、HTMX にしませんか？

## はじめに

バックエンドに関連する技術で素振りをしたいって時にバックエンドとフロントエンド分けて構築するのは面倒だなって思う時ありますよね。
とはいえRailsは今更触りたくない、phpやpythonも宗教上の理由で嫌。JavaやKotlinはInteliJが重いから嫌。素振りでrustを触るのは面倒。という方が多いと思います（本当か？）

そんな方のためにGo, Templ, htmxという組み合わせを紹介します。

この記事では、なぜこの組み合わせが素振りに適しているのか、そして実際に動くサンプルコードを通じてその魅力をお伝えします。

## 閑話：ところで素振りってなんだ？

納得感のある説明がありました。

https://azukiazusa.dev/blog/write-articles-for-technical-practice/

> 技術の素振りを、ここではある特定の言語やフレームワークに対する理解を深めるために、その技術を使って何かしらの成果物を作成することと定義します。素振りの目的としては、ドキュメントからは読み取れない Pro/Con を得ること、その技術が実際のプロジェクトで使えるかどうか調査するといった事項があげられるでしょう。

ここではこれを採用します。

## なぜ Go、Templ、HTMX なのか？

学習コスト、実装コストを最小限に抑えられるので素振りに集中できます。

### Go の魅力

1. **シンプルな構文**: 学習コストが低く、初心者でも理解しやすい
2. **高速なコンパイル**: コードを書いてから実行までが非常に速い
3. **並行処理が簡単**: goroutine を使った並行プログラミングが直感的

### Templ の魅力

1. **Reactに近いテンプレート**: `<%= render :partial=>@list_item %>`や`<%= x %>`じゃなくて、`@ListItem()`と`{x}`で書ける
2. **Go に近い構文で書ける**: jsにjsxが追加された程度の学習で十分

### HTMX の魅力

1. **シンプルな Ajax**: HTML 属性だけで動的な機能を実現
2. **段階的導入**: 既存の HTML に少しずつ追加できる+AIが書いたことを調べればいいので学習コストも抑えられる

## 実際に動くサンプル：タスク管理アプリ

それでは、実際にtemplとhtmxを使ってタスク管理アプリを作ってみましょう。このサンプルを通じて、どれだけ簡単にフロントエンドが実装できるかを体感していただけます。

### 1. Templテンプレートの定義

まず、HTMLテンプレートをGoのコードとして定義します：

```go
package html

import "fmt"

type Task struct {
	ID   int
	Text string
	Done bool
}

templ TodoApp() {
	@Styles()
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
							class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all"
						/>
						<button
							type="submit"
							class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors font-medium"
						>
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

templ TaskList(tasks []Task) {
	for _, task := range tasks {
		@TaskItem(task)
	}
}

templ TaskItem(task Task) {
	{{
		containerClass := "task-item flex items-center p-4 bg-white border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow"
		if task.Done {
			containerClass += " opacity-60"
		}
	}}
	<div class={ containerClass }>
		<input
			type="checkbox"
			checked={ task.Done }
			hx-post={ fmt.Sprintf("/toggle-task/%d", task.ID) }
			hx-target="closest div"
			hx-swap="outerHTML"
			class="w-5 h-5 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 focus:ring-2"
		/>
		{{
		textClass := "flex-1 ml-3 text-gray-800 font-medium"
		if task.Done {
			textClass += " line-through text-gray-500"
		}
		}}
		<span class={ textClass }>
			{ task.Text }
		</span>
		<div class="flex gap-2 ml-4">
			{{
			buttonClass := "px-3 py-1 text-sm rounded-md font-medium transition-colors"
			if task.Done {
				buttonClass += " bg-yellow-100 text-yellow-800 hover:bg-yellow-200"
			} else {
				buttonClass += " bg-green-100 text-green-800 hover:bg-green-200"
			}
			}}
			<button
				class={ buttonClass }
				hx-post={ fmt.Sprintf("/toggle-task/%d", task.ID) }
				hx-target="closest .task-item"
				hx-swap="outerHTML"
			>
				{{
				body := "未完了"
				if task.Done {
					body = "完了"
				}
				}}
				{ body }
			</button>
			<button
				class="px-3 py-1 text-sm bg-red-100 text-red-800 rounded-md hover:bg-red-200 font-medium transition-colors"
				hx-delete={ fmt.Sprintf("/delete-task/%d", task.ID) }
				hx-target="closest .task-item"
				hx-swap="outerHTML"
				hx-confirm="このタスクを削除しますか？"
			>
				削除
			</button>
		</div>
	</div>
}

```

### 2. バックエンドの実装

次に、GoのバックエンドでHTMXのリクエストを処理します：

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"tentashiratori/templ-htmx/html"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

// 簡単なメモリ内タスクストレージ
var tasks = []html.Task{
	{ID: 1, Text: "Goの学習", Done: false},
	{ID: 2, Text: "Templの習得", Done: true},
	{ID: 3, Text: "HTMXの理解", Done: false},
}
var nextID = 4

func main() {
	// 静的ファイルの配信
	fs := http.FileServer(http.Dir("static/"))
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// ホームページ
	router.Handle("/", templ.Handler(html.Hello("World")))
	// タスク管理アプリ
	router.Handle("/todo", templ.Handler(html.TodoApp()))
	// HTMX API エンドポイント
	router.HandleFunc("/add-task", addTask)
	router.HandleFunc("/toggle-task/{id}", toggleTask)
	router.HandleFunc("/delete-task/{id}", deleteTask)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", router)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	taskText := r.FormValue("task")
	if taskText == "" {
		http.Error(w, "Task text is required", http.StatusBadRequest)
		return
	}
	
	newTask := html.Task{
		ID:   nextID,
		Text: taskText,
		Done: false,
	}
	nextID++
	tasks = append(tasks, newTask)
	
	// 新しいタスクアイテムを返す
	component := html.TaskItem(newTask)
	component.Render(r.Context(), w)
}

func toggleTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// URLからIDを抽出
	idStr := r.URL.Path[len("/toggle-task/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	
	// タスクを検索してトグル
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = !tasks[i].Done
			component := html.TaskItem(tasks[i])
			component.Render(r.Context(), w)
			return
		}
	}
	
	http.Error(w, "Task not found", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// URLからIDを抽出
	idStr := r.URL.Path[len("/delete-task/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	
	// タスクを削除
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	
	http.Error(w, "Task not found", http.StatusNotFound)
}

```

### 3. 細かいコードはサンプルリポジトリを参考にしてください

[サンプルリポジトリ](https://github.com/tentaShiratori/templ-htmx)

## なぜこれが素晴らしいのか？

### 1. JavaScriptを一切書かない

すべての動的な機能（タスクの追加、完了切り替え、削除）がHTMXの属性だけで実現されています：

- `hx-post="/add-task"` - フォーム送信時にPOSTリクエスト
- `hx-target="#task-list"` - レスポンスを特定の要素に挿入
- `hx-swap="beforeend"` - 要素の末尾に追加
- `hx-delete="/delete-task/1"` - 削除リクエスト
- `hx-confirm="このタスクを削除しますか？"` - 確認ダイアログ

### 2. cssを一切書かない

tailwindの設定方法は省きましたが、tailwindを使うことでcssも一行も書いていません。

### 3. 型安全性

TemplはGoの型システムを活用するため、エディタの拡張でテンプレートのエラーを検出できます。実行時エラーを大幅に減らせます。

### 4. 学習コストの低さ

HTMLとプログラミングの基本だけを知っていれば、すぐに開発を始められます。新しいフレームワーク特有の複雑な概念や独自記法を覚える必要がほとんどなく、公式ドキュメントやサンプルを見れば直感的に理解できます。  

## まとめ

Go + Templ + HTMXの組み合わせは、素振りに最適です。

この組み合わせなら、バックエンドの学習に集中しながら、モダンなWebアプリケーションを簡単に構築できます。