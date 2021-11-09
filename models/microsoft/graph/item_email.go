package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemEmail struct {
    ItemFacet
    // The email address itself.
    address *string;
    // The name or label a user has associated with a particular email address.
    displayName *string;
    // The type of email address. Possible values are: unknown, work, personal, main, other.
    type_escaped *EmailType;
}
// Instantiates a new itemEmail and sets the default values.
func NewItemEmail()(*ItemEmail) {
    m := &ItemEmail{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the address property value. The email address itself.
func (m *ItemEmail) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the displayName property value. The name or label a user has associated with a particular email address.
func (m *ItemEmail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the type_escaped property value. The type of email address. Possible values are: unknown, work, personal, main, other.
func (m *ItemEmail) GetType_escaped()(*EmailType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *ItemEmail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EmailType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    return res
}
func (m *ItemEmail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ItemEmail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("address", m.GetAddress())
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
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the address property value. The email address itself.
// Parameters:
//  - value : Value to set for the address property.
func (m *ItemEmail) SetAddress(value *string)() {
    m.address = value
}
// Sets the displayName property value. The name or label a user has associated with a particular email address.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ItemEmail) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the type_escaped property value. The type of email address. Possible values are: unknown, work, personal, main, other.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *ItemEmail) SetType_escaped(value *EmailType)() {
    m.type_escaped = value
}
