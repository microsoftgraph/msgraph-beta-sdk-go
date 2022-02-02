package record

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// RecordRequestBody 
type RecordRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    bargeInAllowed *bool;
    // 
    clientContext *string;
    // 
    initialSilenceTimeoutInSeconds *int32;
    // 
    maxRecordDurationInSeconds *int32;
    // 
    maxSilenceTimeoutInSeconds *int32;
    // 
    playBeep *bool;
    // 
    prompts []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt;
    // 
    stopTones []string;
    // 
    streamWhileRecording *bool;
}
// NewRecordRequestBody instantiates a new recordRequestBody and sets the default values.
func NewRecordRequestBody()(*RecordRequestBody) {
    m := &RecordRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecordRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBargeInAllowed gets the bargeInAllowed property value. 
func (m *RecordRequestBody) GetBargeInAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bargeInAllowed
    }
}
// GetClientContext gets the clientContext property value. 
func (m *RecordRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
// GetInitialSilenceTimeoutInSeconds gets the initialSilenceTimeoutInSeconds property value. 
func (m *RecordRequestBody) GetInitialSilenceTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.initialSilenceTimeoutInSeconds
    }
}
// GetMaxRecordDurationInSeconds gets the maxRecordDurationInSeconds property value. 
func (m *RecordRequestBody) GetMaxRecordDurationInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxRecordDurationInSeconds
    }
}
// GetMaxSilenceTimeoutInSeconds gets the maxSilenceTimeoutInSeconds property value. 
func (m *RecordRequestBody) GetMaxSilenceTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxSilenceTimeoutInSeconds
    }
}
// GetPlayBeep gets the playBeep property value. 
func (m *RecordRequestBody) GetPlayBeep()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.playBeep
    }
}
// GetPrompts gets the prompts property value. 
func (m *RecordRequestBody) GetPrompts()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt) {
    if m == nil {
        return nil
    } else {
        return m.prompts
    }
}
// GetStopTones gets the stopTones property value. 
func (m *RecordRequestBody) GetStopTones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.stopTones
    }
}
// GetStreamWhileRecording gets the streamWhileRecording property value. 
func (m *RecordRequestBody) GetStreamWhileRecording()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.streamWhileRecording
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecordRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bargeInAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBargeInAllowed(val)
        }
        return nil
    }
    res["clientContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    res["initialSilenceTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitialSilenceTimeoutInSeconds(val)
        }
        return nil
    }
    res["maxRecordDurationInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxRecordDurationInSeconds(val)
        }
        return nil
    }
    res["maxSilenceTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxSilenceTimeoutInSeconds(val)
        }
        return nil
    }
    res["playBeep"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlayBeep(val)
        }
        return nil
    }
    res["prompts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrompt() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt))
            }
            m.SetPrompts(res)
        }
        return nil
    }
    res["stopTones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["streamWhileRecording"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *RecordRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RecordRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPrompts()))
        for i, v := range m.GetPrompts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *RecordRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBargeInAllowed sets the bargeInAllowed property value. 
func (m *RecordRequestBody) SetBargeInAllowed(value *bool)() {
    if m != nil {
        m.bargeInAllowed = value
    }
}
// SetClientContext sets the clientContext property value. 
func (m *RecordRequestBody) SetClientContext(value *string)() {
    if m != nil {
        m.clientContext = value
    }
}
// SetInitialSilenceTimeoutInSeconds sets the initialSilenceTimeoutInSeconds property value. 
func (m *RecordRequestBody) SetInitialSilenceTimeoutInSeconds(value *int32)() {
    if m != nil {
        m.initialSilenceTimeoutInSeconds = value
    }
}
// SetMaxRecordDurationInSeconds sets the maxRecordDurationInSeconds property value. 
func (m *RecordRequestBody) SetMaxRecordDurationInSeconds(value *int32)() {
    if m != nil {
        m.maxRecordDurationInSeconds = value
    }
}
// SetMaxSilenceTimeoutInSeconds sets the maxSilenceTimeoutInSeconds property value. 
func (m *RecordRequestBody) SetMaxSilenceTimeoutInSeconds(value *int32)() {
    if m != nil {
        m.maxSilenceTimeoutInSeconds = value
    }
}
// SetPlayBeep sets the playBeep property value. 
func (m *RecordRequestBody) SetPlayBeep(value *bool)() {
    if m != nil {
        m.playBeep = value
    }
}
// SetPrompts sets the prompts property value. 
func (m *RecordRequestBody) SetPrompts(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt)() {
    if m != nil {
        m.prompts = value
    }
}
// SetStopTones sets the stopTones property value. 
func (m *RecordRequestBody) SetStopTones(value []string)() {
    if m != nil {
        m.stopTones = value
    }
}
// SetStreamWhileRecording sets the streamWhileRecording property value. 
func (m *RecordRequestBody) SetStreamWhileRecording(value *bool)() {
    if m != nil {
        m.streamWhileRecording = value
    }
}
