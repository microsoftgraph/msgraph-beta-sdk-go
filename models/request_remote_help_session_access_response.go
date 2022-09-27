package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestRemoteHelpSessionAccessResponse remote help - response we provide back to the helper after getting response from pubSub
type RequestRemoteHelpSessionAccessResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // AES encryption Initialization Vector for encrypting client messages sent to PubSub
    pubSubEncryption *string
    // The unique identifier for encrypting client messages sent to PubSub
    pubSubEncryptionKey *string
    // The unique identifier for a session
    sessionKey *string
}
// NewRequestRemoteHelpSessionAccessResponse instantiates a new requestRemoteHelpSessionAccessResponse and sets the default values.
func NewRequestRemoteHelpSessionAccessResponse()(*RequestRemoteHelpSessionAccessResponse) {
    m := &RequestRemoteHelpSessionAccessResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.requestRemoteHelpSessionAccessResponse";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRequestRemoteHelpSessionAccessResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestRemoteHelpSessionAccessResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestRemoteHelpSessionAccessResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestRemoteHelpSessionAccessResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestRemoteHelpSessionAccessResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["pubSubEncryption"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPubSubEncryption)
    res["pubSubEncryptionKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPubSubEncryptionKey)
    res["sessionKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSessionKey)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RequestRemoteHelpSessionAccessResponse) GetOdataType()(*string) {
    return m.odataType
}
// GetPubSubEncryption gets the pubSubEncryption property value. AES encryption Initialization Vector for encrypting client messages sent to PubSub
func (m *RequestRemoteHelpSessionAccessResponse) GetPubSubEncryption()(*string) {
    return m.pubSubEncryption
}
// GetPubSubEncryptionKey gets the pubSubEncryptionKey property value. The unique identifier for encrypting client messages sent to PubSub
func (m *RequestRemoteHelpSessionAccessResponse) GetPubSubEncryptionKey()(*string) {
    return m.pubSubEncryptionKey
}
// GetSessionKey gets the sessionKey property value. The unique identifier for a session
func (m *RequestRemoteHelpSessionAccessResponse) GetSessionKey()(*string) {
    return m.sessionKey
}
// Serialize serializes information the current object
func (m *RequestRemoteHelpSessionAccessResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pubSubEncryption", m.GetPubSubEncryption())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pubSubEncryptionKey", m.GetPubSubEncryptionKey())
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
func (m *RequestRemoteHelpSessionAccessResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RequestRemoteHelpSessionAccessResponse) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPubSubEncryption sets the pubSubEncryption property value. AES encryption Initialization Vector for encrypting client messages sent to PubSub
func (m *RequestRemoteHelpSessionAccessResponse) SetPubSubEncryption(value *string)() {
    m.pubSubEncryption = value
}
// SetPubSubEncryptionKey sets the pubSubEncryptionKey property value. The unique identifier for encrypting client messages sent to PubSub
func (m *RequestRemoteHelpSessionAccessResponse) SetPubSubEncryptionKey(value *string)() {
    m.pubSubEncryptionKey = value
}
// SetSessionKey sets the sessionKey property value. The unique identifier for a session
func (m *RequestRemoteHelpSessionAccessResponse) SetSessionKey(value *string)() {
    m.sessionKey = value
}
