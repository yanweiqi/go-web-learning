package chapter4

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

const (
	url1     = "http://127.0.0.1:8080/process"
	fileName    = "/tmp/a.txt"
)

func TestUpload(t *testing.T) {
	postFile(fileName,url1)
}


func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploaded", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}


	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, bodyWriter.FormDataContentType(), bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(res_body))
	return nil
}