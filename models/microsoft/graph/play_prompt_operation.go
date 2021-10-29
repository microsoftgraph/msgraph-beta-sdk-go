package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlayPromptOperation struct {
    CommsOperation
    // Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
    completionReason *PlayPromptCompletionReason;
}
// Instantiates a new playPromptOperation and sets the default values.
func NewPlayPromptOperation()(*PlayPromptOperation) {
    m := &PlayPromptOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// Gets the completionReason property value. Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
func (m *PlayPromptOperation) GetCompletionReason()(*PlayPromptCompletionReason) {
    if m == nil {
        return nil
    } else {
        return m.completionReason
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the completionReason property value. Possible values are: unknown, completedSuccessfully, mediaOperationCanceled.
// Parameters:
//  - value : Value to set for the completionReason property.
func (m *PlayPromptOperation) SetCompletionReason(value *PlayPromptCompletionReason)() {
    m.completionReason = value
}
