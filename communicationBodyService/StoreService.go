package communicationBodyService

import (
	"github.com/vaetern/moxyServer/servingStrategies/ComLog"
	"time"
	"log"
	"database/sql"
	"os"
	"errors"
)

type StoreService struct {
	db *sql.DB
}

func (ss StoreService) ProcessStoring(ch <-chan ComLog.CommunicationLog) {
	go processStoringRoutine(ch, *ss.db)
}

func NewStoreService() (ss *StoreService) {
	dataSourceName := "./local.db"
	database, err := sql.Open("sqlite3", dataSourceName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ss = &StoreService{db: database}

	return ss
}

func processStoringRoutine(ch <-chan ComLog.CommunicationLog, database sql.DB) {
	log.Println("-> ready to serve")
	for commLog := range ch {

		log.Println(commLog.Target)

		_, err := database.Exec("INSERT OR IGNORE INTO `communication_log`"+
			" (`target`, `responseKey`, `responseBody`, `date`) VALUES($1, $2, $3, $4);",
			commLog.Target, commLog.ResponseKey, commLog.ResponseBody, time.Now())
		if err != nil {
			log.Println("Error writing to db:")
			log.Println(err)
		}
	}
}

func (ss *StoreService) GetBodyByKeyAndTarget(key string, target string) (responseBody string, err error) {

	responseBody = ""

	rows, err := ss.db.Query("SELECT `responseBody` FROM `communication_log` "+
		"WHERE `responseKey` = $1 AND `target`= $2 LIMIT 0,1",
		key, target)

	i := 0
	for rows.Next() {
		i++
		rows.Scan(&responseBody)
	}

	if i == 0{
		err = errors.New("X - Response not found")
	}

	return responseBody, err
}
