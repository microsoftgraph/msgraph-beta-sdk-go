package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCallsItemRedirectPostRequestBody provides operations to call the redirect method.
type AppCallsItemRedirectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The callbackUri property
    callbackUri *string
    // The maskCallee property
    maskCallee *bool
    // The maskCaller property
    maskCaller *bool
    // The targetDisposition property
    targetDisposition *CallDisposition
    // The targets property
    targets []InvitationParticipantInfoable
    // The timeout property
    timeout *int32
}
// NewAppCallsItemRedirectPostRequestBody instantiates a new AppCallsItemRedirectPostRequestBody and sets the default values.
func NewAppCallsItemRedirectPostRequestBody()(*AppCallsItemRedirectPostRequestBody) {
    m := &AppCallsItemRedirectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppCallsItemRedirectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCallsItemRedirectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppCallsItemRedirectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppCallsItemRedirectPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *AppCallsItemRedirectPostRequestBody) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCallsItemRedirectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callbackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["maskCallee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaskCallee(val)
        }
        return nil
    }
    res["maskCaller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaskCaller(val)
        }
        return nil
    }
    res["targetDisposition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallDisposition)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisposition(val.(*CallDisposition))
        }
        return nil
    }
    res["targets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInvitationParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InvitationParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(InvitationParticipantInfoable)
            }
            m.SetTargets(res)
        }
        return nil
    }
    res["timeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeout(val)
        }
        return nil
    }
    return res
}
// GetMaskCallee gets the maskCallee property value. The maskCallee property
func (m *AppCallsItemRedirectPostRequestBody) GetMaskCallee()(*bool) {
    return m.maskCallee
}
// GetMaskCaller gets the maskCaller property value. The maskCaller property
func (m *AppCallsItemRedirectPostRequestBody) GetMaskCaller()(*bool) {
    return m.maskCaller
}
// GetTargetDisposition gets the targetDisposition property value. The targetDisposition property
func (m *AppCallsItemRedirectPostRequestBody) GetTargetDisposition()(*CallDisposition) {
    return m.targetDisposition
}
// GetTargets gets the targets property value. The targets property
func (m *AppCallsItemRedirectPostRequestBody) GetTargets()([]InvitationParticipantInfoable) {
    return m.targets
}
// GetTimeout gets the timeout property value. The timeout property
func (m *AppCallsItemRedirectPostRequestBody) GetTimeout()(*int32) {
    return m.timeout
}
// Serialize serializes information the current object
func (m *AppCallsItemRedirectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maskCallee", m.GetMaskCallee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maskCaller", m.GetMaskCaller())
        if err != nil {
            return err
        }
    }
    if m.GetTargetDisposition() != nil {
        cast := (*m.GetTargetDisposition()).String()
        err := writer.WriteStringValue("targetDisposition", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeout", m.GetTimeout())
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
func (m *AppCallsItemRedirectPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *AppCallsItemRedirectPostRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetMaskCallee sets the maskCallee property value. The maskCallee property
func (m *AppCallsItemRedirectPostRequestBody) SetMaskCallee(value *bool)() {
    m.maskCallee = value
}
// SetMaskCaller sets the maskCaller property value. The maskCaller property
func (m *AppCallsItemRedirectPostRequestBody) SetMaskCaller(value *bool)() {
    m.maskCaller = value
}
// SetTargetDisposition sets the targetDisposition property value. The targetDisposition property
func (m *AppCallsItemRedirectPostRequestBody) SetTargetDisposition(value *CallDisposition)() {
    m.targetDisposition = value
}
// SetTargets sets the targets property value. The targets property
func (m *AppCallsItemRedirectPostRequestBody) SetTargets(value []InvitationParticipantInfoable)() {
    m.targets = value
}
// SetTimeout sets the timeout property value. The timeout property
func (m *AppCallsItemRedirectPostRequestBody) SetTimeout(value *int32)() {
    m.timeout = value
}
