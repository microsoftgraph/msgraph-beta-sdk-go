package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemPatent struct {
    ItemFacet
    // Descpription of the patent or filing.
    description *string;
    // Title of the patent or filing.
    displayName *string;
    // Indicates the patent is pending.
    isPending *bool;
    // The date that the patent was granted.
    issuedDate *string;
    // Authority which granted the patent.
    issuingAuthority *string;
    // The patent number.
    number *string;
    // URL referencing the patent or filing.
    webUrl *string;
}
// Instantiates a new itemPatent and sets the default values.
func NewItemPatent()(*ItemPatent) {
    m := &ItemPatent{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the description property value. Descpription of the patent or filing.
func (m *ItemPatent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Title of the patent or filing.
func (m *ItemPatent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isPending property value. Indicates the patent is pending.
func (m *ItemPatent) GetIsPending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPending
    }
}
// Gets the issuedDate property value. The date that the patent was granted.
func (m *ItemPatent) GetIssuedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuedDate
    }
}
// Gets the issuingAuthority property value. Authority which granted the patent.
func (m *ItemPatent) GetIssuingAuthority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuingAuthority
    }
}
// Gets the number property value. The patent number.
func (m *ItemPatent) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the webUrl property value. URL referencing the patent or filing.
func (m *ItemPatent) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *ItemPatent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["isPending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPending(val)
        }
        return nil
    }
    res["issuedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
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
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
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
func (m *ItemPatent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ItemPatent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("isPending", m.GetIsPending())
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
        err = writer.WriteStringValue("number", m.GetNumber())
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
// Sets the description property value. Descpription of the patent or filing.
// Parameters:
//  - value : Value to set for the description property.
func (m *ItemPatent) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Title of the patent or filing.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ItemPatent) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isPending property value. Indicates the patent is pending.
// Parameters:
//  - value : Value to set for the isPending property.
func (m *ItemPatent) SetIsPending(value *bool)() {
    m.isPending = value
}
// Sets the issuedDate property value. The date that the patent was granted.
// Parameters:
//  - value : Value to set for the issuedDate property.
func (m *ItemPatent) SetIssuedDate(value *string)() {
    m.issuedDate = value
}
// Sets the issuingAuthority property value. Authority which granted the patent.
// Parameters:
//  - value : Value to set for the issuingAuthority property.
func (m *ItemPatent) SetIssuingAuthority(value *string)() {
    m.issuingAuthority = value
}
// Sets the number property value. The patent number.
// Parameters:
//  - value : Value to set for the number property.
func (m *ItemPatent) SetNumber(value *string)() {
    m.number = value
}
// Sets the webUrl property value. URL referencing the patent or filing.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *ItemPatent) SetWebUrl(value *string)() {
    m.webUrl = value
}
