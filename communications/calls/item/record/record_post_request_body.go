package record

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// RecordPostRequestBody provides operations to call the record method.
type RecordPostRequestBody struct {
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
    prompts []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Promptable
    // The stopTones property
    stopTones []string
    // The streamWhileRecording property
    streamWhileRecording *bool
}
// NewRecordPostRequestBody instantiates a new recordPostRequestBody and sets the default values.
func NewRecordPostRequestBody()(*RecordPostRequestBody) {
    m := &RecordPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRecordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecordPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBargeInAllowed gets the bargeInAllowed property value. The bargeInAllowed property
func (m *RecordPostRequestBody) GetBargeInAllowed()(*bool) {
    return m.bargeInAllowed
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *RecordPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bargeInAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBargeInAllowed)
    res["clientContext"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetClientContext)
    res["initialSilenceTimeoutInSeconds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetInitialSilenceTimeoutInSeconds)
    res["maxRecordDurationInSeconds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaxRecordDurationInSeconds)
    res["maxSilenceTimeoutInSeconds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaxSilenceTimeoutInSeconds)
    res["playBeep"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPlayBeep)
    res["prompts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePromptFromDiscriminatorValue , m.SetPrompts)
    res["stopTones"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetStopTones)
    res["streamWhileRecording"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStreamWhileRecording)
    return res
}
// GetInitialSilenceTimeoutInSeconds gets the initialSilenceTimeoutInSeconds property value. The initialSilenceTimeoutInSeconds property
func (m *RecordPostRequestBody) GetInitialSilenceTimeoutInSeconds()(*int32) {
    return m.initialSilenceTimeoutInSeconds
}
// GetMaxRecordDurationInSeconds gets the maxRecordDurationInSeconds property value. The maxRecordDurationInSeconds property
func (m *RecordPostRequestBody) GetMaxRecordDurationInSeconds()(*int32) {
    return m.maxRecordDurationInSeconds
}
// GetMaxSilenceTimeoutInSeconds gets the maxSilenceTimeoutInSeconds property value. The maxSilenceTimeoutInSeconds property
func (m *RecordPostRequestBody) GetMaxSilenceTimeoutInSeconds()(*int32) {
    return m.maxSilenceTimeoutInSeconds
}
// GetPlayBeep gets the playBeep property value. The playBeep property
func (m *RecordPostRequestBody) GetPlayBeep()(*bool) {
    return m.playBeep
}
// GetPrompts gets the prompts property value. The prompts property
func (m *RecordPostRequestBody) GetPrompts()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Promptable) {
    return m.prompts
}
// GetStopTones gets the stopTones property value. The stopTones property
func (m *RecordPostRequestBody) GetStopTones()([]string) {
    return m.stopTones
}
// GetStreamWhileRecording gets the streamWhileRecording property value. The streamWhileRecording property
func (m *RecordPostRequestBody) GetStreamWhileRecording()(*bool) {
    return m.streamWhileRecording
}
// Serialize serializes information the current object
func (m *RecordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPrompts())
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
func (m *RecordPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBargeInAllowed sets the bargeInAllowed property value. The bargeInAllowed property
func (m *RecordPostRequestBody) SetBargeInAllowed(value *bool)() {
    m.bargeInAllowed = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *RecordPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetInitialSilenceTimeoutInSeconds sets the initialSilenceTimeoutInSeconds property value. The initialSilenceTimeoutInSeconds property
func (m *RecordPostRequestBody) SetInitialSilenceTimeoutInSeconds(value *int32)() {
    m.initialSilenceTimeoutInSeconds = value
}
// SetMaxRecordDurationInSeconds sets the maxRecordDurationInSeconds property value. The maxRecordDurationInSeconds property
func (m *RecordPostRequestBody) SetMaxRecordDurationInSeconds(value *int32)() {
    m.maxRecordDurationInSeconds = value
}
// SetMaxSilenceTimeoutInSeconds sets the maxSilenceTimeoutInSeconds property value. The maxSilenceTimeoutInSeconds property
func (m *RecordPostRequestBody) SetMaxSilenceTimeoutInSeconds(value *int32)() {
    m.maxSilenceTimeoutInSeconds = value
}
// SetPlayBeep sets the playBeep property value. The playBeep property
func (m *RecordPostRequestBody) SetPlayBeep(value *bool)() {
    m.playBeep = value
}
// SetPrompts sets the prompts property value. The prompts property
func (m *RecordPostRequestBody) SetPrompts(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Promptable)() {
    m.prompts = value
}
// SetStopTones sets the stopTones property value. The stopTones property
func (m *RecordPostRequestBody) SetStopTones(value []string)() {
    m.stopTones = value
}
// SetStreamWhileRecording sets the streamWhileRecording property value. The streamWhileRecording property
func (m *RecordPostRequestBody) SetStreamWhileRecording(value *bool)() {
    m.streamWhileRecording = value
}
