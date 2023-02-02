package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody 
type AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The productIds property
    productIds []string
}
// NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody instantiates a new AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody()(*AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["productIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetProductIds(res)
        }
        return nil
    }
    return res
}
// GetProductIds gets the productIds property value. The productIds property
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody) GetProductIds()([]string) {
    return m.productIds
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetProductIds() != nil {
        err := writer.WriteCollectionOfStringValues("productIds", m.GetProductIds())
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetProductIds sets the productIds property value. The productIds property
func (m *AndroidManagedStoreAccountEnterpriseSettingsMicrosoftGraphAddAppsAddAppsPostRequestBody) SetProductIds(value []string)() {
    m.productIds = value
}
