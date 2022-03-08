package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemPatent provides operations to manage the compliance singleton.
type ItemPatent struct {
    ItemFacet
    // Descpription of the patent or filing.
    description *string;
    // Title of the patent or filing.
    displayName *string;
    // Indicates the patent is pending.
    isPending *bool;
    // The date that the patent was granted.
    issuedDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // Authority which granted the patent.
    issuingAuthority *string;
    // The patent number.
    number *string;
    // URL referencing the patent or filing.
    webUrl *string;
}
// NewItemPatent instantiates a new itemPatent and sets the default values.
func NewItemPatent()(*ItemPatent) {
    m := &ItemPatent{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreateItemPatentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPatentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewItemPatent(), nil
}
// GetDescription gets the description property value. Descpription of the patent or filing.
func (m *ItemPatent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Title of the patent or filing.
func (m *ItemPatent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetIsPending gets the isPending property value. Indicates the patent is pending.
func (m *ItemPatent) GetIsPending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPending
    }
}
// GetIssuedDate gets the issuedDate property value. The date that the patent was granted.
func (m *ItemPatent) GetIssuedDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.issuedDate
    }
}
// GetIssuingAuthority gets the issuingAuthority property value. Authority which granted the patent.
func (m *ItemPatent) GetIssuingAuthority()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuingAuthority
    }
}
// GetNumber gets the number property value. The patent number.
func (m *ItemPatent) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetWebUrl gets the webUrl property value. URL referencing the patent or filing.
func (m *ItemPatent) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *ItemPatent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDescription sets the description property value. Descpription of the patent or filing.
func (m *ItemPatent) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Title of the patent or filing.
func (m *ItemPatent) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsPending sets the isPending property value. Indicates the patent is pending.
func (m *ItemPatent) SetIsPending(value *bool)() {
    if m != nil {
        m.isPending = value
    }
}
// SetIssuedDate sets the issuedDate property value. The date that the patent was granted.
func (m *ItemPatent) SetIssuedDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.issuedDate = value
    }
}
// SetIssuingAuthority sets the issuingAuthority property value. Authority which granted the patent.
func (m *ItemPatent) SetIssuingAuthority(value *string)() {
    if m != nil {
        m.issuingAuthority = value
    }
}
// SetNumber sets the number property value. The patent number.
func (m *ItemPatent) SetNumber(value *string)() {
    if m != nil {
        m.number = value
    }
}
// SetWebUrl sets the webUrl property value. URL referencing the patent or filing.
func (m *ItemPatent) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
