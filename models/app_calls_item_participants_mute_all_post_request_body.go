package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCallsItemParticipantsMuteAllPostRequestBody provides operations to call the muteAll method.
type AppCallsItemParticipantsMuteAllPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The clientContext property
    clientContext *string
    // The participants property
    participants []string
}
// NewAppCallsItemParticipantsMuteAllPostRequestBody instantiates a new AppCallsItemParticipantsMuteAllPostRequestBody and sets the default values.
func NewAppCallsItemParticipantsMuteAllPostRequestBody()(*AppCallsItemParticipantsMuteAllPostRequestBody) {
    m := &AppCallsItemParticipantsMuteAllPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppCallsItemParticipantsMuteAllPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCallsItemParticipantsMuteAllPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppCallsItemParticipantsMuteAllPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppCallsItemParticipantsMuteAllPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *AppCallsItemParticipantsMuteAllPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCallsItemParticipantsMuteAllPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetParticipants(res)
        }
        return nil
    }
    return res
}
// GetParticipants gets the participants property value. The participants property
func (m *AppCallsItemParticipantsMuteAllPostRequestBody) GetParticipants()([]string) {
    return m.participants
}
// Serialize serializes information the current object
func (m *AppCallsItemParticipantsMuteAllPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    if m.GetParticipants() != nil {
        err := writer.WriteCollectionOfStringValues("participants", m.GetParticipants())
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
func (m *AppCallsItemParticipantsMuteAllPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *AppCallsItemParticipantsMuteAllPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetParticipants sets the participants property value. The participants property
func (m *AppCallsItemParticipantsMuteAllPostRequestBody) SetParticipants(value []string)() {
    m.participants = value
}
