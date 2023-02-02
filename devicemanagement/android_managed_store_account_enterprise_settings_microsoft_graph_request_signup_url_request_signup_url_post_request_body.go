package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody 
type AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The hostName property
    hostName *string
}
// NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody instantiates a new AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hostName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. The hostName property
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody) GetHostName()(*string) {
    return m.hostName
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hostName", m.GetHostName())
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHostName sets the hostName property value. The hostName property
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphRequestSignupUrlRequestSignupUrlPostRequestBody) SetHostName(value *string)() {
    m.hostName = value
}
