package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody provides operations to call the extractLabel method.
type UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentInfo property
    contentInfo ContentInfoable
}
// NewUsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody instantiates a new UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody and sets the default values.
func NewUsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody()(*UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody) {
    m := &UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody) GetContentInfo()(ContentInfoable) {
    return m.contentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(ContentInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
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
func (m *UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *UsersItemInformationProtectionPolicyLabelsExtractLabelPostRequestBody) SetContentInfo(value ContentInfoable)() {
    m.contentInfo = value
}
