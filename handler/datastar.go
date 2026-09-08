package handler

import (
	"embed"
	"path"

	"github.com/sumit-poudel/datastar-lsp/util"
)

//go:embed datastar
var docsFS embed.FS

type ActionDoc struct {
	Action string
	Desc   string
}

type Attribute struct {
	Name    string
	Options map[string]ActionDoc
	Desc    string
}

// Global lookup map: e.g., Attributes["data-on"]
var Attributes = map[string]Attribute{}

func init() {
	attrPath := path.Join("datastar", "attributes")

	attributes := util.GetLS(attrPath, docsFS)
	attrNames := util.ExtractMDBases(attributes)

	for _, attr := range attrNames {
		tempAttribute := Attribute{
			Name:    attr,
			Desc:    util.GetAttributeDoc(attr, docsFS),
			Options: make(map[string]ActionDoc),
		}

		actionPath := path.Join("datastar", attr)
		actionFiles := util.GetLS(actionPath, docsFS)
		actionNames := util.ExtractMDBases(actionFiles)

		for _, actionName := range actionNames {
			tempAttribute.Options[actionName] = ActionDoc{
				Action: actionName,
				Desc:   util.GetActionDoc(attr, actionName, docsFS),
			}
		}

		Attributes[attr] = tempAttribute
	}
}
