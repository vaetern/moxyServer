package communicationBodyService

import (
	"github.com/vaetern/moxyServer/servingStrategies/ComLog"
	"time"
	"log"
	"database/sql"
)

type StoreService struct {
	input     string
	output    string
	hashValue string
}

func (comServ StoreService) ProcessStoring(ch <-chan ComLog.CommunicationLog, database sql.DB) {
	go processStoringRoutine(ch, database)
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

type storageService struct {
}
