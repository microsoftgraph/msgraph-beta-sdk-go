package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PersonAnnotation struct {
    ItemFacet
    // Contains the detail of the note itself.
    detail *ItemBody;
    // Contains a friendly name for the note.
    displayName *string;
    // 
    thumbnailUrl *string;
}
// Instantiates a new personAnnotation and sets the default values.
func NewPersonAnnotation()(*PersonAnnotation) {
    m := &PersonAnnotation{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the detail property value. Contains the detail of the note itself.
func (m *PersonAnnotation) GetDetail()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// Gets the displayName property value. Contains a friendly name for the note.
func (m *PersonAnnotation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the thumbnailUrl property value. 
func (m *PersonAnnotation) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// The deserialization information for the current model
func (m *PersonAnnotation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(*ItemBody))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    return res
}
func (m *PersonAnnotation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PersonAnnotation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the detail property value. Contains the detail of the note itself.
// Parameters:
//  - value : Value to set for the detail property.
func (m *PersonAnnotation) SetDetail(value *ItemBody)() {
    m.detail = value
}
// Sets the displayName property value. Contains a friendly name for the note.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PersonAnnotation) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the thumbnailUrl property value. 
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *PersonAnnotation) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
