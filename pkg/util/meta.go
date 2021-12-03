// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
	"mime/multipart"

	"github.com/spf13/viper"

	rp "github.com/liquanhui01/filestore/internal/apiserver/store/repo"
)

func Meta(r rp.File, folderid, userid uint64, filename string, file multipart.File) (rp.File, error) {
	r.Filename = filename
	r.Location = SplicingString(viper.GetString("file.location"), filename)
	r.Folderid = uint(folderid)
	r.Userid = uint(userid)

	filesha1, size, err := LocalStore(file, r.Location)
	if err != nil {
		return r, err
	}
	r.Filesha1 = filesha1
	r.Filesize = size

	return r, nil
}
