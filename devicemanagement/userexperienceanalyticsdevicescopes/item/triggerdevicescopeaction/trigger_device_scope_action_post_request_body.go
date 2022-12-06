package triggerdevicescopeaction

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TriggerDeviceScopeActionPostRequestBody provides operations to call the triggerDeviceScopeAction method.
type TriggerDeviceScopeActionPostRequestBody struct {
    // Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
    actionName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceScopeId property
    deviceScopeId *string
}
// NewTriggerDeviceScopeActionPostRequestBody instantiates a new triggerDeviceScopeActionPostRequestBody and sets the default values.
func NewTriggerDeviceScopeActionPostRequestBody()(*TriggerDeviceScopeActionPostRequestBody) {
    m := &TriggerDeviceScopeActionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTriggerDeviceScopeActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTriggerDeviceScopeActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTriggerDeviceScopeActionPostRequestBody(), nil
}
// GetActionName gets the actionName property value. Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *TriggerDeviceScopeActionPostRequestBody) GetActionName()(*string) {
    return m.actionName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TriggerDeviceScopeActionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceScopeId gets the deviceScopeId property value. The deviceScopeId property
func (m *TriggerDeviceScopeActionPostRequestBody) GetDeviceScopeId()(*string) {
    return m.deviceScopeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TriggerDeviceScopeActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActionName)
    res["deviceScopeId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceScopeId)
    return res
}
// Serialize serializes information the current object
func (m *TriggerDeviceScopeActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceScopeId", m.GetDeviceScopeId())
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
// SetActionName sets the actionName property value. Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *TriggerDeviceScopeActionPostRequestBody) SetActionName(value *string)() {
    m.actionName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TriggerDeviceScopeActionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceScopeId sets the deviceScopeId property value. The deviceScopeId property
func (m *TriggerDeviceScopeActionPostRequestBody) SetDeviceScopeId(value *string)() {
    m.deviceScopeId = value
}
