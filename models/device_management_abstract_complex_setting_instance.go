package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAbstractComplexSettingInstance a setting instance representing a complex value for an abstract setting
type DeviceManagementAbstractComplexSettingInstance struct {
    DeviceManagementSettingInstance
}
// NewDeviceManagementAbstractComplexSettingInstance instantiates a new deviceManagementAbstractComplexSettingInstance and sets the default values.
func NewDeviceManagementAbstractComplexSettingInstance()(*DeviceManagementAbstractComplexSettingInstance) {
    m := &DeviceManagementAbstractComplexSettingInstance{
        DeviceManagementSettingInstance: *NewDeviceManagementSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementAbstractComplexSettingInstance"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementAbstractComplexSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementAbstractComplexSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAbstractComplexSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAbstractComplexSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingInstance.GetFieldDeserializers()
    res["implementationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImplementationId(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementSettingInstanceable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetImplementationId gets the implementationId property value. The definition ID for the chosen implementation of this complex setting
func (m *DeviceManagementAbstractComplexSettingInstance) GetImplementationId()(*string) {
    val, err := m.GetBackingStore().Get("implementationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValue gets the value property value. The values that make up the complex setting
func (m *DeviceManagementAbstractComplexSettingInstance) GetValue()([]DeviceManagementSettingInstanceable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementSettingInstanceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementAbstractComplexSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("implementationId", m.GetImplementationId())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetImplementationId sets the implementationId property value. The definition ID for the chosen implementation of this complex setting
func (m *DeviceManagementAbstractComplexSettingInstance) SetImplementationId(value *string)() {
    err := m.GetBackingStore().Set("implementationId", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. The values that make up the complex setting
func (m *DeviceManagementAbstractComplexSettingInstance) SetValue(value []DeviceManagementSettingInstanceable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementAbstractComplexSettingInstanceable 
type DeviceManagementAbstractComplexSettingInstanceable interface {
    DeviceManagementSettingInstanceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetImplementationId()(*string)
    GetValue()([]DeviceManagementSettingInstanceable)
    SetImplementationId(value *string)()
    SetValue(value []DeviceManagementSettingInstanceable)()
}
