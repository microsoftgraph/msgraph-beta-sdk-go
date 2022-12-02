package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody provides operations to call the parseExpression method.
type ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The expression property
    expression *string
    // The targetAttributeDefinition property
    targetAttributeDefinition AttributeDefinitionable
    // The testInputObject property
    testInputObject ExpressionInputObjectable
}
// NewApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody instantiates a new ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody and sets the default values.
func NewApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody()(*ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) {
    m := &ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpression gets the expression property value. The expression property
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) GetExpression()(*string) {
    return m.expression
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetAttributeDefinition(val.(AttributeDefinitionable))
        }
        return nil
    }
    res["testInputObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExpressionInputObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestInputObject(val.(ExpressionInputObjectable))
        }
        return nil
    }
    return res
}
// GetTargetAttributeDefinition gets the targetAttributeDefinition property value. The targetAttributeDefinition property
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) GetTargetAttributeDefinition()(AttributeDefinitionable) {
    return m.targetAttributeDefinition
}
// GetTestInputObject gets the testInputObject property value. The testInputObject property
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) GetTestInputObject()(ExpressionInputObjectable) {
    return m.testInputObject
}
// Serialize serializes information the current object
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpression sets the expression property value. The expression property
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) SetExpression(value *string)() {
    m.expression = value
}
// SetTargetAttributeDefinition sets the targetAttributeDefinition property value. The targetAttributeDefinition property
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) SetTargetAttributeDefinition(value AttributeDefinitionable)() {
    m.targetAttributeDefinition = value
}
// SetTestInputObject sets the testInputObject property value. The testInputObject property
func (m *ApplicationsItemSynchronizationTemplatesItemSchemaParseExpressionPostRequestBody) SetTestInputObject(value ExpressionInputObjectable)() {
    m.testInputObject = value
}
