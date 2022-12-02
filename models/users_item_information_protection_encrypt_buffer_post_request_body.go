package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemInformationProtectionEncryptBufferPostRequestBody provides operations to call the encryptBuffer method.
type UsersItemInformationProtectionEncryptBufferPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The buffer property
    buffer []byte
    // The labelId property
    labelId *string
}
// NewUsersItemInformationProtectionEncryptBufferPostRequestBody instantiates a new UsersItemInformationProtectionEncryptBufferPostRequestBody and sets the default values.
func NewUsersItemInformationProtectionEncryptBufferPostRequestBody()(*UsersItemInformationProtectionEncryptBufferPostRequestBody) {
    m := &UsersItemInformationProtectionEncryptBufferPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemInformationProtectionEncryptBufferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemInformationProtectionEncryptBufferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemInformationProtectionEncryptBufferPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemInformationProtectionEncryptBufferPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBuffer gets the buffer property value. The buffer property
func (m *UsersItemInformationProtectionEncryptBufferPostRequestBody) GetBuffer()([]byte) {
    return m.buffer
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemInformationProtectionEncryptBufferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuffer(val)
        }
        return nil
    }
    res["labelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelId(val)
        }
        return nil
    }
    return res
}
// GetLabelId gets the labelId property value. The labelId property
func (m *UsersItemInformationProtectionEncryptBufferPostRequestBody) GetLabelId()(*string) {
    return m.labelId
}
// Serialize serializes information the current object
func (m *UsersItemInformationProtectionEncryptBufferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("buffer", m.GetBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labelId", m.GetLabelId())
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
func (m *UsersItemInformationProtectionEncryptBufferPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBuffer sets the buffer property value. The buffer property
func (m *UsersItemInformationProtectionEncryptBufferPostRequestBody) SetBuffer(value []byte)() {
    m.buffer = value
}
// SetLabelId sets the labelId property value. The labelId property
func (m *UsersItemInformationProtectionEncryptBufferPostRequestBody) SetLabelId(value *string)() {
    m.labelId = value
}
