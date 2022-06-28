package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptRunSummary contains properties for the run summary of a device management script.
type DeviceManagementScriptRunSummary struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Error device count.
    errorDeviceCount *int32
    // Error user count.
    errorUserCount *int32
    // Success device count.
    successDeviceCount *int32
    // Success user count.
    successUserCount *int32
}
// NewDeviceManagementScriptRunSummary instantiates a new deviceManagementScriptRunSummary and sets the default values.
func NewDeviceManagementScriptRunSummary()(*DeviceManagementScriptRunSummary) {
    m := &DeviceManagementScriptRunSummary{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementScriptRunSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptRunSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptRunSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementScriptRunSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Error device count.
func (m *DeviceManagementScriptRunSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetErrorUserCount gets the errorUserCount property value. Error user count.
func (m *DeviceManagementScriptRunSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptRunSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["errorUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorUserCount(val)
        }
        return nil
    }
    res["successDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessDeviceCount(val)
        }
        return nil
    }
    res["successUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessUserCount(val)
        }
        return nil
    }
    return res
}
// GetSuccessDeviceCount gets the successDeviceCount property value. Success device count.
func (m *DeviceManagementScriptRunSummary) GetSuccessDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successDeviceCount
    }
}
// GetSuccessUserCount gets the successUserCount property value. Success user count.
func (m *DeviceManagementScriptRunSummary) GetSuccessUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successUserCount
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptRunSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorUserCount", m.GetErrorUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successDeviceCount", m.GetSuccessDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successUserCount", m.GetSuccessUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementScriptRunSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Error device count.
func (m *DeviceManagementScriptRunSummary) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetErrorUserCount sets the errorUserCount property value. Error user count.
func (m *DeviceManagementScriptRunSummary) SetErrorUserCount(value *int32)() {
    if m != nil {
        m.errorUserCount = value
    }
}
// SetSuccessDeviceCount sets the successDeviceCount property value. Success device count.
func (m *DeviceManagementScriptRunSummary) SetSuccessDeviceCount(value *int32)() {
    if m != nil {
        m.successDeviceCount = value
    }
}
// SetSuccessUserCount sets the successUserCount property value. Success user count.
func (m *DeviceManagementScriptRunSummary) SetSuccessUserCount(value *int32)() {
    if m != nil {
        m.successUserCount = value
    }
}
