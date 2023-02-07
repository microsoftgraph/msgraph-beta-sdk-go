package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody 
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The decision property
    decision *string
    // The justification property
    justification *string
    // The principalId property
    principalId *string
    // The resourceId property
    resourceId *string
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody instantiates a new AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody()(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDecision gets the decision property value. The decision property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) GetDecision()(*string) {
    return m.decision
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecision(val)
        }
        return nil
    }
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
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
// GetJustification gets the justification property value. The justification property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) GetJustification()(*string) {
    return m.justification
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) GetPrincipalId()(*string) {
    return m.principalId
}
// GetResourceId gets the resourceId property value. The resourceId property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) GetResourceId()(*string) {
    return m.resourceId
}
// Serialize serializes information the current object
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("decision", m.GetDecision())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
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
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDecision sets the decision property value. The decision property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) SetDecision(value *string)() {
    m.decision = value
}
// SetJustification sets the justification property value. The justification property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) SetJustification(value *string)() {
    m.justification = value
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetResourceId sets the resourceId property value. The resourceId property
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceMicrosoftGraphBatchRecordDecisionsBatchRecordDecisionsPostRequestBody) SetResourceId(value *string)() {
    m.resourceId = value
}
