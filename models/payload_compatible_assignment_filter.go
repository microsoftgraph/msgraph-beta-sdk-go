package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PayloadCompatibleAssignmentFilter 
type PayloadCompatibleAssignmentFilter struct {
    DeviceAndAppManagementAssignmentFilter
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // PayloadType of the Assignment Filter. Possible values are: notSet, enrollmentRestrictions.
    payloadType *AssignmentFilterPayloadType
}
// NewPayloadCompatibleAssignmentFilter instantiates a new PayloadCompatibleAssignmentFilter and sets the default values.
func NewPayloadCompatibleAssignmentFilter()(*PayloadCompatibleAssignmentFilter) {
    m := &PayloadCompatibleAssignmentFilter{
        DeviceAndAppManagementAssignmentFilter: *NewDeviceAndAppManagementAssignmentFilter(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePayloadCompatibleAssignmentFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadCompatibleAssignmentFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayloadCompatibleAssignmentFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadCompatibleAssignmentFilter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PayloadCompatibleAssignmentFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAndAppManagementAssignmentFilter.GetFieldDeserializers()
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
// GetPayloadType gets the payloadType property value. PayloadType of the Assignment Filter. Possible values are: notSet, enrollmentRestrictions.
func (m *PayloadCompatibleAssignmentFilter) GetPayloadType()(*AssignmentFilterPayloadType) {
    if m == nil {
        return nil
    } else {
        return m.payloadType
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadCompatibleAssignmentFilter) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPayloadType sets the payloadType property value. PayloadType of the Assignment Filter. Possible values are: notSet, enrollmentRestrictions.
func (m *PayloadCompatibleAssignmentFilter) SetPayloadType(value *AssignmentFilterPayloadType)() {
    if m != nil {
        m.payloadType = value
    }
}
