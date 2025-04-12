package internal

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

func LoadFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
}
func ReadCsvFile(file *os.File) (*[]Task, error) {

	csvReader := csv.NewReader(bufio.NewReader(file))
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		deleted, _ := strconv.ParseBool(record[4])
		task := Task{
			Id:          id,
			Description: record[1],
			CreatedAt:   record[2],
			Status:      record[3],
			Deleted:     deleted,
		}

		tasks = append(tasks, task)
	}
	return &tasks, nil
}

func AppendToCsvFile(file *os.File, data []string) error {

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	defer csvWriter.Flush()

	if err := csvWriter.Write(data); err != nil {
		return err
	}
	return nil
}

func WriteCsvFile(file *os.File, tasks []Task) error {

	csvWriter := csv.NewWriter(bufio.NewWriter(file))
	defer csvWriter.Flush()

	for _, task := range tasks {
		file.Seek(0, 0)
		if err := csvWriter.Write(task.ToCSVFormat()); err != nil {
			return err
		}
	}

	return nil
}
