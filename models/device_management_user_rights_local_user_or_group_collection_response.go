package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementUserRightsLocalUserOrGroupCollectionResponse 
type DeviceManagementUserRightsLocalUserOrGroupCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DeviceManagementUserRightsLocalUserOrGroupable
}
// NewDeviceManagementUserRightsLocalUserOrGroupCollectionResponse instantiates a new DeviceManagementUserRightsLocalUserOrGroupCollectionResponse and sets the default values.
func NewDeviceManagementUserRightsLocalUserOrGroupCollectionResponse()(*DeviceManagementUserRightsLocalUserOrGroupCollectionResponse) {
    m := &DeviceManagementUserRightsLocalUserOrGroupCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceManagementUserRightsLocalUserOrGroupCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementUserRightsLocalUserOrGroupCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementUserRightsLocalUserOrGroupCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementUserRightsLocalUserOrGroupCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementUserRightsLocalUserOrGroupFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceManagementUserRightsLocalUserOrGroupCollectionResponse) GetValue()([]DeviceManagementUserRightsLocalUserOrGroupable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementUserRightsLocalUserOrGroupCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DeviceManagementUserRightsLocalUserOrGroupCollectionResponse) SetValue(value []DeviceManagementUserRightsLocalUserOrGroupable)() {
    m.value = value
}
