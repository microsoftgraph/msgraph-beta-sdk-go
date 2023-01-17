package approvalworkflowproviders

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemBusinessFlowsItemRecordDecisionsPostRequestBody 
type ItemBusinessFlowsItemRecordDecisionsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The justification property
    justification *string
    // The reviewResult property
    reviewResult *string
}
// NewItemBusinessFlowsItemRecordDecisionsPostRequestBody instantiates a new ItemBusinessFlowsItemRecordDecisionsPostRequestBody and sets the default values.
func NewItemBusinessFlowsItemRecordDecisionsPostRequestBody()(*ItemBusinessFlowsItemRecordDecisionsPostRequestBody) {
    m := &ItemBusinessFlowsItemRecordDecisionsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemBusinessFlowsItemRecordDecisionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemBusinessFlowsItemRecordDecisionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemBusinessFlowsItemRecordDecisionsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemBusinessFlowsItemRecordDecisionsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemBusinessFlowsItemRecordDecisionsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemBusinessFlowsItemRecordDecisionsPostRequestBody) GetJustification()(*string) {
    return m.justification
}
// GetReviewResult gets the reviewResult property value. The reviewResult property
func (m *ItemBusinessFlowsItemRecordDecisionsPostRequestBody) GetReviewResult()(*string) {
    return m.reviewResult
}
// Serialize serializes information the current object
func (m *ItemBusinessFlowsItemRecordDecisionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemBusinessFlowsItemRecordDecisionsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetJustification sets the justification property value. The justification property
func (m *ItemBusinessFlowsItemRecordDecisionsPostRequestBody) SetJustification(value *string)() {
    m.justification = value
}
// SetReviewResult sets the reviewResult property value. The reviewResult property
func (m *ItemBusinessFlowsItemRecordDecisionsPostRequestBody) SetReviewResult(value *string)() {
    m.reviewResult = value
}
