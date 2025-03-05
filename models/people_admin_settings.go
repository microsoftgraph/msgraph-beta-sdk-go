package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PeopleAdminSettings struct {
    Entity
}
// NewPeopleAdminSettings instantiates a new PeopleAdminSettings and sets the default values.
func NewPeopleAdminSettings()(*PeopleAdminSettings) {
    m := &PeopleAdminSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePeopleAdminSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePeopleAdminSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPeopleAdminSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PeopleAdminSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["itemInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInsightsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemInsights(val.(InsightsSettingsable))
        }
        return nil
    }
    res["namePronunciation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNamePronunciationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamePronunciation(val.(NamePronunciationSettingsable))
        }
        return nil
    }
    res["profileCardProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProfileCardPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfileCardPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProfileCardPropertyable)
                }
            }
            m.SetProfileCardProperties(res)
        }
        return nil
    }
    res["pronouns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePronounsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPronouns(val.(PronounsSettingsable))
        }
        return nil
    }
    return res
}
// GetItemInsights gets the itemInsights property value. Administrator settings that manage the support for item insights in an organization.
// returns a InsightsSettingsable when successful
func (m *PeopleAdminSettings) GetItemInsights()(InsightsSettingsable) {
    val, err := m.GetBackingStore().Get("itemInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InsightsSettingsable)
    }
    return nil
}
// GetNamePronunciation gets the namePronunciation property value. Administrator settings that manage the support of name pronunciation in an organization.
// returns a NamePronunciationSettingsable when successful
func (m *PeopleAdminSettings) GetNamePronunciation()(NamePronunciationSettingsable) {
    val, err := m.GetBackingStore().Get("namePronunciation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(NamePronunciationSettingsable)
    }
    return nil
}
// GetProfileCardProperties gets the profileCardProperties property value. A collection of the properties an administrator defined as visible on the Microsoft 365 profile card.
// returns a []ProfileCardPropertyable when successful
func (m *PeopleAdminSettings) GetProfileCardProperties()([]ProfileCardPropertyable) {
    val, err := m.GetBackingStore().Get("profileCardProperties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ProfileCardPropertyable)
    }
    return nil
}
// GetPronouns gets the pronouns property value. Administrator settings that manage the support of pronouns in an organization.
// returns a PronounsSettingsable when successful
func (m *PeopleAdminSettings) GetPronouns()(PronounsSettingsable) {
    val, err := m.GetBackingStore().Get("pronouns")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PronounsSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PeopleAdminSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("itemInsights", m.GetItemInsights())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("namePronunciation", m.GetNamePronunciation())
        if err != nil {
            return err
        }
    }
    if m.GetProfileCardProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProfileCardProperties()))
        for i, v := range m.GetProfileCardProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("profileCardProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("pronouns", m.GetPronouns())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemInsights sets the itemInsights property value. Administrator settings that manage the support for item insights in an organization.
func (m *PeopleAdminSettings) SetItemInsights(value InsightsSettingsable)() {
    err := m.GetBackingStore().Set("itemInsights", value)
    if err != nil {
        panic(err)
    }
}
// SetNamePronunciation sets the namePronunciation property value. Administrator settings that manage the support of name pronunciation in an organization.
func (m *PeopleAdminSettings) SetNamePronunciation(value NamePronunciationSettingsable)() {
    err := m.GetBackingStore().Set("namePronunciation", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileCardProperties sets the profileCardProperties property value. A collection of the properties an administrator defined as visible on the Microsoft 365 profile card.
func (m *PeopleAdminSettings) SetProfileCardProperties(value []ProfileCardPropertyable)() {
    err := m.GetBackingStore().Set("profileCardProperties", value)
    if err != nil {
        panic(err)
    }
}
// SetPronouns sets the pronouns property value. Administrator settings that manage the support of pronouns in an organization.
func (m *PeopleAdminSettings) SetPronouns(value PronounsSettingsable)() {
    err := m.GetBackingStore().Set("pronouns", value)
    if err != nil {
        panic(err)
    }
}
type PeopleAdminSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemInsights()(InsightsSettingsable)
    GetNamePronunciation()(NamePronunciationSettingsable)
    GetProfileCardProperties()([]ProfileCardPropertyable)
    GetPronouns()(PronounsSettingsable)
    SetItemInsights(value InsightsSettingsable)()
    SetNamePronunciation(value NamePronunciationSettingsable)()
    SetProfileCardProperties(value []ProfileCardPropertyable)()
    SetPronouns(value PronounsSettingsable)()
}
