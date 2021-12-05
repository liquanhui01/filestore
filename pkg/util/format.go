// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"strings"
)

// formatFileSize format file size string.
func formatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

func FormatFileType(filename string) (string, bool) {
	array := strings.Split(filename, ".")
	suffix := array[len(array)-1]

	switch suffix {
	case "mp3", "mp4":
		return "音视频", true
	case "png", "jpg", "jpeg", "gif", "bmp":
		return "图片", true
	case "doc", "docx", "xls", "ppt", "pptx", "pages", "txt":
		return "常规文件", true
	case "rar", "zip", "tar", "gz", "bz2":
		return "压缩文件", true
	default:
		return "", false
	}
}
