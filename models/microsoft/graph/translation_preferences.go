package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TranslationPreferences struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Translation override behavior for languages, if any.Returned by default.
    languageOverrides []TranslationLanguageOverride;
    // The user's preferred translation behavior.Returned by default. Not nullable.
    translationBehavior *TranslationBehavior;
    // The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
    untranslatedLanguages []string;
}
// Instantiates a new translationPreferences and sets the default values.
func NewTranslationPreferences()(*TranslationPreferences) {
    m := &TranslationPreferences{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationPreferences) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the languageOverrides property value. Translation override behavior for languages, if any.Returned by default.
func (m *TranslationPreferences) GetLanguageOverrides()([]TranslationLanguageOverride) {
    if m == nil {
        return nil
    } else {
        return m.languageOverrides
    }
}
// Gets the translationBehavior property value. The user's preferred translation behavior.Returned by default. Not nullable.
func (m *TranslationPreferences) GetTranslationBehavior()(*TranslationBehavior) {
    if m == nil {
        return nil
    } else {
        return m.translationBehavior
    }
}
// Gets the untranslatedLanguages property value. The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
func (m *TranslationPreferences) GetUntranslatedLanguages()([]string) {
    if m == nil {
        return nil
    } else {
        return m.untranslatedLanguages
    }
}
// The deserialization information for the current model
func (m *TranslationPreferences) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["languageOverrides"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTranslationLanguageOverride() })
        if err != nil {
            return err
        }
        res := make([]TranslationLanguageOverride, len(val))
        for i, v := range val {
            res[i] = *(v.(*TranslationLanguageOverride))
        }
        m.SetLanguageOverrides(res)
        return nil
    }
    res["translationBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTranslationBehavior)
        if err != nil {
            return err
        }
        cast := val.(TranslationBehavior)
        m.SetTranslationBehavior(&cast)
        return nil
    }
    res["untranslatedLanguages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetUntranslatedLanguages(res)
        return nil
    }
    return res
}
func (m *TranslationPreferences) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TranslationPreferences) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLanguageOverrides()))
        for i, v := range m.GetLanguageOverrides() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("languageOverrides", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTranslationBehavior() != nil {
        cast := m.GetTranslationBehavior().String()
        err := writer.WriteStringValue("translationBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("untranslatedLanguages", m.GetUntranslatedLanguages())
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
func (m *TranslationPreferences) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the languageOverrides property value. Translation override behavior for languages, if any.Returned by default.
// Parameters:
//  - value : Value to set for the languageOverrides property.
func (m *TranslationPreferences) SetLanguageOverrides(value []TranslationLanguageOverride)() {
    m.languageOverrides = value
}
// Sets the translationBehavior property value. The user's preferred translation behavior.Returned by default. Not nullable.
// Parameters:
//  - value : Value to set for the translationBehavior property.
func (m *TranslationPreferences) SetTranslationBehavior(value *TranslationBehavior)() {
    m.translationBehavior = value
}
// Sets the untranslatedLanguages property value. The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
// Parameters:
//  - value : Value to set for the untranslatedLanguages property.
func (m *TranslationPreferences) SetUntranslatedLanguages(value []string)() {
    m.untranslatedLanguages = value
}
