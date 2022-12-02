package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallsItemRecordPostRequestBody provides operations to call the record method.
type CommunicationsCallsItemRecordPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The bargeInAllowed property
    bargeInAllowed *bool
    // The clientContext property
    clientContext *string
    // The initialSilenceTimeoutInSeconds property
    initialSilenceTimeoutInSeconds *int32
    // The maxRecordDurationInSeconds property
    maxRecordDurationInSeconds *int32
    // The maxSilenceTimeoutInSeconds property
    maxSilenceTimeoutInSeconds *int32
    // The playBeep property
    playBeep *bool
    // The prompts property
    prompts []Promptable
    // The stopTones property
    stopTones []string
    // The streamWhileRecording property
    streamWhileRecording *bool
}
// NewCommunicationsCallsItemRecordPostRequestBody instantiates a new CommunicationsCallsItemRecordPostRequestBody and sets the default values.
func NewCommunicationsCallsItemRecordPostRequestBody()(*CommunicationsCallsItemRecordPostRequestBody) {
    m := &CommunicationsCallsItemRecordPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommunicationsCallsItemRecordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsCallsItemRecordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsCallsItemRecordPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsCallsItemRecordPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBargeInAllowed gets the bargeInAllowed property value. The bargeInAllowed property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetBargeInAllowed()(*bool) {
    return m.bargeInAllowed
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsCallsItemRecordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bargeInAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBargeInAllowed(val)
        }
        return nil
    }
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
    res["initialSilenceTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitialSilenceTimeoutInSeconds(val)
        }
        return nil
    }
    res["maxRecordDurationInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxRecordDurationInSeconds(val)
        }
        return nil
    }
    res["maxSilenceTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSilenceTimeoutInSeconds(val)
        }
        return nil
    }
    res["playBeep"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlayBeep(val)
        }
        return nil
    }
    res["prompts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePromptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Promptable, len(val))
            for i, v := range val {
                res[i] = v.(Promptable)
            }
            m.SetPrompts(res)
        }
        return nil
    }
    res["stopTones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStopTones(res)
        }
        return nil
    }
    res["streamWhileRecording"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreamWhileRecording(val)
        }
        return nil
    }
    return res
}
// GetInitialSilenceTimeoutInSeconds gets the initialSilenceTimeoutInSeconds property value. The initialSilenceTimeoutInSeconds property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetInitialSilenceTimeoutInSeconds()(*int32) {
    return m.initialSilenceTimeoutInSeconds
}
// GetMaxRecordDurationInSeconds gets the maxRecordDurationInSeconds property value. The maxRecordDurationInSeconds property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetMaxRecordDurationInSeconds()(*int32) {
    return m.maxRecordDurationInSeconds
}
// GetMaxSilenceTimeoutInSeconds gets the maxSilenceTimeoutInSeconds property value. The maxSilenceTimeoutInSeconds property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetMaxSilenceTimeoutInSeconds()(*int32) {
    return m.maxSilenceTimeoutInSeconds
}
// GetPlayBeep gets the playBeep property value. The playBeep property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetPlayBeep()(*bool) {
    return m.playBeep
}
// GetPrompts gets the prompts property value. The prompts property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetPrompts()([]Promptable) {
    return m.prompts
}
// GetStopTones gets the stopTones property value. The stopTones property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetStopTones()([]string) {
    return m.stopTones
}
// GetStreamWhileRecording gets the streamWhileRecording property value. The streamWhileRecording property
func (m *CommunicationsCallsItemRecordPostRequestBody) GetStreamWhileRecording()(*bool) {
    return m.streamWhileRecording
}
// Serialize serializes information the current object
func (m *CommunicationsCallsItemRecordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("bargeInAllowed", m.GetBargeInAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("initialSilenceTimeoutInSeconds", m.GetInitialSilenceTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxRecordDurationInSeconds", m.GetMaxRecordDurationInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxSilenceTimeoutInSeconds", m.GetMaxSilenceTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("playBeep", m.GetPlayBeep())
        if err != nil {
            return err
        }
    }
    if m.GetPrompts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrompts()))
        for i, v := range m.GetPrompts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("prompts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStopTones() != nil {
        err := writer.WriteCollectionOfStringValues("stopTones", m.GetStopTones())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("streamWhileRecording", m.GetStreamWhileRecording())
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
func (m *CommunicationsCallsItemRecordPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBargeInAllowed sets the bargeInAllowed property value. The bargeInAllowed property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetBargeInAllowed(value *bool)() {
    m.bargeInAllowed = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetInitialSilenceTimeoutInSeconds sets the initialSilenceTimeoutInSeconds property value. The initialSilenceTimeoutInSeconds property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetInitialSilenceTimeoutInSeconds(value *int32)() {
    m.initialSilenceTimeoutInSeconds = value
}
// SetMaxRecordDurationInSeconds sets the maxRecordDurationInSeconds property value. The maxRecordDurationInSeconds property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetMaxRecordDurationInSeconds(value *int32)() {
    m.maxRecordDurationInSeconds = value
}
// SetMaxSilenceTimeoutInSeconds sets the maxSilenceTimeoutInSeconds property value. The maxSilenceTimeoutInSeconds property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetMaxSilenceTimeoutInSeconds(value *int32)() {
    m.maxSilenceTimeoutInSeconds = value
}
// SetPlayBeep sets the playBeep property value. The playBeep property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetPlayBeep(value *bool)() {
    m.playBeep = value
}
// SetPrompts sets the prompts property value. The prompts property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetPrompts(value []Promptable)() {
    m.prompts = value
}
// SetStopTones sets the stopTones property value. The stopTones property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetStopTones(value []string)() {
    m.stopTones = value
}
// SetStreamWhileRecording sets the streamWhileRecording property value. The streamWhileRecording property
func (m *CommunicationsCallsItemRecordPostRequestBody) SetStreamWhileRecording(value *bool)() {
    m.streamWhileRecording = value
}
