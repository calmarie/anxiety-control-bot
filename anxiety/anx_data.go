package anxiety

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var AnxietyStorage = make(map[int64][]AnxietyEntry)

func SaveAnxietyData(chatID int64, templevel int, tempcause string, longcause string) map[int64][]AnxietyEntry {
	entry := AnxietyEntry{
		Time:           string(time.Now().Format("02/01/2006 15:04")),
		Level:          templevel,
		ShortReason:    tempcause,
		DetailedReason: longcause,
	}

	AnxietyStorage[chatID] = append(AnxietyStorage[chatID], entry)

	return AnxietyStorage
}

func StartLoadAnxietyFromJSON(anxietyStorage *map[int64][]AnxietyEntry) {
	jsonData, err := os.ReadFile("AnxietyData.json")

	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	err = json.Unmarshal(jsonData, &anxietyStorage)
	if err != nil {
		fmt.Println("Ошибка анмаршалинга JSON:", err)
		return
	}

}

func SaveAnxietyToJSON(anxietyStorage *map[int64][]AnxietyEntry) {
	// Здесь реализация сохранения данных в JSON
	data, err := json.MarshalIndent(anxietyStorage, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("AnxietyData.json", data, 0644) // 0644 - права доступа (записываются в восьмиричной СС)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

}
