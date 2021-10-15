package db

func AddOne(history History) {
	Database.Create(history)
}
