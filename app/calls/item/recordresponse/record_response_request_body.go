package recordresponse

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type RecordResponseRequestBody struct {
    additionalData map[string]interface{};
    bargeInAllowed *bool;
    clientContext *string;
    initialSilenceTimeoutInSeconds *int32;
    maxRecordDurationInSeconds *int32;
    maxSilenceTimeoutInSeconds *int32;
    playBeep *bool;
    prompts []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt;
    stopTones []string;
    streamWhileRecording *bool;
}
func NewRecordResponseRequestBody()(*RecordResponseRequestBody) {
    m := &RecordResponseRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RecordResponseRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RecordResponseRequestBody) GetBargeInAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.bargeInAllowed
    }
}
func (m *RecordResponseRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
func (m *RecordResponseRequestBody) GetInitialSilenceTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.initialSilenceTimeoutInSeconds
    }
}
func (m *RecordResponseRequestBody) GetMaxRecordDurationInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxRecordDurationInSeconds
    }
}
func (m *RecordResponseRequestBody) GetMaxSilenceTimeoutInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxSilenceTimeoutInSeconds
    }
}
func (m *RecordResponseRequestBody) GetPlayBeep()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.playBeep
    }
}
func (m *RecordResponseRequestBody) GetPrompts()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt) {
    if m == nil {
        return nil
    } else {
        return m.prompts
    }
}
func (m *RecordResponseRequestBody) GetStopTones()([]string) {
    if m == nil {
        return nil
    } else {
        return m.stopTones
    }
}
func (m *RecordResponseRequestBody) GetStreamWhileRecording()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.streamWhileRecording
    }
}
func (m *RecordResponseRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bargeInAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetBargeInAllowed(val)
        return nil
    }
    res["clientContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientContext(val)
        return nil
    }
    res["initialSilenceTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInitialSilenceTimeoutInSeconds(val)
        return nil
    }
    res["maxRecordDurationInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaxRecordDurationInSeconds(val)
        return nil
    }
    res["maxSilenceTimeoutInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaxSilenceTimeoutInSeconds(val)
        return nil
    }
    res["playBeep"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPlayBeep(val)
        return nil
    }
    res["prompts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrompt() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt))
        }
        m.SetPrompts(res)
        return nil
    }
    res["stopTones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetStopTones(res)
        return nil
    }
    res["streamWhileRecording"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetStreamWhileRecording(val)
        return nil
    }
    return res
}
func (m *RecordResponseRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *RecordResponseRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
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
    {
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
func (m *RecordResponseRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RecordResponseRequestBody) SetBargeInAllowed(value *bool)() {
    m.bargeInAllowed = value
}
func (m *RecordResponseRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
func (m *RecordResponseRequestBody) SetInitialSilenceTimeoutInSeconds(value *int32)() {
    m.initialSilenceTimeoutInSeconds = value
}
func (m *RecordResponseRequestBody) SetMaxRecordDurationInSeconds(value *int32)() {
    m.maxRecordDurationInSeconds = value
}
func (m *RecordResponseRequestBody) SetMaxSilenceTimeoutInSeconds(value *int32)() {
    m.maxSilenceTimeoutInSeconds = value
}
func (m *RecordResponseRequestBody) SetPlayBeep(value *bool)() {
    m.playBeep = value
}
func (m *RecordResponseRequestBody) SetPrompts(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Prompt)() {
    m.prompts = value
}
func (m *RecordResponseRequestBody) SetStopTones(value []string)() {
    m.stopTones = value
}
func (m *RecordResponseRequestBody) SetStreamWhileRecording(value *bool)() {
    m.streamWhileRecording = value
}
