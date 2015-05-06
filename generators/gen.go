package generators

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"odf/model"
	"odf/xmlns"
	"odf/xmlns/office"
	"odf/xmlns/urn"
	"ypk/assert"
	"ypk/halt"
)

type Parts map[string]*bytes.Buffer

type Entry struct {
	MediaType string `xml:"manifest:media-type,attr"`
	FullPath  string `xml:"manifest:full-path,attr"`
}

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

func Generate(m model.Model, out io.Writer, mimetype xmlns.Mime) {
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
	//place for attachements
	if w, err := z.Create(xmlns.Manifest); err == nil {
		w.Write([]byte(xml.Header))
		enc := xml.NewEncoder(w)
		err = enc.Encode(manifest)
		assert.For(err == nil, 60, err)
	}
	z.Close()
}
