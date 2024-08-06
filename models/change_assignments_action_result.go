package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangeAssignmentsActionResult changeAssignmentsActionResult represents the result of executing the changeAssignments action on tracking the live reporting data for applications or configuration regarding their removal or restoration process
type ChangeAssignmentsActionResult struct {
    DeviceActionResult
}
// NewChangeAssignmentsActionResult instantiates a new ChangeAssignmentsActionResult and sets the default values.
func NewChangeAssignmentsActionResult()(*ChangeAssignmentsActionResult) {
    m := &ChangeAssignmentsActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    return m
}
// CreateChangeAssignmentsActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateChangeAssignmentsActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChangeAssignmentsActionResult(), nil
}
// GetDeviceAssignmentItems gets the deviceAssignmentItems property value. Indicates the list of applications or configuration to report live results during their changeAssignments action execution process. The result for each individual application or configuration can contain whether it's being removed or restored, what's the current status with potential message or error code, and when any changes happen on it. Read-Only. This collection can contain a maximum of 30 elements.
// returns a []DeviceAssignmentItemable when successful
func (m *ChangeAssignmentsActionResult) GetDeviceAssignmentItems()([]DeviceAssignmentItemable) {
    val, err := m.GetBackingStore().Get("deviceAssignmentItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceAssignmentItemable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ChangeAssignmentsActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["deviceAssignmentItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAssignmentItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAssignmentItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceAssignmentItemable)
                }
            }
            m.SetDeviceAssignmentItems(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ChangeAssignmentsActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeviceAssignmentItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceAssignmentItems()))
        for i, v := range m.GetDeviceAssignmentItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceAssignmentItems", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceAssignmentItems sets the deviceAssignmentItems property value. Indicates the list of applications or configuration to report live results during their changeAssignments action execution process. The result for each individual application or configuration can contain whether it's being removed or restored, what's the current status with potential message or error code, and when any changes happen on it. Read-Only. This collection can contain a maximum of 30 elements.
func (m *ChangeAssignmentsActionResult) SetDeviceAssignmentItems(value []DeviceAssignmentItemable)() {
    err := m.GetBackingStore().Set("deviceAssignmentItems", value)
    if err != nil {
        panic(err)
    }
}
type ChangeAssignmentsActionResultable interface {
    DeviceActionResultable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceAssignmentItems()([]DeviceAssignmentItemable)
    SetDeviceAssignmentItems(value []DeviceAssignmentItemable)()
}
