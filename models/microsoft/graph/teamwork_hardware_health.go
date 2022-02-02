package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkHardwareHealth 
type TeamworkHardwareHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The system health details for a teamworkDevice.
    computeHealth *TeamworkPeripheralHealth;
    // The health details about the HDMI ingest of a device.
    hdmiIngestHealth *TeamworkPeripheralHealth;
}
// NewTeamworkHardwareHealth instantiates a new teamworkHardwareHealth and sets the default values.
func NewTeamworkHardwareHealth()(*TeamworkHardwareHealth) {
    m := &TeamworkHardwareHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkHardwareHealth) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetComputeHealth gets the computeHealth property value. The system health details for a teamworkDevice.
func (m *TeamworkHardwareHealth) GetComputeHealth()(*TeamworkPeripheralHealth) {
    if m == nil {
        return nil
    } else {
        return m.computeHealth
    }
}
// GetHdmiIngestHealth gets the hdmiIngestHealth property value. The health details about the HDMI ingest of a device.
func (m *TeamworkHardwareHealth) GetHdmiIngestHealth()(*TeamworkPeripheralHealth) {
    if m == nil {
        return nil
    } else {
        return m.hdmiIngestHealth
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkHardwareHealth) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["computeHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeHealth(val.(*TeamworkPeripheralHealth))
        }
        return nil
    }
    res["hdmiIngestHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHdmiIngestHealth(val.(*TeamworkPeripheralHealth))
        }
        return nil
    }
    return res
}
func (m *TeamworkHardwareHealth) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkHardwareHealth) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("computeHealth", m.GetComputeHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hdmiIngestHealth", m.GetHdmiIngestHealth())
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
func (m *TeamworkHardwareHealth) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetComputeHealth sets the computeHealth property value. The system health details for a teamworkDevice.
func (m *TeamworkHardwareHealth) SetComputeHealth(value *TeamworkPeripheralHealth)() {
    if m != nil {
        m.computeHealth = value
    }
}
// SetHdmiIngestHealth sets the hdmiIngestHealth property value. The health details about the HDMI ingest of a device.
func (m *TeamworkHardwareHealth) SetHdmiIngestHealth(value *TeamworkPeripheralHealth)() {
    if m != nil {
        m.hdmiIngestHealth = value
    }
}
