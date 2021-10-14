package web

import (
	"clothshare/db"
	"clothshare/web/api"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductCreate godoc
// @Summary creates a new product and inserts it into DB
//
// @Tags Products
//
// @Accept json
// @Produce json
//
// @Param Product body api.JSONProductInput true "Product data"
//
// @Success 200 "success"
// @Router /api/products [post]
func (app *Application) ProductCreate(w http.ResponseWriter, r *http.Request) {
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

	err = app.db.Products.InsertOne(db.ProductSchema{
		ID:   primitive.NewObjectID(),
		Name: input.Name,
	})
	if err != nil {
		app.log.WithError(err).Error("could not insert product")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}
}

// ProductGet godoc
//
// @Summary returns product by the specified id
//
// @Tags Products
//
// @Produce json
//
// @Param id path string true "id of the product"
//
// @Success 200 {object} api.JSONProductOutput
// @Router /api/products/{id} [get]
func (app *Application) ProductGet(w http.ResponseWriter, r *http.Request) {
	id, found := mux.Vars(r)["id"]
	if !found {
		app.log.Error("could not parse id from request")
		JSON(w, http.StatusBadRequest, nil)
		return
	}

	pID, _ := primitive.ObjectIDFromHex(id)

	products, err := app.db.Products.Find(bson.M{"_id": pID})
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

	productImage := "data:image/jpeg;base64,"
	productFile, err := ioutil.ReadFile("../../images/" + products[0].ID.Hex() + ".png")
	if err != nil {
		app.log.WithError(err).Error("could not read image from db")
		// IMAGE(w, http.StatusInternalServerError, nil)
		// return
	}

	productImage = productImage + base64.StdEncoding.EncodeToString(productFile)

	JSON(w, http.StatusOK, api.JSONProductOutput{
		ID:          products[0].ID.Hex(),
		Name:        products[0].Name,
		Description: products[0].Description,
		Condition:   products[0].Condition,
		Size:        products[0].Size,
		Color:       products[0].Color,
		ViewCount:   products[0].ViewCount,
		Brand:       products[0].Brand,
		Category:    products[0].Category,
		Location:    products[0].Location,
		Price:       products[0].Price,
		Tags:        products[0].Tags,
		CreatedAt:   products[0].CreatedAt,
		Author:      products[0].Author.Hex(),
		Image:       productImage,
	})
}

// ProductGetList ProductGetListByUser
//
// @Summary returns products stored in db by user
//
// @Tags Products
//
// @Produce json
//
// @Success 200 {array} []api.JSONProductOutput
// @Router /api/products/user/{id} [get]
func (app *Application) ProductGetListByUser(w http.ResponseWriter, r *http.Request) {
	id, found := mux.Vars(r)["id"]
	if !found {
		app.log.Error("could not parse id from request")
		JSON(w, http.StatusBadRequest, nil)
		return
	}

	pID, _ := primitive.ObjectIDFromHex(id)

	products, err := app.db.Products.Find(bson.M{"author": pID})
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

	output := []api.JSONProductOutput{}
	for _, v := range products {
		productImage := "data:image/jpeg;base64,"
		productFile, err := ioutil.ReadFile("../../images/" + v.ID.Hex() + ".png")
		if err != nil {
			app.log.WithError(err).Error("could not read image from db")
			// IMAGE(w, http.StatusInternalServerError, nil)
			// return
		}

		productImage = productImage + base64.StdEncoding.EncodeToString(productFile)
		output = append(output, api.JSONProductOutput{
			ID:          v.ID.Hex(),
			Name:        v.Name,
			Description: v.Description,
			Condition:   v.Condition,
			Size:        v.Size,
			Color:       v.Color,
			ViewCount:   v.ViewCount,
			Brand:       v.Brand,
			Category:    v.Category,
			Location:    v.Location,
			Price:       v.Price,
			Tags:        v.Tags,
			CreatedAt:   v.CreatedAt,
			Author:      v.Author.Hex(),
			Image:       productImage,
		})
	}

	JSON(w, http.StatusOK, output)
}

// ProductGetList godoc
//
// @Summary returns all products stored in db
//
// @Tags Products
//
// @Produce json
//
// @Success 200 {array} []api.JSONProductOutput
// @Router /api/products [get]
func (app *Application) ProductGetList(w http.ResponseWriter, r *http.Request) {
	products, err := app.db.Products.Find(bson.M{})
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

	output := []api.JSONProductOutput{}
	for _, v := range products {
		productImage := "data:image/jpeg;base64,"
		productFile, err := ioutil.ReadFile("../../images/" + v.ID.Hex() + ".png")
		if err != nil {
			app.log.WithError(err).Error("could not read image from db")
			// IMAGE(w, http.StatusInternalServerError, nil)
			// return
		}

		productImage = productImage + base64.StdEncoding.EncodeToString(productFile)
		output = append(output, api.JSONProductOutput{
			ID:          v.ID.Hex(),
			Name:        v.Name,
			Description: v.Description,
			Condition:   v.Condition,
			Size:        v.Size,
			Color:       v.Color,
			ViewCount:   v.ViewCount,
			Brand:       v.Brand,
			Category:    v.Category,
			Location:    v.Location,
			Price:       v.Price,
			Tags:        v.Tags,
			CreatedAt:   v.CreatedAt,
			Author:      v.Author.Hex(),
			Image:       productImage,
		})
	}

	JSON(w, http.StatusOK, output)
}

// ProductDelete godoc
// @Summary deletes a product from db
//
// @Tags Products
//
// @Accept json
//
// @Param id path string true "id of the product"
//
// @Success 200 "success"
// @Router /api/products/{id} [delete]
func (app *Application) ProductDelete(w http.ResponseWriter, r *http.Request) {
	id, found := mux.Vars(r)["id"]
	if !found {
		app.log.Error("could not parse id from request")
		JSON(w, http.StatusBadRequest, nil)
		return
	}

	pID, _ := primitive.ObjectIDFromHex(id)

	count, err := app.db.Products.Delete(bson.M{"_id": pID})
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

// ProductStoreImage godoc
// @Summary accepts an image and stores it in DB
//
// @Tags Products
//
// @Accept jpeg
// @Accept png
//
// @Param id path string true "id of the product"
//
// @Success 200 "success"
// @Router /api/products/{id}/image [post]
func (app *Application) ProductStoreImage(w http.ResponseWriter, r *http.Request) {
	var (
		id  string
		err error
	)

	id, found := mux.Vars(r)["id"]
	if !found {
		app.log.Error("could not prase id from request")
		JSON(w, http.StatusBadRequest, nil)
		return
	}

	// switch r.Header.Get("Content-Type") {
	// case "image/jpeg":
	// 	imgType = ".jpeg"
	// case "image/png":
	// 	imgType = ".png"
	// default:
	// 	app.log.Error("wrong content type in request")
	// 	JSON(w, http.StatusBadRequest, nil)
	// 	return
	// }

	_, err = app.db.Products.Find(bson.M{"_id": id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			app.log.WithError(err).Error("specified product was not found")
			JSON(w, http.StatusInternalServerError, nil)
			return
		}
		app.log.WithError(err).Error("could not extract product from db")
		JSON(w, http.StatusInternalServerError, nil)
		return
	}

	// lol := struct {
	// 	File interface{} `json:"file"`
	// }{}

	// err = json.Unmarshal(data, &lol.File)
	// if err != nil {
	// 	app.log.WithError(err).Error("could not unmarshal")
	// 	JSON(w, http.StatusInternalServerError, nil)
	// 	return
	// }

	r.ParseMultipartForm(100)
	mForm := r.MultipartForm
	fmt.Println(len(mForm.File))
	for k, _ := range mForm.File {
		// k is the key of file part
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Println("inovke FormFile error:", err)
			return
		}
		defer file.Close()
		fmt.Printf("the uploaded file: name[%s], size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		// store uploaded file into local path
		localFileName := "../../images/" + id + ".png"
		out, err := os.Create(localFileName)
		if err != nil {
			fmt.Printf("failed to open the file %s for writing", localFileName)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Printf("copy file err:%s\n", err)
			return
		}
		fmt.Printf("file %s uploaded ok\n", fileHeader.Filename)
	}
	pID, _ := primitive.ObjectIDFromHex(id)
	err = app.db.Products.UpdateImgPath(pID, ImageFilePath)
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

// ProductGetImage godoc
// @Summary accepts an image and stores it in DB
//
// @Tags Products
//
// @Produces jpeg
// @Produces png
//
// @Param id path string true "id of the product"
//
// @Success 200 "success"
// @Router /api/products/{id}/image [get]
func (app *Application) ProductGetImage(w http.ResponseWriter, r *http.Request) {
	var (
		payload []byte
		err     error
	)

	id, found := mux.Vars(r)["id"]
	if !found {
		app.log.Error("could not prase id from request")
		IMAGE(w, http.StatusBadRequest, nil)
		return
	}

	_, err = app.db.Products.Find(bson.M{"_id": id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			app.log.WithError(err).Error("specified product was not found")
			IMAGE(w, http.StatusInternalServerError, nil)
			return
		}
		app.log.WithError(err).Error("could not extract product from db")
		IMAGE(w, http.StatusInternalServerError, nil)
		return
	}

	payload, err = ioutil.ReadFile("/Users/dominykasjuscius/Projektai/clothshare/server/images/615390048d7c1e97caabc0c0")
	if err != nil {
		app.log.WithError(err).Error("could not read image from db")
		IMAGE(w, http.StatusInternalServerError, nil)
		return
	}

	IMAGE(w, http.StatusOK, payload)
}
