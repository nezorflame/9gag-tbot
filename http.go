package main

import (
	"strconv"
	"time"

	fh "github.com/valyala/fasthttp"
)

func requestGET(uri string, sign bool) (result []byte, err error) {
	ts := time.Now().Unix() * 1000

	req := fh.AcquireRequest()
	req.SetRequestURI(uri)
	req.Header.SetMethod("GET")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("9GAG-9GAG_TOKEN", appToken)
	req.Header.Set("9GAG-TIMESTAMP", strconv.FormatInt(ts, 10))
	req.Header.Set("9GAG-APP_ID", appID)
	req.Header.Set("X-Package-ID", appID)
	req.Header.Set("9GAG-DEVICE_UUID", deviceUUID)
	req.Header.Set("X-Device-UUID", deviceUUID)
	req.Header.Set("9GAG-DEVICE_TYPE", deviceType)
	req.Header.Set("9GAG-BUCKET_NAME", bucketName)

	if sign {
		req.Header.Set("9GAG-REQUEST-SIGNATURE", formReqSignature(ts))
	}

	defer fh.ReleaseRequest(req)

	resp := fh.AcquireResponse()
	if err = fh.Do(req, resp); err != nil {
		return
	}
	defer fh.ReleaseResponse(resp)

	return resp.BodyGunzip()
}
