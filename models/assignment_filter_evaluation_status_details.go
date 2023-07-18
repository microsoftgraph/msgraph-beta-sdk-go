package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterEvaluationStatusDetails a class containing information about the payloads on which filter has been applied.
type AssignmentFilterEvaluationStatusDetails struct {
    Entity
}
// NewAssignmentFilterEvaluationStatusDetails instantiates a new assignmentFilterEvaluationStatusDetails and sets the default values.
func NewAssignmentFilterEvaluationStatusDetails()(*AssignmentFilterEvaluationStatusDetails) {
    m := &AssignmentFilterEvaluationStatusDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterEvaluationStatusDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterEvaluationStatusDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignmentFilterEvaluationStatusDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPayloadId gets the payloadId property value. PayloadId on which filter has been applied.
func (m *AssignmentFilterEvaluationStatusDetails) GetPayloadId()(*string) {
    val, err := m.GetBackingStore().Get("payloadId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AssignmentFilterEvaluationStatusDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterEvaluationStatusDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadId sets the payloadId property value. PayloadId on which filter has been applied.
func (m *AssignmentFilterEvaluationStatusDetails) SetPayloadId(value *string)() {
    err := m.GetBackingStore().Set("payloadId", value)
    if err != nil {
        panic(err)
    }
}
// AssignmentFilterEvaluationStatusDetailsable 
type AssignmentFilterEvaluationStatusDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPayloadId()(*string)
    SetOdataType(value *string)()
    SetPayloadId(value *string)()
}
