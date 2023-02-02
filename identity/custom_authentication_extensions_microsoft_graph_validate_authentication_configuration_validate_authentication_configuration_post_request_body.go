package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody 
type CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The authenticationConfiguration property
    authenticationConfiguration ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionAuthenticationConfigurationable
    // The endpointConfiguration property
    endpointConfiguration ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionEndpointConfigurationable
}
// NewCustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody instantiates a new CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody and sets the default values.
func NewCustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody()(*CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) {
    m := &CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthenticationConfiguration gets the authenticationConfiguration property value. The authenticationConfiguration property
func (m *CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) GetAuthenticationConfiguration()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionAuthenticationConfigurationable) {
    return m.authenticationConfiguration
}
// GetEndpointConfiguration gets the endpointConfiguration property value. The endpointConfiguration property
func (m *CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) GetEndpointConfiguration()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionEndpointConfigurationable) {
    return m.endpointConfiguration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomExtensionAuthenticationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationConfiguration(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionAuthenticationConfigurationable))
        }
        return nil
    }
    res["endpointConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomExtensionEndpointConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointConfiguration(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionEndpointConfigurationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authenticationConfiguration", m.GetAuthenticationConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("endpointConfiguration", m.GetEndpointConfiguration())
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
func (m *CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthenticationConfiguration sets the authenticationConfiguration property value. The authenticationConfiguration property
func (m *CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) SetAuthenticationConfiguration(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionAuthenticationConfigurationable)() {
    m.authenticationConfiguration = value
}
// SetEndpointConfiguration sets the endpointConfiguration property value. The endpointConfiguration property
func (m *CustomAuthenticationExtensionsMicrosoftGraphValidateAuthenticationConfigurationValidateAuthenticationConfigurationPostRequestBody) SetEndpointConfiguration(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomExtensionEndpointConfigurationable)() {
    m.endpointConfiguration = value
}
