package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse provides operations to call the evaluateClassificationResults method.
type UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse instantiates a new UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse and sets the default values.
func NewUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse()(*UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) {
    m := &UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateClassificationResultsResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
