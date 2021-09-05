// Code generated for package template by go-bindata DO NOT EDIT. (@generated)
// sources:
// index.html.tmpl
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

var _indexHtmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x91\x5f\x6b\xdb\x30\x14\xc5\xdf\x0b\xfd\x0e\x9a\x5e\x47\xa5\x79\x0e\x6b\x56\x64\xc3\x70\x56\x46\x59\x48\xd3\xb9\xb4\xdd\xdb\x8d\x7c\x6b\x69\x91\x25\x23\xdd\xfc\x31\x21\xdf\x7d\x34\x4d\xd3\x3e\x1d\xf8\x9d\x23\x38\x3a\x57\x7d\x9a\xcc\xaa\xfa\xe9\xf6\x27\x33\xd4\xb9\xf2\xfc\x4c\xbd\x28\x73\xe0\xdb\x82\xa3\xe7\x07\x82\xd0\x94\xe7\x67\x8c\x31\xa6\x3a\x24\x60\xda\x40\x4c\x48\x05\xbf\xaf\xaf\x2f\xc6\xfc\xcd\x23\x4b\x0e\xcb\xdf\xd0\x60\x17\xbc\xa5\x10\x95\x7c\x45\x47\xdf\x59\xbf\x64\x11\x5d\xc1\x13\x0d\x0e\x93\x41\x24\xce\x4c\xc4\xe7\x82\x1b\xa2\x3e\x5d\x49\x99\x08\xf4\xb2\x07\x32\x62\x11\x02\x25\x8a\xd0\xeb\xc6\x0b\x1d\x3a\x79\x02\x72\x24\x72\x91\x49\x9d\xd2\x3b\x13\x9d\xf5\x42\xa7\xc4\x99\xf5\x84\x6d\xb4\x34\x14\x3c\x19\xc8\xc7\xa3\x8b\xb6\x9d\x0d\x77\x5f\xec\x63\xb5\x98\xce\xd7\xf9\xa3\xed\x3b\xc8\x47\xd3\xc9\xe7\xe6\x97\xcc\x9e\xe7\x97\xe3\x91\xfc\xf7\x4d\x3f\x49\x7b\x53\xcf\xef\x67\x46\x3f\xc4\xcb\xed\xf7\x9b\x75\xb8\xdb\xd6\x5f\xa7\x7f\x37\x59\xcd\x99\x8e\x21\xa5\x10\x6d\x6b\x7d\xc1\xc1\x07\x3f\x74\x61\x95\x0e\xe3\xc8\xe3\x3a\x6a\x11\x9a\xe1\xed\xa7\x8d\x5d\x33\xed\x20\xa5\x82\x6f\x22\xf4\x3d\xc6\xd3\x48\x26\x2b\x7f\x2c\x69\x85\xce\x21\x9b\x00\xa1\x57\xd2\x64\xa5\x5a\xc4\x53\x20\x2f\x77\x3b\x71\x1b\x36\x18\xaf\xc3\x56\x3c\x00\xd1\x7e\xcf\x5e\x44\x49\x93\x7f\x48\x55\x10\xd9\x1f\x02\x5a\xa5\x2b\xb6\xdb\x89\xca\x40\x6c\x31\x8a\x0a\xe2\x2b\xdd\xef\x3f\x3e\x90\x8d\x5d\x1f\xfa\x1e\x7b\x2a\x79\x38\xf8\xff\x00\x00\x00\xff\xff\x54\x83\x9d\xa1\x00\x02\x00\x00")

func indexHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_indexHtmlTmpl,
		"index.html.tmpl",
	)
}

func indexHtmlTmpl() (*asset, error) {
	bytes, err := indexHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html.tmpl", size: 512, mode: os.FileMode(438), modTime: time.Unix(1630846855, 0)}
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
	"index.html.tmpl": indexHtmlTmpl,
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
	"index.html.tmpl": &bintree{indexHtmlTmpl, map[string]*bintree{}},
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
