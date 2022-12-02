package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody provides operations to call the addMembers method.
type AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody instantiates a new AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody and sets the default values.
func NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody()(*AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody) {
    m := &AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
