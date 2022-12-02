package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody provides operations to call the recordDecisions method.
type ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The justification property
    justification *string
    // The reviewResult property
    reviewResult *string
}
// NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody instantiates a new ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody and sets the default values.
func NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody()(*ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) {
    m := &ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) GetJustification()(*string) {
    return m.justification
}
// GetReviewResult gets the reviewResult property value. The reviewResult property
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) GetReviewResult()(*string) {
    return m.reviewResult
}
// Serialize serializes information the current object
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetJustification sets the justification property value. The justification property
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) SetJustification(value *string)() {
    m.justification = value
}
// SetReviewResult sets the reviewResult property value. The reviewResult property
func (m *ApprovalWorkflowProvidersItemBusinessFlowsWithRequestsAwaitingMyDecisionItemRecordDecisionsPostRequestBody) SetReviewResult(value *string)() {
    m.reviewResult = value
}
