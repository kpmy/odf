package office

import (
	"odf/model"
)

const (
	DocumentStyles  model.LeafName = "office:document-styles"
	AutomaticStyles model.LeafName = "office:automatic-styles"
	MasterStyles    model.LeafName = "office:master-styles"
	FontFaceDecls   model.LeafName = "office:font-face-decls"

	DocumentContent model.LeafName = "office:document-content"
	Body            model.LeafName = "office:body"
	Text            model.LeafName = "office:text"
)

const (
	Version model.AttrName = "office:version"
)
