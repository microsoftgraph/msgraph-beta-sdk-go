package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse provides operations to call the evaluateApplication method.
type MeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewMeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse instantiates a new MeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse and sets the default values.
func NewMeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse()(*MeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse) {
    m := &MeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MeSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
