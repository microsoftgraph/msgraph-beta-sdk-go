package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SearchAlterationOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether spelling modifications are enabled. If enabled, user will get the search results for corrected query when there are no results for the original query with typos and get the spelling modification information in queryAlterationResponse property of the response. Optional.
    enableModification *bool;
    // Indicates whether spelling suggestions are enabled. If enabled, user will get the search results for original search query and suggesting spelling correction in queryAlterationResponse property of the response for typos in query. Optional.
    enableSuggestion *bool;
}
// Instantiates a new searchAlterationOptions and sets the default values.
func NewSearchAlterationOptions()(*SearchAlterationOptions) {
    m := &SearchAlterationOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAlterationOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the enableModification property value. Indicates whether spelling modifications are enabled. If enabled, user will get the search results for corrected query when there are no results for the original query with typos and get the spelling modification information in queryAlterationResponse property of the response. Optional.
func (m *SearchAlterationOptions) GetEnableModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableModification
    }
}
// Gets the enableSuggestion property value. Indicates whether spelling suggestions are enabled. If enabled, user will get the search results for original search query and suggesting spelling correction in queryAlterationResponse property of the response for typos in query. Optional.
func (m *SearchAlterationOptions) GetEnableSuggestion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSuggestion
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SearchAlterationOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the enableModification property value. Indicates whether spelling modifications are enabled. If enabled, user will get the search results for corrected query when there are no results for the original query with typos and get the spelling modification information in queryAlterationResponse property of the response. Optional.
// Parameters:
//  - value : Value to set for the enableModification property.
func (m *SearchAlterationOptions) SetEnableModification(value *bool)() {
    m.enableModification = value
}
// Sets the enableSuggestion property value. Indicates whether spelling suggestions are enabled. If enabled, user will get the search results for original search query and suggesting spelling correction in queryAlterationResponse property of the response for typos in query. Optional.
// Parameters:
//  - value : Value to set for the enableSuggestion property.
func (m *SearchAlterationOptions) SetEnableSuggestion(value *bool)() {
    m.enableSuggestion = value
}
