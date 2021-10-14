package web

import (
	"clothshare/db"
	"clothshare/web/api"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserCreate godoc
// @Summary creates a new user and inserts it into DB
//
// @Tags Users
//
// @Accept json
//
// @Param Product body api.JSONUserInput true "User data"
//
// @Success 200 "success"
// @Router /api/users [post]
func (app *Application) UserCreate(w http.ResponseWriter, r *http.Request) {
	var input api.JSONProductInput

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.log.WithError(err).Error("could not read body")
		JSON(w, http.StatusBadRequest, nil)
		return
	}

	if err := json.Unmarshal(data, &input); err != nil {
		app.log.WithError(err).Error("could not unmarshal input")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	err = app.db.Users.InsertOne(db.UserSchema{
		ID:   primitive.NewObjectID(),
		Name: input.Name,
	})
	if err != nil {
		app.log.WithError(err).Error("could not insert product")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}
}

// UserGet godoc
// @Summary retrieves user from db by id
//
// @Tags Users
//
// @Produce json
//
// @Param id path string true "id of the user"
//
// @Success 200 "success"
// @Router /api/users/{id} [get]
func (app *Application) UserGet(w http.ResponseWriter, r *http.Request) {
	id, found := mux.Vars(r)["id"]
	if !found {
		app.log.Error("could not parse id from request")
		JSON(w, http.StatusBadRequest, nil)
		return
	}

	pID, _ := primitive.ObjectIDFromHex(id)

	users, err := app.db.Users.Find(bson.M{"_id": pID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			app.log.WithError(err).Error("could not find product with this id")
			JSON(w, http.StatusNotFound, nil)
			return
		}
		app.log.WithError(err).Error("could not query database")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	JSON(w, http.StatusOK, api.JSONUserOutput{
		ID:   id,
		Name: users[0].Name,
	})
}

// UserGetList godoc
//
// @Summary returns all users stored in db
//
// @Tags Users
//
// @Produce json
//
// @Success 200 {array} []api.JSONUserOutput
// @Router /api/users [get]
func (app *Application) UserGetList(w http.ResponseWriter, r *http.Request) {
	users, err := app.db.Users.Find(bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			app.log.WithError(err).Error("could not any products")
			JSON(w, http.StatusNotFound, nil)
			return
		}
		app.log.WithError(err).Error("could not query database")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	output := []api.JSONUserOutput{}
	for _, v := range users {
		output = append(output, api.JSONUserOutput{
			ID:   v.ID.Hex(),
			Name: v.Name,
		})
	}

	JSON(w, http.StatusOK, output)
}

// UserDelete godoc
// @Summary deletes a user from db
//
// @Tags Users
//
// @Accept json
//
// @Param id path string true "id of the user"
//
// @Success 200 "success"
// @Router /api/users/{id} [delete]
func (app *Application) UserDelete(w http.ResponseWriter, r *http.Request) {
	id, found := mux.Vars(r)["id"]
	if !found {
		app.log.Error("could not parse id from request")
		JSON(w, http.StatusBadRequest, nil)
		return
	}

	pID, _ := primitive.ObjectIDFromHex(id)

	count, err := app.db.Users.Delete(bson.M{"_id": pID})
	if err != nil {
		app.log.WithError(err).Error("could not delete image from db")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	if count == 0 {
		app.log.Error("the product with specified id could not be found in db")
		JSON(w, http.StatusNotFound, nil)
		return
	}
}

// UserStoreImage godoc
// @Summary accepts an image and stores it in DB
//
// @Tags Users
//
// @Accept jpeg
// @Accept png
//
// @Param id path string true "id of the user"
//
// @Success 200 "success"
// @Router /api/users/image/{id} [post]
func (app *Application) UserStoreImage(w http.ResponseWriter, r *http.Request) {
	var (
		id, imgType string
		data        []byte
		err         error
	)

	id, found := mux.Vars(r)["id"]
	if !found {
		app.log.Error("could not prase id from request")
		JSON(w, http.StatusBadRequest, nil)
		return
	}
	app.log.Info("IMAGE RECEIVED")

	switch r.Header.Get("Content-Type") {
	case "image/jpeg":
		imgType = ".jpeg"
	case "image/png":
		imgType = ".png"
	default:
		app.log.Error("wrong content type in request")
		JSON(w, http.StatusBadRequest, nil)
		return
	}

	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		app.log.WithError(err).Error("could not parse request to image")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	err = ioutil.WriteFile(id+imgType, data, 0666)
	if err != nil {
		app.log.WithError(err).Error("could not write image to file")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	pID, _ := primitive.ObjectIDFromHex(id)
	err = app.db.Users.UpdateImgPath(pID, ImageFilePath)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			app.log.WithError(err).Error("the product with specified id could not be found in db")
			JSON(w, http.StatusNotFound, nil)
			return
		}
		app.log.WithError(err).Error("could not delete product in db")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}
}
