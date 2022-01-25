package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonAward 
type PersonAward struct {
    ItemFacet
    // Descpription of the award or honor.
    description *string;
    // Name of the award or honor.
    displayName *string;
    // The date that the award or honor was granted.
    issuedDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // Authority which granted the award or honor.
    issuingAuthority *string;
    // URL referencing a thumbnail of the award or honor.
    thumbnailUrl *string;
    // URL referencing the award or honor.
    webUrl *string;
}
// NewPersonAward instantiates a new personAward and sets the default values.
func NewPersonAward()(*PersonAward) {
    m := &PersonAward{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetDescription gets the description property value. Descpription of the award or honor.
func (m *PersonAward) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Name of the award or honor.
func (m *PersonAward) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIssuedDate gets the issuedDate property value. The date that the award or honor was granted.
func (m *PersonAward) GetIssuedDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.issuedDate
    }
}
// GetIssuingAuthority gets the issuingAuthority property value. Authority which granted the award or honor.
func (m *PersonAward) GetIssuingAuthority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuingAuthority
    }
}
// GetThumbnailUrl gets the thumbnailUrl property value. URL referencing a thumbnail of the award or honor.
func (m *PersonAward) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// GetWebUrl gets the webUrl property value. URL referencing the award or honor.
func (m *PersonAward) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonAward) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["issuedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuedDate(val)
        }
        return nil
    }
    res["issuingAuthority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuingAuthority(val)
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
func (m *PersonAward) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PersonAward) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteDateOnlyValue("issuedDate", m.GetIssuedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuingAuthority", m.GetIssuingAuthority())
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
// SetDescription sets the description property value. Descpription of the award or honor.
func (m *PersonAward) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Name of the award or honor.
func (m *PersonAward) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIssuedDate sets the issuedDate property value. The date that the award or honor was granted.
func (m *PersonAward) SetIssuedDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.issuedDate = value
    }
}
// SetIssuingAuthority sets the issuingAuthority property value. Authority which granted the award or honor.
func (m *PersonAward) SetIssuingAuthority(value *string)() {
    if m != nil {
        m.issuingAuthority = value
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. URL referencing a thumbnail of the award or honor.
func (m *PersonAward) SetThumbnailUrl(value *string)() {
    if m != nil {
        m.thumbnailUrl = value
    }
}
// SetWebUrl sets the webUrl property value. URL referencing the award or honor.
func (m *PersonAward) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
