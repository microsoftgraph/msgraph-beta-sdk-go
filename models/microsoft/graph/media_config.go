package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MediaConfig struct {
    additionalData map[string]interface{};
    removeFromDefaultAudioGroup *bool;
}
func NewMediaConfig()(*MediaConfig) {
    m := &MediaConfig{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MediaConfig) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MediaConfig) GetRemoveFromDefaultAudioGroup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.removeFromDefaultAudioGroup
    }
}
func (m *MediaConfig) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["removeFromDefaultAudioGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRemoveFromDefaultAudioGroup(val)
        return nil
    }
    return res
}
func (m *MediaConfig) IsNil()(bool) {
    return m == nil
}
func (m *MediaConfig) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("removeFromDefaultAudioGroup", m.GetRemoveFromDefaultAudioGroup())
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
func (m *MediaConfig) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MediaConfig) SetRemoveFromDefaultAudioGroup(value *bool)() {
    m.removeFromDefaultAudioGroup = value
}
