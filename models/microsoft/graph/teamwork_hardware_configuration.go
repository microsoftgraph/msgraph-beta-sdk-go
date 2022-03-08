package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkHardwareConfiguration provides operations to manage the teamwork singleton.
type TeamworkHardwareConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    compute TeamworkPeripheralable;
    // 
    hdmiIngest TeamworkPeripheralable;
    // The CPU model on the device.
    processorModel *string;
}
// NewTeamworkHardwareConfiguration instantiates a new teamworkHardwareConfiguration and sets the default values.
func NewTeamworkHardwareConfiguration()(*TeamworkHardwareConfiguration) {
    m := &TeamworkHardwareConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkHardwareConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkHardwareConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamworkHardwareConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkHardwareConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompute gets the compute property value. 
func (m *TeamworkHardwareConfiguration) GetCompute()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.compute
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkHardwareConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compute"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompute(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["hdmiIngest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHdmiIngest(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["processorModel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorModel(val)
        }
        return nil
    }
    return res
}
// GetHdmiIngest gets the hdmiIngest property value. 
func (m *TeamworkHardwareConfiguration) GetHdmiIngest()(TeamworkPeripheralable) {
    if m == nil {
        return nil
    } else {
        return m.hdmiIngest
    }
}
// GetProcessorModel gets the processorModel property value. The CPU model on the device.
func (m *TeamworkHardwareConfiguration) GetProcessorModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processorModel
    }
}
func (m *TeamworkHardwareConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkHardwareConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("compute", m.GetCompute())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hdmiIngest", m.GetHdmiIngest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("processorModel", m.GetProcessorModel())
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
func (m *TeamworkHardwareConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompute sets the compute property value. 
func (m *TeamworkHardwareConfiguration) SetCompute(value TeamworkPeripheralable)() {
    if m != nil {
        m.compute = value
    }
}
// SetHdmiIngest sets the hdmiIngest property value. 
func (m *TeamworkHardwareConfiguration) SetHdmiIngest(value TeamworkPeripheralable)() {
    if m != nil {
        m.hdmiIngest = value
    }
}
// SetProcessorModel sets the processorModel property value. The CPU model on the device.
func (m *TeamworkHardwareConfiguration) SetProcessorModel(value *string)() {
    if m != nil {
        m.processorModel = value
    }
}
