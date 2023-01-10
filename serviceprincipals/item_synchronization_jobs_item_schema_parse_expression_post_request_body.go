package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody 
type ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The expression property
    expression *string
    // The targetAttributeDefinition property
    targetAttributeDefinition ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable
    // The testInputObject property
    testInputObject ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable
}
// NewItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody instantiates a new ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody and sets the default values.
func NewItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody()(*ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) {
    m := &ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSynchronizationJobsItemSchemaParseExpressionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationJobsItemSchemaParseExpressionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpression gets the expression property value. The expression property
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) GetExpression()(*string) {
    return m.expression
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expression"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpression(val)
        }
        return nil
    }
    res["targetAttributeDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetAttributeDefinition(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable))
        }
        return nil
    }
    res["testInputObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExpressionInputObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestInputObject(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable))
        }
        return nil
    }
    return res
}
// GetTargetAttributeDefinition gets the targetAttributeDefinition property value. The targetAttributeDefinition property
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) GetTargetAttributeDefinition()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable) {
    return m.targetAttributeDefinition
}
// GetTestInputObject gets the testInputObject property value. The testInputObject property
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) GetTestInputObject()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable) {
    return m.testInputObject
}
// Serialize serializes information the current object
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("targetAttributeDefinition", m.GetTargetAttributeDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("testInputObject", m.GetTestInputObject())
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
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpression sets the expression property value. The expression property
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) SetExpression(value *string)() {
    m.expression = value
}
// SetTargetAttributeDefinition sets the targetAttributeDefinition property value. The targetAttributeDefinition property
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) SetTargetAttributeDefinition(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttributeDefinitionable)() {
    m.targetAttributeDefinition = value
}
// SetTestInputObject sets the testInputObject property value. The testInputObject property
func (m *ItemSynchronizationJobsItemSchemaParseExpressionPostRequestBody) SetTestInputObject(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExpressionInputObjectable)() {
    m.testInputObject = value
}
