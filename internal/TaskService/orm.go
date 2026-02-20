package TaskService

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Task   string ` json:"task"`
	Status string ` json:"status"`
}
