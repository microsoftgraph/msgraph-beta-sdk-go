package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody 
type UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody struct {
    // Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
    actionName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceScopeId property
    deviceScopeId *string
}
// NewUserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody instantiates a new UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody and sets the default values.
func NewUserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody()(*UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) {
    m := &UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody(), nil
}
// GetActionName gets the actionName property value. Trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) GetActionName()(*string) {
    return m.actionName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceScopeId gets the deviceScopeId property value. The deviceScopeId property
func (m *UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) GetDeviceScopeId()(*string) {
    return m.deviceScopeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val)
        }
        return nil
    }
    res["deviceScopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceScopeId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) SetActionName(value *string)() {
    m.actionName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceScopeId sets the deviceScopeId property value. The deviceScopeId property
func (m *UserExperienceAnalyticsDeviceScopeMicrosoftGraphTriggerDeviceScopeActionTriggerDeviceScopeActionPostRequestBody) SetDeviceScopeId(value *string)() {
    m.deviceScopeId = value
}
