package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationJobRestartCriteria 
type SynchronizationJobRestartCriteria struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Comma-separated combination of the following values: Full, QuarantineState, Watermark, Escrows, ConnectorDataStore. Use Full if you want all of the options.
    resetScope *SynchronizationJobRestartScope;
}
// NewSynchronizationJobRestartCriteria instantiates a new synchronizationJobRestartCriteria and sets the default values.
func NewSynchronizationJobRestartCriteria()(*SynchronizationJobRestartCriteria) {
    m := &SynchronizationJobRestartCriteria{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationJobRestartCriteria) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetResetScope gets the resetScope property value. Comma-separated combination of the following values: Full, QuarantineState, Watermark, Escrows, ConnectorDataStore. Use Full if you want all of the options.
func (m *SynchronizationJobRestartCriteria) GetResetScope()(*SynchronizationJobRestartScope) {
    if m == nil {
        return nil
    } else {
        return m.resetScope
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobRestartCriteria) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["resetScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationJobRestartScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResetScope(val.(*SynchronizationJobRestartScope))
        }
        return nil
    }
    return res
}
func (m *SynchronizationJobRestartCriteria) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SynchronizationJobRestartCriteria) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetResetScope() != nil {
        cast := (*m.GetResetScope()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationJobRestartCriteria) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResetScope sets the resetScope property value. Comma-separated combination of the following values: Full, QuarantineState, Watermark, Escrows, ConnectorDataStore. Use Full if you want all of the options.
func (m *SynchronizationJobRestartCriteria) SetResetScope(value *SynchronizationJobRestartScope)() {
    if m != nil {
        m.resetScope = value
    }
}
