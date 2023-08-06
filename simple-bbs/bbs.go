package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"time"
)

const logFile = "logs.json"

type Log struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Body  string `json:"body"`
	CTime int64  `json:"ctime"`
}

func loadLogs() []Log {
	text, err := os.ReadFile(logFile)
	if err != nil {
		return make([]Log, 0)
	}
	var logs []Log
	json.Unmarshal([]byte(text), &logs)

	return logs
}

func getFormHtml() string {
	return "<div><form action='/write' method='POST'>" +
		"名前: <input type='text' name='name'><br>" +
		"本文: <input type='text' name='body' style='width:30em;'><br>" +
		"<button type='submit'>書き込み</button>" +
		"</form></div><hr>"
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	logs := loadLogs()
	logHtml := ""
	for i := len(logs) - 1; i >= 0; i-- {
		log := logs[i]
		logHtml += fmt.Sprintf(
			"<p>(%d) <span>%s</span>: %s --- %s</p>",
			log.ID,
			html.EscapeString(log.Name),
			html.EscapeString(log.Body),
			time.Unix(log.CTime, 0).Format("2006/1/2 15:04"),
		)
	}
	html := "<html><head>" +
		"<style> p { border: 1px solid silver; padding: 1em; } span { background-color: #eef; }</style>" +
		"</head><body><h1>Simple BBS</h1>" + getFormHtml() + logHtml + "</body></html>"
	w.Write([]byte(html))
}

func writeLogs(logs []Log) {
	bytes, _ := json.Marshal(logs)
	os.WriteFile(logFile, bytes, 0644)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.Form["name"][0]
	if name == "" {
		name = "名無し"
	}

	logs := loadLogs()
	log := Log{
		ID:    len(logs) + 1,
		Name:  name,
		Body:  r.Form["body"][0],
		CTime: time.Now().Unix(),
	}
	logs = append(logs, log)
	writeLogs(logs)

	http.Redirect(w, r, "/", 302)
}

func main() {
	println("server - http://localhost:8888")

	http.HandleFunc("/", showHandler)
	http.HandleFunc("/write", writeHandler)
	http.ListenAndServe(":8888", nil)
}
