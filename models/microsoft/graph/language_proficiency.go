package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LanguageProficiency 
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
// NewLanguageProficiency instantiates a new languageProficiency and sets the default values.
func NewLanguageProficiency()(*LanguageProficiency) {
    m := &LanguageProficiency{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Contains the long-form name for the language.
func (m *LanguageProficiency) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetProficiency gets the proficiency property value. 
func (m *LanguageProficiency) GetProficiency()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.proficiency
    }
}
// GetReading gets the reading property value. Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetReading()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.reading
    }
}
// GetSpoken gets the spoken property value. Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetSpoken()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.spoken
    }
}
// GetTag gets the tag property value. Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
func (m *LanguageProficiency) GetTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tag
    }
}
// GetThumbnailUrl gets the thumbnailUrl property value. 
func (m *LanguageProficiency) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// GetWritten gets the written property value. Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetWritten()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.written
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetProficiency(val.(*LanguageProficiencyLevel))
        }
        return nil
    }
    res["reading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReading(val.(*LanguageProficiencyLevel))
        }
        return nil
    }
    res["spoken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpoken(val.(*LanguageProficiencyLevel))
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
            m.SetWritten(val.(*LanguageProficiencyLevel))
        }
        return nil
    }
    return res
}
func (m *LanguageProficiency) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetProficiency()).String()
        err = writer.WriteStringValue("proficiency", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReading() != nil {
        cast := (*m.GetReading()).String()
        err = writer.WriteStringValue("reading", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSpoken() != nil {
        cast := (*m.GetSpoken()).String()
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
        cast := (*m.GetWritten()).String()
        err = writer.WriteStringValue("written", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Contains the long-form name for the language.
func (m *LanguageProficiency) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetProficiency sets the proficiency property value. 
func (m *LanguageProficiency) SetProficiency(value *LanguageProficiencyLevel)() {
    if m != nil {
        m.proficiency = value
    }
}
// SetReading sets the reading property value. Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetReading(value *LanguageProficiencyLevel)() {
    if m != nil {
        m.reading = value
    }
}
// SetSpoken sets the spoken property value. Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetSpoken(value *LanguageProficiencyLevel)() {
    if m != nil {
        m.spoken = value
    }
}
// SetTag sets the tag property value. Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
func (m *LanguageProficiency) SetTag(value *string)() {
    if m != nil {
        m.tag = value
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. 
func (m *LanguageProficiency) SetThumbnailUrl(value *string)() {
    if m != nil {
        m.thumbnailUrl = value
    }
}
// SetWritten sets the written property value. Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetWritten(value *LanguageProficiencyLevel)() {
    if m != nil {
        m.written = value
    }
}
