package public_api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChouYunShuo/oapi-demo/idm"

	"github.com/google/uuid"
)

type Service struct {
	Queries *idm.Queries
}

type IdmStore struct {
	IdmService *Service
}

var _ ServerInterface = (*IdmStore)(nil)

func sendIdmStoreError(w http.ResponseWriter, code int, message string) {
	idmErr := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(idmErr)
}

func (h *IdmStore) GetPublic(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("Public route!!")
}

// curl --location -X GET 'localhost:4000/api/get?username=torpago_usr2'
func (h *IdmStore) GetUser(w http.ResponseWriter, r *http.Request, params GetUserParams) {
	user, err := h.IdmService.Queries.FindUserByUsername(context.Background(), params.Username)

	if err != nil {
		sendIdmStoreError(w, http.StatusNotFound, fmt.Sprintf("Could not find User with Username %s", params.Username))
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

/*
	curl --location -X POST 'localhost:4000/api/post' \
	    --header 'Content-Type: application/json' \
	    --data-raw '{
	        "username": "elliot123",
	        "password": "pwd",
			"firstname": "Jessie",
			"lastname": "Lee"
	    }'
*/
func (h *IdmStore) PostUser(w http.ResponseWriter, r *http.Request) {

	var param PostUser
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		sendIdmStoreError(w, http.StatusBadRequest, "Invalid format for new User")
		return
	}

	var lastName sql.NullString
	if param.Lastname != nil {
		lastName = sql.NullString{String: *param.Lastname, Valid: true}
	} else {
		lastName = sql.NullString{Valid: false}
	}

	newUser := idm.CreateUserParams{
		Username:  param.Username,
		Password:  []byte(param.Password),
		FirstName: param.Firstname,
		LastName:  lastName,
	}

	err := h.IdmService.Queries.CreateUser(context.Background(), newUser)

	if err != nil {
		sendIdmStoreError(w, http.StatusNotFound, "Could not create new entry")
		return
	}

	// Respond with the created user data.
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

/*
 curl --location -X PUT 'localhost:4000/api/put' \
    -d "userid=f68e2ce3-bd5f-4172-99ce-1f35ff378c66&username=erichaha&password=tutorial123&firstname=Eric"
*/

func (h *IdmStore) PutUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		sendIdmStoreError(w, http.StatusBadRequest, "Failed to parse form data")
		return
	}
	userID, err := uuid.Parse(r.FormValue("userid"))

	if err != nil {
		sendIdmStoreError(w, http.StatusBadRequest, "Failed to load userID")
		return
	}

	var lastName sql.NullString
	if r.FormValue("lastname") != "" {
		lastName = sql.NullString{String: r.FormValue("lastname"), Valid: true}
	} else {
		lastName = sql.NullString{Valid: false}
	}
	updateUser := idm.UpdateUserParams{
		Uuid:      userID,
		Username:  r.FormValue("username"),
		Password:  []byte(r.FormValue("password")),
		FirstName: r.FormValue("firstname"),
		LastName:  lastName,
	}
	fmt.Printf("%+v\n", updateUser)

	err = h.IdmService.Queries.UpdateUser(context.Background(), updateUser)

	if err != nil {
		sendIdmStoreError(w, http.StatusBadGateway, "Failed to save user")
		return
	}

	// Respond with the created user data.
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(updateUser)
}

func (h *IdmStore) DeleteUser(w http.ResponseWriter, r *http.Request, params DeleteUserParams) {
	err := h.IdmService.Queries.DeleteUserByUsername(context.Background(), params.Username)

	if err != nil {
		sendIdmStoreError(w, http.StatusNotFound, fmt.Sprintf("Could not delete user with name %s", params.Username))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
