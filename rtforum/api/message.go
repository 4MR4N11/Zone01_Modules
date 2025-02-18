package api

import (
	"database/sql"
	"forum/config"
	"forum/utils"
	"net/http"
	"strconv"
)

type Msg struct {
	MsgId int `json:"msg_id"`
	SenderId int `json:"sender_id"`
	ReceiverId int `json:"receiver_id"`
	Data string `json:"data"`
	Timestamp string `json:"timestamp"`
}

func MessageApi(w http.ResponseWriter, r *http.Request) {
	sessionID := utils.GetSessionCookie(r)
	session, err := config.SESSION.GetSession(sessionID)
	if err != nil {
		http.Error(w, "Session error", http.StatusInternalServerError)
		return
	}
	if session == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	senderID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	UserId := session.UserId
	if UserId != int64(senderID) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	messageID := r.URL.Query().Get("msg_id")
	var row *sql.Rows
	if messageID == "" {
		row, err = config.DB.Query(`SELECT * FROM messages WHERE (senderId = ? AND receiverId = ?) 
								  OR (senderId = ? AND receiverId = ?) 
								  ORDER BY timestamp DESC LIMIT 10`, senderID, UserId, UserId, senderID)
	} else {
		row, err = config.DB.Query(`SELECT * FROM messages WHERE (senderId = ? AND receiverId = ?) 
								  OR (senderId = ? AND receiverId = ?) 
								  AND msgId < ? ORDER BY timestamp DESC LIMIT 10`, senderID, messageID)
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer row.Close()
	var messages []Msg
	for row.Next() {
		var msg Msg
		err = row.Scan(&msg.MsgId, &msg.SenderId, &msg.ReceiverId, &msg.Data, &msg.Timestamp)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		messages = append(messages, msg)
	}
	if err = row.Err(); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "Success", messages)
}