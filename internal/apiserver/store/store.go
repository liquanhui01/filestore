// Copyright 2021 QuanhuiLi. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

//go:generate mockgen -self_package=file-store/internal/store -destination mock_store.go -package store file-store/internal/store Factory,UserStore

var client Factory

// Factory defines the platform storage interface.
type Factory interface {
	Users() UserStore
	Folders() FolderStore
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the store client.
func SetClient(factory Factory) {
	client = factory
}
