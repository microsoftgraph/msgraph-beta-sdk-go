package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppCallsItemAnswerPostRequestBody provides operations to call the answer method.
type AppCallsItemAnswerPostRequestBody struct {
    // The acceptedModalities property
    acceptedModalities []Modality
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The callbackUri property
    callbackUri *string
    // The callOptions property
    callOptions IncomingCallOptionsable
    // The mediaConfig property
    mediaConfig MediaConfigable
    // The participantCapacity property
    participantCapacity *int32
}
// NewAppCallsItemAnswerPostRequestBody instantiates a new AppCallsItemAnswerPostRequestBody and sets the default values.
func NewAppCallsItemAnswerPostRequestBody()(*AppCallsItemAnswerPostRequestBody) {
    m := &AppCallsItemAnswerPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppCallsItemAnswerPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCallsItemAnswerPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppCallsItemAnswerPostRequestBody(), nil
}
// GetAcceptedModalities gets the acceptedModalities property value. The acceptedModalities property
func (m *AppCallsItemAnswerPostRequestBody) GetAcceptedModalities()([]Modality) {
    return m.acceptedModalities
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppCallsItemAnswerPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *AppCallsItemAnswerPostRequestBody) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetCallOptions gets the callOptions property value. The callOptions property
func (m *AppCallsItemAnswerPostRequestBody) GetCallOptions()(IncomingCallOptionsable) {
    return m.callOptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCallsItemAnswerPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptedModalities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Modality, len(val))
            for i, v := range val {
                res[i] = *(v.(*Modality))
            }
            m.SetAcceptedModalities(res)
        }
        return nil
    }
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
    res["callOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIncomingCallOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallOptions(val.(IncomingCallOptionsable))
        }
        return nil
    }
    res["mediaConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaConfig(val.(MediaConfigable))
        }
        return nil
    }
    res["participantCapacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantCapacity(val)
        }
        return nil
    }
    return res
}
// GetMediaConfig gets the mediaConfig property value. The mediaConfig property
func (m *AppCallsItemAnswerPostRequestBody) GetMediaConfig()(MediaConfigable) {
    return m.mediaConfig
}
// GetParticipantCapacity gets the participantCapacity property value. The participantCapacity property
func (m *AppCallsItemAnswerPostRequestBody) GetParticipantCapacity()(*int32) {
    return m.participantCapacity
}
// Serialize serializes information the current object
func (m *AppCallsItemAnswerPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAcceptedModalities() != nil {
        err := writer.WriteCollectionOfStringValues("acceptedModalities", SerializeModality(m.GetAcceptedModalities()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("callOptions", m.GetCallOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mediaConfig", m.GetMediaConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("participantCapacity", m.GetParticipantCapacity())
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
// SetAcceptedModalities sets the acceptedModalities property value. The acceptedModalities property
func (m *AppCallsItemAnswerPostRequestBody) SetAcceptedModalities(value []Modality)() {
    m.acceptedModalities = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppCallsItemAnswerPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *AppCallsItemAnswerPostRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetCallOptions sets the callOptions property value. The callOptions property
func (m *AppCallsItemAnswerPostRequestBody) SetCallOptions(value IncomingCallOptionsable)() {
    m.callOptions = value
}
// SetMediaConfig sets the mediaConfig property value. The mediaConfig property
func (m *AppCallsItemAnswerPostRequestBody) SetMediaConfig(value MediaConfigable)() {
    m.mediaConfig = value
}
// SetParticipantCapacity sets the participantCapacity property value. The participantCapacity property
func (m *AppCallsItemAnswerPostRequestBody) SetParticipantCapacity(value *int32)() {
    m.participantCapacity = value
}
