// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import "golang.org/x/crypto/bcrypt"

// GeneratePassword
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword
func ValidatePassword(password, hashed string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return err
	}

	return nil
}
