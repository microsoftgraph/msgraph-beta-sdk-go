package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnmanagedDeviceDiscoveryTask 
type UnmanagedDeviceDiscoveryTask struct {
    DeviceAppManagementTask
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Unmanaged devices discovered in the network.
    unmanagedDevices []UnmanagedDeviceable
}
// NewUnmanagedDeviceDiscoveryTask instantiates a new UnmanagedDeviceDiscoveryTask and sets the default values.
func NewUnmanagedDeviceDiscoveryTask()(*UnmanagedDeviceDiscoveryTask) {
    m := &UnmanagedDeviceDiscoveryTask{
        DeviceAppManagementTask: *NewDeviceAppManagementTask(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnmanagedDeviceDiscoveryTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnmanagedDeviceDiscoveryTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnmanagedDeviceDiscoveryTask(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnmanagedDeviceDiscoveryTask) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnmanagedDeviceDiscoveryTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAppManagementTask.GetFieldDeserializers()
    res["unmanagedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnmanagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnmanagedDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(UnmanagedDeviceable)
            }
            m.SetUnmanagedDevices(res)
        }
        return nil
    }
    return res
}
// GetUnmanagedDevices gets the unmanagedDevices property value. Unmanaged devices discovered in the network.
func (m *UnmanagedDeviceDiscoveryTask) GetUnmanagedDevices()([]UnmanagedDeviceable) {
    if m == nil {
        return nil
    } else {
        return m.unmanagedDevices
    }
}
// Serialize serializes information the current object
func (m *UnmanagedDeviceDiscoveryTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAppManagementTask.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUnmanagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnmanagedDevices()))
        for i, v := range m.GetUnmanagedDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("unmanagedDevices", cast)
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
func (m *UnmanagedDeviceDiscoveryTask) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUnmanagedDevices sets the unmanagedDevices property value. Unmanaged devices discovered in the network.
func (m *UnmanagedDeviceDiscoveryTask) SetUnmanagedDevices(value []UnmanagedDeviceable)() {
    if m != nil {
        m.unmanagedDevices = value
    }
}
