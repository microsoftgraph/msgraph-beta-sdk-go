package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// personAnnualEvent 
type PersonAnnualEvent struct {
    ItemFacet
    // 
    date *string;
    // 
    displayName *string;
    // 
    type_escaped *PersonAnnualEventType;
}
// NewPersonAnnualEvent instantiates a new personAnnualEvent and sets the default values.
func NewPersonAnnualEvent()(*PersonAnnualEvent) {
    m := &PersonAnnualEvent{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetDate gets the date property value. 
func (m *PersonAnnualEvent) GetDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.date
    }
}
// GetDisplayName gets the displayName property value. 
func (m *PersonAnnualEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetType_escaped gets the type_escaped property value. 
func (m *PersonAnnualEvent) GetType_escaped()(*PersonAnnualEventType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetDate sets the date property value. 
func (m *PersonAnnualEvent) SetDate(value *string)() {
    m.date = value
}
// SetDisplayName sets the displayName property value. 
func (m *PersonAnnualEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetType_escaped sets the type_escaped property value. 
func (m *PersonAnnualEvent) SetType_escaped(value *PersonAnnualEventType)() {
    m.type_escaped = value
}
