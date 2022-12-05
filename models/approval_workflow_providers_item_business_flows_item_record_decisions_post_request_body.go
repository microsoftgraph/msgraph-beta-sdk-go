package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody provides operations to call the recordDecisions method.
type ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The justification property
    justification *string
    // The reviewResult property
    reviewResult *string
}
// NewApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody instantiates a new ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody and sets the default values.
func NewApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody()(*ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) {
    m := &ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["reviewResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewResult(val)
        }
        return nil
    }
    return res
}
// GetJustification gets the justification property value. The justification property
func (m *ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) GetJustification()(*string) {
    return m.justification
}
// GetReviewResult gets the reviewResult property value. The reviewResult property
func (m *ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) GetReviewResult()(*string) {
    return m.reviewResult
}
// Serialize serializes information the current object
func (m *ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reviewResult", m.GetReviewResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetJustification sets the justification property value. The justification property
func (m *ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) SetJustification(value *string)() {
    m.justification = value
}
// SetReviewResult sets the reviewResult property value. The reviewResult property
func (m *ApprovalWorkflowProvidersItemBusinessFlowsItemRecordDecisionsPostRequestBody) SetReviewResult(value *string)() {
    m.reviewResult = value
}
