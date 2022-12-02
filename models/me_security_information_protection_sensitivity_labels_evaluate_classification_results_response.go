package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse provides operations to call the evaluateClassificationResults method.
type MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse instantiates a new MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse and sets the default values.
func NewMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse()(*MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) {
    m := &MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MeSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
