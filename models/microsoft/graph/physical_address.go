package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PhysicalAddress struct {
    additionalData map[string]interface{};
    city *string;
    countryOrRegion *string;
    postalCode *string;
    postOfficeBox *string;
    state *string;
    street *string;
    type_escaped *PhysicalAddressType;
}
func NewPhysicalAddress()(*PhysicalAddress) {
    m := &PhysicalAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PhysicalAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PhysicalAddress) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
func (m *PhysicalAddress) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
func (m *PhysicalAddress) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
func (m *PhysicalAddress) GetPostOfficeBox()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postOfficeBox
    }
}
func (m *PhysicalAddress) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *PhysicalAddress) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
func (m *PhysicalAddress) GetType_escaped()(*PhysicalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *PhysicalAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCity(val)
        return nil
    }
    res["countryOrRegion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryOrRegion(val)
        return nil
    }
    res["postalCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostalCode(val)
        return nil
    }
    res["postOfficeBox"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPostOfficeBox(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
        return nil
    }
    res["street"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStreet(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePhysicalAddressType)
        if err != nil {
            return err
        }
        cast := val.(PhysicalAddressType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *PhysicalAddress) IsNil()(bool) {
    return m == nil
}
func (m *PhysicalAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryOrRegion", m.GetCountryOrRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postOfficeBox", m.GetPostOfficeBox())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("street", m.GetStreet())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PhysicalAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PhysicalAddress) SetCity(value *string)() {
    m.city = value
}
func (m *PhysicalAddress) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
func (m *PhysicalAddress) SetPostalCode(value *string)() {
    m.postalCode = value
}
func (m *PhysicalAddress) SetPostOfficeBox(value *string)() {
    m.postOfficeBox = value
}
func (m *PhysicalAddress) SetState(value *string)() {
    m.state = value
}
func (m *PhysicalAddress) SetStreet(value *string)() {
    m.street = value
}
func (m *PhysicalAddress) SetType_escaped(value *PhysicalAddressType)() {
    m.type_escaped = value
}
