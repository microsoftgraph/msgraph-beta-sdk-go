package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingCollectionConstraint constraint that enforces the maximum number of elements a collection
type DeviceManagementSettingCollectionConstraint struct {
    DeviceManagementConstraint
}
// NewDeviceManagementSettingCollectionConstraint instantiates a new DeviceManagementSettingCollectionConstraint and sets the default values.
func NewDeviceManagementSettingCollectionConstraint()(*DeviceManagementSettingCollectionConstraint) {
    m := &DeviceManagementSettingCollectionConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingCollectionConstraint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementSettingCollectionConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementSettingCollectionConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingCollectionConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementSettingCollectionConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["maximumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumLength(val)
        }
        return nil
    }
    res["minimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumLength(val)
        }
        return nil
    }
    return res
}
// GetMaximumLength gets the maximumLength property value. The maximum number of elements in the collection
// returns a *int32 when successful
func (m *DeviceManagementSettingCollectionConstraint) GetMaximumLength()(*int32) {
    val, err := m.GetBackingStore().Get("maximumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumLength gets the minimumLength property value. The minimum number of elements in the collection
// returns a *int32 when successful
func (m *DeviceManagementSettingCollectionConstraint) GetMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("minimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingCollectionConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("maximumLength", m.GetMaximumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumLength", m.GetMinimumLength())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaximumLength sets the maximumLength property value. The maximum number of elements in the collection
func (m *DeviceManagementSettingCollectionConstraint) SetMaximumLength(value *int32)() {
    err := m.GetBackingStore().Set("maximumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumLength sets the minimumLength property value. The minimum number of elements in the collection
func (m *DeviceManagementSettingCollectionConstraint) SetMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("minimumLength", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementSettingCollectionConstraintable interface {
    DeviceManagementConstraintable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumLength()(*int32)
    GetMinimumLength()(*int32)
    SetMaximumLength(value *int32)()
    SetMinimumLength(value *int32)()
}
