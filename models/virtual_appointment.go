package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualAppointment 
type VirtualAppointment struct {
    Entity
}
// NewVirtualAppointment instantiates a new virtualAppointment and sets the default values.
func NewVirtualAppointment()(*VirtualAppointment) {
    m := &VirtualAppointment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualAppointmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualAppointmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualAppointment(), nil
}
// GetAppointmentClientJoinWebUrl gets the appointmentClientJoinWebUrl property value. The join web URL of the virtual appointment for clients with waiting room and browser join. Optional.
func (m *VirtualAppointment) GetAppointmentClientJoinWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("appointmentClientJoinWebUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppointmentClients gets the appointmentClients property value. The client information for the virtual appointment, including name, email, and SMS phone number. Optional.
func (m *VirtualAppointment) GetAppointmentClients()([]VirtualAppointmentUserable) {
    val, err := m.GetBackingStore().Get("appointmentClients")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualAppointmentUserable)
    }
    return nil
}
// GetExternalAppointmentId gets the externalAppointmentId property value. The identifier of the appointment from the scheduling system, associated with the current virtual appointment. Optional.
func (m *VirtualAppointment) GetExternalAppointmentId()(*string) {
    val, err := m.GetBackingStore().Get("externalAppointmentId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExternalAppointmentUrl gets the externalAppointmentUrl property value. The URL of the appointment resource from the scheduling system, associated with the current virtual appointment. Optional.
func (m *VirtualAppointment) GetExternalAppointmentUrl()(*string) {
    val, err := m.GetBackingStore().Get("externalAppointmentUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualAppointment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appointmentClientJoinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppointmentClientJoinWebUrl(val)
        }
        return nil
    }
    res["appointmentClients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualAppointmentUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualAppointmentUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualAppointmentUserable)
                }
            }
            m.SetAppointmentClients(res)
        }
        return nil
    }
    res["externalAppointmentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAppointmentId(val)
        }
        return nil
    }
    res["externalAppointmentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAppointmentUrl(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualAppointmentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(VirtualAppointmentSettingsable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VirtualAppointment) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettings gets the settings property value. The settings associated with the virtual appointment resource. Optional.
func (m *VirtualAppointment) GetSettings()(VirtualAppointmentSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VirtualAppointmentSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualAppointment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appointmentClientJoinWebUrl", m.GetAppointmentClientJoinWebUrl())
        if err != nil {
            return err
        }
    }
    if m.GetAppointmentClients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppointmentClients()))
        for i, v := range m.GetAppointmentClients() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appointmentClients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalAppointmentId", m.GetExternalAppointmentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalAppointmentUrl", m.GetExternalAppointmentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppointmentClientJoinWebUrl sets the appointmentClientJoinWebUrl property value. The join web URL of the virtual appointment for clients with waiting room and browser join. Optional.
func (m *VirtualAppointment) SetAppointmentClientJoinWebUrl(value *string)() {
    err := m.GetBackingStore().Set("appointmentClientJoinWebUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetAppointmentClients sets the appointmentClients property value. The client information for the virtual appointment, including name, email, and SMS phone number. Optional.
func (m *VirtualAppointment) SetAppointmentClients(value []VirtualAppointmentUserable)() {
    err := m.GetBackingStore().Set("appointmentClients", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalAppointmentId sets the externalAppointmentId property value. The identifier of the appointment from the scheduling system, associated with the current virtual appointment. Optional.
func (m *VirtualAppointment) SetExternalAppointmentId(value *string)() {
    err := m.GetBackingStore().Set("externalAppointmentId", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalAppointmentUrl sets the externalAppointmentUrl property value. The URL of the appointment resource from the scheduling system, associated with the current virtual appointment. Optional.
func (m *VirtualAppointment) SetExternalAppointmentUrl(value *string)() {
    err := m.GetBackingStore().Set("externalAppointmentUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VirtualAppointment) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. The settings associated with the virtual appointment resource. Optional.
func (m *VirtualAppointment) SetSettings(value VirtualAppointmentSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// VirtualAppointmentable 
type VirtualAppointmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppointmentClientJoinWebUrl()(*string)
    GetAppointmentClients()([]VirtualAppointmentUserable)
    GetExternalAppointmentId()(*string)
    GetExternalAppointmentUrl()(*string)
    GetOdataType()(*string)
    GetSettings()(VirtualAppointmentSettingsable)
    SetAppointmentClientJoinWebUrl(value *string)()
    SetAppointmentClients(value []VirtualAppointmentUserable)()
    SetExternalAppointmentId(value *string)()
    SetExternalAppointmentUrl(value *string)()
    SetOdataType(value *string)()
    SetSettings(value VirtualAppointmentSettingsable)()
}
