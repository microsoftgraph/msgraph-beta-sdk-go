package adduserinputlog

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddUserInputLogPostRequestBody provides operations to call the addUserInputLog method.
type AddUserInputLogPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The logInformation property
    logInformation *string
}
// NewAddUserInputLogPostRequestBody instantiates a new addUserInputLogPostRequestBody and sets the default values.
func NewAddUserInputLogPostRequestBody()(*AddUserInputLogPostRequestBody) {
    m := &AddUserInputLogPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAddUserInputLogPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddUserInputLogPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddUserInputLogPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AddUserInputLogPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddUserInputLogPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["logInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLogInformation)
    return res
}
// GetLogInformation gets the logInformation property value. The logInformation property
func (m *AddUserInputLogPostRequestBody) GetLogInformation()(*string) {
    return m.logInformation
}
// Serialize serializes information the current object
func (m *AddUserInputLogPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("logInformation", m.GetLogInformation())
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
func (m *AddUserInputLogPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLogInformation sets the logInformation property value. The logInformation property
func (m *AddUserInputLogPostRequestBody) SetLogInformation(value *string)() {
    m.logInformation = value
}
