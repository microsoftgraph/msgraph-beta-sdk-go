package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PeopleAdminSettings 
type PeopleAdminSettings struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewPeopleAdminSettings instantiates a new peopleAdminSettings and sets the default values.
func NewPeopleAdminSettings()(*PeopleAdminSettings) {
    m := &PeopleAdminSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePeopleAdminSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePeopleAdminSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPeopleAdminSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PeopleAdminSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
// GetPronouns gets the pronouns property value. Represents administrator settings that manage the support of pronouns in an organization.
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
        err = writer.WriteObjectValue("pronouns", m.GetPronouns())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPronouns sets the pronouns property value. Represents administrator settings that manage the support of pronouns in an organization.
func (m *PeopleAdminSettings) SetPronouns(value PronounsSettingsable)() {
    err := m.GetBackingStore().Set("pronouns", value)
    if err != nil {
        panic(err)
    }
}
// PeopleAdminSettingsable 
type PeopleAdminSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPronouns()(PronounsSettingsable)
    SetPronouns(value PronounsSettingsable)()
}
