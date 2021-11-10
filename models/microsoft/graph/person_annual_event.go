package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PersonAnnualEvent struct {
    ItemFacet
    // 
    date *string;
    // 
    displayName *string;
    // 
    type_escaped *PersonAnnualEventType;
}
// Instantiates a new personAnnualEvent and sets the default values.
func NewPersonAnnualEvent()(*PersonAnnualEvent) {
    m := &PersonAnnualEvent{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the date property value. 
func (m *PersonAnnualEvent) GetDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.date
    }
}
// Gets the displayName property value. 
func (m *PersonAnnualEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the type_escaped property value. 
func (m *PersonAnnualEvent) GetType_escaped()(*PersonAnnualEventType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *PersonAnnualEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["date"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
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
        val, err := n.GetEnumValue(ParsePersonAnnualEventType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PersonAnnualEventType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    return res
}
func (m *PersonAnnualEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PersonAnnualEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("date", m.GetDate())
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
// Sets the date property value. 
// Parameters:
//  - value : Value to set for the date property.
func (m *PersonAnnualEvent) SetDate(value *string)() {
    m.date = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PersonAnnualEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *PersonAnnualEvent) SetType_escaped(value *PersonAnnualEventType)() {
    m.type_escaped = value
}
