package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExpressionEvaluationDetails 
type ExpressionEvaluationDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents expression which has been evaluated.
    expression *string
    // Represents the details of the evaluation of the expression.
    expressionEvaluationDetails []ExpressionEvaluationDetailsable
    // Represents the value of the result of the current expression.
    expressionResult *bool
    // The OdataType property
    odataType *string
    // Defines the name of the property and the value of that property.
    propertyToEvaluate PropertyToEvaluateable
}
// NewExpressionEvaluationDetails instantiates a new expressionEvaluationDetails and sets the default values.
func NewExpressionEvaluationDetails()(*ExpressionEvaluationDetails) {
    m := &ExpressionEvaluationDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExpressionEvaluationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExpressionEvaluationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExpressionEvaluationDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpressionEvaluationDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpression gets the expression property value. Represents expression which has been evaluated.
func (m *ExpressionEvaluationDetails) GetExpression()(*string) {
    return m.expression
}
// GetExpressionEvaluationDetails gets the expressionEvaluationDetails property value. Represents the details of the evaluation of the expression.
func (m *ExpressionEvaluationDetails) GetExpressionEvaluationDetails()([]ExpressionEvaluationDetailsable) {
    return m.expressionEvaluationDetails
}
// GetExpressionResult gets the expressionResult property value. Represents the value of the result of the current expression.
func (m *ExpressionEvaluationDetails) GetExpressionResult()(*bool) {
    return m.expressionResult
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpressionEvaluationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expression"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExpression)
    res["expressionEvaluationDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExpressionEvaluationDetailsFromDiscriminatorValue , m.SetExpressionEvaluationDetails)
    res["expressionResult"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetExpressionResult)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["propertyToEvaluate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePropertyToEvaluateFromDiscriminatorValue , m.SetPropertyToEvaluate)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExpressionEvaluationDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetPropertyToEvaluate gets the propertyToEvaluate property value. Defines the name of the property and the value of that property.
func (m *ExpressionEvaluationDetails) GetPropertyToEvaluate()(PropertyToEvaluateable) {
    return m.propertyToEvaluate
}
// Serialize serializes information the current object
func (m *ExpressionEvaluationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    if m.GetExpressionEvaluationDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExpressionEvaluationDetails())
        err := writer.WriteCollectionOfObjectValues("expressionEvaluationDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("expressionResult", m.GetExpressionResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("propertyToEvaluate", m.GetPropertyToEvaluate())
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
func (m *ExpressionEvaluationDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpression sets the expression property value. Represents expression which has been evaluated.
func (m *ExpressionEvaluationDetails) SetExpression(value *string)() {
    m.expression = value
}
// SetExpressionEvaluationDetails sets the expressionEvaluationDetails property value. Represents the details of the evaluation of the expression.
func (m *ExpressionEvaluationDetails) SetExpressionEvaluationDetails(value []ExpressionEvaluationDetailsable)() {
    m.expressionEvaluationDetails = value
}
// SetExpressionResult sets the expressionResult property value. Represents the value of the result of the current expression.
func (m *ExpressionEvaluationDetails) SetExpressionResult(value *bool)() {
    m.expressionResult = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExpressionEvaluationDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPropertyToEvaluate sets the propertyToEvaluate property value. Defines the name of the property and the value of that property.
func (m *ExpressionEvaluationDetails) SetPropertyToEvaluate(value PropertyToEvaluateable)() {
    m.propertyToEvaluate = value
}
