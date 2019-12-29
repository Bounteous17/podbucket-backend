package user

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/bounteous/podbucket-backend/podbucket/clioutput"

	"github.com/bounteous/podbucket-backend/podbucket/coredb"
	"github.com/bounteous/podbucket-backend/podbucket/httpod"
)

type mongoUserID struct {
	ID string `json:"id"`
}

// TODO separate into multiple functions
func insertOne(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		body, errMarsh := httpod.NewError(&httpod.PodError{
			Error: httpod.INVALID_BODY,
		})
		if body != nil {
			httpod.StatusBadRequest(res, body)
		} else {
			clioutput.Fatal(errMarsh)
		}
	} else {
		// TODO handle _ error
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
		coreDbClient := coredb.Client()

		result, insertErr := coreDbClient.Database(database).Collection(collection).InsertOne(ctx, user)

		if insertErr != nil {
			body, errMarsh := httpod.NewError(&httpod.PodError{
				Error: insertErr.Error(),
			})
			if body != nil {
				httpod.StatusInternalServerError(res, body)
			} else {
				clioutput.Fatal(errMarsh)
			}
		} else {
			body, err := json.Marshal(&mongoUserID{
				ID: result.InsertedID.(primitive.ObjectID).Hex(),
			})

			if err != nil {
				body, errMash := httpod.NewError(&httpod.PodError{
					Error: httpod.DECODE_FAIL,
				})
				if body != nil {
					httpod.StatusInternalServerError(res, body)
				} else {
					clioutput.Fatal(errMash)
				}
			} else {
				httpod.StatusOK(res, body)
			}
		}
	}
}
