// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
	"io"
	"mime/multipart"
	"os"
)

// LocalStore storage file to the local server.
func LocalStore(file multipart.File, location string) (string, string, error) {
	// create new file
	newFile, err := os.Create(location)
	if err != nil {
		return "", "", err
	}
	defer newFile.Close()

	// copy file to new file
	size, err := io.Copy(newFile, file)
	if err != nil {
		return "", "", err
	}
	newFile.Seek(0, 0)

	filesha1 := FileSha1(newFile)

	return filesha1, formatFileSize(size), nil
}

// func FileIsExit(path string) (bool, error) {
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		return true, nil
// 	} else if os.IsNotExist(err) {
// 		return false, err
// 	} else {
// 		return false, err
// 	}
// }
