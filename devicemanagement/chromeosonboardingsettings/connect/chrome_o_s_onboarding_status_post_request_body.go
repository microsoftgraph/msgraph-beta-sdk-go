package connect

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChromeOSOnboardingStatusPostRequestBody provides operations to call the connect method.
type ChromeOSOnboardingStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ownerAccessToken property
    ownerAccessToken *string
    // The ownerUserPrincipalName property
    ownerUserPrincipalName *string
}
// NewChromeOSOnboardingStatusPostRequestBody instantiates a new ChromeOSOnboardingStatusPostRequestBody and sets the default values.
func NewChromeOSOnboardingStatusPostRequestBody()(*ChromeOSOnboardingStatusPostRequestBody) {
    m := &ChromeOSOnboardingStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChromeOSOnboardingStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChromeOSOnboardingStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChromeOSOnboardingStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChromeOSOnboardingStatusPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChromeOSOnboardingStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ownerAccessToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerAccessToken(val)
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetOwnerAccessToken gets the ownerAccessToken property value. The ownerAccessToken property
func (m *ChromeOSOnboardingStatusPostRequestBody) GetOwnerAccessToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerAccessToken
    }
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. The ownerUserPrincipalName property
func (m *ChromeOSOnboardingStatusPostRequestBody) GetOwnerUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerUserPrincipalName
    }
}
// Serialize serializes information the current object
func (m *ChromeOSOnboardingStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ownerAccessToken", m.GetOwnerAccessToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ownerUserPrincipalName", m.GetOwnerUserPrincipalName())
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
func (m *ChromeOSOnboardingStatusPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOwnerAccessToken sets the ownerAccessToken property value. The ownerAccessToken property
func (m *ChromeOSOnboardingStatusPostRequestBody) SetOwnerAccessToken(value *string)() {
    if m != nil {
        m.ownerAccessToken = value
    }
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. The ownerUserPrincipalName property
func (m *ChromeOSOnboardingStatusPostRequestBody) SetOwnerUserPrincipalName(value *string)() {
    if m != nil {
        m.ownerUserPrincipalName = value
    }
}
