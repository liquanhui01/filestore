// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import "strings"

// Filename rename file name.
func Filename(oldFilename, newFilename string) string {
	array := strings.Split(oldFilename, ".")
	suffix := array[len(array)-1]
	return SplicingString(newFilename, ".", suffix)
}
