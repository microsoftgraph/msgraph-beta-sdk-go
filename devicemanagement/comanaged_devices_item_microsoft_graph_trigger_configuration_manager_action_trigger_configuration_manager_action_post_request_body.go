package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody 
type ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Parameter for action triggerConfigurationManagerAction
    configurationManagerAction ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable
}
// NewComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody instantiates a new ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody and sets the default values.
func NewComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody()(*ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody) {
    m := &ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfigurationManagerAction gets the configurationManagerAction property value. Parameter for action triggerConfigurationManagerAction
func (m *ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody) GetConfigurationManagerAction()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable) {
    return m.configurationManagerAction
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["configurationManagerAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConfigurationManagerActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerAction(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("configurationManagerAction", m.GetConfigurationManagerAction())
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
func (m *ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfigurationManagerAction sets the configurationManagerAction property value. Parameter for action triggerConfigurationManagerAction
func (m *ComanagedDevicesItemMicrosoftGraphTriggerConfigurationManagerActionTriggerConfigurationManagerActionPostRequestBody) SetConfigurationManagerAction(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigurationManagerActionable)() {
    m.configurationManagerAction = value
}
