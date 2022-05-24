package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonName casts the previous resource to user.
type PersonName struct {
    ItemFacet
    // Provides an ordered rendering of firstName and lastName depending on the locale of the user or their device.
    displayName *string
    // First name of the user.
    first *string
    // Initials of the user.
    initials *string
    // Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
    languageTag *string
    // Last name of the user.
    last *string
    // Maiden name of the user.
    maiden *string
    // Middle name of the user.
    middle *string
    // Nickname of the user.
    nickname *string
    // Guidance on how to pronounce the users name.
    pronunciation PersonNamePronounciationable
    // Designators used after the users name (eg: PhD.)
    suffix *string
    // Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
    title *string
}
// NewPersonName instantiates a new personName and sets the default values.
func NewPersonName()(*PersonName) {
    m := &PersonName{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreatePersonNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonName(), nil
}
// GetDisplayName gets the displayName property value. Provides an ordered rendering of firstName and lastName depending on the locale of the user or their device.
func (m *PersonName) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonName) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["first"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirst(val)
        }
        return nil
    }
    res["initials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitials(val)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["last"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLast(val)
        }
        return nil
    }
    res["maiden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaiden(val)
        }
        return nil
    }
    res["middle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddle(val)
        }
        return nil
    }
    res["nickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNickname(val)
        }
        return nil
    }
    res["pronunciation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePersonNamePronounciationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPronunciation(val.(PersonNamePronounciationable))
        }
        return nil
    }
    res["suffix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuffix(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetFirst gets the first property value. First name of the user.
func (m *PersonName) GetFirst()(*string) {
    if m == nil {
        return nil
    } else {
        return m.first
    }
}
// GetInitials gets the initials property value. Initials of the user.
func (m *PersonName) GetInitials()(*string) {
    if m == nil {
        return nil
    } else {
        return m.initials
    }
}
// GetLanguageTag gets the languageTag property value. Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
func (m *PersonName) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// GetLast gets the last property value. Last name of the user.
func (m *PersonName) GetLast()(*string) {
    if m == nil {
        return nil
    } else {
        return m.last
    }
}
// GetMaiden gets the maiden property value. Maiden name of the user.
func (m *PersonName) GetMaiden()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maiden
    }
}
// GetMiddle gets the middle property value. Middle name of the user.
func (m *PersonName) GetMiddle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middle
    }
}
// GetNickname gets the nickname property value. Nickname of the user.
func (m *PersonName) GetNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nickname
    }
}
// GetPronunciation gets the pronunciation property value. Guidance on how to pronounce the users name.
func (m *PersonName) GetPronunciation()(PersonNamePronounciationable) {
    if m == nil {
        return nil
    } else {
        return m.pronunciation
    }
}
// GetSuffix gets the suffix property value. Designators used after the users name (eg: PhD.)
func (m *PersonName) GetSuffix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.suffix
    }
}
// GetTitle gets the title property value. Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
func (m *PersonName) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Serialize serializes information the current object
func (m *PersonName) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetDisplayName sets the displayName property value. Provides an ordered rendering of firstName and lastName depending on the locale of the user or their device.
func (m *PersonName) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFirst sets the first property value. First name of the user.
func (m *PersonName) SetFirst(value *string)() {
    if m != nil {
        m.first = value
    }
}
// SetInitials sets the initials property value. Initials of the user.
func (m *PersonName) SetInitials(value *string)() {
    if m != nil {
        m.initials = value
    }
}
// SetLanguageTag sets the languageTag property value. Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
func (m *PersonName) SetLanguageTag(value *string)() {
    if m != nil {
        m.languageTag = value
    }
}
// SetLast sets the last property value. Last name of the user.
func (m *PersonName) SetLast(value *string)() {
    if m != nil {
        m.last = value
    }
}
// SetMaiden sets the maiden property value. Maiden name of the user.
func (m *PersonName) SetMaiden(value *string)() {
    if m != nil {
        m.maiden = value
    }
}
// SetMiddle sets the middle property value. Middle name of the user.
func (m *PersonName) SetMiddle(value *string)() {
    if m != nil {
        m.middle = value
    }
}
// SetNickname sets the nickname property value. Nickname of the user.
func (m *PersonName) SetNickname(value *string)() {
    if m != nil {
        m.nickname = value
    }
}
// SetPronunciation sets the pronunciation property value. Guidance on how to pronounce the users name.
func (m *PersonName) SetPronunciation(value PersonNamePronounciationable)() {
    if m != nil {
        m.pronunciation = value
    }
}
// SetSuffix sets the suffix property value. Designators used after the users name (eg: PhD.)
func (m *PersonName) SetSuffix(value *string)() {
    if m != nil {
        m.suffix = value
    }
}
// SetTitle sets the title property value. Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
func (m *PersonName) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
