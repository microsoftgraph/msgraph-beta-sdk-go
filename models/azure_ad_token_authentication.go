package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureAdTokenAuthentication 
type AzureAdTokenAuthentication struct {
    CustomExtensionAuthenticationConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appID of the Azure AD application to use to authenticate a logic app with a custom access package workflow extension.
    resourceId *string
}
// NewAzureAdTokenAuthentication instantiates a new AzureAdTokenAuthentication and sets the default values.
func NewAzureAdTokenAuthentication()(*AzureAdTokenAuthentication) {
    m := &AzureAdTokenAuthentication{
        CustomExtensionAuthenticationConfiguration: *NewCustomExtensionAuthenticationConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAzureAdTokenAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureAdTokenAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureAdTokenAuthentication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureAdTokenAuthentication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureAdTokenAuthentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionAuthenticationConfiguration.GetFieldDeserializers()
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
// GetResourceId gets the resourceId property value. The appID of the Azure AD application to use to authenticate a logic app with a custom access package workflow extension.
func (m *AzureAdTokenAuthentication) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// Serialize serializes information the current object
func (m *AzureAdTokenAuthentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionAuthenticationConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AzureAdTokenAuthentication) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResourceId sets the resourceId property value. The appID of the Azure AD application to use to authenticate a logic app with a custom access package workflow extension.
func (m *AzureAdTokenAuthentication) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
