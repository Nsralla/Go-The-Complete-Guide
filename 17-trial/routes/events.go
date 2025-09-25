package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/app/db"
	"example.com/app/models"
	"github.com/gin-gonic/gin"
)

func handleGetEvents(context *gin.Context) {
	query := `
		SELECT * 
		FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println("Error while executing Select events query.", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error while fetching events. %v", err)})
		return 
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil{
			fmt.Println("Error scanning row:", err)
            continue // Skip this row and continue
		}
		events = append(events, event)
	}
	context.JSON(http.StatusOK, gin.H{"events":events})

}

func handleAddEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error while reading new event data. %v",err.Error()) })
		return
	}

	new_event, err := event.Create()
	if err!=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"event": new_event})
}


func handleGetEvent(context *gin.Context){
	id := context.Param("id")
	id_int,err := strconv.ParseInt(id,10,64)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Invalid ID"})
		return
	}
	if id_int == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"MESSAGE":"MISSING OR INVALID ID"})
		return
	}

	query := `
		SELECT *
		FROM events 
		where id = ?
	`

	row := db.DB.QueryRow(query, id_int)
	var event models.Event

	err = row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":fmt.Sprintf("Error while fetching an event. %v", err.Error() )})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event":event})


}