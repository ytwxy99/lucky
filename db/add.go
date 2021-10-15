package db

func AddOne(history History) {
	db, err := GetDB()
	if err != nil {
		Log.Error("Open database error:", err)
	}
	db.Create(history)
}
