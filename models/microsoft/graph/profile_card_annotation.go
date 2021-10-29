package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ProfileCardAnnotation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If present, the value of this field is used by the profile card as the default property label in the experience (for example, 'Cost Center').
    displayName *string;
    // Each resource in this collection represents the localized value of the attribute name for a given language, used as the default label for that locale. For example, a user with a no-NB client gets 'Kostnads Senter' as the attribute label, rather than 'Cost Center.'
    localizations []DisplayNameLocalization;
}
// Instantiates a new profileCardAnnotation and sets the default values.
func NewProfileCardAnnotation()(*ProfileCardAnnotation) {
    m := &ProfileCardAnnotation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProfileCardAnnotation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. If present, the value of this field is used by the profile card as the default property label in the experience (for example, 'Cost Center').
func (m *ProfileCardAnnotation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the localizations property value. Each resource in this collection represents the localized value of the attribute name for a given language, used as the default label for that locale. For example, a user with a no-NB client gets 'Kostnads Senter' as the attribute label, rather than 'Cost Center.'
func (m *ProfileCardAnnotation) GetLocalizations()([]DisplayNameLocalization) {
    if m == nil {
        return nil
    } else {
        return m.localizations
    }
}
// The deserialization information for the current model
func (m *ProfileCardAnnotation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["localizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDisplayNameLocalization() })
        if err != nil {
            return err
        }
        res := make([]DisplayNameLocalization, len(val))
        for i, v := range val {
            res[i] = *(v.(*DisplayNameLocalization))
        }
        m.SetLocalizations(res)
        return nil
    }
    return res
}
func (m *ProfileCardAnnotation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ProfileCardAnnotation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("localizations", cast)
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
func (m *ProfileCardAnnotation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. If present, the value of this field is used by the profile card as the default property label in the experience (for example, 'Cost Center').
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ProfileCardAnnotation) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the localizations property value. Each resource in this collection represents the localized value of the attribute name for a given language, used as the default label for that locale. For example, a user with a no-NB client gets 'Kostnads Senter' as the attribute label, rather than 'Cost Center.'
// Parameters:
//  - value : Value to set for the localizations property.
func (m *ProfileCardAnnotation) SetLocalizations(value []DisplayNameLocalization)() {
    m.localizations = value
}
