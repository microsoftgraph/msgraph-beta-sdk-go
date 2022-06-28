package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteLockActionResult 
type RemoteLockActionResult struct {
    DeviceActionResult
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Pin to unlock the client
    unlockPin *string
}
// NewRemoteLockActionResult instantiates a new RemoteLockActionResult and sets the default values.
func NewRemoteLockActionResult()(*RemoteLockActionResult) {
    m := &RemoteLockActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRemoteLockActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoteLockActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteLockActionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoteLockActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoteLockActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["unlockPin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnlockPin(val)
        }
        return nil
    }
    return res
}
// GetUnlockPin gets the unlockPin property value. Pin to unlock the client
func (m *RemoteLockActionResult) GetUnlockPin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unlockPin
    }
}
// Serialize serializes information the current object
func (m *RemoteLockActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("unlockPin", m.GetUnlockPin())
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
func (m *RemoteLockActionResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUnlockPin sets the unlockPin property value. Pin to unlock the client
func (m *RemoteLockActionResult) SetUnlockPin(value *string)() {
    if m != nil {
        m.unlockPin = value
    }
}
