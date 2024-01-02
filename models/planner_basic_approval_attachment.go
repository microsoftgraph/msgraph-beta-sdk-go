package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerBasicApprovalAttachment 
type PlannerBasicApprovalAttachment struct {
    PlannerBaseApprovalAttachment
}
// NewPlannerBasicApprovalAttachment instantiates a new plannerBasicApprovalAttachment and sets the default values.
func NewPlannerBasicApprovalAttachment()(*PlannerBasicApprovalAttachment) {
    m := &PlannerBasicApprovalAttachment{
        PlannerBaseApprovalAttachment: *NewPlannerBaseApprovalAttachment(),
    }
    odataTypeValue := "#microsoft.graph.plannerBasicApprovalAttachment"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePlannerBasicApprovalAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerBasicApprovalAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerBasicApprovalAttachment(), nil
}
// GetApprovalId gets the approvalId property value. The approvalId property
func (m *PlannerBasicApprovalAttachment) GetApprovalId()(*string) {
    val, err := m.GetBackingStore().Get("approvalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerBasicApprovalAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerBaseApprovalAttachment.GetFieldDeserializers()
    res["approvalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PlannerBasicApprovalAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerBaseApprovalAttachment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("approvalId", m.GetApprovalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApprovalId sets the approvalId property value. The approvalId property
func (m *PlannerBasicApprovalAttachment) SetApprovalId(value *string)() {
    err := m.GetBackingStore().Set("approvalId", value)
    if err != nil {
        panic(err)
    }
}
// PlannerBasicApprovalAttachmentable 
type PlannerBasicApprovalAttachmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerBaseApprovalAttachmentable
    GetApprovalId()(*string)
    SetApprovalId(value *string)()
}
