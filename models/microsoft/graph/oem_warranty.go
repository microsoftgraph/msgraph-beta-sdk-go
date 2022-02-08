package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OemWarranty 
type OemWarranty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of additional warranty offers. This collection can contain a maximum of 100 elements.
    additionalWarranties []WarrantyOffer;
    // List of base warranty offers. This collection can contain a maximum of 100 elements.
    baseWarranties []WarrantyOffer;
    // Device configuration page URL
    deviceConfigurationUrl *string;
    // Device warranty page URL
    deviceWarrantyUrl *string;
}
// NewOemWarranty instantiates a new oemWarranty and sets the default values.
func NewOemWarranty()(*OemWarranty) {
    m := &OemWarranty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OemWarranty) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdditionalWarranties gets the additionalWarranties property value. List of additional warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) GetAdditionalWarranties()([]WarrantyOffer) {
    if m == nil {
        return nil
    } else {
        return m.additionalWarranties
    }
}
// GetBaseWarranties gets the baseWarranties property value. List of base warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) GetBaseWarranties()([]WarrantyOffer) {
    if m == nil {
        return nil
    } else {
        return m.baseWarranties
    }
}
// GetDeviceConfigurationUrl gets the deviceConfigurationUrl property value. Device configuration page URL
func (m *OemWarranty) GetDeviceConfigurationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationUrl
    }
}
// GetDeviceWarrantyUrl gets the deviceWarrantyUrl property value. Device warranty page URL
func (m *OemWarranty) GetDeviceWarrantyUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceWarrantyUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OemWarranty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["additionalWarranties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWarrantyOffer() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WarrantyOffer, len(val))
            for i, v := range val {
                res[i] = *(v.(*WarrantyOffer))
            }
            m.SetAdditionalWarranties(res)
        }
        return nil
    }
    res["baseWarranties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWarrantyOffer() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WarrantyOffer, len(val))
            for i, v := range val {
                res[i] = *(v.(*WarrantyOffer))
            }
            m.SetBaseWarranties(res)
        }
        return nil
    }
    res["deviceConfigurationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationUrl(val)
        }
        return nil
    }
    res["deviceWarrantyUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceWarrantyUrl(val)
        }
        return nil
    }
    return res
}
func (m *OemWarranty) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OemWarranty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAdditionalWarranties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalWarranties()))
        for i, v := range m.GetAdditionalWarranties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("additionalWarranties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBaseWarranties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBaseWarranties()))
        for i, v := range m.GetBaseWarranties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("baseWarranties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceConfigurationUrl", m.GetDeviceConfigurationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceWarrantyUrl", m.GetDeviceWarrantyUrl())
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
func (m *OemWarranty) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdditionalWarranties sets the additionalWarranties property value. List of additional warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) SetAdditionalWarranties(value []WarrantyOffer)() {
    if m != nil {
        m.additionalWarranties = value
    }
}
// SetBaseWarranties sets the baseWarranties property value. List of base warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) SetBaseWarranties(value []WarrantyOffer)() {
    if m != nil {
        m.baseWarranties = value
    }
}
// SetDeviceConfigurationUrl sets the deviceConfigurationUrl property value. Device configuration page URL
func (m *OemWarranty) SetDeviceConfigurationUrl(value *string)() {
    if m != nil {
        m.deviceConfigurationUrl = value
    }
}
// SetDeviceWarrantyUrl sets the deviceWarrantyUrl property value. Device warranty page URL
func (m *OemWarranty) SetDeviceWarrantyUrl(value *string)() {
    if m != nil {
        m.deviceWarrantyUrl = value
    }
}
