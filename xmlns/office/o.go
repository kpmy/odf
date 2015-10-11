package office

import (
	"github.com/kpmy/odf/model"
)

const (
	DocumentStyles  model.LeafName = "office:document-styles"
	AutomaticStyles model.LeafName = "office:automatic-styles"
	MasterStyles    model.LeafName = "office:master-styles"
	FontFaceDecls   model.LeafName = "office:font-face-decls"
	DocumentMeta    model.LeafName = "office:document-meta"
	DocumentContent model.LeafName = "office:document-content"
	Body            model.LeafName = "office:body"
	Text            model.LeafName = "office:text"
	Spreadsheet     model.LeafName = "office:spreadsheet"
	Styles          model.LeafName = "office:styles"
)

const (
	Version model.AttrName = "office:version"
)
