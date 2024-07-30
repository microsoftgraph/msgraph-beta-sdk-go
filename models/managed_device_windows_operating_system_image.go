package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceWindowsOperatingSystemImage this entity defines different Windows Operating System products, like 'Windows 11 22H1', 'Windows 11 22H2' etc., along with their available configurations.
type ManagedDeviceWindowsOperatingSystemImage struct {
    Entity
}
// NewManagedDeviceWindowsOperatingSystemImage instantiates a new ManagedDeviceWindowsOperatingSystemImage and sets the default values.
func NewManagedDeviceWindowsOperatingSystemImage()(*ManagedDeviceWindowsOperatingSystemImage) {
    m := &ManagedDeviceWindowsOperatingSystemImage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceWindowsOperatingSystemImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedDeviceWindowsOperatingSystemImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceWindowsOperatingSystemImage(), nil
}
// GetAvailableUpdates gets the availableUpdates property value. Indicates the available Quality/Security updates for a specific Windows product version (example: Windows 11 22H1), for upto last 3 Patch Tuesdays . This value in the API response would be updated 2-3 days after every Patch Tuesday. Supports: $filter, $select, $top, $skip. Read-only.
// returns a []ManagedDeviceWindowsOperatingSystemUpdateable when successful
func (m *ManagedDeviceWindowsOperatingSystemImage) GetAvailableUpdates()([]ManagedDeviceWindowsOperatingSystemUpdateable) {
    val, err := m.GetBackingStore().Get("availableUpdates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceWindowsOperatingSystemUpdateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManagedDeviceWindowsOperatingSystemImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["availableUpdates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceWindowsOperatingSystemUpdateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceWindowsOperatingSystemUpdateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceWindowsOperatingSystemUpdateable)
                }
            }
            m.SetAvailableUpdates(res)
        }
        return nil
    }
    res["supportedArchitectures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedDeviceArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceArchitecture, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*ManagedDeviceArchitecture))
                }
            }
            m.SetSupportedArchitectures(res)
        }
        return nil
    }
    res["supportedEditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceWindowsOperatingSystemEditionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceWindowsOperatingSystemEditionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceWindowsOperatingSystemEditionable)
                }
            }
            m.SetSupportedEditions(res)
        }
        return nil
    }
    return res
}
// GetSupportedArchitectures gets the supportedArchitectures property value. Indicates the list of architectures supported by the image. E.g. ['ARM64','X86']. Supports: $filter, $select, $top, $skip. Read-only.
// returns a []ManagedDeviceArchitecture when successful
func (m *ManagedDeviceWindowsOperatingSystemImage) GetSupportedArchitectures()([]ManagedDeviceArchitecture) {
    val, err := m.GetBackingStore().Get("supportedArchitectures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceArchitecture)
    }
    return nil
}
// GetSupportedEditions gets the supportedEditions property value. Indicates the list of editions supported by the image along with their support dates. Supports: $filter, $select, $top, $skip. Read-only.
// returns a []ManagedDeviceWindowsOperatingSystemEditionable when successful
func (m *ManagedDeviceWindowsOperatingSystemImage) GetSupportedEditions()([]ManagedDeviceWindowsOperatingSystemEditionable) {
    val, err := m.GetBackingStore().Get("supportedEditions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceWindowsOperatingSystemEditionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceWindowsOperatingSystemImage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAvailableUpdates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAvailableUpdates()))
        for i, v := range m.GetAvailableUpdates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("availableUpdates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedArchitectures() != nil {
        err = writer.WriteCollectionOfStringValues("supportedArchitectures", SerializeManagedDeviceArchitecture(m.GetSupportedArchitectures()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedEditions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSupportedEditions()))
        for i, v := range m.GetSupportedEditions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("supportedEditions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailableUpdates sets the availableUpdates property value. Indicates the available Quality/Security updates for a specific Windows product version (example: Windows 11 22H1), for upto last 3 Patch Tuesdays . This value in the API response would be updated 2-3 days after every Patch Tuesday. Supports: $filter, $select, $top, $skip. Read-only.
func (m *ManagedDeviceWindowsOperatingSystemImage) SetAvailableUpdates(value []ManagedDeviceWindowsOperatingSystemUpdateable)() {
    err := m.GetBackingStore().Set("availableUpdates", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedArchitectures sets the supportedArchitectures property value. Indicates the list of architectures supported by the image. E.g. ['ARM64','X86']. Supports: $filter, $select, $top, $skip. Read-only.
func (m *ManagedDeviceWindowsOperatingSystemImage) SetSupportedArchitectures(value []ManagedDeviceArchitecture)() {
    err := m.GetBackingStore().Set("supportedArchitectures", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedEditions sets the supportedEditions property value. Indicates the list of editions supported by the image along with their support dates. Supports: $filter, $select, $top, $skip. Read-only.
func (m *ManagedDeviceWindowsOperatingSystemImage) SetSupportedEditions(value []ManagedDeviceWindowsOperatingSystemEditionable)() {
    err := m.GetBackingStore().Set("supportedEditions", value)
    if err != nil {
        panic(err)
    }
}
type ManagedDeviceWindowsOperatingSystemImageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableUpdates()([]ManagedDeviceWindowsOperatingSystemUpdateable)
    GetSupportedArchitectures()([]ManagedDeviceArchitecture)
    GetSupportedEditions()([]ManagedDeviceWindowsOperatingSystemEditionable)
    SetAvailableUpdates(value []ManagedDeviceWindowsOperatingSystemUpdateable)()
    SetSupportedArchitectures(value []ManagedDeviceArchitecture)()
    SetSupportedEditions(value []ManagedDeviceWindowsOperatingSystemEditionable)()
}
