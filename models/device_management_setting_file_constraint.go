package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingFileConstraint constraint enforcing the file extension is acceptable for a given setting
type DeviceManagementSettingFileConstraint struct {
    DeviceManagementConstraint
}
// NewDeviceManagementSettingFileConstraint instantiates a new DeviceManagementSettingFileConstraint and sets the default values.
func NewDeviceManagementSettingFileConstraint()(*DeviceManagementSettingFileConstraint) {
    m := &DeviceManagementSettingFileConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingFileConstraint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementSettingFileConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementSettingFileConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingFileConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementSettingFileConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["supportedExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSupportedExtensions(res)
        }
        return nil
    }
    return res
}
// GetSupportedExtensions gets the supportedExtensions property value. Acceptable file extensions to upload for this setting
// returns a []string when successful
func (m *DeviceManagementSettingFileConstraint) GetSupportedExtensions()([]string) {
    val, err := m.GetBackingStore().Get("supportedExtensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingFileConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSupportedExtensions() != nil {
        err = writer.WriteCollectionOfStringValues("supportedExtensions", m.GetSupportedExtensions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSupportedExtensions sets the supportedExtensions property value. Acceptable file extensions to upload for this setting
func (m *DeviceManagementSettingFileConstraint) SetSupportedExtensions(value []string)() {
    err := m.GetBackingStore().Set("supportedExtensions", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementSettingFileConstraintable interface {
    DeviceManagementConstraintable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSupportedExtensions()([]string)
    SetSupportedExtensions(value []string)()
}
