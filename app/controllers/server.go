package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"github.com/aihou/todo/config"
	"github.com/aihou/todo/app/models"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	//template.Mustは、第一引数にテンプレートのポインタを、第二引数にエラーを返す
	templates := template.Must(template.ParseFiles(files...))
	//ExecuteTemplateは、第一引数に書き込み先のio.Writerを、第二引数にテンプレート名を、第三引数にデータを渡す
	templates.ExecuteTemplate(w, "layout", data)
}

//クッキーからセッションを取得する
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	//クッキーの取得
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		//クッキーがあれば、クッキーの値を元にセッションを取得
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	//第一引数にレスポンスライター、第二引数にリクエスト、第三引数にURLのパスを受け取る
	return func(w http.ResponseWriter, r *http.Request) {
		//URLのパスを取得
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		//URLのパスを取得
		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	return http.ListenAndServe(":"+config.Config.Port, nil)
}