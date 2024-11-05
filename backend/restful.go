package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack/v5"
)

var newUUID = uuid.New()

var apis = []func(r *gin.Engine){
	registerSystemJson,
	registerListPartyQos,
	registerV1Api,
	registerAuth,
	registerOthers,
}

func RegisterHandler() *gin.Engine {
	r := gin.Default()
	for _, api := range apis {
		api(r)
	}
	return r
}

func registerSystemJson(r *gin.Engine) {
	r.GET("/systems/EAR-B-WW/00001/system.json", func(c *gin.Context) {
		m, err := filenameToMap("system.json")
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
		c.JSON(200, m)
	})
}

func registerListPartyQos(r *gin.Engine) {
	r.POST("/v1/steam-steam/sign/EAR-B-WW", func(c *gin.Context) {
		m, err := filenameToMap("steam_sign_ear-b-ww.json")
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
		c.JSON(200, m)
	})
	r.POST("/MultiplayerServer/ListPartyQosServers", func(c *gin.Context) {
		m, err := filenameToMap("list_party_qos_servers.json")
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
		c.JSON(200, m)
	})
}

func registerV1Api(r *gin.Engine) {
	g := r.Group("/v1")
	g.GET("/consent/restrictions/:country_code", func(c *gin.Context) {
		m, err := filenameToMap("restrictions.json")
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
		c.JSON(200, m)
	})
	g.GET("/consent/countries/:country_code", func(c *gin.Context) {
		m, err := filenameToMap("countries.json")
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
		c.JSON(200, m)
	})
	g.GET("/consent/documents/EAR-B-WW/:restriction/:lang/over", func(c *gin.Context) {
		m, err := filenameToMap("over.json")
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
		c.JSON(200, m)
	})
	g.POST("/projects/*junk", func(c *gin.Context) {
		m, err := filenameToMap("projects.json")
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
		}
		c.JSON(200, m)
	})
}

func registerAuth(r *gin.Engine) {
	g := r.Group("/auth")

	g.POST("/login", func(c *gin.Context) {
		data := "\x83\xa9" + "SessionId" +
			"\xd9\x24" + newUUID.String() +
			"\xa6" + "UserId" +
			"\xd9\x24" + newUUID.String() +
			"\xb0" + "IsInCommunityBan" +
			"\xc2"

		c.Header("x-session-nonce", newUUID.String())
		c.Data(200, "application/octct-stream", []byte(data))
	})
	g.POST("/ticket", func(c *gin.Context) {
		data := "\x81\xa6" + "Ticket" + "\xd9\x24" + newUUID.String()

		c.Header("x-session-nonce", newUUID.String())
		c.Data(200, "application/octct-stream", []byte(data))
	})
}

// --------------------------
// todo:

func registerOthers(r *gin.Engine) {
	r.POST("/delivery_data/get", func(c *gin.Context) {
		c.File("asserts/delivery_data_get.bin")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("x-session-nonce", uuid.New().String())
	})

	hunterG := r.Group("/hunter")
	hunterG.POST("/sync", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}

		var hunter Hunter
		if err := msgpack.Unmarshal(body, &hunter); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("%v\n", hunter)
		uid := hunter.HunterSaveList[0].HunterId
		if uid == "" {
			uid = newUUID.String()
		}

		data := "\x83\xb7" + "InvalidSaveSlotInfoList" +
			"\xc0\xb9" + "InvalidClientHunterIdList" +
			"\xc0\xb0" + "SaveSlotInfoList" +
			"\x91\x85\xa8" + "HunterId" +
			"\xd9\x24" + uid +
			"\xaa" + "HunterName" +
			"\xac" + hunter.HunterSaveList[0].HunterName +
			"\xa9" + "OtomoName" +
			"\xa6" + hunter.HunterSaveList[0].OtomoName +
			"\xa8" + "SaveSlot" +
			"\x00\xa7" + "ShortId" +
			"\xa8" + "1A2B3C4D"

		c.Header("x-session-nonce", newUUID.String())
		c.Data(200, "application/octet-stream", []byte(data))
	})
	hunterG.POST("/character_creation/upload", func(c *gin.Context) {
		// todo:
	})
	hunterG.POST("/profile/update", func(c *gin.Context) {
		// todo:
	})
	hunterG.POST("/update/rank", func(c *gin.Context) {
		// todo:
	})

	r.POST("/obt/play", func(c *gin.Context) {
		// todo:
	})
	r.PUT("/character-creation/*junk", func(c *gin.Context) {
		// todo:
	})
	r.PUT("/hunter-profile/*junk", func(c *gin.Context) {
		// todo:
	})
}

func registerInGame(r *gin.Engine) {
	r.POST("/follow/total_list", func(c *gin.Context) {})
	r.POST("/offline/notification_list", func(c *gin.Context) {})
	r.POST("/community/invitation/received_list", func(c *gin.Context) {})
	r.POST("/block/list", func(c *gin.Context) {})
	r.POST("/friend/list", func(c *gin.Context) {})

	r.POST("/lobby/auto_join", func(c *gin.Context) {})
}

func registerWssHandler(r *gin.Engine) {
	r.GET("/ws", func(c *gin.Context) {
		// todo:
	})

	r.GET("/socket", func(c *gin.Context) {
		// todo:
	})
}

// --------------------------

func filenameToMap(filename string) (map[string]interface{}, error) {
	data, err := os.ReadFile("asserts/" + filename)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}
