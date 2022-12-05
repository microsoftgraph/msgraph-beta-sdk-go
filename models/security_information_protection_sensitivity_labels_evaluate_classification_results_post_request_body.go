package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody provides operations to call the evaluateClassificationResults method.
type SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody instantiates a new SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody and sets the default values.
func NewSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody()(*SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody) {
    m := &SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
