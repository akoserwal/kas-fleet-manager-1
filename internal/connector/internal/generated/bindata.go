// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package generated generated by go-bindata.// sources:
// .generate/openapi/connector_mgmt.yaml
package generated

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

// ModTime return file modify time
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

var _connector_mgmtYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x7b\x73\xdb\x38\x92\xff\x5f\x9f\xa2\x4f\xb9\x2b\xef\xde\xc5\x7a\xf8\x19\xab\x6a\xb6\xca\xb1\x9d\x8c\x67\x12\xc7\xe3\xc7\x24\xd9\xab\x2d\x07\x22\x5b\x22\x2c\x12\xa0\x01\x50\xb6\xbc\x35\xdf\xfd\x0a\x00\x29\xf1\x29\x51\x8e\x2f\x76\x12\xb9\x2a\x7f\x88\x6c\x34\x1b\x8d\xc6\xaf\x5f\x20\xc3\x43\x64\x24\xa4\x3d\xd8\x6c\x75\x5a\x1d\x78\x01\x0c\xd1\x05\xe5\x51\x09\x44\xc2\x80\x0a\xa9\xc0\xa7\x0c\x41\x71\x20\xbe\xcf\x6f\x41\xf2\x00\xe1\xf8\xf0\x48\xea\x4b\x23\xc6\x6f\x2d\xb5\x1e\xc0\x20\x66\x07\x2e\x77\xa2\x00\x99\x6a\x35\x5e\xc0\xbe\xef\x03\x32\x37\xe4\x94\x29\x09\x2e\x0e\x28\x43\x17\x3c\x14\x08\xb7\xd4\xf7\xa1\x8f\xe0\x52\xe9\xf0\x31\x0a\xd2\xf7\x11\xfa\x13\xfd\x24\x88\x24\x0a\xd9\x82\xe3\x01\x28\x43\xab\x1f\x10\x4b\xc7\x61\x84\x18\x5a\x49\x66\x9c\x9b\xa1\xa0\x63\xa2\xb0\xf9\x12\x88\xab\xe7\x80\x81\x26\x55\x1e\x42\x73\x44\xe4\xfa\xc0\x47\x54\xeb\x01\x61\x64\x88\x62\x3d\x26\x6e\x4d\x48\xe0\x37\x61\x40\x7d\x6c\x50\x36\xe0\xbd\x06\x80\xa2\xca\xc7\x1e\x1c\x70\xc6\xd0\x51\x5c\xc0\x39\x8a\x31\x75\x10\xde\x68\x0e\xf0\xde\x72\x68\x00\x8c\x51\x48\xca\x59\x0f\x3a\xad\x4e\x6b\xb3\x01\xe0\xa2\x74\x04\x0d\x95\xb9\xb8\x60\xbc\x9d\xd0\x19\x4a\x05\xfb\xa7\xc7\x5a\x52\x2b\x1a\x38\xc9\x38\xd9\x6a\x48\x14\xfa\x21\x5a\xaa\x75\x88\x84\xdf\x03\x4f\xa9\x50\xf6\xda\x6d\x12\xd2\x96\xd6\xb6\xf4\xe8\x40\xb5\x1c\x1e\x34\x00\x72\x02\xbc\x27\x94\xc1\xdf\x42\xc1\xdd\xc8\xd1\x57\xfe\x0e\x96\x5d\x39\x33\xa9\xc8\x10\x17\xb1\x3c\x57\x64\x48\xd9\xb0\x94\x51\xaf\xdd\xf6\xb9\x43\x7c\x8f\x4b\xd5\x7b\xd5\xe9\x74\x8a\xc3\xa7\xf7\x67\x23\xdb\x45\x2a\x27\x12\x02\x99\x02\x97\x07\x84\xb2\x86\x22\xc3\x58\x01\x8c\x04\x99\x75\xb9\x98\x84\x28\x8b\xe3\x9b\xcd\x32\xea\xda\x84\x70\xe0\x47\x52\x61\xc5\x80\x46\x48\x94\x67\xe4\x79\xa1\xff\xc1\x85\x87\x12\x81\x08\x34\x86\x36\x5d\x3b\x10\xe8\x13\x85\xae\x5e\x5b\x19\x13\x37\xb5\x9e\xdb\x53\x92\xab\x60\x18\xa8\xf6\xb8\xdb\x1e\x91\xc1\x88\x5c\xcd\xae\x2b\x3d\xad\xf6\xbf\xb3\x17\xae\xa8\xfb\x57\xb3\x67\x44\x0a\x89\x20\x01\xaa\xd8\x2e\xf4\x5f\x32\x87\xc2\x90\xf8\x7e\x6e\x1a\x17\x1e\x02\x75\x81\x0f\x72\x32\xeb\x41\xd3\x11\xd2\xf1\x30\x20\xbd\xe9\x6f\x30\xb7\x7b\x20\x95\xa0\x6c\x38\xbd\x4c\x59\x0f\xb4\x4a\xa6\x17\x04\xde\x44\x54\xa0\xdb\x03\x25\x22\xcb\x6e\x88\x2a\xe1\x93\x2c\x66\x22\x77\xd9\x62\x6a\x1e\x32\xe4\x4c\x62\x8a\xb4\xb9\xd1\xe9\x34\xd3\xd2\x38\x9c\x29\x64\x2a\x7d\x09\x80\x84\xa1\x4f\x1d\xa2\xa7\xd9\xbe\x96\x9c\x65\xef\x96\x4d\xca\xfe\xfd\xa7\xc0\x41\x0f\x9a\x2f\xda\x0e\x0f\x42\xce\x90\x29\xd9\xb6\xb4\xb2\x3d\x15\x51\x4b\xd8\x4c\x0d\x2d\xe8\x34\xab\x49\x08\x88\x72\x3c\xbd\x5b\xb4\x96\xb5\x5a\xd0\x18\x7e\x3c\x9f\xad\x4e\xf7\x69\xe6\x73\x24\x04\x17\xcd\xdc\x10\xbc\x23\x41\xe8\xa7\x15\x9e\xfc\x6d\x75\xba\x47\xf6\x66\xf1\x5e\xf9\x83\x12\x5e\xed\xd9\xd0\x4a\xb5\xed\x47\xca\x03\xc5\x47\xc8\x34\x1e\x52\x36\x26\x7e\xca\x6a\x9b\x5b\x9d\xad\xef\x44\x49\x5b\x0f\x57\xd2\xd6\x22\x25\x9d\xf0\x99\x2d\xe5\x6c\x0c\xef\xa8\x54\x72\xa6\xb0\xed\xa7\xda\x25\x4b\x2a\x6c\xbb\xd3\x79\xa8\xc2\x66\x43\x2b\x15\x76\xc9\xf0\x2e\x44\x47\xe3\x2f\x6a\xb9\x80\x3b\xc6\xa9\x24\x96\x25\xd1\x89\x04\x55\x93\x34\x12\xbd\x46\x22\x50\xf4\xe0\x7f\xe1\x5f\xf1\x55\x1e\xa2\x30\x4a\x3a\x76\x7b\x1a\xc3\x32\x40\xf0\x7a\x72\x7c\x98\x70\x8b\x82\x80\x88\x49\x0f\xde\xa2\x02\x92\x5f\xa1\xfe\x04\xa8\xdb\x68\x00\x2c\x83\xfe\xbd\x87\xe1\xe6\x03\xe6\xe5\x53\x99\x9d\x98\xcc\xcf\xea\x0c\x55\x24\x98\x0e\x56\x34\xad\xf6\x1a\xd9\x19\x26\x03\x8a\x6e\x49\x3f\xbe\x6c\x29\x67\x94\xed\x90\x0c\x53\xcb\xb8\x90\x5c\xd2\xfb\x29\xf9\xf3\x75\x12\xef\xa8\x54\xd5\x88\xb7\x40\x8d\x2b\xd7\x50\xcf\x35\xac\x90\x6e\x11\xd2\xd5\x03\x9c\x58\xe4\x50\x47\xed\x8b\xc0\x46\x96\x21\x88\x23\x90\x28\x9c\xd2\xcc\x45\x03\x1d\x2d\xde\x44\x28\x26\xa9\xf9\xd8\xd0\x95\xc8\x09\x73\xaa\x66\x79\x8a\x62\xc0\x45\x60\x02\x29\x62\x72\x19\xa0\x4c\xe7\x9b\x66\x94\x27\x38\xe3\x91\xd4\xf9\x13\x43\xd1\x98\xbf\xba\x36\x88\xed\x73\xee\x23\x61\xa9\x3b\x25\x61\x2b\x24\x41\xdb\x6b\xee\xa6\x20\xb5\x22\xc9\x73\x89\x22\x53\x9a\x12\x63\x9c\x6f\x8a\xe5\x86\x58\x0b\x71\x9a\xf3\x42\xef\x2a\x98\xdc\x78\x62\x98\xac\xde\xf5\x8e\x83\xa1\xc2\x4c\x0c\xf8\x7d\x6c\xf4\xad\x4e\xe7\x40\x6f\x05\xca\xd9\xc3\x51\x31\xcf\xa2\x52\x4f\x7f\x6a\x34\x34\x94\x76\xe3\xcb\x7c\x8c\xb3\xf2\x23\xab\x14\xa3\x7e\x8a\x71\x31\x4b\x51\xd1\xd5\x98\xc1\x23\xe1\x20\xb8\x1c\x25\x5b\x53\x36\xcd\x58\xf9\xde\x9c\x61\x31\x88\xaa\xdc\xaf\xf5\x8a\x49\xf2\xef\xe4\x9c\x63\xbd\x38\x7d\x1a\x81\x1b\x4c\x40\x53\xfe\xbc\xcd\xf1\xaa\x93\x22\x3c\x55\x80\x3e\xdf\xe1\xdb\x30\x24\xb5\x13\x4b\x2c\xd2\xd0\x80\x63\xeb\x71\x90\xa1\xad\xf6\xed\x99\x02\xd5\xf3\x4c\x12\x96\x4d\x10\x56\xb9\xc1\x2a\x37\x78\x16\x55\x90\x4c\xb5\xe0\x41\x95\x82\x9a\x45\x70\xd9\xfe\xf7\xfc\x82\xf7\x02\x1c\xa2\x6e\x73\x4a\x5a\xc4\xa0\x0a\x04\xaa\x8f\x3f\x35\x8a\xe3\xcb\x20\xf3\xf3\x44\xa9\x9a\xb5\xee\x55\x99\x7b\x15\x83\xd6\x51\xd2\xaa\xcc\xbd\x94\xc2\x9e\xbe\xcc\x9d\x87\xf7\x5c\x79\xbb\x61\xe5\xf1\x51\xe1\x57\x02\xdd\x73\x31\xfe\x8c\x72\x0f\xcd\xcc\x56\xd9\xf4\x0f\x8b\x64\x76\x81\xbf\x02\xcf\x32\x0c\xe6\xa1\x9a\x0d\x23\x62\xd7\x08\xb7\x54\x79\x20\x43\x74\xe8\x80\xa2\x0b\xc7\x87\xdf\x33\xba\x7d\x9d\x12\xf3\x0c\xbe\x25\xd2\x59\xe4\xaa\x04\x3b\x2b\x57\x01\xef\x42\xed\xa0\xbe\x12\xee\x9e\x6d\xed\x75\x15\xd7\xfd\xbc\x68\xb8\x8a\xeb\x7e\xe8\xb8\xce\xe0\x56\x25\xd8\x99\xbb\x05\xac\xab\xd3\xff\x3a\x24\x8a\x80\xe2\x31\x87\xec\x71\x32\xed\xe8\x1a\x73\x56\xf2\x1b\x74\xc4\xb2\x0f\x09\x50\x0c\x71\xdd\x88\xfa\x3f\xdf\xe4\x81\xfa\x21\x4b\x3e\x2f\x57\x0a\xfc\xed\xfc\xc3\x09\x9c\x6a\x0e\x2f\xe1\xec\xcd\x01\xec\xec\x75\x36\x60\x7d\x7a\x70\x53\x71\xee\xcb\x16\x45\x35\x68\x71\x31\x6c\x7b\x2a\xf0\xdb\x62\xe0\x68\xaa\x1c\x5f\x5b\x9f\xe0\xfd\x6b\x74\x66\x28\x9e\x6b\x18\x9a\x03\x8d\x47\xd3\x03\xb5\x03\xe3\x04\x98\x3d\xf1\x39\x5b\xd6\xb8\x1c\x22\x81\x30\x37\x75\x99\x0c\x91\x29\x88\xbb\xb8\x35\xcf\xba\x24\xac\xe2\x32\x4f\x9d\x16\x74\xf6\x74\xe6\xfc\x56\x74\x4c\x1a\x53\xfe\x74\x1d\xe9\xa4\x6e\xf5\x64\x9d\xe9\x58\xff\xdf\x65\x83\xba\x20\xfb\xaa\x4f\xbd\xea\x53\xaf\x62\xc9\x3a\x4a\x5a\xf5\xa9\x9f\x57\x38\xf9\x80\x3e\x75\xe2\x3b\x96\x0a\x38\x17\xf4\xab\x33\x3c\x6b\x1d\x6d\xcd\xb9\xfa\xd5\x01\xd3\x94\x63\x5a\xfa\x8c\xa9\x93\x55\xe6\x0a\x76\x57\xad\xe4\xe7\xd3\x4a\xce\xed\xf4\x5a\x1d\xe5\x94\x41\x2f\x99\x6e\xa4\xdf\xad\x8a\xaf\x5d\x51\xf7\xaf\x38\x0b\xa9\xf1\x76\xd5\x6c\x50\x79\x0c\x5e\xf5\x82\x55\x16\x55\xbf\xfd\x3b\x56\x39\x35\x3f\x6b\x80\xab\x59\xa5\x4c\x92\x9c\x55\xb5\x72\x15\x61\xd6\x51\xd2\x03\xab\x95\x89\x99\xad\xaa\x96\x4f\xd5\x8d\xce\x16\x73\x2a\xdf\xb9\x4a\x43\x6c\xcd\xe6\x74\x3d\x58\x7c\x2e\x7b\x63\xd5\xa4\xfe\x99\x80\xee\xdb\x35\xa9\xa7\xb9\xf4\xaa\x3f\xfd\x1c\xfb\xd3\x15\xf0\x57\x6c\x53\xa7\x10\xf0\x71\xc2\xe2\x36\x71\x5d\xce\xae\xf2\x71\xf1\x2a\x4e\x7e\x9a\x3d\xb0\xaf\x57\xe3\x34\x51\xfe\xdc\x32\x80\x56\xef\x6c\x99\x40\x79\x44\x81\xf4\x78\xe4\xbb\xd0\x47\x88\xa4\xfd\x96\x8a\xc3\xd9\x80\x0e\xa3\xf8\x0b\x16\xf6\x23\x24\x99\xc6\x8e\x7e\x20\x70\x66\xd7\xc8\x6a\xa6\xb5\xf2\x38\x3f\xaa\xc7\x59\x85\xd6\x3f\x6b\x68\x9d\x45\x96\x42\x2d\xa6\x22\xd2\x5e\x93\x31\x42\xcc\x90\xa6\xd1\x98\xcd\x53\x8b\x14\xeb\xd4\x4a\x97\x79\x53\x3e\x11\x38\x33\xdb\x7f\xac\x4f\xa7\x71\x86\xa1\x40\xa9\xf9\x14\x3f\xac\x20\xa3\x30\xe4\x42\xeb\xa4\x3f\x31\xd8\xb4\x7f\x7a\x3c\x05\xec\x18\xf7\x53\xfa\xd0\xbe\x29\xf5\x33\xfe\x76\x53\xea\x8a\xb6\x9c\x2b\x2b\x6a\x7c\x95\xf8\xfe\x87\xc1\xa2\x5a\x73\x62\x30\x1f\x4c\x6b\xfd\x0c\x07\x28\x90\x39\x99\x9a\x73\x69\xef\x1d\x20\x14\x7a\x45\x14\xcd\xdb\x93\xf1\xa2\x39\x2b\xca\xee\x23\x12\x60\xf9\xd7\x7b\x5a\xb9\x61\xa5\x6e\x52\xff\x25\x9f\xae\x9a\xf7\x98\x3f\x2d\xcd\x57\x3e\xc9\xf1\x08\x63\xe8\x17\xf6\xcc\xd7\x3e\x8a\x08\x41\x26\xb9\x3b\x54\x61\x50\xb2\x39\x2b\x65\x4b\xcb\x30\x4f\xbc\xfd\xf4\xcf\x82\x90\xb5\x55\x41\x1d\xce\xae\x3c\x6d\x43\xf3\x1e\x76\x79\xf6\xce\x7c\x6d\x8d\x19\xfa\x87\x3f\xcd\x27\xfd\x45\x6a\xb7\x24\xb3\x48\x80\x28\x1c\x72\x41\xef\xb1\xf4\xad\xc5\xc7\xd7\xbf\xfe\x43\x16\x05\x1a\xa5\x24\x65\xa3\x97\x10\xe7\x20\xff\xca\x90\xa5\x76\xe6\xdc\xe9\xa4\x80\x23\xf9\xdb\x37\x83\x63\x04\xb2\x11\x90\x43\x58\x3a\xfc\x19\xdb\x06\x75\x26\x86\x96\x05\x3e\xb3\x98\x56\x67\x46\x30\xa0\xe8\xbb\xe5\x2b\x61\xb7\x39\xbc\x00\xc5\x5d\xde\x03\x81\xa1\x4f\x92\xac\xaa\x8f\x2a\xf1\x74\x45\x2c\xd4\xa1\x5c\xef\x21\xc8\x93\x8d\x01\x97\x86\x9b\xd2\x45\x5b\x7a\x8d\x97\xf8\x82\x56\x76\xea\x3f\xc1\xb4\xf3\x53\x2e\xf5\x7a\xfb\x29\xd4\xf3\xb8\xef\xca\x64\x0f\x9a\xe0\xdc\x9e\xa0\xb0\xd1\xba\x26\x02\x02\xbf\x9b\x53\xe5\x8a\x87\xd4\xb1\x78\xc1\x95\x87\x02\xe4\x44\x2a\x0c\x5a\x4f\xeb\xc2\x02\x54\xc4\x25\xaa\xb0\x5d\x2b\xd8\xcc\x63\xa5\xff\x92\x37\xf5\xca\x62\xab\xb9\xe0\xc2\x6f\x19\x8a\xa5\x47\x95\x39\xe0\x85\x83\xec\xd9\x33\xf7\x8a\xa8\xb2\xa1\x03\x2e\x02\xa2\x7a\xa0\x91\x66\x5d\xd1\x54\x28\x52\x93\x7d\x14\xba\xff\x9f\xec\x93\xea\xcf\x55\x45\x60\x30\x63\x40\x99\xc2\x21\xe6\x7d\x42\x5a\x08\xca\xd4\xce\x56\xce\xcf\x86\x3e\x9f\x04\xc8\xd4\x95\xcf\x6d\xa0\x9e\x7f\xc0\xdc\x8d\x64\x23\xcd\x0b\x22\x86\xa8\xb2\x31\xb9\xb1\x8c\x65\x78\x99\x5d\x13\xef\x44\xca\xd9\x39\x2a\x45\xd9\x50\x66\xb9\x16\xbe\xa3\x58\x6e\xc7\x65\xd1\x4e\xc6\x4f\xd4\x36\xff\x38\x48\xaa\xfd\x18\x17\xa5\x8e\x6f\xaf\xa4\x22\xaa\x60\xaa\x95\xa3\x34\x75\x54\x81\x79\x19\xf2\xbb\x75\x9f\xb2\x51\x8a\x32\xab\x90\x34\x87\xba\x1f\x49\x83\x8a\x03\x1c\xe5\xea\x86\xb5\xc2\xb5\xb5\xaa\xe4\xe8\x1f\x73\x78\x69\x89\x64\x49\x48\xa9\x53\x86\x48\xc6\xed\xc2\xcc\x78\x4d\xfb\xa5\xf0\xf0\x2f\x33\xd9\x81\xb2\xe9\xab\x3c\x8a\x67\xc6\x7e\x79\x7b\x74\xb1\xd4\x67\xde\xca\x3e\xf2\xf9\xa5\x15\x7b\x8b\xb4\xd1\x4f\x3d\x06\xd5\x13\x0f\x28\x23\x29\x37\x32\xc5\xce\xc9\x89\x7d\xaf\x99\xb2\x59\xe9\x21\x20\x61\x48\xd9\x30\xad\x70\x93\xae\xcd\x2d\x2f\x55\x6e\x38\xc7\xe7\x91\x7b\x15\x0a\x3e\xa6\xae\x4e\x2e\xab\xf6\x2b\x8f\xdc\xd3\x98\xa8\x94\x17\x67\x58\xdf\x2d\xcd\x13\x69\xfe\xc8\x79\x82\x18\x16\x45\xd6\xa5\xbe\xb9\x69\xef\x49\xb8\xe5\x62\xe4\x73\xe2\xca\x38\x42\xb7\x99\xaf\x93\xed\x58\x97\xec\xf2\xb2\x84\x34\xb5\x4c\x65\xbe\x4f\xdf\x5e\x58\xf1\x9c\x15\x58\x2b\x49\x63\x63\xaa\x52\xc4\x32\xf3\xb5\xcb\x0f\xc9\xf2\x3f\xc9\x7c\x33\xf6\xb7\x88\x5c\xe0\x30\xe7\x6a\x4a\xc9\x82\xc8\x57\xf4\x8a\xdc\x17\x09\x93\x53\xd9\x33\x63\xc9\x54\x5e\xcb\x95\x37\x3b\x7d\x95\x2f\x8c\x64\x35\x96\x0e\x33\x73\xe1\x65\xfd\x02\x70\xb3\x4c\xb6\x2a\xb9\x72\xf2\xcc\x59\xc0\xb2\x15\x9a\x63\x64\xc9\xc5\x31\xf1\x23\x5c\x60\x8a\xb9\xaa\x53\xb9\xac\xe7\x36\x61\x1b\x68\xbc\x9e\x1d\xe8\xd0\x56\x68\x03\x21\x20\xe6\x58\x3b\x84\x3e\x61\x98\x2a\x45\x59\x17\xd7\xfc\xa1\xa2\xdf\x55\x10\x5b\x83\x7d\x8d\xd8\xa6\xd4\xfc\x7e\x82\xf4\x73\x7a\xa0\xca\x0c\xaf\x08\x80\x7b\x73\x36\x7e\x9f\x73\x25\x95\x20\xe1\x95\xfd\xee\x7b\x0d\x98\xa6\x3a\xd6\xaf\x01\x18\x31\xa5\x44\x47\xa0\x9a\x0f\x1c\xb9\xfd\xd9\x7b\x74\xf0\xaa\xe5\x81\xf2\xe5\xbb\xa2\x9c\x69\x8b\x5a\xe0\x07\xf5\xcf\x90\x0c\xd3\x75\x69\x49\xef\xd3\x3f\x15\x57\xc4\x4f\xfd\x36\x76\xb0\xdc\xcc\x6b\x4d\x4b\x4b\x51\x24\xca\xe7\x7a\x5a\xb8\xc5\x54\x46\xe6\x6a\x32\x73\xc3\xb4\x57\x1e\xb4\xef\x1e\x0f\xa6\x1d\xee\xd6\xcf\x9d\x04\x92\x92\xce\x52\x25\xf9\x34\x2d\x5a\x94\x3c\x36\x8a\x69\xd1\x6c\x84\xed\xa7\x4f\x1b\x89\x85\xe6\xee\xf1\xa1\x8e\x32\x04\x3a\x5c\x4c\x5b\x41\xb9\x32\x69\x89\x84\xb9\x36\x79\x49\x93\x3c\x6d\x0d\x56\x86\x94\x95\xe6\xdf\x98\xcb\xbe\x17\x47\x86\x08\x94\xb9\x78\x57\xe0\x3e\x20\xbe\xc4\xfa\x52\x16\x7b\x6d\x79\x1b\xb5\xc1\x06\x34\xbb\xd6\x06\xd2\xc6\x69\x85\x4e\xed\xa5\xb9\x42\x9f\x44\x41\x1f\x85\x56\xa5\xd9\x5e\x3a\xc3\x43\xe2\x78\xe9\x49\x3f\xe2\x34\xf2\x9b\x68\x3a\x8d\x4e\xc7\x40\x74\xd2\xcc\x33\x21\x50\x32\x30\xee\xe6\xa5\x9f\xab\xc3\x53\x73\x35\xbe\x68\x7f\xbc\x89\x9d\xe8\x6f\x1f\x2f\x32\x30\xe1\x29\x15\x6a\xee\x59\x79\xaa\xdf\x05\xcb\x45\x72\x3a\x27\x6f\x76\x3b\x9b\xcd\x2c\xac\x40\x33\xd7\x27\xf5\xec\xce\xad\x48\x82\xed\x1b\x60\xed\x0c\x1f\xb3\x0b\xa1\x79\xf0\xe1\xe4\xe4\xe8\xe0\xe2\xc3\xd9\xfa\xfb\xb7\xef\x2f\xd6\x33\x24\xf1\xde\x83\xe6\x79\xea\xdd\xca\xe4\xad\x4b\x09\x8c\xab\x59\x33\xf0\x25\x44\x12\xed\x5b\x98\xbf\x68\x7b\x2e\x46\xbb\xb9\xcd\x09\xcd\x2e\xfd\x78\x4c\x83\x9b\xb7\x8e\x38\x8c\xde\xed\x74\xc9\xe5\xdd\xf1\x3f\x6f\x5e\x5f\xdc\x9c\x9c\x91\x66\xa2\xa5\x63\xdb\xd1\xff\x43\x9b\x50\x0d\x4d\x6d\x3c\x92\xa6\x36\x16\x2a\x6a\xa3\x4c\x4f\x6f\x08\xf5\x6d\x77\x23\x24\x42\x22\x48\x24\xc2\xf1\xec\x06\xe8\xc1\x25\x33\xff\xcf\x8e\xe2\x36\x4f\xf9\x3d\xfd\x75\x12\xfb\xd6\x2f\x09\xe9\x95\x2d\x58\x48\xfb\x5f\xd7\xf4\xa0\xf0\xd8\x1e\x2c\x7a\xca\xac\x41\xeb\x70\x3f\x0a\x98\xd9\x95\x86\xbf\xa5\xec\xc1\x1a\x75\xd7\x5a\x70\x5e\x46\x27\x81\x08\xec\xc5\x19\xdc\x4b\x73\xe9\x65\x2e\xfd\x4b\xae\xda\xd0\xaf\x05\x66\x71\x92\xb3\x17\x1a\x36\xe1\x17\xe8\xa6\xf5\x93\x5f\x77\xff\xe3\xe1\xdb\x68\xd2\x3f\x16\x47\xec\x4e\xec\x63\xb0\xbb\xb1\x35\xbc\x19\x8d\xe8\xe1\x38\x59\xf7\xfc\x69\x88\xb2\xb5\xde\xea\x6c\x3d\xca\x5a\xef\x2e\x5a\xea\xdd\x92\x95\xae\xf3\x3e\xdc\x74\x32\xa5\x87\xe1\xca\xa6\xb4\xfb\x74\x13\x32\xd6\x78\x96\x39\x8b\x48\xdd\x5f\xd6\xba\xf4\xf7\x4d\x37\xfa\xf3\xf3\xf1\x78\xbc\xfd\x79\xfc\xce\x9f\xdc\x77\x83\xb7\x67\x9b\xbf\x4d\x6e\x4e\xd6\x0c\x00\x0c\x78\xc4\xdc\x39\x5b\xfc\xf3\x87\xdd\xe1\xc6\x70\xe7\xd7\x0b\xf7\xf2\xf7\x4b\xb2\x31\x92\xbf\xbe\xda\x18\xfd\x71\xb8\x39\x49\xb4\x93\x3f\x1d\x54\x0a\x80\xdd\xc7\xc1\xbf\xee\x42\xf8\xeb\x96\xa8\x66\xb6\x6f\xc7\x28\xe8\x60\xa2\x31\xde\x9e\x39\xea\xc1\x59\xec\x9f\x80\x44\xca\xe3\x82\xde\x27\x0d\xa3\x11\xb2\x7a\xfa\xd9\xbc\xf4\x8e\xbc\xdb\xe0\xd3\xeb\xf0\xe3\xe9\xe0\x78\xc3\x3f\xc1\x51\xe8\x6e\xfd\xf3\x30\xd1\xcf\xde\x09\x09\xf0\x80\xb3\x81\x4f\x1d\x55\x43\x57\x9b\x3b\x8f\xa2\xab\x34\x9b\x72\x5d\xa5\x29\xb2\x66\x34\x3d\x5a\x64\x80\x87\x4a\x20\xbe\x40\xe2\x4e\x4c\xef\xb7\x52\x17\x3b\xa3\xcf\x9d\x4b\x7a\x34\xba\x1f\x7d\x3a\xb8\xff\x78\x8a\xc7\x1b\xfc\x33\x7a\xee\xe6\x51\xac\x8a\xe2\x51\x9f\xb2\xe9\xef\x3d\xca\xec\xf7\x16\x4d\x7e\xaf\xd4\x4e\x66\xa7\x77\x31\xfb\xd0\xc2\xb2\xe3\xd1\xbb\xf1\x9b\xbd\xeb\xf7\x7f\x7c\xde\xf9\x3c\xf4\x06\xef\xf7\x86\x6f\xcf\xe4\xaf\xe3\xa3\x8f\xd3\xb9\xd6\x06\x8d\x27\x9b\x71\x2a\x96\x6a\xda\xe3\x51\xe6\xec\xac\xf9\x7f\xca\x26\xcc\x91\xa8\x7a\xf0\xe1\xe0\xfd\xfa\xd1\xa7\xf5\x3d\x1d\xbf\x25\x5e\xcb\x9e\xb0\x9d\xd1\xe0\x9d\x5a\x8f\xfd\x1d\x09\xe9\x7a\x97\xde\x75\x36\x7d\xe6\xfa\xc1\x4d\xe7\x66\xe0\xec\x4a\xaa\xc8\xb6\xf4\xaf\xc7\xaf\xd2\x85\x89\x41\xea\x04\x9b\x56\x43\x77\xb8\xed\xbe\x7a\x75\xd3\xf1\x85\xe3\x8e\xb7\x86\xbb\xc4\xef\xef\x4a\x7f\x30\x64\xd7\x9b\xae\xd7\x97\xd7\xff\xf5\x1f\x7f\x3b\xfa\x74\x71\xb6\x0f\xff\x6d\x27\xdc\x32\x4a\xfa\x85\xba\xc8\x94\x5e\xb2\x74\x63\x8d\x4a\x58\xdb\xea\x6c\xad\xbd\x34\xaa\x30\x3f\x0f\xde\x5d\x9e\x5f\x1c\x9d\x9d\x5b\x5d\xe8\x9b\xe6\x33\x1c\xd3\x75\x85\x19\x23\x43\xdf\x1d\x6e\x73\xb1\xdd\x19\xd3\xa8\xb3\xcb\x51\xaf\x9a\x27\x46\xce\xc6\x8e\x3b\x1c\xa8\xeb\x2e\x71\xd6\xd2\xda\x8b\x0b\x05\x66\xd4\xdc\x49\xa4\x20\xf7\xef\xf3\x30\xe5\x42\x7e\x14\x93\x1d\x26\x6f\xfa\x1b\xf2\x24\x78\x73\xbd\xdd\xff\x14\x1e\xee\x1e\x90\x66\xe3\xff\x02\x00\x00\xff\xff\xf1\x09\x57\xfb\x78\x70\x00\x00")

func connector_mgmtYamlBytes() ([]byte, error) {
	return bindataRead(
		_connector_mgmtYaml,
		"connector_mgmt.yaml",
	)
}

func connector_mgmtYaml() (*asset, error) {
	bytes, err := connector_mgmtYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "connector_mgmt.yaml", size: 28792, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"connector_mgmt.yaml": connector_mgmtYaml,
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
	"connector_mgmt.yaml": &bintree{connector_mgmtYaml, map[string]*bintree{}},
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
