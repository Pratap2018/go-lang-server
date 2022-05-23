package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Pratap2018/go-lang-server/entity"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

type SignatureGenerator interface {
	GenerateSign() entity.RequestBodyRedirection
}

type UserRedicrectionBodyType struct {
	Message entity.UserRedicrectionMessageType `json:"message"`
}

func (requestData UserRedicrectionBodyType) GenerateSign() entity.RequestBodyRedirection {
	private_key := os.Getenv("HYPERFYRE_PRIVATE_KEY")
	edcsaPrivateKey, err := crypto.HexToECDSA(private_key)
	if err != nil {
		panic(err)
	}
	requestData.Message.Iat = time.Now().UnixMilli()
	requestData.Message.Exp = time.Now().UnixMilli() + (1000 * 24 * 60 * 60)
	message := requestData.Message
	byteArray, err := json.Marshal(message)

	stringmsg := string(byteArray)
	fmt.Println(stringmsg)

	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(stringmsg), stringmsg)
	data := []byte(msg)
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), edcsaPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	var resBody entity.RequestBodyRedirection = entity.RequestBodyRedirection{Message: message, Fyresign: hexutil.Encode(signature), MessageHash: hash.Hex()}
	fmt.Println(resBody)
	return resBody
}

func AuxController(ctx *gin.Context) {
	var data entity.UserRedicrectionMessageType
	ctx.ShouldBindJSON(&data)
	userData := UserRedicrectionBodyType{Message: data}
	resBody := SignatureGenerator.GenerateSign(userData)
	ctx.JSON(200, gin.H{
		"Data": resBody,
	})
}
