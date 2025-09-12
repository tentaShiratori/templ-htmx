package main

import (
	"fmt"
	"net/http"
	"strconv"
	"tentashiratori/templ-htmx/html"

	"github.com/a-h/templ"
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
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// ホームページ
	http.Handle("/", templ.Handler(html.Hello("World")))
	
	// カウンター
	http.Handle("/counter", templ.Handler(html.Counter(0)))
	
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
