package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DlpWindowsDevicesNotification 
type DlpWindowsDevicesNotification struct {
    DlpNotification
}
// NewDlpWindowsDevicesNotification instantiates a new dlpWindowsDevicesNotification and sets the default values.
func NewDlpWindowsDevicesNotification()(*DlpWindowsDevicesNotification) {
    m := &DlpWindowsDevicesNotification{
        DlpNotification: *NewDlpNotification(),
    }
    odataTypeValue := "#microsoft.graph.dlpWindowsDevicesNotification"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDlpWindowsDevicesNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpWindowsDevicesNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDlpWindowsDevicesNotification(), nil
}
// GetContentName gets the contentName property value. The contentName property
func (m *DlpWindowsDevicesNotification) GetContentName()(*string) {
    val, err := m.GetBackingStore().Get("contentName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpWindowsDevicesNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DlpNotification.GetFieldDeserializers()
    res["contentName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentName(val)
        }
        return nil
    }
    res["lastModfiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModfiedBy(val)
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
    return res
}
// GetLastModfiedBy gets the lastModfiedBy property value. The lastModfiedBy property
func (m *DlpWindowsDevicesNotification) GetLastModfiedBy()(*string) {
    val, err := m.GetBackingStore().Get("lastModfiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DlpWindowsDevicesNotification) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DlpWindowsDevicesNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DlpNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentName", m.GetContentName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastModfiedBy", m.GetLastModfiedBy())
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
    return nil
}
// SetContentName sets the contentName property value. The contentName property
func (m *DlpWindowsDevicesNotification) SetContentName(value *string)() {
    err := m.GetBackingStore().Set("contentName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModfiedBy sets the lastModfiedBy property value. The lastModfiedBy property
func (m *DlpWindowsDevicesNotification) SetLastModfiedBy(value *string)() {
    err := m.GetBackingStore().Set("lastModfiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DlpWindowsDevicesNotification) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DlpWindowsDevicesNotificationable 
type DlpWindowsDevicesNotificationable interface {
    DlpNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentName()(*string)
    GetLastModfiedBy()(*string)
    GetOdataType()(*string)
    SetContentName(value *string)()
    SetLastModfiedBy(value *string)()
    SetOdataType(value *string)()
}
