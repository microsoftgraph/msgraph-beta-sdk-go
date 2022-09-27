package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnmanagedDeviceDiscoveryTask 
type UnmanagedDeviceDiscoveryTask struct {
    DeviceAppManagementTask
    // Unmanaged devices discovered in the network.
    unmanagedDevices []UnmanagedDeviceable
}
// NewUnmanagedDeviceDiscoveryTask instantiates a new UnmanagedDeviceDiscoveryTask and sets the default values.
func NewUnmanagedDeviceDiscoveryTask()(*UnmanagedDeviceDiscoveryTask) {
    m := &UnmanagedDeviceDiscoveryTask{
        DeviceAppManagementTask: *NewDeviceAppManagementTask(),
    }
    odataTypeValue := "#microsoft.graph.unmanagedDeviceDiscoveryTask";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUnmanagedDeviceDiscoveryTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnmanagedDeviceDiscoveryTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnmanagedDeviceDiscoveryTask(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnmanagedDeviceDiscoveryTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAppManagementTask.GetFieldDeserializers()
    res["unmanagedDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnmanagedDeviceFromDiscriminatorValue , m.SetUnmanagedDevices)
    return res
}
// GetUnmanagedDevices gets the unmanagedDevices property value. Unmanaged devices discovered in the network.
func (m *UnmanagedDeviceDiscoveryTask) GetUnmanagedDevices()([]UnmanagedDeviceable) {
    return m.unmanagedDevices
}
// Serialize serializes information the current object
func (m *UnmanagedDeviceDiscoveryTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAppManagementTask.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUnmanagedDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUnmanagedDevices())
        err = writer.WriteCollectionOfObjectValues("unmanagedDevices", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUnmanagedDevices sets the unmanagedDevices property value. Unmanaged devices discovered in the network.
func (m *UnmanagedDeviceDiscoveryTask) SetUnmanagedDevices(value []UnmanagedDeviceable)() {
    m.unmanagedDevices = value
}
