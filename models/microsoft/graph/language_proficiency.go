package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LanguageProficiency struct {
    ItemFacet
    // Contains the long-form name for the language.
    displayName *string;
    // 
    proficiency *LanguageProficiencyLevel;
    // Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
    reading *LanguageProficiencyLevel;
    // Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
    spoken *LanguageProficiencyLevel;
    // Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
    tag *string;
    // 
    thumbnailUrl *string;
    // Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
    written *LanguageProficiencyLevel;
}
// Instantiates a new languageProficiency and sets the default values.
func NewLanguageProficiency()(*LanguageProficiency) {
    m := &LanguageProficiency{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the displayName property value. Contains the long-form name for the language.
func (m *LanguageProficiency) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the proficiency property value. 
func (m *LanguageProficiency) GetProficiency()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.proficiency
    }
}
// Gets the reading property value. Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetReading()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.reading
    }
}
// Gets the spoken property value. Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetSpoken()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.spoken
    }
}
// Gets the tag property value. Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
func (m *LanguageProficiency) GetTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tag
    }
}
// Gets the thumbnailUrl property value. 
func (m *LanguageProficiency) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// Gets the written property value. Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetWritten()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.written
    }
}
// The deserialization information for the current model
func (m *LanguageProficiency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["proficiency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(LanguageProficiencyLevel)
            m.SetProficiency(&cast)
        }
        return nil
    }
    res["reading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(LanguageProficiencyLevel)
            m.SetReading(&cast)
        }
        return nil
    }
    res["spoken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(LanguageProficiencyLevel)
            m.SetSpoken(&cast)
        }
        return nil
    }
    res["tag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["written"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(LanguageProficiencyLevel)
            m.SetWritten(&cast)
        }
        return nil
    }
    return res
}
func (m *LanguageProficiency) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LanguageProficiency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetProficiency() != nil {
        cast := m.GetProficiency().String()
        err = writer.WriteStringValue("proficiency", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReading() != nil {
        cast := m.GetReading().String()
        err = writer.WriteStringValue("reading", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSpoken() != nil {
        cast := m.GetSpoken().String()
        err = writer.WriteStringValue("spoken", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tag", m.GetTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    if m.GetWritten() != nil {
        cast := m.GetWritten().String()
        err = writer.WriteStringValue("written", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. Contains the long-form name for the language.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *LanguageProficiency) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the proficiency property value. 
// Parameters:
//  - value : Value to set for the proficiency property.
func (m *LanguageProficiency) SetProficiency(value *LanguageProficiencyLevel)() {
    m.proficiency = value
}
// Sets the reading property value. Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
// Parameters:
//  - value : Value to set for the reading property.
func (m *LanguageProficiency) SetReading(value *LanguageProficiencyLevel)() {
    m.reading = value
}
// Sets the spoken property value. Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
// Parameters:
//  - value : Value to set for the spoken property.
func (m *LanguageProficiency) SetSpoken(value *LanguageProficiencyLevel)() {
    m.spoken = value
}
// Sets the tag property value. Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
// Parameters:
//  - value : Value to set for the tag property.
func (m *LanguageProficiency) SetTag(value *string)() {
    m.tag = value
}
// Sets the thumbnailUrl property value. 
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *LanguageProficiency) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// Sets the written property value. Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
// Parameters:
//  - value : Value to set for the written property.
func (m *LanguageProficiency) SetWritten(value *LanguageProficiencyLevel)() {
    m.written = value
}
