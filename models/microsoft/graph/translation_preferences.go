package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TranslationPreferences provides operations to manage the deviceManagement singleton.
type TranslationPreferences struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Translation override behavior for languages, if any.Returned by default.
    languageOverrides []TranslationLanguageOverrideable;
    // The user's preferred translation behavior.Returned by default. Not nullable.
    translationBehavior *TranslationBehavior;
    // The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
    untranslatedLanguages []string;
}
// NewTranslationPreferences instantiates a new translationPreferences and sets the default values.
func NewTranslationPreferences()(*TranslationPreferences) {
    m := &TranslationPreferences{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTranslationPreferencesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTranslationPreferencesFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTranslationPreferences(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationPreferences) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TranslationPreferences) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["languageOverrides"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTranslationLanguageOverrideFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TranslationLanguageOverrideable, len(val))
            for i, v := range val {
                res[i] = v.(TranslationLanguageOverrideable)
            }
            m.SetLanguageOverrides(res)
        }
        return nil
    }
    res["translationBehavior"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTranslationBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranslationBehavior(val.(*TranslationBehavior))
        }
        return nil
    }
    res["untranslatedLanguages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUntranslatedLanguages(res)
        }
        return nil
    }
    return res
}
// GetLanguageOverrides gets the languageOverrides property value. Translation override behavior for languages, if any.Returned by default.
func (m *TranslationPreferences) GetLanguageOverrides()([]TranslationLanguageOverrideable) {
    if m == nil {
        return nil
    } else {
        return m.languageOverrides
    }
}
// GetTranslationBehavior gets the translationBehavior property value. The user's preferred translation behavior.Returned by default. Not nullable.
func (m *TranslationPreferences) GetTranslationBehavior()(*TranslationBehavior) {
    if m == nil {
        return nil
    } else {
        return m.translationBehavior
    }
}
// GetUntranslatedLanguages gets the untranslatedLanguages property value. The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
func (m *TranslationPreferences) GetUntranslatedLanguages()([]string) {
    if m == nil {
        return nil
    } else {
        return m.untranslatedLanguages
    }
}
func (m *TranslationPreferences) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TranslationPreferences) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetLanguageOverrides() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLanguageOverrides()))
        for i, v := range m.GetLanguageOverrides() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("languageOverrides", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTranslationBehavior() != nil {
        cast := (*m.GetTranslationBehavior()).String()
        err := writer.WriteStringValue("translationBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUntranslatedLanguages() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationPreferences) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLanguageOverrides sets the languageOverrides property value. Translation override behavior for languages, if any.Returned by default.
func (m *TranslationPreferences) SetLanguageOverrides(value []TranslationLanguageOverrideable)() {
    if m != nil {
        m.languageOverrides = value
    }
}
// SetTranslationBehavior sets the translationBehavior property value. The user's preferred translation behavior.Returned by default. Not nullable.
func (m *TranslationPreferences) SetTranslationBehavior(value *TranslationBehavior)() {
    if m != nil {
        m.translationBehavior = value
    }
}
// SetUntranslatedLanguages sets the untranslatedLanguages property value. The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
func (m *TranslationPreferences) SetUntranslatedLanguages(value []string)() {
    if m != nil {
        m.untranslatedLanguages = value
    }
}
