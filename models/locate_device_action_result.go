package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocateDeviceActionResult 
type LocateDeviceActionResult struct {
    DeviceActionResult
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // device location
    deviceLocation DeviceGeoLocationable
}
// NewLocateDeviceActionResult instantiates a new LocateDeviceActionResult and sets the default values.
func NewLocateDeviceActionResult()(*LocateDeviceActionResult) {
    m := &LocateDeviceActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLocateDeviceActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocateDeviceActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocateDeviceActionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocateDeviceActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceLocation gets the deviceLocation property value. device location
func (m *LocateDeviceActionResult) GetDeviceLocation()(DeviceGeoLocationable) {
    if m == nil {
        return nil
    } else {
        return m.deviceLocation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocateDeviceActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["deviceLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceGeoLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLocation(val.(DeviceGeoLocationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *LocateDeviceActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceLocation", m.GetDeviceLocation())
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
func (m *LocateDeviceActionResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceLocation sets the deviceLocation property value. device location
func (m *LocateDeviceActionResult) SetDeviceLocation(value DeviceGeoLocationable)() {
    if m != nil {
        m.deviceLocation = value
    }
}
