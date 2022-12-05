package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse provides operations to call the evaluateClassificationResults method.
type SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse instantiates a new SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse and sets the default values.
func NewSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse()(*SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) {
    m := &SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
