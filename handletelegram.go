package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/NicoNex/echotron/v3"
	// "time"
)

const MaxUploadSize = 50 << 20 // 50 MB

type ResponseParam struct {
	Url string `json:"response"`
}

func (botcfg *BotApiConfig) HandleTelegramUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(MaxUploadSize)
	botcfg.Bot.Debug =  true
	if err != nil {
		RespondWithError(w, http.StatusRequestEntityTooLarge, err.Error())
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("error with file %v", err))
		return
	}
	defer file.Close()
	if fileHeader.Size > MaxUploadSize {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("File %s is too big. Max size: %d bytes", fileHeader.Filename, MaxUploadSize))
		return
	}
	//send message using uploader bot
	fileBytes, err  := io.ReadAll(file)
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error converting file to bytes %v", err))
		return
	}

	fileUpload := echotron.NewInputFileBytes(fileHeader.Filename, fileBytes)
	
	documentMsg, err := botcfg.BotUploader.SendDocument(fileUpload, botcfg.ChannelID, &echotron.DocumentOptions{})
	if err != nil{
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error sending document with bot uploader %v", err))
		return
	}

	var fileID string
	if documentMsg.Result.Document != nil {
		fileID = documentMsg.Result.Document.FileID
	} else if documentMsg.Result.Video != nil {
		fileID = documentMsg.Result.Video.FileID
	} else if documentMsg.Result.Audio != nil {
		fileID = documentMsg.Result.Audio.FileID
	} else if len(documentMsg.Result.Photo) > 0 {
		fileID = documentMsg.Result.Photo[len(documentMsg.Result.Photo)-1].FileID
	} else if documentMsg.Result.Voice != nil {
		fileID = documentMsg.Result.Voice.FileID
	} else if documentMsg.Result.VideoNote != nil {
		fileID = documentMsg.Result.VideoNote.FileID
	} else if documentMsg.Result.Sticker != nil {
		fileID = documentMsg.Result.Sticker.FileID
	} else if documentMsg.Result.Animation != nil {
		fileID = documentMsg.Result.Animation.FileID
	} else {
		RespondWithError(w, http.StatusInternalServerError, "Unknown file type")
		return
	}

    
	fileURL, err := botcfg.Bot.GetFileDirectURL(fileID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("error getting file URL %v", err))
		return
	}

	res := ResponseParam{
		Url: fileURL,
	}
	RespondWithJSON(w, http.StatusOK, res)
}