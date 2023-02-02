package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody 
type AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The parentUri property
    parentUri *string
}
// NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody instantiates a new AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody) GetParentUri()(*string) {
    return m.parentUri
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetParentUri sets the parentUri property value. The parentUri property
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphCreateGooglePlayWebTokenCreateGooglePlayWebTokenPostRequestBody) SetParentUri(value *string)() {
    m.parentUri = value
}
