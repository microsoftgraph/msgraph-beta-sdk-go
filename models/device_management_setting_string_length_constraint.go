package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingStringLengthConstraint constraint enforcing a given string length range
type DeviceManagementSettingStringLengthConstraint struct {
    DeviceManagementConstraint
    // The OdataType property
    OdataType *string
}
// NewDeviceManagementSettingStringLengthConstraint instantiates a new deviceManagementSettingStringLengthConstraint and sets the default values.
func NewDeviceManagementSettingStringLengthConstraint()(*DeviceManagementSettingStringLengthConstraint) {
    m := &DeviceManagementSettingStringLengthConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingStringLengthConstraint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementSettingStringLengthConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingStringLengthConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingStringLengthConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingStringLengthConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetMaximumLength gets the maximumLength property value. The maximum permitted string length
func (m *DeviceManagementSettingStringLengthConstraint) GetMaximumLength()(*int32) {
    val, err := m.GetBackingStore().Get("maximumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumLength gets the minimumLength property value. The minimum permitted string length
func (m *DeviceManagementSettingStringLengthConstraint) GetMinimumLength()(*int32) {
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
func (m *DeviceManagementSettingStringLengthConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetMaximumLength sets the maximumLength property value. The maximum permitted string length
func (m *DeviceManagementSettingStringLengthConstraint) SetMaximumLength(value *int32)() {
    err := m.GetBackingStore().Set("maximumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumLength sets the minimumLength property value. The minimum permitted string length
func (m *DeviceManagementSettingStringLengthConstraint) SetMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("minimumLength", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementSettingStringLengthConstraintable 
type DeviceManagementSettingStringLengthConstraintable interface {
    DeviceManagementConstraintable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumLength()(*int32)
    GetMinimumLength()(*int32)
    SetMaximumLength(value *int32)()
    SetMinimumLength(value *int32)()
}
