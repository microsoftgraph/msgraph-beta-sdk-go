package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PayloadCompatibleAssignmentFilter a class containing the properties used for Payload Compatible Assignment Filter.
type PayloadCompatibleAssignmentFilter struct {
    DeviceAndAppManagementAssignmentFilter
}
// NewPayloadCompatibleAssignmentFilter instantiates a new payloadCompatibleAssignmentFilter and sets the default values.
func NewPayloadCompatibleAssignmentFilter()(*PayloadCompatibleAssignmentFilter) {
    m := &PayloadCompatibleAssignmentFilter{
        DeviceAndAppManagementAssignmentFilter: *NewDeviceAndAppManagementAssignmentFilter(),
    }
    odataTypeValue := "#microsoft.graph.payloadCompatibleAssignmentFilter"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePayloadCompatibleAssignmentFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadCompatibleAssignmentFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayloadCompatibleAssignmentFilter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PayloadCompatibleAssignmentFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAndAppManagementAssignmentFilter.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["payloadType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentFilterPayloadType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadType(val.(*AssignmentFilterPayloadType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PayloadCompatibleAssignmentFilter) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPayloadType gets the payloadType property value. Represents the payload type AssignmentFilter is being assigned to.
func (m *PayloadCompatibleAssignmentFilter) GetPayloadType()(*AssignmentFilterPayloadType) {
    val, err := m.GetBackingStore().Get("payloadType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AssignmentFilterPayloadType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PayloadCompatibleAssignmentFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAndAppManagementAssignmentFilter.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
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
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PayloadCompatibleAssignmentFilter) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadType sets the payloadType property value. Represents the payload type AssignmentFilter is being assigned to.
func (m *PayloadCompatibleAssignmentFilter) SetPayloadType(value *AssignmentFilterPayloadType)() {
    err := m.GetBackingStore().Set("payloadType", value)
    if err != nil {
        panic(err)
    }
}
// PayloadCompatibleAssignmentFilterable 
type PayloadCompatibleAssignmentFilterable interface {
    DeviceAndAppManagementAssignmentFilterable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPayloadType()(*AssignmentFilterPayloadType)
    SetOdataType(value *string)()
    SetPayloadType(value *AssignmentFilterPayloadType)()
}
