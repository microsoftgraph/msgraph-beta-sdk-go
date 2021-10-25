package muteall

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MuteAllRequestBody struct {
    additionalData map[string]interface{};
    clientContext *string;
    participants []string;
}
func NewMuteAllRequestBody()(*MuteAllRequestBody) {
    m := &MuteAllRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MuteAllRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MuteAllRequestBody) GetClientContext()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientContext
    }
}
func (m *MuteAllRequestBody) GetParticipants()([]string) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
func (m *MuteAllRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clientContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientContext(val)
        return nil
    }
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetParticipants(res)
        return nil
    }
    return res
}
func (m *MuteAllRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *MuteAllRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
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
func (m *MuteAllRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MuteAllRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
func (m *MuteAllRequestBody) SetParticipants(value []string)() {
    m.participants = value
}
