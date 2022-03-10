package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonAnnotation provides operations to manage the compliance singleton.
type PersonAnnotation struct {
    ItemFacet
    // Contains the detail of the note itself.
    detail ItemBodyable;
    // Contains a friendly name for the note.
    displayName *string;
    // 
    thumbnailUrl *string;
}
// NewPersonAnnotation instantiates a new personAnnotation and sets the default values.
func NewPersonAnnotation()(*PersonAnnotation) {
    m := &PersonAnnotation{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreatePersonAnnotationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonAnnotationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPersonAnnotation(), nil
}
// GetDetail gets the detail property value. Contains the detail of the note itself.
func (m *PersonAnnotation) GetDetail()(ItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the note.
func (m *PersonAnnotation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonAnnotation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(ItemBodyable))
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
// GetThumbnailUrl gets the thumbnailUrl property value. 
func (m *PersonAnnotation) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
func (m *PersonAnnotation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDetail sets the detail property value. Contains the detail of the note itself.
func (m *PersonAnnotation) SetDetail(value ItemBodyable)() {
    if m != nil {
        m.detail = value
    }
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the note.
func (m *PersonAnnotation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. 
func (m *PersonAnnotation) SetThumbnailUrl(value *string)() {
    if m != nil {
        m.thumbnailUrl = value
    }
}
