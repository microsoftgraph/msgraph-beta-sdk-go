package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody 
type IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The payloadIds property
    payloadIds []string
}
// NewIosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody instantiates a new IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody and sets the default values.
func NewIosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody()(*IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody) {
    m := &IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["payloadIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPayloadIds(res)
        }
        return nil
    }
    return res
}
// GetPayloadIds gets the payloadIds property value. The payloadIds property
func (m *IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody) GetPayloadIds()([]string) {
    return m.payloadIds
}
// Serialize serializes information the current object
func (m *IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPayloadIds() != nil {
        err := writer.WriteCollectionOfStringValues("payloadIds", m.GetPayloadIds())
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
func (m *IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPayloadIds sets the payloadIds property value. The payloadIds property
func (m *IosLobAppProvisioningConfigurationsMicrosoftGraphHasPayloadLinksHasPayloadLinksPostRequestBody) SetPayloadIds(value []string)() {
    m.payloadIds = value
}
