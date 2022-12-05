package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeCloudPCsItemChangeUserAccountTypePostRequestBody provides operations to call the changeUserAccountType method.
type MeCloudPCsItemChangeUserAccountTypePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The userAccountType property
    userAccountType *CloudPcUserAccountType
}
// NewMeCloudPCsItemChangeUserAccountTypePostRequestBody instantiates a new MeCloudPCsItemChangeUserAccountTypePostRequestBody and sets the default values.
func NewMeCloudPCsItemChangeUserAccountTypePostRequestBody()(*MeCloudPCsItemChangeUserAccountTypePostRequestBody) {
    m := &MeCloudPCsItemChangeUserAccountTypePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeCloudPCsItemChangeUserAccountTypePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeCloudPCsItemChangeUserAccountTypePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeCloudPCsItemChangeUserAccountTypePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeCloudPCsItemChangeUserAccountTypePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeCloudPCsItemChangeUserAccountTypePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["userAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcUserAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*CloudPcUserAccountType))
        }
        return nil
    }
    return res
}
// GetUserAccountType gets the userAccountType property value. The userAccountType property
func (m *MeCloudPCsItemChangeUserAccountTypePostRequestBody) GetUserAccountType()(*CloudPcUserAccountType) {
    return m.userAccountType
}
// Serialize serializes information the current object
func (m *MeCloudPCsItemChangeUserAccountTypePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUserAccountType() != nil {
        cast := (*m.GetUserAccountType()).String()
        err := writer.WriteStringValue("userAccountType", &cast)
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
func (m *MeCloudPCsItemChangeUserAccountTypePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetUserAccountType sets the userAccountType property value. The userAccountType property
func (m *MeCloudPCsItemChangeUserAccountTypePostRequestBody) SetUserAccountType(value *CloudPcUserAccountType)() {
    m.userAccountType = value
}
