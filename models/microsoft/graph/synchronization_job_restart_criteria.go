package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationJobRestartCriteria struct {
    additionalData map[string]interface{};
    resetScope *SynchronizationJobRestartScope;
}
func NewSynchronizationJobRestartCriteria()(*SynchronizationJobRestartCriteria) {
    m := &SynchronizationJobRestartCriteria{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationJobRestartCriteria) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationJobRestartCriteria) GetResetScope()(*SynchronizationJobRestartScope) {
    if m == nil {
        return nil
    } else {
        return m.resetScope
    }
}
func (m *SynchronizationJobRestartCriteria) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resetScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationJobRestartScope)
        if err != nil {
            return err
        }
        cast := val.(SynchronizationJobRestartScope)
        m.SetResetScope(&cast)
        return nil
    }
    return res
}
func (m *SynchronizationJobRestartCriteria) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationJobRestartCriteria) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetResetScope() != nil {
        cast := m.GetResetScope().String()
        err := writer.WriteStringValue("resetScope", &cast)
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
func (m *SynchronizationJobRestartCriteria) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationJobRestartCriteria) SetResetScope(value *SynchronizationJobRestartScope)() {
    m.resetScope = value
}
