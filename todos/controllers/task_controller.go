package controllers

import (
    "todos/models"
    "database/sql"
    "net/http"
)

type TaskHandler struct {
    Db *sql.DB
}

// 一覧
func (handler *TaskHandler) Index(w http.ResponseWriter, r *http.Request) {
    var tasks []models.Task                                    // データ一覧を格納するため、Task構造体のスライスを変数宣言
    getall := handler.Db.Find(&tasks)                          // DBから全てのデータを取得する
    w.HTML(http.StatusOK, "index.html", {"tasks": getall})     // index.htmlに全てのデータを渡す
}


// 新規
func (handler *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
    text, _ := r.GetPostForm("text")            // index.htmlからtextを取得
    handler.Db.Create(&models.Task{Text: text}) // レコードを挿入する
    http.Redirect(http.StatusMovedPermanently, "/")
}


// 更新
func (handler *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
    task := models.Task{}            // Task構造体の変数宣言
    id := r.Param("id")              // edit.htmlからidを取得
    text, _ := r.GetPostForm("text") // edit.htmlからtextを取得
    handler.Db.First(&task, id)      // idに一致するレコードを取得する
    task.Text = text                 // textを上書きする
    handler.Db.Save(&task)           // 指定のレコードを更新する
    http.Redirect(http.StatusMovedPermanently, "/")
}

// 削除
func (handler *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
    task := models.Task{}       // Task構造体の変数宣言
    id := r.Param("id")         // index.htmlからidを取得
    handler.Db.First(&task, id) // idに一致するレコードを取得する
    handler.Db.Delete(&task)    // 指定のレコードを削除する
    http.Redirect(http.StatusMovedPermanently, "/")
}