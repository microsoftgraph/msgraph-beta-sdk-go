package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParseExpressionResponse 
type ParseExpressionResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Error details, if expression evaluation resulted in an error.
    error PublicErrorable
    // A collection of values produced by the evaluation of the expression.
    evaluationResult []string
    // true if the evaluation was successful.
    evaluationSucceeded *bool
    // The OdataType property
    odataType *string
    // An attributeMappingSource object representing the parsed expression.
    parsedExpression AttributeMappingSourceable
    // true if the expression was parsed successfully.
    parsingSucceeded *bool
}
// NewParseExpressionResponse instantiates a new parseExpressionResponse and sets the default values.
func NewParseExpressionResponse()(*ParseExpressionResponse) {
    m := &ParseExpressionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateParseExpressionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParseExpressionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParseExpressionResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParseExpressionResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetError gets the error property value. Error details, if expression evaluation resulted in an error.
func (m *ParseExpressionResponse) GetError()(PublicErrorable) {
    return m.error
}
// GetEvaluationResult gets the evaluationResult property value. A collection of values produced by the evaluation of the expression.
func (m *ParseExpressionResponse) GetEvaluationResult()([]string) {
    return m.evaluationResult
}
// GetEvaluationSucceeded gets the evaluationSucceeded property value. true if the evaluation was successful.
func (m *ParseExpressionResponse) GetEvaluationSucceeded()(*bool) {
    return m.evaluationSucceeded
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParseExpressionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePublicErrorFromDiscriminatorValue , m.SetError)
    res["evaluationResult"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetEvaluationResult)
    res["evaluationSucceeded"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEvaluationSucceeded)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["parsedExpression"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAttributeMappingSourceFromDiscriminatorValue , m.SetParsedExpression)
    res["parsingSucceeded"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetParsingSucceeded)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ParseExpressionResponse) GetOdataType()(*string) {
    return m.odataType
}
// GetParsedExpression gets the parsedExpression property value. An attributeMappingSource object representing the parsed expression.
func (m *ParseExpressionResponse) GetParsedExpression()(AttributeMappingSourceable) {
    return m.parsedExpression
}
// GetParsingSucceeded gets the parsingSucceeded property value. true if the expression was parsed successfully.
func (m *ParseExpressionResponse) GetParsingSucceeded()(*bool) {
    return m.parsingSucceeded
}
// Serialize serializes information the current object
func (m *ParseExpressionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    if m.GetEvaluationResult() != nil {
        err := writer.WriteCollectionOfStringValues("evaluationResult", m.GetEvaluationResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("evaluationSucceeded", m.GetEvaluationSucceeded())
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
        err := writer.WriteObjectValue("parsedExpression", m.GetParsedExpression())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("parsingSucceeded", m.GetParsingSucceeded())
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
func (m *ParseExpressionResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetError sets the error property value. Error details, if expression evaluation resulted in an error.
func (m *ParseExpressionResponse) SetError(value PublicErrorable)() {
    m.error = value
}
// SetEvaluationResult sets the evaluationResult property value. A collection of values produced by the evaluation of the expression.
func (m *ParseExpressionResponse) SetEvaluationResult(value []string)() {
    m.evaluationResult = value
}
// SetEvaluationSucceeded sets the evaluationSucceeded property value. true if the evaluation was successful.
func (m *ParseExpressionResponse) SetEvaluationSucceeded(value *bool)() {
    m.evaluationSucceeded = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ParseExpressionResponse) SetOdataType(value *string)() {
    m.odataType = value
}
// SetParsedExpression sets the parsedExpression property value. An attributeMappingSource object representing the parsed expression.
func (m *ParseExpressionResponse) SetParsedExpression(value AttributeMappingSourceable)() {
    m.parsedExpression = value
}
// SetParsingSucceeded sets the parsingSucceeded property value. true if the expression was parsed successfully.
func (m *ParseExpressionResponse) SetParsingSucceeded(value *bool)() {
    m.parsingSucceeded = value
}
