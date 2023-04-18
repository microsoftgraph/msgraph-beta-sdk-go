package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Command 
type Command struct {
    Entity
}
// NewCommand instantiates a new command and sets the default values.
func NewCommand()(*Command) {
    m := &Command{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCommandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommandFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommand(), nil
}
// GetAppServiceName gets the appServiceName property value. The appServiceName property
func (m *Command) GetAppServiceName()(*string) {
    val, err := m.GetBackingStore().Get("appServiceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetError gets the error property value. The error property
func (m *Command) GetError()(*string) {
    val, err := m.GetBackingStore().Get("error")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Command) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appServiceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppServiceName(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["packageFamilyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageFamilyName(val)
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(PayloadRequestable))
        }
        return nil
    }
    res["permissionTicket"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionTicket(val)
        }
        return nil
    }
    res["postBackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostBackUri(val)
        }
        return nil
    }
    res["responsepayload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsepayload(val.(PayloadResponseable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetPackageFamilyName gets the packageFamilyName property value. The packageFamilyName property
func (m *Command) GetPackageFamilyName()(*string) {
    val, err := m.GetBackingStore().Get("packageFamilyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPayload gets the payload property value. The payload property
func (m *Command) GetPayload()(PayloadRequestable) {
    val, err := m.GetBackingStore().Get("payload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PayloadRequestable)
    }
    return nil
}
// GetPermissionTicket gets the permissionTicket property value. The permissionTicket property
func (m *Command) GetPermissionTicket()(*string) {
    val, err := m.GetBackingStore().Get("permissionTicket")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPostBackUri gets the postBackUri property value. The postBackUri property
func (m *Command) GetPostBackUri()(*string) {
    val, err := m.GetBackingStore().Get("postBackUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResponsepayload gets the responsepayload property value. The responsepayload property
func (m *Command) GetResponsepayload()(PayloadResponseable) {
    val, err := m.GetBackingStore().Get("responsepayload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PayloadResponseable)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *Command) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetType gets the type property value. The type property
func (m *Command) GetType()(*string) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Command) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appServiceName", m.GetAppServiceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageFamilyName", m.GetPackageFamilyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionTicket", m.GetPermissionTicket())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postBackUri", m.GetPostBackUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("responsepayload", m.GetResponsepayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppServiceName sets the appServiceName property value. The appServiceName property
func (m *Command) SetAppServiceName(value *string)() {
    err := m.GetBackingStore().Set("appServiceName", value)
    if err != nil {
        panic(err)
    }
}
// SetError sets the error property value. The error property
func (m *Command) SetError(value *string)() {
    err := m.GetBackingStore().Set("error", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageFamilyName sets the packageFamilyName property value. The packageFamilyName property
func (m *Command) SetPackageFamilyName(value *string)() {
    err := m.GetBackingStore().Set("packageFamilyName", value)
    if err != nil {
        panic(err)
    }
}
// SetPayload sets the payload property value. The payload property
func (m *Command) SetPayload(value PayloadRequestable)() {
    err := m.GetBackingStore().Set("payload", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionTicket sets the permissionTicket property value. The permissionTicket property
func (m *Command) SetPermissionTicket(value *string)() {
    err := m.GetBackingStore().Set("permissionTicket", value)
    if err != nil {
        panic(err)
    }
}
// SetPostBackUri sets the postBackUri property value. The postBackUri property
func (m *Command) SetPostBackUri(value *string)() {
    err := m.GetBackingStore().Set("postBackUri", value)
    if err != nil {
        panic(err)
    }
}
// SetResponsepayload sets the responsepayload property value. The responsepayload property
func (m *Command) SetResponsepayload(value PayloadResponseable)() {
    err := m.GetBackingStore().Set("responsepayload", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *Command) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetType sets the type property value. The type property
func (m *Command) SetType(value *string)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// Commandable 
type Commandable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppServiceName()(*string)
    GetError()(*string)
    GetPackageFamilyName()(*string)
    GetPayload()(PayloadRequestable)
    GetPermissionTicket()(*string)
    GetPostBackUri()(*string)
    GetResponsepayload()(PayloadResponseable)
    GetStatus()(*string)
    GetType()(*string)
    SetAppServiceName(value *string)()
    SetError(value *string)()
    SetPackageFamilyName(value *string)()
    SetPayload(value PayloadRequestable)()
    SetPermissionTicket(value *string)()
    SetPostBackUri(value *string)()
    SetResponsepayload(value PayloadResponseable)()
    SetStatus(value *string)()
    SetType(value *string)()
}
