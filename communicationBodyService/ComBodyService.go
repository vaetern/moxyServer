package communicationBodyService

import (
	"github.com/vaetern/moxyServer/servingStrategies/ComLog"
	"time"
	"log"
	"database/sql"
	"os"
)

type StoreService struct {
	db *sql.DB
}

func (ss StoreService) ProcessStoring(ch <-chan ComLog.CommunicationLog) {
	go processStoringRoutine(ch, *ss.db)
}

func NewStoreService()(ss *StoreService){
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
	log.Println("Routine start")
	for commLog := range ch {

		log.Println("<-ch")
		_, err := database.Exec("INSERT INTO `communication_log`"+
			" (`target`, `responseKey`, `responseBody`, `date`) VALUES($1, $2, $3, $4)",
			commLog.Target, commLog.ResponseKey, commLog.ResponseBody, time.Now())
		if err != nil {
			log.Println("Error writing to db")
		}
	}
}