package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceSummarizedAppState event representing a user's devices with failed or pending apps.
type ManagedDeviceSummarizedAppState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // DeviceId of device represented by this object
    deviceId *string
    // The OdataType property
    odataType *string
    // Indicates the type of execution status of the device management script.
    summarizedAppState *RunState
}
// NewManagedDeviceSummarizedAppState instantiates a new managedDeviceSummarizedAppState and sets the default values.
func NewManagedDeviceSummarizedAppState()(*ManagedDeviceSummarizedAppState) {
    m := &ManagedDeviceSummarizedAppState{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateManagedDeviceSummarizedAppStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceSummarizedAppStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceSummarizedAppState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceSummarizedAppState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceId gets the deviceId property value. DeviceId of device represented by this object
func (m *ManagedDeviceSummarizedAppState) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceSummarizedAppState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["summarizedAppState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedDeviceSummarizedAppState) GetOdataType()(*string) {
    return m.odataType
}
// GetSummarizedAppState gets the summarizedAppState property value. Indicates the type of execution status of the device management script.
func (m *ManagedDeviceSummarizedAppState) GetSummarizedAppState()(*RunState) {
    return m.summarizedAppState
}
// Serialize serializes information the current object
func (m *ManagedDeviceSummarizedAppState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *ManagedDeviceSummarizedAppState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceId sets the deviceId property value. DeviceId of device represented by this object
func (m *ManagedDeviceSummarizedAppState) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedDeviceSummarizedAppState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSummarizedAppState sets the summarizedAppState property value. Indicates the type of execution status of the device management script.
func (m *ManagedDeviceSummarizedAppState) SetSummarizedAppState(value *RunState)() {
    m.summarizedAppState = value
}
