func log(w http.ResponseWriter, r *http.Request) {
    FormEmail := r.FormValue("email")
    FormPassword := r.FormValue("password")
    // подключение
    db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:8889)/service")
    if err != nil {
        panic(err)
    }

    defer db.Close()

    // проверяем есть ли пользователь с заданным email + password
    res, err := db.Query("SELECT * FROM `users` WHERE `email` = '?' and password = '?'", FormEmail, FormPassword)
    if err != nil {
        fmt.Printf("%w", err)
        JSONError(500, "DB100 error code example", "DB error", w)
        return
    }

    // проверяем есть ли записи в БД.
    if !res.Next() {
        JSONError(413, "trololo", "user not found", w)
        return
    }

    // если дошли сюда, значит все ок, значит можно пропустить его дальше
    http.Redirect(w, r, "/", http.StatusSeeOther)
}


func JSONError(httpcode int, code, msg string, w http.ResponseWriter) {
    type Error struct {
        Code      *string `json:"code,omitempty"`
        Message   *string `json:"message,omitempty"`
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.WriteHeader(httpcode)
    json.NewEncoder(w).Encode(
        Error{
            Code:      &code,
            Message:   &msg,
        },
    )
}
