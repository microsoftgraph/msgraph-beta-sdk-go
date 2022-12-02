package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeGetLoggedOnManagedDevicesResponse provides operations to call the getLoggedOnManagedDevices method.
type MeGetLoggedOnManagedDevicesResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ManagedDeviceable
}
// NewMeGetLoggedOnManagedDevicesResponse instantiates a new MeGetLoggedOnManagedDevicesResponse and sets the default values.
func NewMeGetLoggedOnManagedDevicesResponse()(*MeGetLoggedOnManagedDevicesResponse) {
    m := &MeGetLoggedOnManagedDevicesResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMeGetLoggedOnManagedDevicesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeGetLoggedOnManagedDevicesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeGetLoggedOnManagedDevicesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeGetLoggedOnManagedDevicesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MeGetLoggedOnManagedDevicesResponse) GetValue()([]ManagedDeviceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeGetLoggedOnManagedDevicesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MeGetLoggedOnManagedDevicesResponse) SetValue(value []ManagedDeviceable)() {
    m.value = value
}
