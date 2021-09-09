// Code generated for package template by go-bindata DO NOT EDIT. (@generated)
// sources:
// index.tmpl.html
package template

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexTmplHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\xdf\x6f\xdb\x46\x0c\x7e\xef\x5f\xc1\x1d\xf6\xb0\x61\xb0\x14\xd7\x41\xdb\x05\x92\x01\x23\x59\xb1\x15\x29\x92\xd6\xce\xda\xee\xed\xa4\xa3\xa4\x5b\xef\x87\x7a\x47\xb9\x71\x0c\xfd\xef\xc3\xc9\x72\x62\x3b\xb6\x8b\x35\x7b\x9b\x5f\x64\xf1\xc8\x8f\x1f\xf9\x51\xa2\x92\x1f\x2e\xae\xce\x67\x9f\xae\x7f\x83\x8a\xb4\x1a\x3f\x4b\xc2\x05\x14\x37\x65\xca\xd0\xb0\x60\x40\x2e\xc6\xcf\x00\x00\x12\x8d\xc4\x21\xaf\xb8\xf3\x48\x29\xbb\x99\xbd\x1e\xbc\x62\x9b\x47\x15\x51\x3d\xc0\x2f\x8d\x9c\xa7\xcc\x61\xe1\xd0\x57\x0c\x72\x6b\x08\x0d\xa5\x6c\x74\xc2\xe2\xde\x9b\x24\x29\x1c\x5f\x72\x81\xda\x1a\x49\xd6\x25\xf1\xca\xb4\x3a\x56\xd2\x7c\x06\x87\x2a\x65\x9e\x16\x0a\x7d\x85\x48\x0c\x2a\x87\x45\xca\x42\x0e\x7f\x16\xc7\x9e\x78\xfe\xb9\xe6\x54\x45\x99\xb5\xe4\xc9\xf1\x3a\x17\x26\xca\xad\x8e\xef\x0d\xf1\x69\x34\x8a\x86\x71\xee\xfd\x83\x2d\xd2\xd2\x44\xb9\xf7\x0c\xa4\x21\x2c\x9d\xa4\x45\xca\x7c\xc5\x47\xaf\x4e\x07\x65\x79\xb5\x78\x7f\x22\x3f\x9e\x67\x6f\xdf\xcd\x47\x1f\x65\xad\xf9\xe8\xf4\xed\xc5\x2f\xe2\xf7\x78\x58\xbc\x7b\xf9\xea\x34\xfe\xfb\x45\xfe\x29\x96\x6f\x66\xef\x6e\xae\xaa\xfc\x83\x7b\x79\xfb\xeb\x9b\xb9\x7d\x7f\x3b\x7b\xfe\xf6\xaf\xaf\xc3\x19\x83\xdc\x59\xef\xad\x93\xa5\x34\x29\xe3\xc6\x9a\x85\xb6\x8d\x0f\x6d\x8c\x57\x7d\x4c\x32\x2b\x16\x7d\x99\x42\xce\x21\x57\xdc\xfb\x94\x85\x26\x71\x69\xd0\x0d\x0a\xd5\x48\xd1\x77\x75\xd7\xcb\xd9\xaf\x1b\x27\x8f\x31\xd4\x40\x8b\x1d\x87\xce\xa9\x1a\x8e\x27\x9f\xa9\x41\xa5\x10\x2e\x38\xa1\x49\xe2\x6a\x38\x4e\x32\xb7\xcf\x77\x34\xbe\xb6\x5f\xd1\x15\xf6\x16\x2e\x64\x51\xa0\x43\x73\x77\x06\xcb\x65\xd4\x99\x5f\xdb\xdb\xe8\x03\x27\x6a\x5b\x08\x97\x24\xae\x46\xfb\x41\xfe\x44\x97\x49\x23\x1a\x53\xc2\x5d\xa3\x61\xd2\x90\x0d\x28\x20\x0b\xc0\x2f\x10\x9d\x57\xdc\x95\xe8\xa2\x73\xee\xa6\xc4\xa9\xf1\x30\x6c\x5b\x23\xf3\x8a\x60\xb9\x44\x23\xda\x76\x8e\x2e\x6b\x8c\xe8\xc8\xee\xc9\xb1\x82\xfa\xa9\xa4\xbd\x58\x3f\x43\xdb\xee\xa5\x15\x66\xce\x93\xb3\xba\x23\x73\x1f\x3a\xd1\x35\x3a\x5e\x22\xb4\xed\xe4\x70\x4d\x21\x58\xa1\xf4\xd4\x98\x72\x3b\x3e\xe4\xf5\xd1\xb4\xd1\x97\x96\x0b\x68\xdb\x83\x9c\xd1\x88\xbd\xd4\xf6\x69\x31\x25\x6e\x44\x97\xa7\x70\x56\xdf\x18\x79\x0b\x0f\x22\xcc\xa4\x46\x4f\x5c\xd7\x70\xb2\x0b\x98\xc4\x42\xce\x8f\xcf\x89\xd7\xfb\xe6\xe4\xe8\xa8\xad\x7f\xa1\x09\x28\x8d\x27\x54\xaa\x31\x25\x9a\xb3\xc7\x40\x8f\x09\x74\xe6\xc2\x3a\x0d\x3c\x27\x69\x4d\xca\x18\x68\xa4\xca\x8a\x94\x5d\x5f\x4d\x67\x07\x92\x6d\x52\x0a\xd1\x83\xd2\xd9\xa6\x3e\xe0\xdc\x05\x28\x9e\xa1\x82\xc2\xba\x94\x71\x5d\x5f\xcd\xd1\x39\x29\x90\x3d\x08\x0f\xd2\xc0\x24\x89\x3b\xbf\x23\x38\x1e\x15\xe6\xb4\x95\x3b\x3c\xa4\xce\x2a\x06\x52\x6c\x83\x83\xe1\x1a\x77\xf2\x1d\x44\x86\xd5\x20\x38\x6e\x4a\x84\x1f\xe7\x5c\xc1\x59\x0a\x7f\x10\x3a\x4e\x08\xc3\x17\x30\x78\xbe\x6f\x40\xb6\xb8\xd9\x3a\xb4\x70\xbc\x5c\x86\xf0\x30\x6c\xbd\xe1\x5b\x39\x0f\x0c\xdf\x3d\x6e\xbc\x2a\xfa\x80\x14\xfb\x45\x85\xa7\xaa\x54\x57\xdc\xe3\x43\xdf\x26\xe6\x8e\x57\x0a\xae\x83\xd5\x3c\x5d\xa6\x6d\xf4\x5e\xa8\x9d\x94\xdf\x27\xd5\x08\x06\xc3\xff\x97\x52\xca\x72\x31\xab\xc2\x56\xb7\x4a\xb0\xf1\x65\xff\x26\xf4\x5a\x1a\xa9\x1b\xdd\xbd\xe7\xa7\xc4\x1d\x7d\x5b\x36\x69\xea\xe6\x88\x6a\xdb\x99\x7a\xd5\x76\x8c\x73\xae\x1a\x4c\x97\xcb\x68\x2d\xa4\x8f\x2e\x37\x3d\xda\xf6\x3f\xea\x0e\x64\x64\xfa\x3e\x81\xb3\x0a\x53\xd6\xdf\x70\x27\xf9\xa0\xab\x34\x65\xe1\xab\xa8\x94\xa6\x3c\xd6\xca\xac\x21\xb2\x06\x68\x51\x63\xca\x7c\x93\x69\x49\xab\x72\x7d\x68\xda\xba\xcc\xc2\xba\x1c\xd7\xe5\xb1\x21\x5b\x13\xca\xc8\x74\x4c\x7c\x93\xe7\xe8\x7d\xbf\x00\x37\xaa\x7f\x1d\x02\x57\x0b\x09\xda\x56\x48\xcf\x33\x85\xa2\xdf\xa6\xeb\x17\x60\xa7\xce\x8a\xc8\xf7\x31\xb5\xf5\x7e\xa2\x27\x8f\x88\x8a\xf0\xd8\xb8\x9e\xa7\xb1\xf4\x2f\xb8\x4e\xc9\xd6\x4f\xa1\xd9\xd4\x82\xd3\xfd\xe3\xbe\xbe\xdb\xe1\x57\x3b\xa9\xb9\x5b\xb0\xf1\x4d\x77\x7e\x3c\xdf\xa1\x95\x16\x87\x39\x19\x1f\x5b\xbf\x1b\xb7\xfd\xdf\x24\x5e\x7d\x04\x26\x71\xf7\xcd\xfd\x4f\x00\x00\x00\xff\xff\x37\x52\x44\xaf\x83\x0b\x00\x00")

func indexTmplHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexTmplHtml,
		"index.tmpl.html",
	)
}

func indexTmplHtml() (*asset, error) {
	bytes, err := indexTmplHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.tmpl.html", size: 2947, mode: os.FileMode(436), modTime: time.Unix(1631205620, 0)}
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
	"index.tmpl.html": indexTmplHtml,
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
	"index.tmpl.html": &bintree{indexTmplHtml, map[string]*bintree{}},
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
