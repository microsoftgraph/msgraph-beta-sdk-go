package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody 
type MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The azureAdDeviceId property
    azureAdDeviceId *string
}
// NewMicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody instantiates a new MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody and sets the default values.
func NewMicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody()(*MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) {
    m := &MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. The azureAdDeviceId property
func (m *MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) GetAzureAdDeviceId()(*string) {
    return m.azureAdDeviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. The azureAdDeviceId property
func (m *MicrosoftGraphWipeManagedAppRegistrationsByAzureAdDeviceIdWipeManagedAppRegistrationsByAzureAdDeviceIdPostRequestBody) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
