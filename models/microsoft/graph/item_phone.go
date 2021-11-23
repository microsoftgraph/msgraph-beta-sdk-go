package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// itemPhone 
type ItemPhone struct {
    ItemFacet
    // Friendly name the user has assigned this phone number.
    displayName *string;
    // Phone number provided by the user.
    number *string;
    // The type of phone number within the object. Possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio.
    type_escaped *PhoneType;
}
// NewItemPhone instantiates a new itemPhone and sets the default values.
func NewItemPhone()(*ItemPhone) {
    m := &ItemPhone{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Friendly name the user has assigned this phone number.
func (m *ItemPhone) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetNumber gets the number property value. Phone number provided by the user.
func (m *ItemPhone) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetType_escaped gets the type_escaped property value. The type of phone number within the object. Possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio.
func (m *ItemPhone) GetType_escaped()(*PhoneType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPhone) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePhoneType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PhoneType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    return res
}
func (m *ItemPhone) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ItemPhone) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Friendly name the user has assigned this phone number.
func (m *ItemPhone) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetNumber sets the number property value. Phone number provided by the user.
func (m *ItemPhone) SetNumber(value *string)() {
    m.number = value
}
// SetType_escaped sets the type_escaped property value. The type of phone number within the object. Possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio.
func (m *ItemPhone) SetType_escaped(value *PhoneType)() {
    m.type_escaped = value
}
