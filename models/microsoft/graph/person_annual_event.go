package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PersonAnnualEvent struct {
    ItemFacet
    date *string;
    displayName *string;
    type_escpaped *PersonAnnualEventType;
}
func NewPersonAnnualEvent()(*PersonAnnualEvent) {
    m := &PersonAnnualEvent{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
func (m *PersonAnnualEvent) GetDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.date
    }
}
func (m *PersonAnnualEvent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *PersonAnnualEvent) GetType_escpaped()(*PersonAnnualEventType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *PersonAnnualEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["date"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDate(val)
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
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersonAnnualEventType)
        if err != nil {
            return err
        }
        cast := val.(PersonAnnualEventType)
        m.SetType_escpaped(&cast)
        return nil
    }
    return res
}
func (m *PersonAnnualEvent) IsNil()(bool) {
    return m == nil
}
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
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err = writer.WriteStringValue("type_escpaped", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PersonAnnualEvent) SetDate(value *string)() {
    m.date = value
}
func (m *PersonAnnualEvent) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *PersonAnnualEvent) SetType_escpaped(value *PersonAnnualEventType)() {
    m.type_escpaped = value
}
