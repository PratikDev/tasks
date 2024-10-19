package cmdUtils

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

var file *os.File

// New method creates a new task with the given title and status.
func (t *Task) New() error {
	defer file.Close()

	id, err := getNewId()
	if err != nil {
		return err
	}

	createdAt := time.Now().Format(time.DateTime)

	record := []string{
		strconv.Itoa(id),
		t.Title,
		t.Status,
		createdAt,
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(record)
	if err != nil {
		return err
	}

	return nil
}

// getNewId function reads the last record from the csv file and returns the new id by incrementing the last id by 1.
func getNewId() (int, error) {
	reader := csv.NewReader(file)
	record, _ := reader.ReadAll()
	if len(record) <= 1 {
		return 1, nil
	}

	id, err := strconv.ParseInt(record[len(record)-1][0], 10, 64)
	if err != nil {
		return 0, err
	}

	return int(id) + 1, nil
}

func init() {
	// open the csv file with required permissions
	_file, err := os.OpenFile("./data/tasks.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	file = _file
}
