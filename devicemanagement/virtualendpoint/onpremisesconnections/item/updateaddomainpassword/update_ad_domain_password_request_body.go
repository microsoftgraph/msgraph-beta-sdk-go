package updateaddomainpassword

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateAdDomainPasswordRequestBody provides operations to call the updateAdDomainPassword method.
type UpdateAdDomainPasswordRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The adDomainPassword property
    adDomainPassword *string
}
// NewUpdateAdDomainPasswordRequestBody instantiates a new updateAdDomainPasswordRequestBody and sets the default values.
func NewUpdateAdDomainPasswordRequestBody()(*UpdateAdDomainPasswordRequestBody) {
    m := &UpdateAdDomainPasswordRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdateAdDomainPasswordRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateAdDomainPasswordRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateAdDomainPasswordRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAdDomainPasswordRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdDomainPassword gets the adDomainPassword property value. The adDomainPassword property
func (m *UpdateAdDomainPasswordRequestBody) GetAdDomainPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainPassword
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateAdDomainPasswordRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adDomainPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainPassword(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UpdateAdDomainPasswordRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UpdateAdDomainPasswordRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdDomainPassword sets the adDomainPassword property value. The adDomainPassword property
func (m *UpdateAdDomainPasswordRequestBody) SetAdDomainPassword(value *string)() {
    if m != nil {
        m.adDomainPassword = value
    }
}
