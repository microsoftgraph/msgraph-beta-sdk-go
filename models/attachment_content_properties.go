package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttachmentContentProperties 
type AttachmentContentProperties struct {
    ContentProperties
}
// NewAttachmentContentProperties instantiates a new attachmentContentProperties and sets the default values.
func NewAttachmentContentProperties()(*AttachmentContentProperties) {
    m := &AttachmentContentProperties{
        ContentProperties: *NewContentProperties(),
    }
    odataTypeValue := "#microsoft.graph.attachmentContentProperties"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAttachmentContentPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttachmentContentPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttachmentContentProperties(), nil
}
// GetCurrentLabel gets the currentLabel property value. The currentLabel property
func (m *AttachmentContentProperties) GetCurrentLabel()(CurrentLabelable) {
    val, err := m.GetBackingStore().Get("currentLabel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CurrentLabelable)
    }
    return nil
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *AttachmentContentProperties) GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveTypeable) {
    val, err := m.GetBackingStore().Get("discoveredSensitiveTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DiscoveredSensitiveTypeable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttachmentContentProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ContentProperties.GetFieldDeserializers()
    res["currentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCurrentLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(CurrentLabelable))
        }
        return nil
    }
    res["discoveredSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiscoveredSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiscoveredSensitiveTypeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DiscoveredSensitiveTypeable)
                }
            }
            m.SetDiscoveredSensitiveTypes(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AttachmentContentProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ContentProperties.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    if m.GetDiscoveredSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDiscoveredSensitiveTypes()))
        for i, v := range m.GetDiscoveredSensitiveTypes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCurrentLabel sets the currentLabel property value. The currentLabel property
func (m *AttachmentContentProperties) SetCurrentLabel(value CurrentLabelable)() {
    err := m.GetBackingStore().Set("currentLabel", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *AttachmentContentProperties) SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveTypeable)() {
    err := m.GetBackingStore().Set("discoveredSensitiveTypes", value)
    if err != nil {
        panic(err)
    }
}
// AttachmentContentPropertiesable 
type AttachmentContentPropertiesable interface {
    ContentPropertiesable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentLabel()(CurrentLabelable)
    GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveTypeable)
    SetCurrentLabel(value CurrentLabelable)()
    SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveTypeable)()
}
