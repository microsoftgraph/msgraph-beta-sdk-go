package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkHardwareHealth 
type TeamworkHardwareHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The system health details for a teamworkDevice.
    computeHealth TeamworkPeripheralHealthable;
    // The health details about the HDMI ingest of a device.
    hdmiIngestHealth TeamworkPeripheralHealthable;
}
// NewTeamworkHardwareHealth instantiates a new teamworkHardwareHealth and sets the default values.
func NewTeamworkHardwareHealth()(*TeamworkHardwareHealth) {
    m := &TeamworkHardwareHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkHardwareHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkHardwareHealthFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamworkHardwareHealth(), nil
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
func (m *TeamworkHardwareHealth) GetComputeHealth()(TeamworkPeripheralHealthable) {
    if m == nil {
        return nil
    } else {
        return m.computeHealth
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkHardwareHealth) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["computeHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["hdmiIngestHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHdmiIngestHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    return res
}
// GetHdmiIngestHealth gets the hdmiIngestHealth property value. The health details about the HDMI ingest of a device.
func (m *TeamworkHardwareHealth) GetHdmiIngestHealth()(TeamworkPeripheralHealthable) {
    if m == nil {
        return nil
    } else {
        return m.hdmiIngestHealth
    }
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
func (m *TeamworkHardwareHealth) SetComputeHealth(value TeamworkPeripheralHealthable)() {
    if m != nil {
        m.computeHealth = value
    }
}
// SetHdmiIngestHealth sets the hdmiIngestHealth property value. The health details about the HDMI ingest of a device.
func (m *TeamworkHardwareHealth) SetHdmiIngestHealth(value TeamworkPeripheralHealthable)() {
    if m != nil {
        m.hdmiIngestHealth = value
    }
}
