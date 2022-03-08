package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceSummarizedAppState provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
type ManagedDeviceSummarizedAppState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // DeviceId of device represented by this object
    deviceId *string;
    // runState for the object. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
    summarizedAppState *RunState;
}
// NewManagedDeviceSummarizedAppState instantiates a new managedDeviceSummarizedAppState and sets the default values.
func NewManagedDeviceSummarizedAppState()(*ManagedDeviceSummarizedAppState) {
    m := &ManagedDeviceSummarizedAppState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagedDeviceSummarizedAppStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceSummarizedAppStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagedDeviceSummarizedAppState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceSummarizedAppState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceId gets the deviceId property value. DeviceId of device represented by this object
func (m *ManagedDeviceSummarizedAppState) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceSummarizedAppState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["summarizedAppState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummarizedAppState(val.(*RunState))
        }
        return nil
    }
    return res
}
// GetSummarizedAppState gets the summarizedAppState property value. runState for the object. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
func (m *ManagedDeviceSummarizedAppState) GetSummarizedAppState()(*RunState) {
    if m == nil {
        return nil
    } else {
        return m.summarizedAppState
    }
}
func (m *ManagedDeviceSummarizedAppState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceSummarizedAppState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetSummarizedAppState() != nil {
        cast := (*m.GetSummarizedAppState()).String()
        err := writer.WriteStringValue("summarizedAppState", &cast)
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
func (m *ManagedDeviceSummarizedAppState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceId sets the deviceId property value. DeviceId of device represented by this object
func (m *ManagedDeviceSummarizedAppState) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetSummarizedAppState sets the summarizedAppState property value. runState for the object. Possible values are: unknown, success, fail, scriptError, pending, notApplicable.
func (m *ManagedDeviceSummarizedAppState) SetSummarizedAppState(value *RunState)() {
    if m != nil {
        m.summarizedAppState = value
    }
}
