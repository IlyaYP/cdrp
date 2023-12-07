package model

import (
	"encoding/xml"

	"github.com/google/uuid"
)

type (
	Device struct {
		XMLName     xml.Name  `xml:"row"`
		Pkid        uuid.UUID `xml:"pkid"`
		Name        string    `xml:"name"`
		Description string    `xml:"description"`
	}

	Env struct {
		Return []Device `xml:"return>row"`
	}
)

var QueryDevices = `
<executeSQLQuery>
	<sql>select pkid,name,description from device</sql>
</executeSQLQuery>`
