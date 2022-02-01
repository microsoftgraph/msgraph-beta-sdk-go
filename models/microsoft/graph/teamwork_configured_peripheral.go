package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkConfiguredPeripheral 
type TeamworkConfiguredPeripheral struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // True if the current peripheral is optional. If set to false, this property is also used as part of the calculation of the health state for the device.
    isOptional *bool;
    // 
    peripheral *TeamworkPeripheral;
}
// NewTeamworkConfiguredPeripheral instantiates a new teamworkConfiguredPeripheral and sets the default values.
func NewTeamworkConfiguredPeripheral()(*TeamworkConfiguredPeripheral) {
    m := &TeamworkConfiguredPeripheral{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkConfiguredPeripheral) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsOptional gets the isOptional property value. True if the current peripheral is optional. If set to false, this property is also used as part of the calculation of the health state for the device.
func (m *TeamworkConfiguredPeripheral) GetIsOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOptional
    }
}
// GetPeripheral gets the peripheral property value. 
func (m *TeamworkConfiguredPeripheral) GetPeripheral()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.peripheral
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkConfiguredPeripheral) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOptional(val)
        }
        return nil
    }
    res["peripheral"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeripheral(val.(*TeamworkPeripheral))
        }
        return nil
    }
    return res
}
func (m *TeamworkConfiguredPeripheral) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkConfiguredPeripheral) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isOptional", m.GetIsOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("peripheral", m.GetPeripheral())
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
func (m *TeamworkConfiguredPeripheral) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsOptional sets the isOptional property value. True if the current peripheral is optional. If set to false, this property is also used as part of the calculation of the health state for the device.
func (m *TeamworkConfiguredPeripheral) SetIsOptional(value *bool)() {
    if m != nil {
        m.isOptional = value
    }
}
// SetPeripheral sets the peripheral property value. 
func (m *TeamworkConfiguredPeripheral) SetPeripheral(value *TeamworkPeripheral)() {
    if m != nil {
        m.peripheral = value
    }
}
