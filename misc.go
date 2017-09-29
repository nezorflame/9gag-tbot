package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash"
	"log"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/encoding/charmap"
)

func urlArgsStr(args map[string]string) (result string) {
	for k, v := range args {
		result += "/" + k + "/" + v
	}
	return
}

func formReqSignature(ts int64) string {
	return getSHA1HexStr(fmt.Sprintf("*%d_._%s._.%s9GAG", ts, appID, deviceUUID))
}

func randomUUIDHexStr() string {
	return fmt.Sprintf("%x", uuid.NewV4().Bytes())
}

func randomSHA1HexStr() string {
	return getHashedHexStr(sha1.New(), strconv.FormatInt(time.Now().Unix(), 10))
}

func getSHA1HexStr(str string) string {
	return getHashedHexStr(sha1.New(), str)
}

func getMD5HexStr(str string) string {
	return getHashedHexStr(md5.New(), str)
}

func getHashedHexStr(h hash.Hash, str string) string {
	var err error

	if str, err = charmap.ISO8859_1.NewEncoder().String(str); err != nil {
		log.Fatal(err)
	}

	h.Write([]byte(str))

	return fmt.Sprintf("%x", h.Sum(nil))
}
