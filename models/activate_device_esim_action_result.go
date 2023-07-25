package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActivateDeviceEsimActionResult activate device eSIM action result
type ActivateDeviceEsimActionResult struct {
    DeviceActionResult
}
// NewActivateDeviceEsimActionResult instantiates a new activateDeviceEsimActionResult and sets the default values.
func NewActivateDeviceEsimActionResult()(*ActivateDeviceEsimActionResult) {
    m := &ActivateDeviceEsimActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    return m
}
// CreateActivateDeviceEsimActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActivateDeviceEsimActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActivateDeviceEsimActionResult(), nil
}
// GetCarrierUrl gets the carrierUrl property value. Carrier Url to activate the device eSIM
func (m *ActivateDeviceEsimActionResult) GetCarrierUrl()(*string) {
    val, err := m.GetBackingStore().Get("carrierUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivateDeviceEsimActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["carrierUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ActivateDeviceEsimActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("carrierUrl", m.GetCarrierUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCarrierUrl sets the carrierUrl property value. Carrier Url to activate the device eSIM
func (m *ActivateDeviceEsimActionResult) SetCarrierUrl(value *string)() {
    err := m.GetBackingStore().Set("carrierUrl", value)
    if err != nil {
        panic(err)
    }
}
// ActivateDeviceEsimActionResultable 
type ActivateDeviceEsimActionResultable interface {
    DeviceActionResultable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCarrierUrl()(*string)
    SetCarrierUrl(value *string)()
}
