package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse provides operations to call the evaluateRemoval method.
type SecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse instantiates a new SecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse and sets the default values.
func NewSecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse()(*SecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse) {
    m := &SecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecurityInformationProtectionSensitivityLabelsEvaluateRemovalResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
