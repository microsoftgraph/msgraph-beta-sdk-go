package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Pictureable 
type Pictureable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContent()([]byte)
    GetContentType()(*string)
    GetHeight()(*int32)
    GetWidth()(*int32)
    SetContent(value []byte)()
    SetContentType(value *string)()
    SetHeight(value *int32)()
    SetWidth(value *int32)()
}
