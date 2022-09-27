package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PayloadCompatibleAssignmentFilter 
type PayloadCompatibleAssignmentFilter struct {
    DeviceAndAppManagementAssignmentFilter
    // Represents the payload type AssignmentFilter is being assigned to.
    payloadType *AssignmentFilterPayloadType
}
// NewPayloadCompatibleAssignmentFilter instantiates a new PayloadCompatibleAssignmentFilter and sets the default values.
func NewPayloadCompatibleAssignmentFilter()(*PayloadCompatibleAssignmentFilter) {
    m := &PayloadCompatibleAssignmentFilter{
        DeviceAndAppManagementAssignmentFilter: *NewDeviceAndAppManagementAssignmentFilter(),
    }
    odataTypeValue := "#microsoft.graph.payloadCompatibleAssignmentFilter";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePayloadCompatibleAssignmentFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadCompatibleAssignmentFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayloadCompatibleAssignmentFilter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PayloadCompatibleAssignmentFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAndAppManagementAssignmentFilter.GetFieldDeserializers()
    res["payloadType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAssignmentFilterPayloadType , m.SetPayloadType)
    return res
}
// GetPayloadType gets the payloadType property value. Represents the payload type AssignmentFilter is being assigned to.
func (m *PayloadCompatibleAssignmentFilter) GetPayloadType()(*AssignmentFilterPayloadType) {
    return m.payloadType
}
// Serialize serializes information the current object
func (m *PayloadCompatibleAssignmentFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAndAppManagementAssignmentFilter.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPayloadType() != nil {
        cast := (*m.GetPayloadType()).String()
        err = writer.WriteStringValue("payloadType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPayloadType sets the payloadType property value. Represents the payload type AssignmentFilter is being assigned to.
func (m *PayloadCompatibleAssignmentFilter) SetPayloadType(value *AssignmentFilterPayloadType)() {
    m.payloadType = value
}
