package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PostalAddressType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    city *string;
    // 
    countryLetterCode *string;
    // 
    postalCode *string;
    // 
    state *string;
    // 
    street *string;
}
// Instantiates a new postalAddressType and sets the default values.
func NewPostalAddressType()(*PostalAddressType) {
    m := &PostalAddressType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PostalAddressType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the city property value. 
func (m *PostalAddressType) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// Gets the countryLetterCode property value. 
func (m *PostalAddressType) GetCountryLetterCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryLetterCode
    }
}
// Gets the postalCode property value. 
func (m *PostalAddressType) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// Gets the state property value. 
func (m *PostalAddressType) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the street property value. 
func (m *PostalAddressType) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
// The deserialization information for the current model
func (m *PostalAddressType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCity(val)
        return nil
    }
    res["countryLetterCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryLetterCode(val)
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
    return res
}
func (m *PostalAddressType) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PostalAddressType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryLetterCode", m.GetCountryLetterCode())
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
func (m *PostalAddressType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the city property value. 
// Parameters:
//  - value : Value to set for the city property.
func (m *PostalAddressType) SetCity(value *string)() {
    m.city = value
}
// Sets the countryLetterCode property value. 
// Parameters:
//  - value : Value to set for the countryLetterCode property.
func (m *PostalAddressType) SetCountryLetterCode(value *string)() {
    m.countryLetterCode = value
}
// Sets the postalCode property value. 
// Parameters:
//  - value : Value to set for the postalCode property.
func (m *PostalAddressType) SetPostalCode(value *string)() {
    m.postalCode = value
}
// Sets the state property value. 
// Parameters:
//  - value : Value to set for the state property.
func (m *PostalAddressType) SetState(value *string)() {
    m.state = value
}
// Sets the street property value. 
// Parameters:
//  - value : Value to set for the street property.
func (m *PostalAddressType) SetStreet(value *string)() {
    m.street = value
}
