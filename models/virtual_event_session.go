package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEventSession struct {
    OnlineMeetingBase
}
// NewVirtualEventSession instantiates a new VirtualEventSession and sets the default values.
func NewVirtualEventSession()(*VirtualEventSession) {
    m := &VirtualEventSession{
        OnlineMeetingBase: *NewOnlineMeetingBase(),
    }
    odataTypeValue := "#microsoft.graph.virtualEventSession"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateVirtualEventSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventSession(), nil
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
// returns a DateTimeTimeZoneable when successful
func (m *VirtualEventSession) GetEndDateTime()(DateTimeTimeZoneable) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DateTimeTimeZoneable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEventSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnlineMeetingBase.GetFieldDeserializers()
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["presenters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventPresenterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventPresenterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventPresenterable)
                }
            }
            m.SetPresenters(res)
        }
        return nil
    }
    res["registrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventRegistrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventRegistrationable)
                }
            }
            m.SetRegistrations(res)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetPresenters gets the presenters property value. The presenters property
// returns a []VirtualEventPresenterable when successful
func (m *VirtualEventSession) GetPresenters()([]VirtualEventPresenterable) {
    val, err := m.GetBackingStore().Get("presenters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventPresenterable)
    }
    return nil
}
// GetRegistrations gets the registrations property value. The registrations property
// returns a []VirtualEventRegistrationable when successful
func (m *VirtualEventSession) GetRegistrations()([]VirtualEventRegistrationable) {
    val, err := m.GetBackingStore().Get("registrations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventRegistrationable)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
// returns a DateTimeTimeZoneable when successful
func (m *VirtualEventSession) GetStartDateTime()(DateTimeTimeZoneable) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DateTimeTimeZoneable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEventSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnlineMeetingBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPresenters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPresenters()))
        for i, v := range m.GetPresenters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("presenters", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegistrations()))
        for i, v := range m.GetRegistrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("registrations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *VirtualEventSession) SetEndDateTime(value DateTimeTimeZoneable)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPresenters sets the presenters property value. The presenters property
func (m *VirtualEventSession) SetPresenters(value []VirtualEventPresenterable)() {
    err := m.GetBackingStore().Set("presenters", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrations sets the registrations property value. The registrations property
func (m *VirtualEventSession) SetRegistrations(value []VirtualEventRegistrationable)() {
    err := m.GetBackingStore().Set("registrations", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *VirtualEventSession) SetStartDateTime(value DateTimeTimeZoneable)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
type VirtualEventSessionable interface {
    OnlineMeetingBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDateTime()(DateTimeTimeZoneable)
    GetPresenters()([]VirtualEventPresenterable)
    GetRegistrations()([]VirtualEventRegistrationable)
    GetStartDateTime()(DateTimeTimeZoneable)
    SetEndDateTime(value DateTimeTimeZoneable)()
    SetPresenters(value []VirtualEventPresenterable)()
    SetRegistrations(value []VirtualEventRegistrationable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
}
