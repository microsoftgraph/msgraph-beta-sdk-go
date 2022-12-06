package updateaddomainpassword

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateAdDomainPasswordPostRequestBody provides operations to call the updateAdDomainPassword method.
type UpdateAdDomainPasswordPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The adDomainPassword property
    adDomainPassword *string
}
// NewUpdateAdDomainPasswordPostRequestBody instantiates a new updateAdDomainPasswordPostRequestBody and sets the default values.
func NewUpdateAdDomainPasswordPostRequestBody()(*UpdateAdDomainPasswordPostRequestBody) {
    m := &UpdateAdDomainPasswordPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateAdDomainPasswordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateAdDomainPasswordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateAdDomainPasswordPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAdDomainPasswordPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdDomainPassword gets the adDomainPassword property value. The adDomainPassword property
func (m *UpdateAdDomainPasswordPostRequestBody) GetAdDomainPassword()(*string) {
    return m.adDomainPassword
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateAdDomainPasswordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adDomainPassword"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAdDomainPassword)
    return res
}
// Serialize serializes information the current object
func (m *UpdateAdDomainPasswordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("adDomainPassword", m.GetAdDomainPassword())
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
func (m *UpdateAdDomainPasswordPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdDomainPassword sets the adDomainPassword property value. The adDomainPassword property
func (m *UpdateAdDomainPasswordPostRequestBody) SetAdDomainPassword(value *string)() {
    m.adDomainPassword = value
}
