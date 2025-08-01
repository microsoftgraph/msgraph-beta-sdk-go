// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AssignmentRequestApprovalStageCallbackData struct {
    AccessPackageAssignmentRequestCallbackData
}
// NewAssignmentRequestApprovalStageCallbackData instantiates a new AssignmentRequestApprovalStageCallbackData and sets the default values.
func NewAssignmentRequestApprovalStageCallbackData()(*AssignmentRequestApprovalStageCallbackData) {
    m := &AssignmentRequestApprovalStageCallbackData{
        AccessPackageAssignmentRequestCallbackData: *NewAccessPackageAssignmentRequestCallbackData(),
    }
    odataTypeValue := "#microsoft.graph.assignmentRequestApprovalStageCallbackData"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAssignmentRequestApprovalStageCallbackDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAssignmentRequestApprovalStageCallbackDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentRequestApprovalStageCallbackData(), nil
}
// GetApprovalStage gets the approvalStage property value. The stage in the approval decision.
// returns a AccessPackageApprovalStageable when successful
func (m *AssignmentRequestApprovalStageCallbackData) GetApprovalStage()(AccessPackageApprovalStageable) {
    val, err := m.GetBackingStore().Get("approvalStage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageApprovalStageable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AssignmentRequestApprovalStageCallbackData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageAssignmentRequestCallbackData.GetFieldDeserializers()
    res["approvalStage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageApprovalStageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalStage(val.(AccessPackageApprovalStageable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AssignmentRequestApprovalStageCallbackData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessPackageAssignmentRequestCallbackData.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("approvalStage", m.GetApprovalStage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApprovalStage sets the approvalStage property value. The stage in the approval decision.
func (m *AssignmentRequestApprovalStageCallbackData) SetApprovalStage(value AccessPackageApprovalStageable)() {
    err := m.GetBackingStore().Set("approvalStage", value)
    if err != nil {
        panic(err)
    }
}
type AssignmentRequestApprovalStageCallbackDataable interface {
    AccessPackageAssignmentRequestCallbackDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalStage()(AccessPackageApprovalStageable)
    SetApprovalStage(value AccessPackageApprovalStageable)()
}
