package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PersonName struct {
    ItemFacet
    // Provides an ordered rendering of firstName and lastName depending on the locale of the user or their device.
    displayName *string;
    // First name of the user.
    first *string;
    // Initials of the user.
    initials *string;
    // Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
    languageTag *string;
    // Last name of the user.
    last *string;
    // Maiden name of the user.
    maiden *string;
    // Middle name of the user.
    middle *string;
    // Nickname of the user.
    nickname *string;
    // Guidance on how to pronounce the users name.
    pronunciation *PersonNamePronounciation;
    // Designators used after the users name (eg: PhD.)
    suffix *string;
    // Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
    title *string;
}
// Instantiates a new personName and sets the default values.
func NewPersonName()(*PersonName) {
    m := &PersonName{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the displayName property value. Provides an ordered rendering of firstName and lastName depending on the locale of the user or their device.
func (m *PersonName) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the first property value. First name of the user.
func (m *PersonName) GetFirst()(*string) {
    if m == nil {
        return nil
    } else {
        return m.first
    }
}
// Gets the initials property value. Initials of the user.
func (m *PersonName) GetInitials()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initials
    }
}
// Gets the languageTag property value. Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
func (m *PersonName) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// Gets the last property value. Last name of the user.
func (m *PersonName) GetLast()(*string) {
    if m == nil {
        return nil
    } else {
        return m.last
    }
}
// Gets the maiden property value. Maiden name of the user.
func (m *PersonName) GetMaiden()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maiden
    }
}
// Gets the middle property value. Middle name of the user.
func (m *PersonName) GetMiddle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middle
    }
}
// Gets the nickname property value. Nickname of the user.
func (m *PersonName) GetNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nickname
    }
}
// Gets the pronunciation property value. Guidance on how to pronounce the users name.
func (m *PersonName) GetPronunciation()(*PersonNamePronounciation) {
    if m == nil {
        return nil
    } else {
        return m.pronunciation
    }
}
// Gets the suffix property value. Designators used after the users name (eg: PhD.)
func (m *PersonName) GetSuffix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.suffix
    }
}
// Gets the title property value. Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
func (m *PersonName) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// The deserialization information for the current model
func (m *PersonName) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["first"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFirst(val)
        return nil
    }
    res["initials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInitials(val)
        return nil
    }
    res["languageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLanguageTag(val)
        return nil
    }
    res["last"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLast(val)
        return nil
    }
    res["maiden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaiden(val)
        return nil
    }
    res["middle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMiddle(val)
        return nil
    }
    res["nickname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNickname(val)
        return nil
    }
    res["pronunciation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonNamePronounciation() })
        if err != nil {
            return err
        }
        m.SetPronunciation(val.(*PersonNamePronounciation))
        return nil
    }
    res["suffix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSuffix(val)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    return res
}
func (m *PersonName) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PersonName) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("first", m.GetFirst())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initials", m.GetInitials())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("last", m.GetLast())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maiden", m.GetMaiden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("middle", m.GetMiddle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickname", m.GetNickname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("pronunciation", m.GetPronunciation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("suffix", m.GetSuffix())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. Provides an ordered rendering of firstName and lastName depending on the locale of the user or their device.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PersonName) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the first property value. First name of the user.
// Parameters:
//  - value : Value to set for the first property.
func (m *PersonName) SetFirst(value *string)() {
    m.first = value
}
// Sets the initials property value. Initials of the user.
// Parameters:
//  - value : Value to set for the initials property.
func (m *PersonName) SetInitials(value *string)() {
    m.initials = value
}
// Sets the languageTag property value. Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
// Parameters:
//  - value : Value to set for the languageTag property.
func (m *PersonName) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// Sets the last property value. Last name of the user.
// Parameters:
//  - value : Value to set for the last property.
func (m *PersonName) SetLast(value *string)() {
    m.last = value
}
// Sets the maiden property value. Maiden name of the user.
// Parameters:
//  - value : Value to set for the maiden property.
func (m *PersonName) SetMaiden(value *string)() {
    m.maiden = value
}
// Sets the middle property value. Middle name of the user.
// Parameters:
//  - value : Value to set for the middle property.
func (m *PersonName) SetMiddle(value *string)() {
    m.middle = value
}
// Sets the nickname property value. Nickname of the user.
// Parameters:
//  - value : Value to set for the nickname property.
func (m *PersonName) SetNickname(value *string)() {
    m.nickname = value
}
// Sets the pronunciation property value. Guidance on how to pronounce the users name.
// Parameters:
//  - value : Value to set for the pronunciation property.
func (m *PersonName) SetPronunciation(value *PersonNamePronounciation)() {
    m.pronunciation = value
}
// Sets the suffix property value. Designators used after the users name (eg: PhD.)
// Parameters:
//  - value : Value to set for the suffix property.
func (m *PersonName) SetSuffix(value *string)() {
    m.suffix = value
}
// Sets the title property value. Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
// Parameters:
//  - value : Value to set for the title property.
func (m *PersonName) SetTitle(value *string)() {
    m.title = value
}
