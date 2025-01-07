package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Template struct {
    Entity
}
// NewTemplate instantiates a new Template and sets the default values.
func NewTemplate()(*Template) {
    m := &Template{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemplate(), nil
}
// GetDeviceTemplates gets the deviceTemplates property value. Defines the templates that are common to a set of device objects, such as IoT devices.
// returns a []DeviceTemplateable when successful
func (m *Template) GetDeviceTemplates()([]DeviceTemplateable) {
    val, err := m.GetBackingStore().Get("deviceTemplates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceTemplateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Template) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceTemplateable)
                }
            }
            m.SetDeviceTemplates(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Template) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeviceTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceTemplates()))
        for i, v := range m.GetDeviceTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceTemplates sets the deviceTemplates property value. Defines the templates that are common to a set of device objects, such as IoT devices.
func (m *Template) SetDeviceTemplates(value []DeviceTemplateable)() {
    err := m.GetBackingStore().Set("deviceTemplates", value)
    if err != nil {
        panic(err)
    }
}
type Templateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceTemplates()([]DeviceTemplateable)
    SetDeviceTemplates(value []DeviceTemplateable)()
}
