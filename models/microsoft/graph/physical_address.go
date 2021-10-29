package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PhysicalAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The city.
    city *string;
    // The country or region. It's a free-format string value, for example, 'United States'.
    countryOrRegion *string;
    // The postal code.
    postalCode *string;
    // The post office box number.
    postOfficeBox *string;
    // The state.
    state *string;
    // The street.
    street *string;
    // The type of address. Possible values are: unknown, home, business, other.
    type_escaped *PhysicalAddressType;
}
// Instantiates a new physicalAddress and sets the default values.
func NewPhysicalAddress()(*PhysicalAddress) {
    m := &PhysicalAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PhysicalAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the city property value. The city.
func (m *PhysicalAddress) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// Gets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalAddress) GetCountryOrRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryOrRegion
    }
}
// Gets the postalCode property value. The postal code.
func (m *PhysicalAddress) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// Gets the postOfficeBox property value. The post office box number.
func (m *PhysicalAddress) GetPostOfficeBox()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postOfficeBox
    }
}
// Gets the state property value. The state.
func (m *PhysicalAddress) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the street property value. The street.
func (m *PhysicalAddress) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
// Gets the type_escaped property value. The type of address. Possible values are: unknown, home, business, other.
func (m *PhysicalAddress) GetType_escaped()(*PhysicalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PhysicalAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the city property value. The city.
// Parameters:
//  - value : Value to set for the city property.
func (m *PhysicalAddress) SetCity(value *string)() {
    m.city = value
}
// Sets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
// Parameters:
//  - value : Value to set for the countryOrRegion property.
func (m *PhysicalAddress) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// Sets the postalCode property value. The postal code.
// Parameters:
//  - value : Value to set for the postalCode property.
func (m *PhysicalAddress) SetPostalCode(value *string)() {
    m.postalCode = value
}
// Sets the postOfficeBox property value. The post office box number.
// Parameters:
//  - value : Value to set for the postOfficeBox property.
func (m *PhysicalAddress) SetPostOfficeBox(value *string)() {
    m.postOfficeBox = value
}
// Sets the state property value. The state.
// Parameters:
//  - value : Value to set for the state property.
func (m *PhysicalAddress) SetState(value *string)() {
    m.state = value
}
// Sets the street property value. The street.
// Parameters:
//  - value : Value to set for the street property.
func (m *PhysicalAddress) SetStreet(value *string)() {
    m.street = value
}
// Sets the type_escaped property value. The type of address. Possible values are: unknown, home, business, other.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *PhysicalAddress) SetType_escaped(value *PhysicalAddressType)() {
    m.type_escaped = value
}
