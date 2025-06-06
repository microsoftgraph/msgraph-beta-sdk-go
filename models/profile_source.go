// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProfileSource struct {
    Entity
}
// NewProfileSource instantiates a new ProfileSource and sets the default values.
func NewProfileSource()(*ProfileSource) {
    m := &ProfileSource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProfileSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProfileSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfileSource(), nil
}
// GetDisplayName gets the displayName property value. Name of the profile source intended to inform users about the profile source name.
// returns a *string when successful
func (m *ProfileSource) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProfileSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["kind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKind(val)
        }
        return nil
    }
    res["localizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProfileSourceLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfileSourceLocalizationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProfileSourceLocalizationable)
                }
            }
            m.SetLocalizations(res)
        }
        return nil
    }
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetKind gets the kind property value. Type of the profile source.
// returns a *string when successful
func (m *ProfileSource) GetKind()(*string) {
    val, err := m.GetBackingStore().Get("kind")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalizations gets the localizations property value. Alternative localized labels specified by an administrator.
// returns a []ProfileSourceLocalizationable when successful
func (m *ProfileSource) GetLocalizations()([]ProfileSourceLocalizationable) {
    val, err := m.GetBackingStore().Get("localizations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ProfileSourceLocalizationable)
    }
    return nil
}
// GetSourceId gets the sourceId property value. Profile source identifier used as an alternate key.
// returns a *string when successful
func (m *ProfileSource) GetSourceId()(*string) {
    val, err := m.GetBackingStore().Get("sourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWebUrl gets the webUrl property value. Web URL of the profile source that directs users to the page view of profile data.
// returns a *string when successful
func (m *ProfileSource) GetWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("webUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProfileSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteStringValue("kind", m.GetKind())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("localizations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Name of the profile source intended to inform users about the profile source name.
func (m *ProfileSource) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetKind sets the kind property value. Type of the profile source.
func (m *ProfileSource) SetKind(value *string)() {
    err := m.GetBackingStore().Set("kind", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalizations sets the localizations property value. Alternative localized labels specified by an administrator.
func (m *ProfileSource) SetLocalizations(value []ProfileSourceLocalizationable)() {
    err := m.GetBackingStore().Set("localizations", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceId sets the sourceId property value. Profile source identifier used as an alternate key.
func (m *ProfileSource) SetSourceId(value *string)() {
    err := m.GetBackingStore().Set("sourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetWebUrl sets the webUrl property value. Web URL of the profile source that directs users to the page view of profile data.
func (m *ProfileSource) SetWebUrl(value *string)() {
    err := m.GetBackingStore().Set("webUrl", value)
    if err != nil {
        panic(err)
    }
}
type ProfileSourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetKind()(*string)
    GetLocalizations()([]ProfileSourceLocalizationable)
    GetSourceId()(*string)
    GetWebUrl()(*string)
    SetDisplayName(value *string)()
    SetKind(value *string)()
    SetLocalizations(value []ProfileSourceLocalizationable)()
    SetSourceId(value *string)()
    SetWebUrl(value *string)()
}
