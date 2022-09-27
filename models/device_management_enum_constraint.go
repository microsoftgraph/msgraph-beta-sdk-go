package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementEnumConstraint 
type DeviceManagementEnumConstraint struct {
    DeviceManagementConstraint
    // List of valid values for this string
    values []DeviceManagementEnumValueable
}
// NewDeviceManagementEnumConstraint instantiates a new DeviceManagementEnumConstraint and sets the default values.
func NewDeviceManagementEnumConstraint()(*DeviceManagementEnumConstraint) {
    m := &DeviceManagementEnumConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementEnumConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementEnumConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementEnumConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementEnumConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementEnumConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["values"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementEnumValueFromDiscriminatorValue , m.SetValues)
    return res
}
// GetValues gets the values property value. List of valid values for this string
func (m *DeviceManagementEnumConstraint) GetValues()([]DeviceManagementEnumValueable) {
    return m.values
}
// Serialize serializes information the current object
func (m *DeviceManagementEnumConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValues() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValues())
        err = writer.WriteCollectionOfObjectValues("values", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValues sets the values property value. List of valid values for this string
func (m *DeviceManagementEnumConstraint) SetValues(value []DeviceManagementEnumValueable)() {
    m.values = value
}
