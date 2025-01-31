// Code generated by go-bindata.
// sources:
// templates/json-schema-resourcename.gotpl
// templates/json-schema-restname.gotpl
// templates/json-schema.gotpl
// templates/spec-md.gotpl
// templates/toc-md.gotpl
// specset/.gitignore
// specset/Gopkg.toml
// specset/specs/.regolithe-gen-cmd
// specset/specs/@identifiable.abs
// specset/specs/_api.info
// specset/specs/_parameter.mapping
// specset/specs/_type.mapping
// specset/specs/object.spec
// specset/specs/regolithe.ini
// specset/specs/root.spec
// DO NOT EDIT!

package static

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
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
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

var _templatesJsonSchemaResourcenameGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x6a\xc3\x30\x10\x44\xef\xfe\x8a\x45\xf8\xe8\xe8\x03\x0c\x3d\xf4\xde\xf6\x50\x43\xef\x8a\xbc\x2e\x1b\x1c\x4b\xec\x6e\xa0\x41\xe8\xdf\x8b\x5d\x25\xb5\x5a\xe2\xe3\xcc\xec\x8c\xf5\x52\x03\x00\x60\x94\x74\x46\xd3\x83\x49\x09\xec\x9b\x3b\x23\xe4\x0c\x28\xe1\xc2\x1e\x61\x26\x51\x31\x5d\x49\x5e\xe3\x16\x0c\xc7\x13\x7a\xbd\xa9\x91\x43\x44\x56\x42\x31\x3d\xfc\x74\xae\x5f\x4a\x07\x60\xb7\x7c\x22\xb4\x34\x7e\x0d\x11\x7d\x07\xad\x44\xf4\x34\x91\x77\x4a\x61\x81\xfe\x09\xec\x80\x6a\x87\xbd\x2a\x70\xc8\xb9\x6a\xa1\xe9\x5e\xb1\x7a\xdd\xaa\xe1\x32\x42\x15\x83\x76\x76\x8a\xa2\x1f\xc8\x52\xba\xeb\x35\xfb\xb2\xf9\xcf\xaa\x4c\xc7\x8b\xa2\xdc\x92\xfb\xb9\x95\xc1\x9f\xb3\xd7\x30\xe2\x6c\xdf\x0b\x90\xc2\xa7\x7a\x69\xc5\xc6\x31\xbb\x6b\x41\x73\x37\x49\xf1\x2c\xff\x6e\x36\xab\x65\x9c\x0a\xfc\x47\xc3\x5a\x46\xed\x49\xc2\x62\xaa\x8a\xdf\x3f\xaf\x91\xed\xf0\xe4\x26\x37\xdf\x01\x00\x00\xff\xff\x0c\x51\x73\xb6\xea\x01\x00\x00")

func templatesJsonSchemaResourcenameGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesJsonSchemaResourcenameGotpl,
		"templates/json-schema-resourcename.gotpl",
	)
}

func templatesJsonSchemaResourcenameGotpl() (*asset, error) {
	bytes, err := templatesJsonSchemaResourcenameGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/json-schema-resourcename.gotpl", size: 490, mode: os.FileMode(420), modTime: time.Unix(1552522176, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJsonSchemaRestnameGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xb1\x6a\xc4\x30\x10\x44\x7b\x7f\xc5\x20\x5c\xfa\xf4\x01\x86\x14\xe9\x93\x14\x31\xa4\xf7\xd9\xe3\xa0\xc3\xb1\x84\x76\x0f\x12\x84\xfe\x3d\xd8\xf1\x1d\x56\x9a\xdb\x72\x76\x66\x96\xb7\xa9\x02\x00\xa3\x4e\x67\x9a\x16\x26\x25\xd8\xb7\xfe\x8b\xc8\x19\x91\xe2\xaf\x71\xa0\x98\x66\x77\xfd\x84\xcd\xe4\xcf\x17\x0e\x7a\x53\x43\xf4\x81\x51\x1d\xc5\xb4\xf8\xeb\x5b\x27\xa5\x13\x62\xbf\x7c\x12\xb5\x1b\xbf\xbb\xc0\xa1\x41\x2d\x81\x83\x9b\xdc\xd0\xab\xf3\x0b\xda\x27\xd8\x8e\x6a\xbb\xa3\x2a\x38\xe5\x5c\xb4\xb8\xe9\x5e\xb1\xee\x9a\x55\xe3\x32\xa2\xb0\xa1\x9e\x7b\xa5\xe8\x07\xa3\xec\xdd\xe5\x35\xfb\xb2\xed\x9f\x55\xa3\x3b\x5f\x95\x72\x73\x1e\xcf\xad\xfc\xff\x62\xaf\x7e\xe4\x6c\xdf\x29\xba\xff\xa5\xa0\xdc\x42\x75\xe4\xb4\x3f\xef\x51\xd8\x5e\xc4\x2f\xe6\x1e\x2f\x41\x0f\x50\xb9\xca\xd5\x6f\x00\x00\x00\xff\xff\x3e\x41\x49\xb4\x9c\x01\x00\x00")

func templatesJsonSchemaRestnameGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesJsonSchemaRestnameGotpl,
		"templates/json-schema-restname.gotpl",
	)
}

func templatesJsonSchemaRestnameGotpl() (*asset, error) {
	bytes, err := templatesJsonSchemaRestnameGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/json-schema-restname.gotpl", size: 412, mode: os.FileMode(420), modTime: time.Unix(1552522171, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJsonSchemaGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x4f\x6f\xe3\xb8\x0f\xbd\xe7\x53\x10\x42\x0e\xbf\xdf\x20\xc9\xdc\x0b\xe4\x30\x3b\xdb\x3d\x2c\xfa\x67\xd0\x2e\xe6\xb2\xe8\x81\x89\x99\x44\x53\x5b\xf6\x48\xf4\x34\x41\xd6\xdf\x7d\x21\x59\xb2\x65\x5b\x69\xa7\x8b\xc5\xde\x5c\xea\x3d\xf2\x91\xa2\x18\xf6\x7c\x5e\x42\x46\x3b\xa9\x08\x44\x85\xda\xd0\x92\x4f\x15\x09\x58\x36\xcd\xcc\x9e\xc9\x1d\xd0\x77\x58\x81\x60\x59\x78\xb3\x70\x88\x2b\x10\x86\xb5\x54\x7b\xb1\x98\x89\x5d\xa9\x0b\x64\x6b\xcb\x90\x69\xe9\xb0\x8e\x4e\xb9\xa1\xde\x87\x27\x24\xbd\x24\xe0\x9b\xb2\xcc\x09\xd5\x08\x1f\xac\x09\x82\x54\x4c\x7b\xd2\x23\x42\xb0\x26\x08\xbb\xbc\x44\x1e\xc1\x55\x5d\x6c\xd2\xe8\x72\xf3\x8d\xb6\xdc\xd7\xc6\x9d\x3a\xee\x5c\xd3\xce\x72\xcf\x67\x58\x41\xd3\xac\xbe\x99\x32\x08\x54\x59\x8f\xf7\xdf\xb3\xf3\x0c\x00\xe0\x7c\x86\x79\x8e\x4c\x86\xbf\x92\x36\xb2\x54\x70\xb5\x86\xd5\x63\x45\xdb\xd5\x8d\x33\x7f\x62\xd6\x72\x53\x33\x99\x00\xb0\x6c\x4b\x15\x2c\x39\xa7\x10\xd1\x51\x6e\xcb\x8c\xf2\xd5\xb5\x62\xc9\xa7\x3b\x2c\x08\x9a\x46\x2c\x3c\xd8\x67\xe6\xf5\x7b\x6b\xa5\xcb\x8a\x34\x4b\x32\xe2\x0a\x82\xa2\x25\x68\x54\x7b\x82\xb9\xcc\x8e\x36\xfc\x02\xe6\xc8\xac\x7b\x65\xd7\xc7\xaa\x34\x94\xf5\xd2\xc6\x39\x04\x89\xbe\x79\x82\x23\x6b\x5f\x84\x22\x84\x2c\x6c\x05\xac\xfb\xd5\xe7\x52\xfd\x20\xcd\x94\x05\xe5\x41\x91\x40\x75\xba\xb7\xa5\xfd\xd3\xfd\x19\x39\x56\x25\xc3\xea\x81\xbe\xd7\x52\x53\xe7\xd1\x9d\x43\x74\x95\x79\x2e\xa0\x59\x0c\xb8\x91\x00\x67\x1a\x1c\x7e\xfc\x00\xeb\xf5\x7a\x0d\x7f\x9c\x2a\x6a\xbf\x3e\x7c\xec\x52\x0a\xa0\xb9\xf5\x6f\x4b\xb2\x6d\x65\x3b\x70\x9b\x88\xfb\x1c\xe2\x5d\x15\xdc\xe1\xa7\x3c\x2f\x5f\x28\xfb\x7c\x28\xe5\x96\x4c\x2c\x42\x90\xaa\x8b\x41\x96\xd3\xeb\x68\x69\x0b\x98\x6f\xdd\x87\x8d\x9f\x74\x1b\x47\x1f\x5d\x44\x0b\x71\x57\x61\x85\x85\x7e\x8c\xd1\xee\x4e\x7c\x84\xa6\x11\x13\x4f\x63\xce\xd3\xb0\xb6\xfd\x7b\x69\x8b\x24\xe6\x74\x64\xd2\x0a\xed\x3d\x8c\x75\x81\x7d\xfa\xd5\x6f\x52\x1b\xbe\xa1\x1f\x94\xff\xa2\x71\xfb\x4c\x6c\xc6\x4d\xe1\x8a\x3a\xba\x83\x44\xa4\x5c\x1a\x1e\x44\xe9\xfa\x00\xb5\xc6\x93\x18\x77\x50\x1b\xe5\xb1\xde\x78\xff\x7d\x9b\x08\xc9\x54\xf4\xcf\x22\x52\xcc\x54\x54\xb6\xdf\x87\xc3\x72\xec\xa9\x63\x35\xaf\x76\x5e\x3a\x8d\x02\xab\x74\x16\xfe\xf9\xbe\x9e\xc6\x72\x90\x07\x66\x99\x64\x59\x2a\xcc\xbf\x4c\x5f\xfb\xbf\x96\x96\x1f\x83\x6f\xc7\x65\x5d\xd3\x3f\xa8\x88\x1d\xae\x43\xd8\x7b\x05\x77\x8e\x7f\xca\x0b\xa7\xd8\xad\xd2\xe4\xb0\xb8\xe7\x03\x69\xe8\xe7\xe9\xc5\xc1\x31\x9d\x04\xa8\xcd\xa8\x74\x15\xb2\x7d\x30\x7e\xb6\xfb\x09\xf3\x40\x7b\x3a\x56\x29\x72\x6b\x8a\xe6\xa0\x98\xc8\xbe\xa8\xe2\x16\x8f\x37\xa4\xf6\x7c\x18\x49\x28\x82\xdd\xb6\xca\x79\x02\x4e\x54\xe6\x72\x08\xa9\xd2\x21\x82\x7d\x10\xa2\x03\xbf\x2b\x04\x1e\xbf\x62\x5e\x8f\x5b\xb0\xc0\xa3\x2c\xdc\x50\x8d\x53\x68\x91\xef\xcc\x20\xe9\x5e\xaa\x89\xfb\x80\x7c\xd3\x3d\x08\x31\x7d\x4e\x4f\xb3\xd9\xb8\xb1\x6e\x89\x31\x43\xc6\x2b\xb8\x2b\x19\x10\x7e\x7f\xbc\xbf\x03\xb3\x3d\x50\x81\xf0\x4c\xa7\x97\x52\x67\xb0\xa9\x19\x9e\xa9\x62\xd8\x95\x1a\xa4\x6a\x37\x31\xfb\x43\x3c\xe9\xc2\x41\x5a\x5d\xc7\x74\x69\xd9\x77\xd6\xda\xfc\x4b\x0d\x52\xe2\xa8\xb5\x21\x03\xa8\xc0\x4d\x54\x40\x06\xd1\x91\x3a\x45\xc8\xc0\x07\x82\x0a\x35\x29\x1e\x0b\x88\x8b\x31\x12\x84\xd9\xbd\xca\x4f\x43\x41\xad\x2d\x1e\x1d\x17\x9c\xd8\x75\xe0\x7f\xd2\xdc\xc9\xdc\xfb\xfb\x95\x76\x58\xe7\xec\x6e\xe4\xff\xb1\xd3\x2c\x3a\x88\xaf\x2f\x26\xc0\x5f\x60\x57\xb8\x47\xb7\x9a\xca\xdd\x09\x2e\x66\x10\x5d\xe5\x10\x12\x8c\x3e\xec\x85\x89\xb8\xc3\xdc\x90\x87\x44\xe5\x1f\x6e\x3b\x2f\x92\x0f\xef\x5d\xbf\x02\x77\xbe\xb3\x3f\xaf\x76\x59\x70\x57\x3a\x3e\xf6\xeb\x05\x06\x7f\x6e\xd1\x1b\x37\x70\xb8\x23\x87\x18\x74\x4e\x6a\xd1\xb0\x37\xe1\xa3\xbe\xb6\x68\x44\xe2\xd6\x6d\x19\xd2\xbb\x48\x1f\x78\xb2\x24\xbe\xf9\x88\xdf\xb2\x3d\x2d\xfe\x8b\x37\xd7\x77\x88\x98\x17\x76\x4b\xb7\xfa\x7f\x76\x77\x6f\x19\x0f\x64\xf8\x02\x2b\x1c\xa5\x38\x65\xad\xb7\x74\x99\xd7\x1d\x4f\xb8\x5f\x70\xfb\x8c\xfb\x14\xcd\x9f\xc4\x0c\xbb\xe1\xa9\xcc\x4c\xda\xb6\x5f\x5d\xaf\x1d\x62\x01\x1e\xda\xff\x37\xe1\x13\x6f\x1d\x24\x67\x70\xc7\xbe\xd8\x4c\xae\x4b\xbc\xe3\xd7\x9a\xe2\x69\xd6\xcc\xfe\x0e\x00\x00\xff\xff\xda\x1d\x70\x5c\xee\x0e\x00\x00")

func templatesJsonSchemaGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesJsonSchemaGotpl,
		"templates/json-schema.gotpl",
	)
}

func templatesJsonSchemaGotpl() (*asset, error) {
	bytes, err := templatesJsonSchemaGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/json-schema.gotpl", size: 3822, mode: os.FileMode(420), modTime: time.Unix(1549567555, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSpecMdGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xcd\x6a\xc3\x30\x10\x84\xef\x7e\x8a\x05\xf7\x5a\x3d\x40\xa1\x87\x42\x72\xeb\x0f\xb4\x25\xd7\x78\x2b\x6f\x1a\xb5\xb6\x25\xb4\x1b\x48\x10\x7a\xf7\x22\xd9\xb1\x1c\xd3\x9b\x3c\xcc\xcc\xe7\xdd\xad\xeb\x1a\x42\x00\xf5\xe1\x48\xab\x17\xdb\x52\xa7\xb6\x83\x18\xb9\xbc\x62\x4f\x10\x63\x15\x02\xdc\x75\x28\xc4\xb2\x23\xcf\xc6\x0e\xf0\xf0\x38\xd9\x9f\xb3\xfc\x24\xe2\xcd\xd7\x49\x88\xaf\x86\x31\x25\xde\xf4\xec\x50\xd3\x4d\xf9\x86\x58\x7b\xe3\x64\xf2\xe5\x7a\x3a\x63\xef\x3a\x4a\xc5\xd7\x67\x8e\xac\xc1\x63\xaf\x39\x94\x44\x8c\x55\x9d\x26\xd8\x8e\xdf\x55\xd5\x34\xcd\x0f\xdb\xe1\xa6\x36\xc6\x24\x27\x89\x86\x76\x86\x5a\xc7\x09\x68\x1d\x79\x4c\x7f\xc3\x13\x53\xbd\x53\x37\x0a\x47\xe3\x92\x48\xb2\x00\x5b\xc7\x57\xe6\xec\x2b\x7d\xa3\x6f\x49\xc1\x79\x39\x65\x6d\xdb\xb3\xb3\x4c\x6d\xd9\xdb\xbf\x73\xde\x67\xde\x22\x9f\x2a\x33\xb8\x04\x13\xc2\xe3\xf0\x4d\x6b\x63\x9d\x8d\x4d\x3a\xec\x74\xc7\x26\x5d\x59\x2e\x8e\xde\x0e\xa0\x20\xc6\x10\x40\x1f\xd1\xa3\x16\xf2\x86\xc5\x68\xce\xf2\x34\xa7\xda\x90\xf3\xa4\x51\x28\x8f\xb2\xff\x3c\x1a\x86\x19\x01\x86\xa1\x9d\x0d\x7b\xb5\x18\xfa\xf6\xee\xab\x63\x87\x00\x3d\xfe\xd2\x86\x0e\x78\xea\x64\x87\xdd\x89\x66\x68\xc9\x4f\xaf\xbf\x00\x00\x00\xff\xff\xda\xfc\xa8\x5e\x99\x02\x00\x00")

func templatesSpecMdGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSpecMdGotpl,
		"templates/spec-md.gotpl",
	)
}

func templatesSpecMdGotpl() (*asset, error) {
	bytes, err := templatesSpecMdGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/spec-md.gotpl", size: 665, mode: os.FileMode(420), modTime: time.Unix(1550885181, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTocMdGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x56\xa8\xae\x56\xd0\x73\x2f\xca\x2f\x2d\xf0\x4b\xcc\x4d\x55\xa8\xad\xe5\x02\x04\x00\x00\xff\xff\x0a\x31\xfc\xe1\x14\x00\x00\x00")

func templatesTocMdGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesTocMdGotpl,
		"templates/toc-md.gotpl",
	)
}

func templatesTocMdGotpl() (*asset, error) {
	bytes, err := templatesTocMdGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/toc-md.gotpl", size: 20, mode: os.FileMode(420), modTime: time.Unix(1562023324, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetGitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xca\x41\x0a\x84\x30\x0c\x05\xd0\xfd\x3f\x4a\x61\x72\xa1\x61\x16\x43\xf2\x0d\xc5\x9a\x48\x2d\x45\x6f\x2f\xe2\xe6\xad\x5e\x91\xfd\xfa\x6a\xda\x0f\x45\xe8\xfe\xfa\xa9\xb1\x24\x34\x8d\xce\xc0\x53\x14\xde\xc9\x51\xc3\xb7\x34\xb6\x03\x93\x61\xd9\x51\xa4\xa5\xae\xd0\x9c\xec\x7f\xa7\x8c\x73\xe0\x0e\x00\x00\xff\xff\x66\xf6\x5a\x96\x53\x00\x00\x00")

func specsetGitignoreBytes() ([]byte, error) {
	return bindataRead(
		_specsetGitignore,
		"specset/.gitignore",
	)
}

func specsetGitignore() (*asset, error) {
	bytes, err := specsetGitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/.gitignore", size: 83, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetGopkgToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x2e\x28\x2a\xcd\x4b\x8d\xe5\x52\x50\x48\xcf\xd7\x2d\x49\x2d\x2e\x29\x56\xb0\x55\x28\x29\x2a\x4d\xe5\x52\x50\x28\xcd\x2b\x2d\x4e\x4d\xd1\x2d\x48\x4c\xce\x4e\x4c\x4f\x85\x4b\x00\x02\x00\x00\xff\xff\x34\x36\xb3\x89\x33\x00\x00\x00")

func specsetGopkgTomlBytes() ([]byte, error) {
	return bindataRead(
		_specsetGopkgToml,
		"specset/Gopkg.toml",
	)
}

func specsetGopkgToml() (*asset, error) {
	bytes, err := specsetGopkgTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/Gopkg.toml", size: 51, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsRegolitheGenCmd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x41\x4e\xc5\x30\x0c\x44\xf7\x3e\xc5\xf0\xf9\x0b\x40\x4a\x22\x0e\x00\x77\x49\x63\x37\x8d\x48\xe2\x2a\x0d\x88\x45\x0f\x8f\x5a\x5a\xca\xc2\x0b\x6b\xde\xb3\xe7\xf1\xc1\x0d\xa9\xba\xc1\x2f\x13\x95\x0f\x4e\x0d\x66\x86\xb5\x8e\x35\x10\xfb\xee\xdf\xee\x4f\x4d\xa2\x82\x35\xc0\x30\x2c\xd6\x15\xf2\x9d\x3a\x5e\x9f\x49\xc2\xa4\xb8\xdd\x37\xec\x86\xf7\xc3\xda\xe6\xb3\x48\xed\xbe\x27\xad\xb6\x30\x05\x86\xfd\xe7\x91\x64\x89\x52\x31\x6a\x66\x69\xdb\xd1\x65\x96\xb0\xc0\x28\x82\xf2\x1e\x5d\xec\xd5\x29\x6a\xf6\x35\x52\x2b\x30\x6d\x84\x75\xbf\xbb\x7b\xb1\x51\xa9\x7c\x9d\xaa\x93\x2c\xfb\xf3\xbc\x27\x7f\xdc\xe9\x1d\x18\xfd\x04\x00\x00\xff\xff\x25\x5b\x5e\xc5\xf7\x00\x00\x00")

func specsetSpecsRegolitheGenCmdBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsRegolitheGenCmd,
		"specset/specs/.regolithe-gen-cmd",
	)
}

func specsetSpecsRegolitheGenCmd() (*asset, error) {
	bytes, err := specsetSpecsRegolitheGenCmdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/.regolithe-gen-cmd", size: 247, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsIdentifiableAbs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\xb1\x0a\xc2\x40\x0c\xc6\xf1\xbd\x4f\xf1\x81\xb3\x82\xeb\x6d\x82\x8b\x4f\x21\xd7\xde\xd7\x1a\xa9\x97\x92\x4b\xc5\xbe\xbd\xf4\x06\x6d\x97\x10\x7e\xfc\x03\x39\xe0\xe2\x6e\xd2\xce\xce\xd2\xc4\xdf\x1a\x1a\xe0\x7d\x5e\xe7\x11\x39\xbe\x18\x70\xbb\x36\x00\x90\x58\x3a\x93\xc9\x45\xf3\x6a\x90\x02\x7f\x10\x92\x98\x5d\x7a\xa1\x41\xfb\x2a\xda\x3e\xd9\xf9\xa9\x1e\xf9\x32\x31\xa0\xb8\x49\x1e\x2a\xf0\x33\x69\x61\x0a\x70\x9b\x59\xa5\xb8\xda\x0e\x8c\x31\xdd\x35\x8f\xcb\xc6\xe2\xec\x3a\x30\xd3\xa2\xef\xda\x5e\x46\xa7\xc5\x76\xe4\x06\xff\x2f\x6d\x50\x2d\xed\xc2\x6f\x00\x00\x00\xff\xff\x62\x4c\x45\x33\xff\x00\x00\x00")

func specsetSpecsIdentifiableAbsBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsIdentifiableAbs,
		"specset/specs/@identifiable.abs",
	)
}

func specsetSpecsIdentifiableAbs() (*asset, error) {
	bytes, err := specsetSpecsIdentifiableAbsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/@identifiable.abs", size: 255, mode: os.FileMode(420), modTime: time.Unix(1533330802, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecs_apiInfo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x28\x4a\x4d\xcb\xac\xb0\x52\x48\x2c\xc8\xe4\x2a\xca\xcf\x2f\xb1\x52\x00\x91\x5c\x65\xa9\x45\xc5\x99\xf9\x79\x56\x0a\x86\x5c\x80\x00\x00\x00\xff\xff\x0c\x97\x42\xd8\x22\x00\x00\x00")

func specsetSpecs_apiInfoBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecs_apiInfo,
		"specset/specs/_api.info",
	)
}

func specsetSpecs_apiInfo() (*asset, error) {
	bytes, err := specsetSpecs_apiInfoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/_api.info", size: 34, mode: os.FileMode(420), modTime: time.Unix(1528494787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecs_parameterMapping = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8a\x31\x0e\xc2\x40\x0c\x04\xfb\xbc\x62\x0b\x5a\x78\xc0\x49\xf9\x05\x3d\x32\x64\x45\x2c\x39\xbe\x8b\x7d\x01\xf2\x7b\x44\x44\x37\x9a\x99\xd3\xba\x31\x76\xb9\x1b\xcb\x00\xd0\x7b\x28\xf3\x87\x67\xb8\x2c\x2c\x58\x07\x00\x98\x98\x8f\xd0\xd6\xb5\x7a\xc1\x75\xd6\x84\x26\xc4\xc1\x8f\x2c\xcd\x78\x39\xa6\xbe\x37\x16\x64\x0f\xf5\xe7\x21\xfe\xf5\xf6\x12\xdb\x58\x30\xd3\xac\x62\x1c\xf1\xae\x61\xd3\xf0\x0d\x00\x00\xff\xff\x6d\xff\xeb\x46\x7b\x00\x00\x00")

func specsetSpecs_parameterMappingBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecs_parameterMapping,
		"specset/specs/_parameter.mapping",
	)
}

func specsetSpecs_parameterMapping() (*asset, error) {
	bytes, err := specsetSpecs_parameterMappingBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/_parameter.mapping", size: 123, mode: os.FileMode(420), modTime: time.Unix(1551136779, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecs_typeMapping = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x8f\x8e\x8d\x8e\x2d\x2e\x29\xca\xcc\x4b\x57\xb7\xe2\x52\x50\x48\xcd\x49\xcd\x4d\xcd\x2b\x49\xcc\x01\x71\x14\x14\x4a\x2a\x0b\x52\xad\x14\x90\x15\x81\x85\x33\xf3\x32\x4b\x50\x84\xab\x6b\x41\x12\x59\xc5\xf9\x79\xc5\xc9\x19\xa9\xb9\x89\xc8\xba\x6b\x74\xc1\x1c\x05\x85\x6a\x28\xad\xa0\xa0\x04\x92\x51\xb2\x52\x50\x4a\x2c\x2a\x4a\xac\x54\xd2\x41\x48\x64\x96\xa4\xe6\x16\x2b\x59\x21\xa9\xc5\xa3\x1a\x87\x7a\x24\x1d\x10\xd7\x29\x21\x49\xd6\x72\xa1\xb3\x6a\xb9\x00\x01\x00\x00\xff\xff\x2c\xe3\x9a\x72\x05\x01\x00\x00")

func specsetSpecs_typeMappingBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecs_typeMapping,
		"specset/specs/_type.mapping",
	)
}

func specsetSpecs_typeMapping() (*asset, error) {
	bytes, err := specsetSpecs_typeMappingBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/_type.mapping", size: 261, mode: os.FileMode(420), modTime: time.Unix(1551136705, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsObjectSpec = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x31\x4f\xc3\x30\x10\x85\x77\xff\x8a\x27\x75\xe8\x04\x12\xab\x27\x90\x90\x98\x10\x4b\x99\x91\x1b\xbf\x26\x86\xc4\xb6\x7c\x67\xd4\xfe\x7b\xe4\xa4\x8a\x20\x74\x39\x59\xdf\x7d\xf7\xce\xf6\x0e\xaf\xc9\x73\x34\x53\xab\xd6\x00\x85\xa2\x1f\xd1\x4d\xb4\x48\xc7\x4f\x76\xba\xb0\x54\x4b\xc7\x3f\x5c\x0c\xc0\xa8\x41\x2f\x57\xfc\xb6\xe2\xec\xba\x2f\xd7\xd3\xc2\xf3\xe4\xea\xd8\x22\xfa\x92\x6a\xb6\xe8\x52\xa1\x01\x3c\xa5\x2b\x21\x6b\x48\xd1\xe2\x30\x04\x41\x10\x14\x17\x7d\x9a\xae\xe9\xf7\x6d\x86\xda\x6e\xb4\xd1\x5f\xa8\x02\x1d\xf8\x4b\xac\xd9\x3b\xe5\x0d\xf7\x7d\x6e\x6c\x75\xcf\x91\x37\xf5\xe7\xb9\xb1\xd5\x79\x56\x46\x2f\xcd\xbf\xc3\xfe\x31\xf8\xf6\xea\x53\x70\xc7\x91\x7b\x63\x76\x78\x52\x2d\xe1\x58\x95\x62\xdc\x7a\x6c\xf6\xf7\xc3\x32\xb3\x7c\x4f\xab\xff\x57\x1e\x06\xce\x1d\xa4\xd3\x66\x2d\xa0\x97\x4c\x0b\xd1\x12\x62\x3f\x03\x9e\x73\x12\x7a\x0b\x2d\x75\xc9\x12\x4d\x65\x05\x3f\x01\x00\x00\xff\xff\x1d\xac\xcb\xd0\xcc\x01\x00\x00")

func specsetSpecsObjectSpecBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsObjectSpec,
		"specset/specs/object.spec",
	)
}

func specsetSpecsObjectSpec() (*asset, error) {
	bytes, err := specsetSpecsObjectSpecBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/object.spec", size: 460, mode: os.FileMode(420), modTime: time.Unix(1551136826, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsRegolitheIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4e\x91\x13\x54\x3d\x40\xc1\x7b\x0c\x83\x84\x1a\x67\x0a\x93\xa6\x64\x52\x4b\x6f\xef\xa2\x2e\x74\xf7\xe1\xff\xff\x16\xe3\x4d\x8f\xec\x3b\xaf\x50\x4d\x9f\x2d\xf9\xa3\x90\x30\x46\x94\x41\x35\x43\xd2\x3a\x2c\x6f\xbb\x63\xc4\xa1\xcd\x92\x4a\xa5\x32\x00\x16\x37\x2a\xe7\x4b\x4d\xd8\x56\xf8\xfb\x34\x3b\x30\x62\xef\x3d\xc8\xf8\xee\x43\x61\xbf\xcc\x96\x9a\xef\x6a\x93\x03\x16\xca\xc7\xcc\xf7\x1f\x3e\x24\x15\x78\xb3\x9d\x59\x0b\x46\xbc\x85\x2b\x7c\x02\x00\x00\xff\xff\x0e\x59\x87\x8e\xaa\x00\x00\x00")

func specsetSpecsRegolitheIniBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsRegolitheIni,
		"specset/specs/regolithe.ini",
	)
}

func specsetSpecsRegolitheIni() (*asset, error) {
	bytes, err := specsetSpecsRegolitheIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/regolithe.ini", size: 170, mode: os.FileMode(420), modTime: time.Unix(1533330787, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _specsetSpecsRootSpec = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x8e\xc2\x30\x0c\x45\xf7\x39\x85\xa5\xae\x67\x0e\x90\xed\xac\x67\xd3\x0b\xa0\x90\x7e\x4a\xa0\x8d\x2b\xdb\x05\x71\x7b\x14\x52\xa0\x20\xd8\x44\xd1\xf7\xfb\x7a\x76\x43\xff\xdc\x61\x70\x63\x79\xbd\x23\x12\xa8\x6d\x72\x18\xe1\x49\x98\xad\x26\x3c\x4b\xc4\x6b\x8a\x6c\xc9\x2e\x4b\xd6\xd6\x6c\x0a\xf1\x18\xfa\x27\xd3\x0b\xcf\x93\xa7\xc8\x02\x47\xd4\x41\xa3\xa4\xc9\x12\xe7\x4a\x10\x6f\x0f\x88\xf6\x5b\x48\x58\x91\xbf\x41\x3d\x4c\xc9\xf6\x58\x81\xa5\xe7\xc9\x64\x86\x73\x0d\xb5\x18\x42\x41\xd5\xc9\xfd\xe7\xdd\xcf\xfa\x86\xda\xfc\x6a\x68\x61\x92\x70\x42\xd5\x0c\x49\x8d\x78\xb7\x94\xb4\xf8\xa2\x20\x18\x3e\x34\xff\x6e\x03\xa5\x40\x19\xe7\xc7\x82\xd7\x00\x00\x00\xff\xff\x65\x54\xa4\xfe\x4f\x01\x00\x00")

func specsetSpecsRootSpecBytes() ([]byte, error) {
	return bindataRead(
		_specsetSpecsRootSpec,
		"specset/specs/root.spec",
	)
}

func specsetSpecsRootSpec() (*asset, error) {
	bytes, err := specsetSpecsRootSpecBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "specset/specs/root.spec", size: 335, mode: os.FileMode(420), modTime: time.Unix(1551136643, 0)}
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
	"templates/json-schema-resourcename.gotpl": templatesJsonSchemaResourcenameGotpl,
	"templates/json-schema-restname.gotpl": templatesJsonSchemaRestnameGotpl,
	"templates/json-schema.gotpl": templatesJsonSchemaGotpl,
	"templates/spec-md.gotpl": templatesSpecMdGotpl,
	"templates/toc-md.gotpl": templatesTocMdGotpl,
	"specset/.gitignore": specsetGitignore,
	"specset/Gopkg.toml": specsetGopkgToml,
	"specset/specs/.regolithe-gen-cmd": specsetSpecsRegolitheGenCmd,
	"specset/specs/@identifiable.abs": specsetSpecsIdentifiableAbs,
	"specset/specs/_api.info": specsetSpecs_apiInfo,
	"specset/specs/_parameter.mapping": specsetSpecs_parameterMapping,
	"specset/specs/_type.mapping": specsetSpecs_typeMapping,
	"specset/specs/object.spec": specsetSpecsObjectSpec,
	"specset/specs/regolithe.ini": specsetSpecsRegolitheIni,
	"specset/specs/root.spec": specsetSpecsRootSpec,
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
	"specset": &bintree{nil, map[string]*bintree{
		".gitignore": &bintree{specsetGitignore, map[string]*bintree{}},
		"Gopkg.toml": &bintree{specsetGopkgToml, map[string]*bintree{}},
		"specs": &bintree{nil, map[string]*bintree{
			".regolithe-gen-cmd": &bintree{specsetSpecsRegolitheGenCmd, map[string]*bintree{}},
			"@identifiable.abs": &bintree{specsetSpecsIdentifiableAbs, map[string]*bintree{}},
			"_api.info": &bintree{specsetSpecs_apiInfo, map[string]*bintree{}},
			"_parameter.mapping": &bintree{specsetSpecs_parameterMapping, map[string]*bintree{}},
			"_type.mapping": &bintree{specsetSpecs_typeMapping, map[string]*bintree{}},
			"object.spec": &bintree{specsetSpecsObjectSpec, map[string]*bintree{}},
			"regolithe.ini": &bintree{specsetSpecsRegolitheIni, map[string]*bintree{}},
			"root.spec": &bintree{specsetSpecsRootSpec, map[string]*bintree{}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"json-schema-resourcename.gotpl": &bintree{templatesJsonSchemaResourcenameGotpl, map[string]*bintree{}},
		"json-schema-restname.gotpl": &bintree{templatesJsonSchemaRestnameGotpl, map[string]*bintree{}},
		"json-schema.gotpl": &bintree{templatesJsonSchemaGotpl, map[string]*bintree{}},
		"spec-md.gotpl": &bintree{templatesSpecMdGotpl, map[string]*bintree{}},
		"toc-md.gotpl": &bintree{templatesTocMdGotpl, map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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

