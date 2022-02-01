package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkHardwareDetail 
type TeamworkHardwareDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // MAC address.
    macAddresses []string;
    // Device manufacturer.
    manufacturer *string;
    // Devie model.
    model *string;
    // Device serial number.
    serialNumber *string;
    // The unique identifier for the device.
    uniqueId *string;
}
// NewTeamworkHardwareDetail instantiates a new teamworkHardwareDetail and sets the default values.
func NewTeamworkHardwareDetail()(*TeamworkHardwareDetail) {
    m := &TeamworkHardwareDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkHardwareDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMacAddresses gets the macAddresses property value. MAC address.
func (m *TeamworkHardwareDetail) GetMacAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.macAddresses
    }
}
// GetManufacturer gets the manufacturer property value. Device manufacturer.
func (m *TeamworkHardwareDetail) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. Devie model.
func (m *TeamworkHardwareDetail) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetSerialNumber gets the serialNumber property value. Device serial number.
func (m *TeamworkHardwareDetail) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// GetUniqueId gets the uniqueId property value. The unique identifier for the device.
func (m *TeamworkHardwareDetail) GetUniqueId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkHardwareDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["macAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMacAddresses(res)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["uniqueId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueId(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkHardwareDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkHardwareDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMacAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("macAddresses", m.GetMacAddresses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uniqueId", m.GetUniqueId())
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
func (m *TeamworkHardwareDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMacAddresses sets the macAddresses property value. MAC address.
func (m *TeamworkHardwareDetail) SetMacAddresses(value []string)() {
    if m != nil {
        m.macAddresses = value
    }
}
// SetManufacturer sets the manufacturer property value. Device manufacturer.
func (m *TeamworkHardwareDetail) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. Devie model.
func (m *TeamworkHardwareDetail) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetSerialNumber sets the serialNumber property value. Device serial number.
func (m *TeamworkHardwareDetail) SetSerialNumber(value *string)() {
    if m != nil {
        m.serialNumber = value
    }
}
// SetUniqueId sets the uniqueId property value. The unique identifier for the device.
func (m *TeamworkHardwareDetail) SetUniqueId(value *string)() {
    if m != nil {
        m.uniqueId = value
    }
}
