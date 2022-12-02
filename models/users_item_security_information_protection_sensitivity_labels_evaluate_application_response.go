package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse provides operations to call the evaluateApplication method.
type UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse instantiates a new UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse and sets the default values.
func NewUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse()(*UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse) {
    m := &UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UsersItemSecurityInformationProtectionSensitivityLabelsEvaluateApplicationResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
