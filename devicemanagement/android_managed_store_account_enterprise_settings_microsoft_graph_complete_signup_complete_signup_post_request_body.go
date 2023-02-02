package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody 
type AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enterpriseToken property
    enterpriseToken *string
}
// NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody instantiates a new AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnterpriseToken gets the enterpriseToken property value. The enterpriseToken property
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody) GetEnterpriseToken()(*string) {
    return m.enterpriseToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enterpriseToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseToken(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("enterpriseToken", m.GetEnterpriseToken())
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnterpriseToken sets the enterpriseToken property value. The enterpriseToken property
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCompleteSignupCompleteSignupPostRequestBody) SetEnterpriseToken(value *string)() {
    m.enterpriseToken = value
}
