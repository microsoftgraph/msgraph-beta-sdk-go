package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemPublication struct {
    ItemFacet
    // Description of the publication.
    description *string;
    // Title of the publication.
    displayName *string;
    // The date that the publication was published.
    publishedDate *string;
    // Publication or publisher for the publication.
    publisher *string;
    // URL referencing a thumbnail of the publication.
    thumbnailUrl *string;
    // URL referencing the publication.
    webUrl *string;
}
// Instantiates a new itemPublication and sets the default values.
func NewItemPublication()(*ItemPublication) {
    m := &ItemPublication{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the description property value. Description of the publication.
func (m *ItemPublication) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Title of the publication.
func (m *ItemPublication) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the publishedDate property value. The date that the publication was published.
func (m *ItemPublication) GetPublishedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publishedDate
    }
}
// Gets the publisher property value. Publication or publisher for the publication.
func (m *ItemPublication) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the thumbnailUrl property value. URL referencing a thumbnail of the publication.
func (m *ItemPublication) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// Gets the webUrl property value. URL referencing the publication.
func (m *ItemPublication) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *ItemPublication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["publishedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublishedDate(val)
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublisher(val)
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThumbnailUrl(val)
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *ItemPublication) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ItemPublication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteStringValue("publishedDate", m.GetPublishedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
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
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. Description of the publication.
// Parameters:
//  - value : Value to set for the description property.
func (m *ItemPublication) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Title of the publication.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ItemPublication) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the publishedDate property value. The date that the publication was published.
// Parameters:
//  - value : Value to set for the publishedDate property.
func (m *ItemPublication) SetPublishedDate(value *string)() {
    m.publishedDate = value
}
// Sets the publisher property value. Publication or publisher for the publication.
// Parameters:
//  - value : Value to set for the publisher property.
func (m *ItemPublication) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the thumbnailUrl property value. URL referencing a thumbnail of the publication.
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *ItemPublication) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// Sets the webUrl property value. URL referencing the publication.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *ItemPublication) SetWebUrl(value *string)() {
    m.webUrl = value
}
