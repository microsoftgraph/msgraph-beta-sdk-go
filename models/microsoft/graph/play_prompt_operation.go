package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlayPromptOperation struct {
    CommsOperation
    completionReason *PlayPromptCompletionReason;
}
func NewPlayPromptOperation()(*PlayPromptOperation) {
    m := &PlayPromptOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
func (m *PlayPromptOperation) GetCompletionReason()(*PlayPromptCompletionReason) {
    if m == nil {
        return nil
    } else {
        return m.completionReason
    }
}
func (m *PlayPromptOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    res["completionReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlayPromptCompletionReason)
        if err != nil {
            return err
        }
        cast := val.(PlayPromptCompletionReason)
        m.SetCompletionReason(&cast)
        return nil
    }
    return res
}
func (m *PlayPromptOperation) IsNil()(bool) {
    return m == nil
}
func (m *PlayPromptOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCompletionReason() != nil {
        cast := m.GetCompletionReason().String()
        err = writer.WriteStringValue("completionReason", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PlayPromptOperation) SetCompletionReason(value *PlayPromptCompletionReason)() {
    m.completionReason = value
}
