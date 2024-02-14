package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TranslationPreferences struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTranslationPreferences instantiates a new TranslationPreferences and sets the default values.
func NewTranslationPreferences()(*TranslationPreferences) {
    m := &TranslationPreferences{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTranslationPreferencesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTranslationPreferencesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTranslationPreferences(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TranslationPreferences) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *TranslationPreferences) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TranslationPreferences) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["languageOverrides"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTranslationLanguageOverrideFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TranslationLanguageOverrideable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TranslationLanguageOverrideable)
                }
            }
            m.SetLanguageOverrides(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["translationBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTranslationBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranslationBehavior(val.(*TranslationBehavior))
        }
        return nil
    }
    res["untranslatedLanguages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetUntranslatedLanguages(res)
        }
        return nil
    }
    return res
}
// GetLanguageOverrides gets the languageOverrides property value. Translation override behavior for languages, if any.Returned by default.
// returns a []TranslationLanguageOverrideable when successful
func (m *TranslationPreferences) GetLanguageOverrides()([]TranslationLanguageOverrideable) {
    val, err := m.GetBackingStore().Get("languageOverrides")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TranslationLanguageOverrideable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TranslationPreferences) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTranslationBehavior gets the translationBehavior property value. The user's preferred translation behavior.Returned by default. Not nullable.
// returns a *TranslationBehavior when successful
func (m *TranslationPreferences) GetTranslationBehavior()(*TranslationBehavior) {
    val, err := m.GetBackingStore().Get("translationBehavior")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TranslationBehavior)
    }
    return nil
}
// GetUntranslatedLanguages gets the untranslatedLanguages property value. The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
// returns a []string when successful
func (m *TranslationPreferences) GetUntranslatedLanguages()([]string) {
    val, err := m.GetBackingStore().Get("untranslatedLanguages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TranslationPreferences) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLanguageOverrides() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLanguageOverrides()))
        for i, v := range m.GetLanguageOverrides() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("languageOverrides", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TranslationPreferences) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TranslationPreferences) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetLanguageOverrides sets the languageOverrides property value. Translation override behavior for languages, if any.Returned by default.
func (m *TranslationPreferences) SetLanguageOverrides(value []TranslationLanguageOverrideable)() {
    err := m.GetBackingStore().Set("languageOverrides", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TranslationPreferences) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTranslationBehavior sets the translationBehavior property value. The user's preferred translation behavior.Returned by default. Not nullable.
func (m *TranslationPreferences) SetTranslationBehavior(value *TranslationBehavior)() {
    err := m.GetBackingStore().Set("translationBehavior", value)
    if err != nil {
        panic(err)
    }
}
// SetUntranslatedLanguages sets the untranslatedLanguages property value. The list of languages the user does not need translated. This is computed from the authoringLanguages collection in regionalAndLanguageSettings, and the languageOverrides collection in translationPreferences. The list specifies neutral culture values that include the language code without any country or region association. For example, it would specify 'fr' for the neutral French culture, but not 'fr-FR' for the French culture in France. Returned by default. Read only.
func (m *TranslationPreferences) SetUntranslatedLanguages(value []string)() {
    err := m.GetBackingStore().Set("untranslatedLanguages", value)
    if err != nil {
        panic(err)
    }
}
type TranslationPreferencesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetLanguageOverrides()([]TranslationLanguageOverrideable)
    GetOdataType()(*string)
    GetTranslationBehavior()(*TranslationBehavior)
    GetUntranslatedLanguages()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetLanguageOverrides(value []TranslationLanguageOverrideable)()
    SetOdataType(value *string)()
    SetTranslationBehavior(value *TranslationBehavior)()
    SetUntranslatedLanguages(value []string)()
}
