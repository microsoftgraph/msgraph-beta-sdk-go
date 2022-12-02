package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody provides operations to call the wipeManagedAppRegistrationsByAzureAdDeviceId method.
type MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The azureAdDeviceId property
    azureAdDeviceId *string
}
// NewMeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody instantiates a new MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody and sets the default values.
func NewMeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody()(*MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) {
    m := &MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. The azureAdDeviceId property
func (m *MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) GetAzureAdDeviceId()(*string) {
    return m.azureAdDeviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
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
func (m *MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. The azureAdDeviceId property
func (m *MeWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
