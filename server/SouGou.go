package server

import (
	"bytes"
	"encoding/base64"
	"go-figure-bed/pkg/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

//上传到搜狗，免登陆
func UpLoadToSouGou(img []byte) string {
	preStr := "LS0tLS0tV2ViS2l0Rm9ybUJvdW5kYXJ5R0xmR0IwSGdVTnRwVFQxaw0KQ29udGVudC1EaXNwb3NpdGlvbjogZm9ybS1kYXRhOyBuYW1lPSJwaWNfcGF0aCI7IGZpbGVuYW1lPSIxMS5wbmciDQpDb250ZW50LVR5cGU6IGltYWdlL3BuZw0KDQo="
	sufStr := "DQotLS0tLS1XZWJLaXRGb3JtQm91bmRhcnlHTGZHQjBIZ1VOdHBUVDFrLS0NCg=="
	preStr = utils.Decode(base64.StdEncoding, preStr)
	sufStr = utils.Decode(base64.StdEncoding, sufStr)
	imgStr := string(img)
	data := []byte(preStr + string(img) + sufStr)
	url := "http://pic.sogou.com/pic/upload_pic.jsp"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", " multipart/form-data; boundary=----WebKitFormBoundaryGLfGB0HgUNtpTT1k")
	req.Header.Add("Content-Length", string(strings.Count(imgStr, "")))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	respUrl := string(body)
	respUrl = strings.Replace(respUrl, "http", "https", -1)
	return respUrl
}
