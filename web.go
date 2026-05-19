package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func startWebServer() {
	var tasks []Task

	r := gin.Default()

	tmpl := template.Must(template.New("index.html").Parse(htmlTemplate))
	r.SetHTMLTemplate(tmpl)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Tasks": tasks,
		})
	})

	r.POST("/add", func(c *gin.Context) {
		title := c.PostForm("title")
		desc := c.PostForm("description")

		if title != "" {
			newTask := Task{
				Title:       title,
				Description: desc,
				Done:        false,
				CreatedAt:   time.Now(),
			}
			tasks = append(tasks, newTask)
		}

		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/done/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id >= 0 && id < len(tasks) {
			tasks[id].Done = !tasks[id].Done
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/delete/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id >= 0 && id < len(tasks) {
			tasks = append(tasks[:id], tasks[id+1:]...)
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":8080")
}

var htmlTemplate = `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Трекер задач</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
            background: #1a1a2e;
            color: #eee;
            max-width: 600px;
            margin: 40px auto;
            padding: 20px;
        }
        h1 { text-align: center; color: #e94560; margin-bottom: 20px; }
        .add-form { display: flex; gap: 10px; margin-bottom: 20px; flex-wrap: wrap; }
        .add-form input {
            flex: 1; padding: 10px; border: none; border-radius: 6px;
            background: #16213e; color: #eee; font-size: 14px; min-width: 120px;
        }
        .add-form button {
            padding: 10px 20px; border: none; border-radius: 6px;
            background: #e94560; color: white; font-size: 14px; cursor: pointer; font-weight: bold;
        }
        .add-form button:hover { background: #c23152; }
        .task {
            background: #16213e; padding: 15px; border-radius: 8px;
            margin-bottom: 10px; display: flex; justify-content: space-between;
            align-items: center; flex-wrap: wrap; gap: 10px;
        }
        .task.done { opacity: 0.6; }
        .task.done .task-title { text-decoration: line-through; }
        .task-info { flex: 1; min-width: 150px; }
        .task-title { font-weight: bold; font-size: 16px; }
        .task-desc { font-size: 13px; color: #aaa; margin-top: 4px; }
        .task-actions { display: flex; gap: 8px; }
        .btn {
            padding: 6px 14px; border: none; border-radius: 4px;
            font-size: 13px; cursor: pointer; text-decoration: none; font-weight: bold;
        }
        .btn-done { background: #0f3460; color: #eee; }
        .btn-done:hover { background: #1a4a7a; }
        .btn-delete { background: #e94560; color: white; }
        .btn-delete:hover { background: #c23152; }
        .empty { text-align: center; color: #666; margin-top: 40px; font-size: 18px; }
    </style>
</head>
<body>
    <h1>📋 Трекер задач</h1>

    <form class="add-form" action="/add" method="POST">
        <input type="text" name="title" placeholder="Название задачи" required>
        <input type="text" name="description" placeholder="Описание (необязательно)">
        <button type="submit">➕ Добавить</button>
    </form>

    {{if .Tasks}}
        {{range $i, $task := .Tasks}}
        <div class="task {{if .Done}}done{{end}}">
            <div class="task-info">
                <div class="task-title">
                    {{if .Done}}✅{{else}}⬜{{end}} {{.Title}}
                </div>
                {{if .Description}}
                <div class="task-desc">{{.Description}}</div>
                {{end}}
            </div>
            <div class="task-actions">
                <a class="btn btn-done" href="/done/{{$i}}">
                    {{if .Done}}↩ Отменить{{else}}✓ Выполнить{{end}}
                </a>
                <a class="btn btn-delete" href="/delete/{{$i}}">🗑 Удалить</a>
            </div>
        </div>
        {{end}}
    {{else}}
        <div class="empty">📭 Список задач пуст</div>
    {{end}}
</body>
</html>
`
