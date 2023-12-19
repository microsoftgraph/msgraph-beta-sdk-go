package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LanguageProficiency 
type LanguageProficiency struct {
    ItemFacet
}
// NewLanguageProficiency instantiates a new languageProficiency and sets the default values.
func NewLanguageProficiency()(*LanguageProficiency) {
    m := &LanguageProficiency{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.languageProficiency"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateLanguageProficiencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLanguageProficiencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLanguageProficiency(), nil
}
// GetDisplayName gets the displayName property value. Contains the long-form name for the language.
func (m *LanguageProficiency) GetDisplayName()(*string) {
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
func (m *LanguageProficiency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["proficiency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiency_proficiency)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProficiency(val.(*LanguageProficiency_proficiency))
        }
        return nil
    }
    res["reading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiency_reading)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReading(val.(*LanguageProficiency_reading))
        }
        return nil
    }
    res["spoken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiency_spoken)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpoken(val.(*LanguageProficiency_spoken))
        }
        return nil
    }
    res["tag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["written"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiency_written)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWritten(val.(*LanguageProficiency_written))
        }
        return nil
    }
    return res
}
// GetProficiency gets the proficiency property value. The proficiency property
func (m *LanguageProficiency) GetProficiency()(*LanguageProficiency_proficiency) {
    val, err := m.GetBackingStore().Get("proficiency")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LanguageProficiency_proficiency)
    }
    return nil
}
// GetReading gets the reading property value. Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetReading()(*LanguageProficiency_reading) {
    val, err := m.GetBackingStore().Get("reading")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LanguageProficiency_reading)
    }
    return nil
}
// GetSpoken gets the spoken property value. Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetSpoken()(*LanguageProficiency_spoken) {
    val, err := m.GetBackingStore().Get("spoken")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LanguageProficiency_spoken)
    }
    return nil
}
// GetTag gets the tag property value. Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
func (m *LanguageProficiency) GetTag()(*string) {
    val, err := m.GetBackingStore().Get("tag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *LanguageProficiency) GetThumbnailUrl()(*string) {
    val, err := m.GetBackingStore().Get("thumbnailUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWritten gets the written property value. Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) GetWritten()(*LanguageProficiency_written) {
    val, err := m.GetBackingStore().Get("written")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LanguageProficiency_written)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LanguageProficiency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetProficiency sets the proficiency property value. The proficiency property
func (m *LanguageProficiency) SetProficiency(value *LanguageProficiency_proficiency)() {
    err := m.GetBackingStore().Set("proficiency", value)
    if err != nil {
        panic(err)
    }
}
// SetReading sets the reading property value. Represents the users reading comprehension for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetReading(value *LanguageProficiency_reading)() {
    err := m.GetBackingStore().Set("reading", value)
    if err != nil {
        panic(err)
    }
}
// SetSpoken sets the spoken property value. Represents the users spoken proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetSpoken(value *LanguageProficiency_spoken)() {
    err := m.GetBackingStore().Set("spoken", value)
    if err != nil {
        panic(err)
    }
}
// SetTag sets the tag property value. Contains the four-character BCP47 name for the language (en-US, no-NB, en-AU).
func (m *LanguageProficiency) SetTag(value *string)() {
    err := m.GetBackingStore().Set("tag", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *LanguageProficiency) SetThumbnailUrl(value *string)() {
    err := m.GetBackingStore().Set("thumbnailUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetWritten sets the written property value. Represents the users written proficiency for the language represented by the object. Possible values are: elementary, conversational, limitedWorking, professionalWorking, fullProfessional, nativeOrBilingual, unknownFutureValue.
func (m *LanguageProficiency) SetWritten(value *LanguageProficiency_written)() {
    err := m.GetBackingStore().Set("written", value)
    if err != nil {
        panic(err)
    }
}
// LanguageProficiencyable 
type LanguageProficiencyable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetProficiency()(*LanguageProficiency_proficiency)
    GetReading()(*LanguageProficiency_reading)
    GetSpoken()(*LanguageProficiency_spoken)
    GetTag()(*string)
    GetThumbnailUrl()(*string)
    GetWritten()(*LanguageProficiency_written)
    SetDisplayName(value *string)()
    SetProficiency(value *LanguageProficiency_proficiency)()
    SetReading(value *LanguageProficiency_reading)()
    SetSpoken(value *LanguageProficiency_spoken)()
    SetTag(value *string)()
    SetThumbnailUrl(value *string)()
    SetWritten(value *LanguageProficiency_written)()
}
