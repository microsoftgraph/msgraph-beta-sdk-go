package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PostalAddressType 
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
// NewPostalAddressType instantiates a new postalAddressType and sets the default values.
func NewPostalAddressType()(*PostalAddressType) {
    m := &PostalAddressType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PostalAddressType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCity gets the city property value. 
func (m *PostalAddressType) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCountryLetterCode gets the countryLetterCode property value. 
func (m *PostalAddressType) GetCountryLetterCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryLetterCode
    }
}
// GetPostalCode gets the postalCode property value. 
func (m *PostalAddressType) GetPostalCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.postalCode
    }
}
// GetState gets the state property value. 
func (m *PostalAddressType) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStreet gets the street property value. 
func (m *PostalAddressType) GetStreet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.street
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PostalAddressType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryLetterCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryLetterCode(val)
        }
        return nil
    }
    res["postalCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["street"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreet(val)
        }
        return nil
    }
    return res
}
func (m *PostalAddressType) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PostalAddressType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCity sets the city property value. 
func (m *PostalAddressType) SetCity(value *string)() {
    if m != nil {
        m.city = value
    }
}
// SetCountryLetterCode sets the countryLetterCode property value. 
func (m *PostalAddressType) SetCountryLetterCode(value *string)() {
    if m != nil {
        m.countryLetterCode = value
    }
}
// SetPostalCode sets the postalCode property value. 
func (m *PostalAddressType) SetPostalCode(value *string)() {
    if m != nil {
        m.postalCode = value
    }
}
// SetState sets the state property value. 
func (m *PostalAddressType) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetStreet sets the street property value. 
func (m *PostalAddressType) SetStreet(value *string)() {
    if m != nil {
        m.street = value
    }
}
