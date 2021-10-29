package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PersonAward struct {
    ItemFacet
    // Descpription of the award or honor.
    description *string;
    // Name of the award or honor.
    displayName *string;
    // The date that the award or honor was granted.
    issuedDate *string;
    // Authority which granted the award or honor.
    issuingAuthority *string;
    // URL referencing a thumbnail of the award or honor.
    thumbnailUrl *string;
    // URL referencing the award or honor.
    webUrl *string;
}
// Instantiates a new personAward and sets the default values.
func NewPersonAward()(*PersonAward) {
    m := &PersonAward{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the description property value. Descpription of the award or honor.
func (m *PersonAward) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Name of the award or honor.
func (m *PersonAward) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the issuedDate property value. The date that the award or honor was granted.
func (m *PersonAward) GetIssuedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuedDate
    }
}
// Gets the issuingAuthority property value. Authority which granted the award or honor.
func (m *PersonAward) GetIssuingAuthority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuingAuthority
    }
}
// Gets the thumbnailUrl property value. URL referencing a thumbnail of the award or honor.
func (m *PersonAward) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// Gets the webUrl property value. URL referencing the award or honor.
func (m *PersonAward) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *PersonAward) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["issuedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuedDate(val)
        return nil
    }
    res["issuingAuthority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIssuingAuthority(val)
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
func (m *PersonAward) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err = writer.WriteStringValue("issuedDate", m.GetIssuedDate())
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
// Sets the description property value. Descpription of the award or honor.
// Parameters:
//  - value : Value to set for the description property.
func (m *PersonAward) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Name of the award or honor.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PersonAward) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the issuedDate property value. The date that the award or honor was granted.
// Parameters:
//  - value : Value to set for the issuedDate property.
func (m *PersonAward) SetIssuedDate(value *string)() {
    m.issuedDate = value
}
// Sets the issuingAuthority property value. Authority which granted the award or honor.
// Parameters:
//  - value : Value to set for the issuingAuthority property.
func (m *PersonAward) SetIssuingAuthority(value *string)() {
    m.issuingAuthority = value
}
// Sets the thumbnailUrl property value. URL referencing a thumbnail of the award or honor.
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *PersonAward) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// Sets the webUrl property value. URL referencing the award or honor.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *PersonAward) SetWebUrl(value *string)() {
    m.webUrl = value
}
