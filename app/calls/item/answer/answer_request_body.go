package answer

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AnswerRequestBody provides operations to call the answer method.
type AnswerRequestBody struct {
    // The acceptedModalities property
    acceptedModalities []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Modality
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The callbackUri property
    callbackUri *string
    // The callOptions property
    callOptions ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IncomingCallOptionsable
    // The mediaConfig property
    mediaConfig ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MediaConfigable
    // The participantCapacity property
    participantCapacity *int32
}
// NewAnswerRequestBody instantiates a new answerRequestBody and sets the default values.
func NewAnswerRequestBody()(*AnswerRequestBody) {
    m := &AnswerRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAnswerRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAnswerRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnswerRequestBody(), nil
}
// GetAcceptedModalities gets the acceptedModalities property value. The acceptedModalities property
func (m *AnswerRequestBody) GetAcceptedModalities()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Modality) {
    if m == nil {
        return nil
    } else {
        return m.acceptedModalities
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *AnswerRequestBody) GetCallbackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callbackUri
    }
}
// GetCallOptions gets the callOptions property value. The callOptions property
func (m *AnswerRequestBody) GetCallOptions()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IncomingCallOptionsable) {
    if m == nil {
        return nil
    } else {
        return m.callOptions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AnswerRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptedModalities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Modality, len(val))
            for i, v := range val {
                res[i] = *(v.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Modality))
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
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIncomingCallOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallOptions(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IncomingCallOptionsable))
        }
        return nil
    }
    res["mediaConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMediaConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaConfig(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MediaConfigable))
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
func (m *AnswerRequestBody) GetMediaConfig()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MediaConfigable) {
    if m == nil {
        return nil
    } else {
        return m.mediaConfig
    }
}
// GetParticipantCapacity gets the participantCapacity property value. The participantCapacity property
func (m *AnswerRequestBody) GetParticipantCapacity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.participantCapacity
    }
}
// Serialize serializes information the current object
func (m *AnswerRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAcceptedModalities() != nil {
        err := writer.WriteCollectionOfStringValues("acceptedModalities", ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SerializeModality(m.GetAcceptedModalities()))
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
func (m *AnswerRequestBody) SetAcceptedModalities(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Modality)() {
    if m != nil {
        m.acceptedModalities = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *AnswerRequestBody) SetCallbackUri(value *string)() {
    if m != nil {
        m.callbackUri = value
    }
}
// SetCallOptions sets the callOptions property value. The callOptions property
func (m *AnswerRequestBody) SetCallOptions(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IncomingCallOptionsable)() {
    if m != nil {
        m.callOptions = value
    }
}
// SetMediaConfig sets the mediaConfig property value. The mediaConfig property
func (m *AnswerRequestBody) SetMediaConfig(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MediaConfigable)() {
    if m != nil {
        m.mediaConfig = value
    }
}
// SetParticipantCapacity sets the participantCapacity property value. The participantCapacity property
func (m *AnswerRequestBody) SetParticipantCapacity(value *int32)() {
    if m != nil {
        m.participantCapacity = value
    }
}
