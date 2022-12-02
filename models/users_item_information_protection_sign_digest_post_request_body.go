package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemInformationProtectionSignDigestPostRequestBody provides operations to call the signDigest method.
type UsersItemInformationProtectionSignDigestPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The digest property
    digest []byte
}
// NewUsersItemInformationProtectionSignDigestPostRequestBody instantiates a new UsersItemInformationProtectionSignDigestPostRequestBody and sets the default values.
func NewUsersItemInformationProtectionSignDigestPostRequestBody()(*UsersItemInformationProtectionSignDigestPostRequestBody) {
    m := &UsersItemInformationProtectionSignDigestPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemInformationProtectionSignDigestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemInformationProtectionSignDigestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemInformationProtectionSignDigestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemInformationProtectionSignDigestPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDigest gets the digest property value. The digest property
func (m *UsersItemInformationProtectionSignDigestPostRequestBody) GetDigest()([]byte) {
    return m.digest
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemInformationProtectionSignDigestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["digest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDigest(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UsersItemInformationProtectionSignDigestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("digest", m.GetDigest())
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
func (m *UsersItemInformationProtectionSignDigestPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDigest sets the digest property value. The digest property
func (m *UsersItemInformationProtectionSignDigestPostRequestBody) SetDigest(value []byte)() {
    m.digest = value
}
