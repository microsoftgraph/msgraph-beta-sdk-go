package requestremotehelpsessionaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestRemoteHelpSessionAccessPostRequestBody provides operations to call the requestRemoteHelpSessionAccess method.
type RequestRemoteHelpSessionAccessPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The pubSubConnectionId property
    pubSubConnectionId *string
    // The sessionKey property
    sessionKey *string
}
// NewRequestRemoteHelpSessionAccessPostRequestBody instantiates a new requestRemoteHelpSessionAccessPostRequestBody and sets the default values.
func NewRequestRemoteHelpSessionAccessPostRequestBody()(*RequestRemoteHelpSessionAccessPostRequestBody) {
    m := &RequestRemoteHelpSessionAccessPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRequestRemoteHelpSessionAccessPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestRemoteHelpSessionAccessPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestRemoteHelpSessionAccessPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestRemoteHelpSessionAccessPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestRemoteHelpSessionAccessPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pubSubConnectionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPubSubConnectionId)
    res["sessionKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSessionKey)
    return res
}
// GetPubSubConnectionId gets the pubSubConnectionId property value. The pubSubConnectionId property
func (m *RequestRemoteHelpSessionAccessPostRequestBody) GetPubSubConnectionId()(*string) {
    return m.pubSubConnectionId
}
// GetSessionKey gets the sessionKey property value. The sessionKey property
func (m *RequestRemoteHelpSessionAccessPostRequestBody) GetSessionKey()(*string) {
    return m.sessionKey
}
// Serialize serializes information the current object
func (m *RequestRemoteHelpSessionAccessPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("pubSubConnectionId", m.GetPubSubConnectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionKey", m.GetSessionKey())
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
func (m *RequestRemoteHelpSessionAccessPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPubSubConnectionId sets the pubSubConnectionId property value. The pubSubConnectionId property
func (m *RequestRemoteHelpSessionAccessPostRequestBody) SetPubSubConnectionId(value *string)() {
    m.pubSubConnectionId = value
}
// SetSessionKey sets the sessionKey property value. The sessionKey property
func (m *RequestRemoteHelpSessionAccessPostRequestBody) SetSessionKey(value *string)() {
    m.sessionKey = value
}
