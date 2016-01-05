package main

import (
	"io"
	"os"
)

/**
*
**/
const (
	FILE_REEOR_CREATE_FAILD    int = 400
	FILE_REEOR_CREATE_SUCCESS  int = 100
	FILE_REEOR_WRITE_FAILD     int = 401
	FILE_REEOR_UPLOADE_SUCCESS int = 101
	FILE_REEOR_UPLOADE_FAILD int = 402
)

/**
*create and write a file to localhost
**/
func write(b []byte, fileName string) (int, error) {

	file, err := os.Create(fileName)
	if err != nil {
		return FILE_REEOR_CREATE_FAILD, err
	}
	_ , err = file.Write(b)
	if err != nil {
		return FILE_REEOR_WRITE_FAILD, err
	} else {
		return FILE_REEOR_CREATE_SUCCESS, nil
	}
}

/**
*create a local file and uploade it to server
**/
func upload(url_upload string, url_local string) (int, error) {
	//1.greate  local file and uploade file
	file_upload, err := os.OpenFile(url_upload, os.O_RDWR, os.ModeType) //a file need upload
	//2.check if the file can be created
	if err != nil {
		return FILE_REEOR_CREATE_FAILD, err
	}
	file_local, err := os.Create(url_local) //a file
	//2.check if the file can be created
	if err != nil {
		return FILE_REEOR_CREATE_FAILD, err
	}

	//3.uploade the local file to a url
	_, err = io.Copy(file_local, file_upload)
	if err != nil {
		return FILE_REEOR_UPLOADE_FAILD, err
	} else {
		return FILE_REEOR_UPLOADE_SUCCESS, nil
	}
}
