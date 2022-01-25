package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemPublication 
type ItemPublication struct {
    ItemFacet
    // Description of the publication.
    description *string;
    // Title of the publication.
    displayName *string;
    // The date that the publication was published.
    publishedDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // Publication or publisher for the publication.
    publisher *string;
    // URL referencing a thumbnail of the publication.
    thumbnailUrl *string;
    // URL referencing the publication.
    webUrl *string;
}
// NewItemPublication instantiates a new itemPublication and sets the default values.
func NewItemPublication()(*ItemPublication) {
    m := &ItemPublication{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetDescription gets the description property value. Description of the publication.
func (m *ItemPublication) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Title of the publication.
func (m *ItemPublication) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetPublishedDate gets the publishedDate property value. The date that the publication was published.
func (m *ItemPublication) GetPublishedDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.publishedDate
    }
}
// GetPublisher gets the publisher property value. Publication or publisher for the publication.
func (m *ItemPublication) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetThumbnailUrl gets the thumbnailUrl property value. URL referencing a thumbnail of the publication.
func (m *ItemPublication) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// GetWebUrl gets the webUrl property value. URL referencing the publication.
func (m *ItemPublication) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPublication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["publishedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDate(val)
        }
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
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
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
func (m *ItemPublication) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteDateOnlyValue("publishedDate", m.GetPublishedDate())
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
// SetDescription sets the description property value. Description of the publication.
func (m *ItemPublication) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Title of the publication.
func (m *ItemPublication) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPublishedDate sets the publishedDate property value. The date that the publication was published.
func (m *ItemPublication) SetPublishedDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.publishedDate = value
    }
}
// SetPublisher sets the publisher property value. Publication or publisher for the publication.
func (m *ItemPublication) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. URL referencing a thumbnail of the publication.
func (m *ItemPublication) SetThumbnailUrl(value *string)() {
    if m != nil {
        m.thumbnailUrl = value
    }
}
// SetWebUrl sets the webUrl property value. URL referencing the publication.
func (m *ItemPublication) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
