package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonName 
type PersonName struct {
    ItemFacet
    // The OdataType property
    OdataType *string
}
// NewPersonName instantiates a new personName and sets the default values.
func NewPersonName()(*PersonName) {
    m := &PersonName{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personName"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePersonNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonName(), nil
}
// GetDisplayName gets the displayName property value. Provides an ordered rendering of firstName and lastName depending on the locale of the user or their device.
func (m *PersonName) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    val, err := m.GetBackingStore().Get("first")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInitials gets the initials property value. Initials of the user.
func (m *PersonName) GetInitials()(*string) {
    val, err := m.GetBackingStore().Get("initials")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLanguageTag gets the languageTag property value. Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
func (m *PersonName) GetLanguageTag()(*string) {
    val, err := m.GetBackingStore().Get("languageTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLast gets the last property value. Last name of the user.
func (m *PersonName) GetLast()(*string) {
    val, err := m.GetBackingStore().Get("last")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMaiden gets the maiden property value. Maiden name of the user.
func (m *PersonName) GetMaiden()(*string) {
    val, err := m.GetBackingStore().Get("maiden")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMiddle gets the middle property value. Middle name of the user.
func (m *PersonName) GetMiddle()(*string) {
    val, err := m.GetBackingStore().Get("middle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNickname gets the nickname property value. Nickname of the user.
func (m *PersonName) GetNickname()(*string) {
    val, err := m.GetBackingStore().Get("nickname")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPronunciation gets the pronunciation property value. Guidance on how to pronounce the users name.
func (m *PersonName) GetPronunciation()(PersonNamePronounciationable) {
    val, err := m.GetBackingStore().Get("pronunciation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PersonNamePronounciationable)
    }
    return nil
}
// GetSuffix gets the suffix property value. Designators used after the users name (eg: PhD.)
func (m *PersonName) GetSuffix()(*string) {
    val, err := m.GetBackingStore().Get("suffix")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTitle gets the title property value. Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
func (m *PersonName) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFirst sets the first property value. First name of the user.
func (m *PersonName) SetFirst(value *string)() {
    err := m.GetBackingStore().Set("first", value)
    if err != nil {
        panic(err)
    }
}
// SetInitials sets the initials property value. Initials of the user.
func (m *PersonName) SetInitials(value *string)() {
    err := m.GetBackingStore().Set("initials", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguageTag sets the languageTag property value. Contains the name for the language (en-US, no-NB, en-AU) following IETF BCP47 format.
func (m *PersonName) SetLanguageTag(value *string)() {
    err := m.GetBackingStore().Set("languageTag", value)
    if err != nil {
        panic(err)
    }
}
// SetLast sets the last property value. Last name of the user.
func (m *PersonName) SetLast(value *string)() {
    err := m.GetBackingStore().Set("last", value)
    if err != nil {
        panic(err)
    }
}
// SetMaiden sets the maiden property value. Maiden name of the user.
func (m *PersonName) SetMaiden(value *string)() {
    err := m.GetBackingStore().Set("maiden", value)
    if err != nil {
        panic(err)
    }
}
// SetMiddle sets the middle property value. Middle name of the user.
func (m *PersonName) SetMiddle(value *string)() {
    err := m.GetBackingStore().Set("middle", value)
    if err != nil {
        panic(err)
    }
}
// SetNickname sets the nickname property value. Nickname of the user.
func (m *PersonName) SetNickname(value *string)() {
    err := m.GetBackingStore().Set("nickname", value)
    if err != nil {
        panic(err)
    }
}
// SetPronunciation sets the pronunciation property value. Guidance on how to pronounce the users name.
func (m *PersonName) SetPronunciation(value PersonNamePronounciationable)() {
    err := m.GetBackingStore().Set("pronunciation", value)
    if err != nil {
        panic(err)
    }
}
// SetSuffix sets the suffix property value. Designators used after the users name (eg: PhD.)
func (m *PersonName) SetSuffix(value *string)() {
    err := m.GetBackingStore().Set("suffix", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. Honorifics used to prefix a users name (eg: Dr, Sir, Madam, Mrs.)
func (m *PersonName) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
// PersonNameable 
type PersonNameable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetFirst()(*string)
    GetInitials()(*string)
    GetLanguageTag()(*string)
    GetLast()(*string)
    GetMaiden()(*string)
    GetMiddle()(*string)
    GetNickname()(*string)
    GetPronunciation()(PersonNamePronounciationable)
    GetSuffix()(*string)
    GetTitle()(*string)
    SetDisplayName(value *string)()
    SetFirst(value *string)()
    SetInitials(value *string)()
    SetLanguageTag(value *string)()
    SetLast(value *string)()
    SetMaiden(value *string)()
    SetMiddle(value *string)()
    SetNickname(value *string)()
    SetPronunciation(value PersonNamePronounciationable)()
    SetSuffix(value *string)()
    SetTitle(value *string)()
}
