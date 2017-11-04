// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package nvm

/*
void V8Log(int level, const char *msg);
char *StorageGetFunc(void *handler, const char *key);
int StoragePutFunc(void *handler, const char *key, const char *value);
int StorageDelFunc(void *handler, const char *key);

// The gateway function
void V8Log_cgo(int level, const char *msg) {
	V8Log(level, msg);
};
char *StorageGetFunc_cgo(void *handler, const char *key) {
	return StorageGetFunc(handler, key);
};
int StoragePutFunc_cgo(void *handler, const char *key, const char *value) {
	return StoragePutFunc(handler, key, value);
};
int StorageDelFunc_cgo(void *handler, const char *key) {
	return StorageDelFunc(handler, key);
};
*/
import "C"