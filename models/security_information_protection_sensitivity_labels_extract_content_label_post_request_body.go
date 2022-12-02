package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody provides operations to call the extractContentLabel method.
type SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewSecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody instantiates a new SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody and sets the default values.
func NewSecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody()(*SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody) {
    m := &SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityInformationProtectionSensitivityLabelsExtractContentLabelPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
