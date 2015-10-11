package generators

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"github.com/kpmy/odf/model"
	"github.com/kpmy/odf/xmlns"
	"github.com/kpmy/odf/xmlns/office"
	"github.com/kpmy/odf/xmlns/urn"
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"io"
)

//Embeddable is any object that can provide []byte data for embedding in document package file (pictures for example)
type Embeddable interface {
	MimeType() xmlns.Mime
	Reader() io.Reader
}

//Parts is a list of root "files" of document package
type Parts map[string]*bytes.Buffer

//Entry is a part of Manifest
type Entry struct {
	MediaType string `xml:"manifest:media-type,attr"`
	FullPath  string `xml:"manifest:full-path,attr"`
}

//Manifest file structure, contains descriptors for any parts of document
type Manifest struct {
	XMLName xml.Name
	NS      string  `xml:"xmlns:manifest,attr"`
	Entries []Entry `xml:"manifest:file-entry"`
}

func (m *Manifest) init(mimetype xmlns.Mime) {
	m.XMLName.Local = "manifest:manifest"
	m.NS = urn.Manifest
	m.Entries = append(m.Entries, Entry{MediaType: string(mimetype), FullPath: "/"})
}

func docParts(m model.Model) (ret Parts) {
	ret = make(map[string]*bytes.Buffer)
	rd := m.NewReader()
	rd.Pos(m.Root())
	for !rd.Eol() {
		l := rd.Read()
		buf := new(bytes.Buffer)
		buf.WriteString(xml.Header)
		switch l.Name() {
		case office.DocumentContent:
			ret[xmlns.Content] = buf
		case office.DocumentStyles:
			ret[xmlns.Styles] = buf
		case office.DocumentMeta:
			ret[xmlns.Meta] = buf
		default:
			halt.As(100, l.Name())
		}
		enc := xml.NewEncoder(buf)
		err := enc.Encode(l)
		assert.For(err == nil, 60, err)
	}
	return
}

//GeneratePackage builds a zip-archived content of document model and embedded files and writes content to target Writer
func GeneratePackage(m model.Model, embed map[string]Embeddable, out io.Writer, mimetype xmlns.Mime) {
	z := zip.NewWriter(out)
	mime := &zip.FileHeader{Name: xmlns.Mimetype, Method: zip.Store} //файл mimetype не надо сжимать, режим Store
	if w, err := z.CreateHeader(mime); err == nil {
		bytes.NewBufferString(string(mimetype)).WriteTo(w)
	} else {
		halt.As(100, err)
	}
	manifest := &Manifest{}
	manifest.init(mimetype)
	for k, v := range docParts(m) {
		if w, err := z.Create(k); err == nil {
			v.WriteTo(w)
			manifest.Entries = append(manifest.Entries, Entry{MediaType: xmlns.MimeDefault, FullPath: k})
		} else {
			halt.As(100, err)
		}
	}
	for k, v := range embed {
		if w, err := z.Create(k); err == nil {
			_, err = io.Copy(w, v.Reader())
			assert.For(err == nil, 60)
			manifest.Entries = append(manifest.Entries, Entry{MediaType: string(v.MimeType()), FullPath: k})
		} else {
			halt.As(100, err)
		}
	}
	if w, err := z.Create(xmlns.Manifest); err == nil {
		w.Write([]byte(xml.Header))
		enc := xml.NewEncoder(w)
		err = enc.Encode(manifest)
		assert.For(err == nil, 60, err)
	}
	z.Close()
}
