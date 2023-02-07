package approvalworkflowproviders

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody 
type ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The justification property
    justification *string
    // The reviewResult property
    reviewResult *string
}
// NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody instantiates a new ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody and sets the default values.
func NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody()(*ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) {
    m := &ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) GetJustification()(*string) {
    return m.justification
}
// GetReviewResult gets the reviewResult property value. The reviewResult property
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) GetReviewResult()(*string) {
    return m.reviewResult
}
// Serialize serializes information the current object
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetJustification sets the justification property value. The justification property
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) SetJustification(value *string)() {
    m.justification = value
}
// SetReviewResult sets the reviewResult property value. The reviewResult property
func (m *ItemBusinessFlowsWithRequestsAwaitingMyDecisionItemMicrosoftGraphRecordDecisionsRecordDecisionsPostRequestBody) SetReviewResult(value *string)() {
    m.reviewResult = value
}
