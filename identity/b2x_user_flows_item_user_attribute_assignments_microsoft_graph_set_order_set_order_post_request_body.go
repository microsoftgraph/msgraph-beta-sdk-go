package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody 
type B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The newAssignmentOrder property
    newAssignmentOrder ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable
}
// NewB2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody instantiates a new B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody and sets the default values.
func NewB2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody()(*B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody) {
    m := &B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateB2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newAssignmentOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentOrderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewAssignmentOrder(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable))
        }
        return nil
    }
    return res
}
// GetNewAssignmentOrder gets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody) GetNewAssignmentOrder()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable) {
    return m.newAssignmentOrder
}
// Serialize serializes information the current object
func (m *B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newAssignmentOrder", m.GetNewAssignmentOrder())
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
func (m *B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewAssignmentOrder sets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *B2xUserFlowsItemUserAttributeAssignmentsMicrosoftGraphSetOrderSetOrderPostRequestBody) SetNewAssignmentOrder(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable)() {
    m.newAssignmentOrder = value
}
