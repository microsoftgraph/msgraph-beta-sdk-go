package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody provides operations to call the validateAuthenticationConfiguration method.
type IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The authenticationConfiguration property
    authenticationConfiguration CustomExtensionAuthenticationConfigurationable
    // The endpointConfiguration property
    endpointConfiguration CustomExtensionEndpointConfigurationable
}
// NewIdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody instantiates a new IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody and sets the default values.
func NewIdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody()(*IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) {
    m := &IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAuthenticationConfiguration gets the authenticationConfiguration property value. The authenticationConfiguration property
func (m *IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) GetAuthenticationConfiguration()(CustomExtensionAuthenticationConfigurationable) {
    return m.authenticationConfiguration
}
// GetEndpointConfiguration gets the endpointConfiguration property value. The endpointConfiguration property
func (m *IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) GetEndpointConfiguration()(CustomExtensionEndpointConfigurationable) {
    return m.endpointConfiguration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomExtensionAuthenticationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationConfiguration(val.(CustomExtensionAuthenticationConfigurationable))
        }
        return nil
    }
    res["endpointConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomExtensionEndpointConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointConfiguration(val.(CustomExtensionEndpointConfigurationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAuthenticationConfiguration sets the authenticationConfiguration property value. The authenticationConfiguration property
func (m *IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) SetAuthenticationConfiguration(value CustomExtensionAuthenticationConfigurationable)() {
    m.authenticationConfiguration = value
}
// SetEndpointConfiguration sets the endpointConfiguration property value. The endpointConfiguration property
func (m *IdentityCustomAuthenticationExtensionsValidateAuthenticationConfigurationPostRequestBody) SetEndpointConfiguration(value CustomExtensionEndpointConfigurationable)() {
    m.endpointConfiguration = value
}
