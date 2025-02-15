package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ak-er/golang-playground/internal/schema"
	"github.com/ak-er/golang-playground/internal/storage"
	"github.com/ak-er/golang-playground/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func List(storage storage.Storager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := storage.GetStudentList()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralMessage(response.StatusError, err.Error(), nil))
			return
		}
		response.WriteJson(w, http.StatusOK, response.GeneralMessage(response.StatusError, "Successfully Fetched data", students))
		return
	}
}

func Get(storage storage.Storager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id from url
		requestId := r.PathValue("id")
		convId, err := strconv.ParseInt(requestId, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralMessage(response.StatusError, err.Error(), nil))
			return
		}
		student, err := storage.GetStudentById(convId)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralMessage(response.StatusError, err.Error(), nil))
			return
		}
		response.WriteJson(w, http.StatusOK, response.GeneralMessage(response.StatusOK, "Successfully Executed", student))
		return
	}
}

func Delete(storage storage.Storager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get id from url
		requestId := r.PathValue("id")
		convId, err := strconv.ParseInt(requestId, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralMessage(response.StatusError, err.Error(), nil))
			return
		}
		rowsAffected, err := storage.DeleteStudent(convId)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralMessage(response.StatusError, err.Error(), nil))
			return
		}
		message := fmt.Sprintf("Successfully %v items deleted", rowsAffected)
		response.WriteJson(w, http.StatusOK, response.GeneralMessage(response.StatusOK, message, nil))
		return
	}
}

func Create(storage storage.Storager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		student := schema.Student{}
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralMessage(response.StatusError, err.Error(), nil))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralMessage(response.StatusError, err.Error(), nil))
			return
		}

		// request validate
		if err := validator.New().Struct(student); err != nil {
			validateErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErr))
			return
		}

		lastInsertedId, err := storage.CreateStudent(student.Name, student.Course, student.City, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError,
				response.GeneralMessage(response.StatusError, err.Error(), nil))
			return
		}
		message := "successfully created"
		data := map[string]int64{"id": lastInsertedId}
		response.WriteJson(w, http.StatusCreated, response.GeneralMessage(response.StatusOK, message, data))
	}
}
