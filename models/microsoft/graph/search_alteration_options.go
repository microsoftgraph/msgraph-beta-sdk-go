package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchAlterationOptions struct {
    additionalData map[string]interface{};
    enableModification *bool;
    enableSuggestion *bool;
}
func NewSearchAlterationOptions()(*SearchAlterationOptions) {
    m := &SearchAlterationOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchAlterationOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchAlterationOptions) GetEnableModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableModification
    }
}
func (m *SearchAlterationOptions) GetEnableSuggestion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSuggestion
    }
}
func (m *SearchAlterationOptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enableModification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableModification(val)
        return nil
    }
    res["enableSuggestion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableSuggestion(val)
        return nil
    }
    return res
}
func (m *SearchAlterationOptions) IsNil()(bool) {
    return m == nil
}
func (m *SearchAlterationOptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enableModification", m.GetEnableModification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableSuggestion", m.GetEnableSuggestion())
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
func (m *SearchAlterationOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchAlterationOptions) SetEnableModification(value *bool)() {
    m.enableModification = value
}
func (m *SearchAlterationOptions) SetEnableSuggestion(value *bool)() {
    m.enableSuggestion = value
}
