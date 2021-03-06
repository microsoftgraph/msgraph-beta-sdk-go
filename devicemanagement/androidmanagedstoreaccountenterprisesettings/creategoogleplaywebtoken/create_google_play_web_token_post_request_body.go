package creategoogleplaywebtoken

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateGooglePlayWebTokenPostRequestBody provides operations to call the createGooglePlayWebToken method.
type CreateGooglePlayWebTokenPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The parentUri property
    parentUri *string
}
// NewCreateGooglePlayWebTokenPostRequestBody instantiates a new createGooglePlayWebTokenPostRequestBody and sets the default values.
func NewCreateGooglePlayWebTokenPostRequestBody()(*CreateGooglePlayWebTokenPostRequestBody) {
    m := &CreateGooglePlayWebTokenPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCreateGooglePlayWebTokenPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreateGooglePlayWebTokenPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateGooglePlayWebTokenPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateGooglePlayWebTokenPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateGooglePlayWebTokenPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parentUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentUri(val)
        }
        return nil
    }
    return res
}
// GetParentUri gets the parentUri property value. The parentUri property
func (m *CreateGooglePlayWebTokenPostRequestBody) GetParentUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentUri
    }
}
// Serialize serializes information the current object
func (m *CreateGooglePlayWebTokenPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("parentUri", m.GetParentUri())
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
func (m *CreateGooglePlayWebTokenPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetParentUri sets the parentUri property value. The parentUri property
func (m *CreateGooglePlayWebTokenPostRequestBody) SetParentUri(value *string)() {
    if m != nil {
        m.parentUri = value
    }
}
