Golangを練習したくてアプリケーションを作りました。

Todoアプリケーション

ユーザーのログイン，ログアウト，ユーザー削除など
TodoのCREATE,DELETE,EIDT

/main.go:サーバーを起動する
/config.ini:設定ファイル
/app/models/base.go:データベースの作成
/app/models/users.go:userのcreate,deleteなどの設計
/app/models/todos.go:todoのcreate,deleteなどの設計
/app/controllers/server.go:サーバー側
/app/controllers/route_main.go:todo側のハンドル関数を記入
/app/controllers/route_auth.go:user側のハンドル関数を記入
/app/utils/logging.go:ログの設定

