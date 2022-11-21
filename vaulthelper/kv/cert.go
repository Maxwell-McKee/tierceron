// Code generated by go-bindata.
// sources:
// ../../certs/cert_files/dcipublic.pem
// ../../certs/cert_files/dcidevpublic.pem
// DO NOT EDIT!

package kv

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if gz != nil {
		defer gz.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _CertsCert_filesDcipublicPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x55\xc9\x92\xab\x38\x16\xdd\xf3\x15\xb9\x6e\x47\x97\xf1\x0c\x8b\x5c\x48\x20\x40\x66\x70\x8a\x19\x76\x4c\x29\xc0\x60\x6c\x26\x61\x7f\x7d\x47\xfa\xf5\xab\xae\x78\xcd\x0a\x1d\x9d\x3b\xc6\x3d\xba\xc3\x94\xd6\x45\x36\x7e\xae\x25\xeb\xf3\x5f\x7f\xe5\xc5\x32\x16\x7d\x56\x26\xb7\xe2\xf9\x57\xd6\xb5\x5c\x35\x0c\x53\xd1\x7f\xae\xa5\x4f\x15\xae\x1d\xf7\x53\xed\x8b\x64\x2c\xfa\x0f\x33\xb9\x65\x65\x31\x8c\x45\xbf\x36\x3e\x9d\xa4\xf9\xee\xfa\x7c\x7d\xf9\x74\x8a\x6c\xac\x68\xf7\x61\x54\x6d\x35\x16\xf9\x8f\xd3\xdf\x90\xed\x80\x0f\xb9\x6b\x93\xea\xf6\xe1\x27\x4d\x95\x27\x63\xd5\xdd\x3e\x9c\x22\x9b\xfa\xe2\xc3\x29\xfa\xb9\xe8\x3f\x24\xc0\xfd\xfb\xe7\x83\x48\xc5\xd6\x87\x84\x6c\x17\x2b\x58\x02\x2e\x7a\xa3\x9c\x89\xb1\xb2\xc8\x92\x04\xf5\x27\x05\x0c\x43\x40\xb1\x0d\xcc\x4b\xae\xa8\x4d\x18\x2a\xcf\xbb\xa5\x0f\x4c\xdb\xf0\xbb\x76\x4d\x99\x4c\xa2\xb3\xde\xc5\xb8\x9c\x33\x0b\x10\x64\x40\x02\x18\x47\x23\x61\x91\x5e\xe0\x0c\xa9\xe5\x43\x10\xb9\xe0\x9a\x4b\xa6\x3d\x30\x95\x44\xb2\x4f\x08\x46\xcb\x59\xcb\x5a\xbf\xcc\x55\xff\x89\xd1\xa6\x4c\x5b\xab\x8b\x43\x8b\x8f\x43\xbc\x20\x19\x5c\xb8\x5f\x86\x99\x0b\x37\x56\x99\xaa\xf1\x9c\xb5\x64\x51\x65\x10\xfc\xc2\x3b\x57\xde\x58\x4d\xb4\xb3\xef\xf1\x56\xa0\xae\xda\x8c\x49\x68\x37\xb1\x8c\x76\xa6\xec\xa9\x60\xe3\x21\x0e\x2c\xe6\xe4\x6d\xfd\x3a\x57\x9b\x5b\xfa\x84\x8e\xc7\x23\x6a\xab\xe2\x18\x05\xcd\x84\x95\xb8\x4c\xd5\xe6\x1a\x85\xf6\x3d\xdd\xee\xe9\x9b\x17\x9e\x1b\xac\x58\x4d\x76\x8b\x9b\xac\x82\x32\x47\x5c\x50\x28\x8c\x7f\x9a\x32\x62\xa6\x6b\x32\x53\x06\xcc\x94\x61\xf2\xc6\x5c\xb4\x98\xae\xf9\x34\x5f\xde\xc1\x72\x9b\xc4\x84\xfc\xa2\xbe\x40\xfc\x2b\x3b\xd3\x44\x55\x37\x71\xb1\xea\xef\x7f\xea\x8b\xb6\x65\x99\xb6\xfe\xc1\x68\xad\x39\x75\x25\x09\x38\xf8\xcf\xa6\x41\x48\x80\x4c\x29\xfa\x02\xb2\x24\x01\xd2\x49\x94\x22\xc8\x81\x73\xdb\x4c\xed\x64\xaf\x50\xef\xdc\xe0\xe6\x94\x78\xfb\x74\x6f\xfb\xd7\xc7\xb6\x48\xcc\x5d\xf3\x0a\x55\x3d\x3b\x3e\x75\x99\xf8\x6b\x7b\xde\xa7\xb1\xbc\xb1\x0f\x86\x77\x0f\xc7\xc0\x9d\xa7\x79\xbf\xe2\x66\x7f\xed\x3a\x71\x7f\x83\x29\xca\x1e\xeb\x58\xaf\xc4\xf9\x14\x4b\xcb\x71\x6e\x8e\x6a\x2f\x34\xba\xe8\x7b\x8f\x7b\x79\x46\xf8\x74\xdb\xc4\xa8\x73\xef\xba\x1f\xb8\x87\x57\x71\x44\xe2\x98\x7f\x03\x8f\xb3\x46\xef\x50\x27\x9b\x49\x8b\xe8\x76\x16\x15\x7e\x69\x4f\x53\x0b\xfc\xb4\x8d\xe1\x55\xf3\x71\x38\x13\x9d\xd0\xfd\xb9\x6d\xe9\xf1\xe0\x5a\xae\x12\x41\xb5\x4e\x76\x37\x7d\x1b\x16\x53\x20\xa8\x2f\x6e\x89\x9f\xaa\xe0\x3e\x8d\x2f\x5b\xb9\xdc\x42\xcf\xa9\xd2\xd5\x2e\x40\x24\xb8\x68\x71\x23\x48\x7e\x2e\x4f\xa6\x50\x4a\x01\x54\x7a\xb1\x32\xa6\xbc\x7d\x05\x5a\xd0\x6b\xbd\x11\xae\xe3\x3e\x0a\x07\xc8\x59\x8a\xda\xa1\x66\xa8\x27\x1e\xf6\x8b\x6c\x45\x62\x64\x4b\x7b\x77\xc0\xcf\x69\x3d\x64\x69\x78\x5e\x6f\x5a\xe4\xad\xaa\xce\xda\x67\xaf\x25\x8c\xc4\xc1\x02\xe2\xb2\x6e\x89\xed\xdf\x4f\x9a\xd7\x10\x6e\xf4\xe5\xd5\xd9\x13\x67\xe9\xce\x70\x3b\x9f\xea\xfb\x6a\x8d\x33\x0d\x49\x80\x21\x00\x92\x8b\x04\xba\x8e\x51\xaa\xab\x26\x14\x7e\x06\x27\xc7\x8c\x44\x26\x4c\x80\x82\xb7\x66\x38\xd8\x1e\xd7\x47\x7d\x99\xaf\xda\x74\x15\x0f\xca\x3e\xa5\xb0\x0e\xb4\xd2\x84\xfc\x9b\x2c\x53\x12\x40\xe8\xcc\x75\x72\x09\xcc\x65\x3d\x5d\xd3\x08\xcb\x7c\x51\x4d\xdb\xe1\x3a\x28\x57\x5f\xfb\x3d\xca\x1a\x11\x20\xf8\x16\x10\x04\xa6\x04\x13\xc0\x64\x10\xc9\xbe\xcd\xbb\x80\x68\x6b\x08\x30\x03\x32\xc8\xdf\x3c\xc7\x43\x4a\x0d\x3c\x48\x69\x0f\x29\x52\x20\xc9\x64\x0e\x90\x08\xeb\x2c\x82\x90\x78\x1a\x60\x98\x39\xe4\x6d\x4c\x21\xc2\x8c\xc8\x80\x87\x74\xf8\x21\x23\x90\xe2\x05\x10\x2c\xc1\x17\x68\x4c\xc9\x54\x25\x69\x50\x01\xf1\x14\x0e\x32\x0c\x95\x32\xef\x72\xcd\x66\xd9\xab\x9b\x8d\xdd\xdf\xe2\x99\xa2\xad\x38\x1a\xbc\x45\xbc\x17\xc0\x90\xc6\x37\x4a\x18\x04\x14\x31\x1a\x91\x7f\x3a\x40\x10\x6a\x94\xe5\x35\xfc\xfa\x47\x6a\x0c\x74\xb1\x9c\x68\x36\x9f\xc9\xdd\x6c\x6c\xad\x67\x2e\x1d\x5e\x71\x60\xf1\x49\x90\xcf\xef\x71\x77\x44\xf7\xe7\xcc\xfd\x00\x5e\x63\xc1\xdf\x0a\xf4\x5b\x65\x48\x02\xbb\xcc\xd5\x66\x4e\x9b\x9f\x6c\xfc\x67\xec\xff\x57\x7d\x57\x0b\x1a\xad\xf5\xcc\x65\x50\xff\x0e\xc6\xbd\xa3\x45\x51\xf8\xbf\x68\x62\x9d\xed\xc0\x94\xfd\x2d\xf1\x43\x9d\x6e\x79\x66\xb0\x77\x6f\x6c\x28\x51\x76\xee\xb0\xa3\x57\x87\x6b\x1c\x96\x3c\x17\x87\xe7\x3a\x51\x95\x29\x0e\xaf\xef\x9a\x69\x09\xdf\x17\x7f\xe2\x26\xfe\x79\xee\x22\xfd\xdd\x70\xa0\x05\x05\xc1\x08\x50\x07\x72\x02\x73\xa0\x00\xe4\x09\x68\x1e\xf8\x76\x4f\xcf\x15\x5e\xaf\x2b\xc5\xef\xce\xa6\x01\x9e\xf7\x83\x53\x85\xd7\x7e\x21\x87\xbd\x14\x0a\x53\x72\xcf\xbb\x36\xdc\x57\x82\x95\x01\x00\x94\x0d\xee\x9c\x92\x2b\x08\x00\x10\x98\xc0\xae\x21\x02\x15\x50\xda\x2d\xd0\xf4\xfa\xd0\x3f\x4a\x75\xf0\x2f\x41\x6f\x3c\xd6\xdf\x73\xb0\x31\x84\x41\x17\x2e\x19\x9e\x13\xa9\x1d\x14\xdd\xf8\x16\x29\xa6\x17\xbe\x33\x39\x43\x61\xde\xdc\x39\x1b\xb6\x43\xf5\x50\x08\xa1\xf8\x9d\xeb\xae\x88\xfc\x50\xa1\xb7\x28\x4f\xe1\x88\x0a\xdb\xdf\x53\x90\x13\xc9\xc3\x06\x2b\xea\x31\xb0\x52\x5c\xbe\xb4\xbd\x6e\xe0\x8a\x59\x1c\x9f\xdf\xad\xb0\x5d\xbe\x1a\x79\x53\x6e\xf9\xfd\x1c\xa0\x6d\xc5\x28\x00\x20\xf4\xaa\x52\x57\x7f\xfe\x10\x60\x50\x35\x11\x91\xb0\x44\x1e\x0b\x4d\x2b\x15\xcb\x8d\xde\x90\xc3\xcc\x1e\x19\x77\x0a\x2d\x25\x33\x8c\xaf\x21\x68\xce\x29\x15\xc0\xd6\xe9\xe2\x5b\x85\x5f\xa0\x02\x87\xea\x99\x1c\x1e\x62\xdd\xa0\x31\x05\x57\x65\xdf\x1c\xec\x00\x8e\xfb\x2f\x45\x3d\xb0\x52\x7c\xe4\xc7\x3a\x35\x09\x77\xb3\x4f\x26\x7a\x01\x0b\xd2\xeb\xa3\xbc\x56\xaa\xc8\x78\x08\xc8\xa0\x00\x70\x91\x00\x41\xe0\x32\xee\x5e\xd7\x7e\x97\xef\x36\xcc\x58\x5e\x8e\x5e\xfb\xb7\x4d\x2a\xae\x06\xed\xf2\x38\x76\x85\x73\xe2\xfc\xf5\x29\x56\xac\x53\x9b\xb4\x73\xd0\x8d\x99\xb0\x12\x4f\xab\xf0\xb1\xab\xda\xf5\xf7\xec\x6f\xac\x63\x51\x37\x03\x7a\xc4\xb6\x75\x84\xf0\x59\x19\x24\x91\xaf\x45\x7d\x6a\xa7\xa8\x27\x67\xc1\xf5\xbe\x39\x35\x39\x1e\xdc\xa7\x95\xe7\x59\x62\xed\x0b\x0f\xf6\xc6\x35\x0c\xd6\xa6\x3f\x27\x3a\xf1\x40\x56\x53\xf9\x6b\xd1\x80\xb6\xa4\xa2\xa6\xa2\x7b\x34\x5c\xc7\x70\x6f\xf4\x5d\x0d\x59\x2f\xf3\xfb\xc7\xc4\xd1\xea\x7a\x36\xac\x50\x52\x8f\x4c\xf7\xbd\x1b\x91\xfb\xa8\x0e\x8f\xfd\xf8\x12\xef\xf5\x2a\xdc\x99\x77\x4f\xd0\x14\x7f\x25\x23\x7f\x4b\x74\x7b\x94\x64\xe3\x3c\x47\xf8\x14\xfb\xe5\xf3\xb6\x4f\xbe\xaf\xdc\xac\x15\xab\x5a\xd5\x76\x3d\x29\xe6\xe3\x6e\x2c\x9f\x90\xc6\x97\x2e\xad\xc1\xf7\x78\x4b\x62\xd3\x34\xe6\xa5\x9e\xbb\x6a\x5e\xd3\x6e\xf4\xc7\x53\xca\x62\xe9\x78\x69\xdb\x78\xf7\xdc\x6c\xbd\x20\x1f\xb9\xd7\x2d\xb5\x4c\xfb\xd0\x57\x11\x0d\x77\xba\xd8\xb0\x58\x20\xbc\x6c\xd9\xc1\x56\x68\x5f\x5f\xf8\xf6\x22\xdf\x11\x8b\x56\xe1\x89\x80\x8e\x27\x5e\x77\x45\xe7\xc2\x60\x9f\x9f\xbf\x56\x3e\xb2\xe4\xff\x5f\xf8\xdc\x7f\x02\x00\x00\xff\xff\xf9\x46\x7c\x49\x9e\x08\x00\x00")

func CertsCert_filesDcipublicPemBytes() ([]byte, error) {
	return bindataRead(
		_CertsCert_filesDcipublicPem,
		"../../certs/cert_files/dcipublic.pem",
	)
}

func CertsCert_filesDcipublicPem() (*asset, error) {
	bytes, err := CertsCert_filesDcipublicPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../certs/cert_files/dcipublic.pem", size: 2206, mode: os.FileMode(420), modTime: time.Unix(1605581277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _CertsCert_filesDcidevpublicPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x56\x49\xd3\xa3\x38\x16\xbc\xf3\x2b\x7c\x1e\xd7\xb4\x01\x6f\xd0\x11\xdf\xe1\x09\xc4\x66\x83\xcd\x6e\xb8\xb1\xef\x60\x03\x36\xe0\x5f\x3f\xe1\xaf\xba\xaa\x3b\x66\x86\xe0\x80\xf2\x3d\xa5\x92\x0c\x65\x48\x28\xc8\x56\x30\x8e\x7d\x11\x3e\xc7\x64\x20\x56\xab\xd5\xaa\xee\xa2\xa0\x3e\x25\x8b\xcc\xff\xb9\x22\xa9\x15\x49\xfe\xf5\x7e\x17\xd3\xbe\x48\xda\xb8\x5e\xb4\xa0\x49\xfe\x5c\xc5\xc9\x1c\xe5\x41\x9c\xbc\x56\x53\x51\xc7\x51\xd0\xc7\xc4\xf0\x0c\xcb\x24\x1a\xbf\x38\x6d\xf5\xb5\xfa\xd7\x1f\xbf\x3b\xfe\x88\xba\x86\x20\x8a\x61\x78\x26\xfd\x17\xb7\xfa\x5a\x89\xe8\xc7\xca\xb4\x3e\x1f\x7d\x12\x8c\x49\xbf\x52\x83\x36\xca\x93\x61\x4c\xfa\x1f\xab\xf3\xea\x6b\x65\x06\x75\xda\xf5\xf1\x8f\xd5\xe5\x33\x48\xa2\xb1\xc8\xba\xd5\xb9\x68\x8a\x31\x89\x7f\xac\xbe\xf9\x7f\xa1\x86\x09\x2b\xbe\x6b\x82\xa2\x5d\x39\x41\x5d\xc4\xc1\x58\x74\xed\xa7\xfa\xec\x93\x95\x99\xf4\xaf\xa4\x5f\x71\x40\x10\xff\xfe\x3c\x08\x8b\xb2\xb6\xe2\xb0\x61\xc9\x82\xcc\x81\x85\xbf\x51\x42\x95\x65\x71\xb0\x38\x0e\xf9\x4d\x06\x93\x8c\x20\x93\xf5\xb4\xbd\x85\x90\x6c\x2f\x86\xec\x14\x4e\xba\x67\xb2\x96\x35\xc0\x02\x0d\x65\xd5\x23\xaf\x0a\x91\x9d\x48\x04\xfa\x20\x00\xcf\x21\xa2\x7c\xe3\xb3\x0a\x95\x08\x94\x8d\x51\xae\x72\x06\x29\xcf\xe2\x1b\x7c\x94\x69\x0e\x82\xcc\xc2\x55\xbc\xf8\xae\x40\xfa\x37\x39\xb3\x5c\xe1\xe9\xd1\x79\x1d\x6d\x8d\x3a\x2a\xb1\xae\xc2\x4e\x24\xbe\x27\xce\xaa\x64\xd3\xc2\xe0\x37\xec\xe2\xf3\xd8\x53\x91\xf7\x4d\xc8\xe5\xea\xd5\xa6\x9d\x32\x16\xeb\x36\x5c\x90\x1a\xb8\xd4\x3d\x16\x9d\x4a\xb5\xa2\x49\xd3\x3d\xde\xd1\x75\x9e\xc0\xcb\xde\xf2\x5d\x8d\x0c\xdc\xf8\x25\x0b\x8a\xa5\x9b\x08\x87\x34\x95\x07\xee\x2e\x73\x1a\x61\x08\x5c\x23\x8f\xc5\xfa\x15\x16\xe8\xd3\x47\x45\x8d\x9d\xd9\xb4\xb3\xc4\x8d\xb3\xc8\x58\x43\x84\x8a\x76\x37\xde\x92\x27\x95\x57\x67\xad\x84\x49\xe5\x61\x52\x85\xee\x83\x2d\x2a\xef\xcd\x17\xfe\x37\x36\x89\xe5\xdf\xea\x60\x9a\xae\xa7\x62\x5f\x11\xfe\x2d\x2f\x03\x51\xa8\xfc\x9b\xf7\xf4\x68\x76\x54\x65\x19\xc9\xe5\x7f\x1b\x86\x05\x80\x0b\x07\x3a\x03\x9f\x3a\x97\x9d\x38\xd0\x31\xb0\x46\x97\x12\xc7\x52\x7b\xe8\x6b\x47\x12\xc5\xd1\x3f\x66\xca\xde\x9f\xae\xc7\x61\x7d\xd1\x90\x2b\xf1\x9b\xa6\x4e\x51\x17\x5c\x13\x2a\x80\xec\x80\x79\xc1\xdc\xe2\x86\x93\xf5\x93\x12\xdc\xef\x64\xc3\x29\x4d\x1a\x10\x4f\x6a\xdb\xb0\xde\xda\xc1\x89\xe7\xd9\x89\x60\xda\xb7\xd3\xb8\xf0\xe5\x7b\xa7\xbf\x1f\xe9\x71\xad\x9e\xee\x0f\x7c\xeb\x75\x4c\x6e\x62\x64\xbe\x1d\x56\x8c\x19\x65\x3d\xec\x60\x36\x39\x31\x7f\x6e\x4f\x04\xef\x25\xbb\x43\x73\x67\x92\xd7\xda\x14\x39\x3d\x17\xe9\x37\xbb\xbd\xc0\x06\x19\xa0\x0a\x46\xbb\xbb\xe7\x8b\x72\x1f\x73\xb8\xf9\xbc\x53\x6c\x5e\xb3\xef\x72\xd7\x1b\xdf\xe2\xb0\x5b\x63\xeb\xda\x74\x44\x8a\x82\x05\x8e\xc1\x36\x75\xdc\x42\xb9\xec\xd6\x56\x22\xd9\x1b\x3e\xd2\x90\xfd\x78\xdd\x7b\xd2\x48\xef\xe9\x96\xab\x9e\x25\xf3\x50\xd8\xad\x54\x85\xed\xa9\x75\x19\xa4\x86\xe5\xee\x20\x1c\x10\xce\x09\xf3\x80\x18\x1a\xb0\x29\x17\x3e\x5b\x84\xac\xe7\x9f\x61\x4a\x9e\xde\x3e\x89\x71\xad\x3e\xac\xc1\x09\xd1\x30\x3f\xae\xa3\x31\x52\xd4\x03\x41\x88\x99\x23\x8b\x34\x1b\xc1\x81\x7a\x47\xe0\x05\x84\xc6\xf2\x1d\x13\xf8\x63\xa6\xd3\x6d\x3d\x5f\x5e\x6f\x5d\xe6\x41\x07\xd4\xed\x64\x3e\x79\x73\x1c\x6c\xa3\x49\x9a\x3c\xde\x31\xc8\x12\xa1\x6c\x12\x3a\xb0\x4b\x6f\x4e\x66\xc1\x1c\x8b\xa7\xb0\x25\x0e\xfe\x6b\xdf\x4c\xb7\xfc\x09\xa2\xe6\x25\x78\x92\xf4\xef\xe6\x0b\x42\x1e\x16\x84\x0e\xc9\x6b\x2b\x8d\x5d\xf7\xa4\x9d\xaa\xe5\x7d\xa0\xf3\x74\x64\x14\xb5\xa6\xc6\xcf\x36\x06\xca\x8e\x09\x7e\xc2\x68\x33\xe9\x18\x26\x59\xe8\x78\x50\x3f\x19\x90\x0c\x15\x41\xca\x60\x28\x01\x54\x44\x7e\x37\x2a\xba\xee\xaa\x48\x17\x39\x6e\x10\x41\xb7\x05\x34\xa9\x08\x65\x59\x4f\xa0\x0c\x0b\x48\x8f\x78\x28\x91\xf2\x3d\xd9\x04\xac\x97\x08\x54\x5e\x17\xb9\xe5\xd3\xac\xa3\xa1\xc4\x9f\x94\x4a\x2a\x67\x4f\xf2\xe4\xc9\xa7\xc9\x43\x48\xb7\x25\xc8\xb0\x4b\x08\x74\x4e\xc6\x12\x7a\x5f\x0a\xe6\x15\xfd\x0e\xce\xbe\x0c\x69\xf2\xa5\x53\xc8\x52\x21\x13\x51\x93\x20\x1e\x30\x07\x16\x87\x72\xf8\x07\x81\x8e\x31\x91\xf0\x88\x56\x31\xf3\xb7\xb4\x37\x70\x79\xa5\x75\xb1\x64\x4c\x1f\x52\x6f\xab\x90\xe7\x56\xab\xbd\xad\x71\xf7\x69\xe6\x7b\xbb\x9f\xa9\x5f\x63\xd6\x24\x6c\x52\xf8\x2b\x7d\x7b\xd7\x73\xe7\xbb\x2f\x0a\x64\xe0\xb2\xcf\xef\x18\xdf\x94\xfa\x57\xf2\x74\x12\x3f\x3f\x64\x2a\xa7\xfe\x73\x31\x44\xe4\x79\xfc\x7b\xb5\x90\xd6\xde\x11\xb7\x7f\xff\x8a\xf7\xb9\xd1\x5e\xa1\x05\xf7\x9f\xc6\x62\x2c\x97\x90\x65\x19\xf3\x38\x37\x46\x9d\x88\x5a\xe7\xb9\x46\x4d\xc4\xc5\xe7\x7f\x69\x8e\x77\xff\x06\x7f\x62\xe4\x94\x65\xd2\x1e\x65\x5d\x8f\x32\x8c\x21\xf6\xf7\x90\xe9\x1c\x92\x65\x74\xd0\x4d\x0e\x12\x1b\xed\x26\xb4\x25\x00\x07\xb5\x93\x8c\xd4\xfa\x8e\x33\xf5\xec\x16\x85\xdb\x92\xcc\x96\x34\xce\x58\x20\x5f\x2f\x4a\x79\xca\x6e\xcf\xbc\xe6\x69\xd3\x50\x12\x00\xa0\xe8\x72\xe5\x06\x93\x01\x00\x9d\x07\x9c\x4d\x84\x91\xc9\x39\xa8\xf3\xd1\xda\xe7\xfd\xc1\x51\x36\x26\x45\xdd\x7c\x5f\xbf\xd0\xee\x95\x32\xb4\x71\x14\x30\xc7\x55\xb7\xa5\xd5\x65\x2a\xd9\xcb\xf5\x1d\x0a\x0c\xb9\xd6\xf2\xe3\x0d\xed\x17\xcb\x6c\x89\x9c\x3e\x1d\x8f\x01\xd5\xa6\x9b\x5c\x88\x8f\x3e\x05\xaf\xc7\xce\xbf\xb9\x07\x7d\xaf\xea\xc1\x71\x80\x78\xe2\xd3\xbb\x73\xe8\x03\x59\xb9\x4a\xf4\x32\x1e\xfb\x3c\xb5\xca\x7d\x70\x18\xe8\x02\x3f\x8c\x07\x71\xeb\x76\x47\x3c\x40\x66\x08\xd3\x23\x9a\x00\xe0\xc6\x97\x53\x2f\x08\x00\x80\x61\x42\xb2\x8a\x3d\x4e\xd6\xb9\x36\x05\x4e\xbb\x9a\xf2\x9b\x95\x98\x6d\x86\x06\xcb\x88\x74\xee\xc2\xd4\x0a\x51\xd9\x4c\xe3\x6c\x8a\x34\xed\x0a\x92\x7e\x6f\xe4\x72\x92\x73\xb8\xe8\xc2\xee\x58\xee\x0d\x6b\xc1\xa9\x77\x6c\xb6\xeb\x4e\xdd\x49\xbb\x4c\x37\xe8\x17\x8e\x84\x25\xde\x29\xd0\x34\x94\x79\x9e\xd7\x24\x01\x92\x07\x7a\x54\x9e\x86\x98\x29\x8c\xaa\xd3\xe7\x07\x3e\x70\xf6\x49\xba\x55\xbb\xb9\x98\x07\xfe\xb0\x1e\xcf\x33\x5d\x4e\x95\x78\x72\xd1\xcb\x03\x00\x61\xda\x31\xa7\xf9\xfc\x71\x14\x54\x20\x8c\x37\x12\x3e\xc6\x1c\xe5\x39\xae\x10\x1e\xa3\x93\x28\xc3\x46\xf1\x98\x5d\xbf\xdb\x3c\x26\xc5\x3d\x77\xb3\xf1\xba\xf1\x42\xa7\xfb\x95\x58\xb7\x7c\x8b\x39\x59\xb8\xef\x8f\xd8\x07\xa3\xf2\x3a\x62\x84\x21\x5f\xd7\x6c\x1c\xa4\xc1\xfa\xd8\xa2\x8a\x6d\x2b\x7c\xf7\xdf\xeb\xe6\x5d\xc6\xe4\xf6\x14\x82\x14\x41\x48\x69\x74\xcf\x4b\x93\x6a\xb4\x5e\xa3\x73\x95\x6d\xdc\x36\xf1\x6c\x47\x38\xae\x38\x62\xd2\xe1\x8e\x3a\x7a\xe1\x94\x6e\x4b\x1b\xea\x2f\x89\x0b\xaf\xff\x94\x68\xf2\x48\xfc\x28\x9c\x46\xbc\x96\x5c\x2e\x19\x4e\xcc\xeb\xf5\x88\xf8\x4c\x8a\x59\xf7\x6a\xd3\xc4\x3d\x0c\x17\x01\x99\xde\xf5\xe9\x74\xd1\xab\x77\xd6\x16\xf9\x71\xdd\xf1\x2f\xcf\xd3\xb8\x78\xf3\x7c\x88\x85\x82\x65\x77\x27\xf7\x9a\x3a\xf7\xe7\x86\xe1\x14\x7d\xb9\x57\xf7\x52\xb5\xcc\xcd\x25\x7d\x13\xf6\xfb\xff\x9c\xba\xdf\x87\x08\x86\xea\x6e\xb3\x86\x4a\xb2\xd6\x96\xe1\xe2\x8e\x6d\x6d\xcd\xbb\x5c\xaa\x43\xb5\xdd\xcc\xd8\x4e\xd3\x4b\xdb\x3e\x03\x62\xa3\x1f\x33\xe3\x92\x6a\xb7\xea\x5a\x8e\x89\x5f\x2f\x4c\xa6\xcc\xc1\xbd\xe6\x2f\xe2\xd0\x16\x55\x0f\xeb\x24\x56\x20\x4d\xf0\xf0\x20\x59\x54\x38\x46\xbf\x33\xb4\xb8\xe6\x9e\x42\x26\xcd\xaa\x74\x23\xfc\xa3\xb9\x6b\x86\x6c\x8f\xcc\x25\xaf\x73\x4d\x69\xc7\xf9\xa1\x42\x78\x2a\x68\x25\xca\x11\x1e\x5f\xd7\xf6\x5d\x4f\xed\x45\xb8\x1e\xa3\xa3\x22\xcb\xaa\x74\x1b\x4b\x6f\x48\xba\x3e\x5f\xf7\x7c\x91\x11\xcb\x31\x34\x0f\x61\x18\x3d\xd2\x66\xa1\x36\x91\x74\x61\x0c\xe3\xed\xed\x34\x39\x17\x5e\x2f\xfa\xb0\x8b\xeb\xb2\x3e\xfb\xd7\x75\xed\xaa\xd1\x9e\xc1\x61\xc8\x9e\xcf\x1b\xd3\x8e\x04\x3b\xe9\x8d\x7d\xad\x12\xe5\xa4\x22\xee\xb1\xa8\xe6\xec\xd0\x5b\xda\x34\xcb\xbb\xe1\x9d\x78\x92\x6d\x99\x45\xc9\xd1\x43\x43\xf1\xa1\xab\xa2\x69\xd0\x2e\xf1\xd3\x34\x90\xea\x0b\x70\x5b\x87\x15\xbf\xf6\x3a\x52\x71\x1e\x3c\x41\x6d\xec\xc0\xdf\xf5\xc5\x5b\xb4\x7a\xa1\x3a\x53\x5e\x76\x5c\x4b\x54\x4f\x49\x4c\x18\x92\xa8\x8c\xf7\x8c\xc3\xb1\x76\xee\xb4\xd9\x29\x7f\x6c\x96\xa8\x99\xbe\xbe\x7e\xde\x99\xb0\xc6\xff\xef\x8d\xe9\x3f\x01\x00\x00\xff\xff\x4b\x6f\xb7\x22\x3c\x0a\x00\x00")

func CertsCert_filesDcidevpublicPemBytes() ([]byte, error) {
	return bindataRead(
		_CertsCert_filesDcidevpublicPem,
		"../../certs/cert_files/dcidevpublic.pem",
	)
}

func CertsCert_filesDcidevpublicPem() (*asset, error) {
	bytes, err := CertsCert_filesDcidevpublicPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../certs/cert_files/dcidevpublic.pem", size: 2620, mode: os.FileMode(420), modTime: time.Unix(1605579396, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"../../certs/cert_files/dcipublic.pem": CertsCert_filesDcipublicPem,
	"../../certs/cert_files/dcidevpublic.pem": CertsCert_filesDcidevpublicPem,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"certs": &bintree{nil, map[string]*bintree{
				"cert_files": &bintree{nil, map[string]*bintree{
					"dcidevpublic.pem": &bintree{CertsCert_filesDcidevpublicPem, map[string]*bintree{}},
					"dcipublic.pem": &bintree{CertsCert_filesDcipublicPem, map[string]*bintree{}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

