package cmdUtils

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

var file *os.File

// New method creates a new task with the given title and status.
func (t *Task) New() error {
	defer file.Close()

	id, err := getNewId()
	if err != nil {
		return err
	}

	createdAt := time.Now().Format(time.RFC3339Nano)

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
	if len(record) == 0 {
		return 1, nil
	}

	id, err := strconv.Atoi(record[len(record)-1][0])
	if err != nil {
		return 0, err
	}

	return id + 1, nil
}

// List method returns the list of tasks based on the given flag.
func (t *Task) List(statusFlag string) ([]Task, error) {
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	for _, record := range records {
		if statusFlag == "all" || strings.ToLower(record[2]) == statusFlag {
			id, err := strconv.Atoi(record[0])
			if err != nil {
				return nil, err
			}
			tasks = append(tasks, Task{
				ID:        id,
				Title:     record[1],
				Status:    record[2],
				CreatedAt: record[3],
			})
		}
	}

	return tasks, nil
}

// Remove method removes the task with the given id.
func (t *Task) Remove(id string) error {
	defer file.Close()

	// need to read the file before opening it with truncating flag
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// open the csv file with required permissions
	_file, err := os.OpenFile("./data/tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer _file.Close()

	writer := csv.NewWriter(_file)
	defer writer.Flush()

	for _, record := range records {
		if record[0] == id {
			continue
		}

		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}

// Edit method edits the task with the given id
func (t *Task) Edit(id string, title string, statusFlag string) error {
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// open the csv file with required permissions
	_file, err := os.OpenFile("./data/tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer _file.Close()

	writer := csv.NewWriter(_file)
	defer writer.Flush()

	found := false

	newTitle := title
	newStatusFlag := statusFlag
	for _, record := range records {
		// if id is not found, append all the old values from record and continue
		if record[0] != id {
			writer.Write([]string{
				record[0],
				record[1],
				record[2],
				record[3],
			})
			continue
		}

		found = true

		// if title is empty, set new title to the old one from record
		if title == "" {
			newTitle = record[1]
		}

		// if flag is empty, set new flag to the old one from record
		if statusFlag == "" {
			newStatusFlag = record[2]
		}

		writer.Write([]string{
			id,
			newTitle,
			newStatusFlag,
			record[3],
		})
	}

	if !found {
		return FlagErrorf("Couldn't find a task with the given id")
	}

	return nil
}

func (t *Task) Prune(statusFlag string) error {
	defer file.Close()

	// need to read the file before opening it with truncating flag
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// open the csv file with required permissions
	_file, err := os.OpenFile("./data/tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer _file.Close()

	writer := csv.NewWriter(_file)
	defer writer.Flush()

	for _, record := range records {
		if strings.ToLower(record[2]) == statusFlag {
			continue
		}

		writer.Write(record)
	}

	return nil
}

func init() {
	// open the csv file with required permissions
	_file, err := os.OpenFile("./data/tasks.csv", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	file = _file
}
