package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper/at"
	"github.com/gocroot/helper/atdb"

	// "github.com/gocroot/helper/watoken"
	"github.com/gocroot/helper/watoken"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRegion(respw http.ResponseWriter, req *http.Request) {
	_, err := watoken.Decode(config.PublicKeyWhatsAuth, at.GetLoginFromHeader(req))
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Token Tidak Valid"
		respn.Location = "Decode Token Error: " + at.GetLoginFromHeader(req)
		respn.Response = err.Error()
		at.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}

	var longlat model.LongLat
	json.NewDecoder(req.Body).Decode(&longlat)

	filter := bson.M{
		"border": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{ // Menggunakan $geometry, bukan $border
					"type":        "Point",
					"coordinates": []float64{longlat.Longitude, longlat.Latitude},
				},
			},
		},
	}

	region, err := atdb.GetOneDoc[model.Region](config.MongoconnGeo, "region", filter)
	if err != nil {
		at.WriteJSON(respw, http.StatusNotFound, region)
		return
	}

	geoJSON := bson.M{
		"type": "FeatureCollection",
		"features": []bson.M{
			{
				"type": "Feature",
				"geometry": bson.M{
					"type":        region.Border.Type,
					"coordinates": region.Border.Coordinates,
				},
				"properties": bson.M{
					"province":     region.Province,
					"district":     region.District,
					"sub_district": region.SubDistrict,
					"village":      region.Village,
				},
			},
		},
	}

	// Kirim respon dalam format GeoJSON
	at.WriteJSON(respw, http.StatusOK, geoJSON)
}

func GetRoads(respw http.ResponseWriter, req *http.Request) {
	_, err := watoken.Decode(config.PublicKeyWhatsAuth, at.GetLoginFromHeader(req))

	if err != nil {
		_, err = watoken.Decode(config.PUBLICKEY, at.GetLoginFromHeader(req))

		if err != nil {
			var respn model.Response
			respn.Status = "Error: Token Tidak Valid"
			respn.Info = at.GetSecretFromHeader(req)
			respn.Location = "Decode Token Error"
			respn.Response = err.Error()
			at.WriteJSON(respw, http.StatusForbidden, respn)
			return
		}
	}

	var longlat model.LongLat
	json.NewDecoder(req.Body).Decode(&longlat)

	filter := bson.M{
		"geometry": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longlat.Longitude, longlat.Latitude},
				},
				"$maxDistance": longlat.MaxDistance,
			},
		},
	}

	roads, err := atdb.GetAllDoc[[]model.Roads](config.MongoconnGeo, "roads", filter)
	if err != nil {
		at.WriteJSON(respw, http.StatusNotFound, roads)
		return
	}
	at.WriteJSON(respw, http.StatusOK, roads)
}
