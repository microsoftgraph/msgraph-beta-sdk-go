package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaArtifact 
type ZebraFotaArtifact struct {
    Entity
}
// NewZebraFotaArtifact instantiates a new ZebraFotaArtifact and sets the default values.
func NewZebraFotaArtifact()(*ZebraFotaArtifact) {
    m := &ZebraFotaArtifact{
        Entity: *NewEntity(),
    }
    return m
}
// CreateZebraFotaArtifactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaArtifactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaArtifact(), nil
}
// GetBoardSupportPackageVersion gets the boardSupportPackageVersion property value. The version of the Board Support Package (BSP. E.g.: 01.18.02.00)
func (m *ZebraFotaArtifact) GetBoardSupportPackageVersion()(*string) {
    val, err := m.GetBackingStore().Get("boardSupportPackageVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Artifact description. (e.g.: `LifeGuard Update 98 (released 24-September-2021)
func (m *ZebraFotaArtifact) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceModel gets the deviceModel property value. Applicable device model (e.g.: TC8300)
func (m *ZebraFotaArtifact) GetDeviceModel()(*string) {
    val, err := m.GetBackingStore().Get("deviceModel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ZebraFotaArtifact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["boardSupportPackageVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBoardSupportPackageVersion(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["patchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatchVersion(val)
        }
        return nil
    }
    res["releaseNotesUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseNotesUrl(val)
        }
        return nil
    }
    return res
}
// GetOsVersion gets the osVersion property value. Artifact OS version (e.g.: 8.1.0)
func (m *ZebraFotaArtifact) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPatchVersion gets the patchVersion property value. Artifact patch version (e.g.: U00)
func (m *ZebraFotaArtifact) GetPatchVersion()(*string) {
    val, err := m.GetBackingStore().Get("patchVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReleaseNotesUrl gets the releaseNotesUrl property value. Artifact release notes URL (e.g.: https://www.zebra.com/<filename.pdf>)
func (m *ZebraFotaArtifact) GetReleaseNotesUrl()(*string) {
    val, err := m.GetBackingStore().Get("releaseNotesUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ZebraFotaArtifact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("boardSupportPackageVersion", m.GetBoardSupportPackageVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("patchVersion", m.GetPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("releaseNotesUrl", m.GetReleaseNotesUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBoardSupportPackageVersion sets the boardSupportPackageVersion property value. The version of the Board Support Package (BSP. E.g.: 01.18.02.00)
func (m *ZebraFotaArtifact) SetBoardSupportPackageVersion(value *string)() {
    err := m.GetBackingStore().Set("boardSupportPackageVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Artifact description. (e.g.: `LifeGuard Update 98 (released 24-September-2021)
func (m *ZebraFotaArtifact) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceModel sets the deviceModel property value. Applicable device model (e.g.: TC8300)
func (m *ZebraFotaArtifact) SetDeviceModel(value *string)() {
    err := m.GetBackingStore().Set("deviceModel", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. Artifact OS version (e.g.: 8.1.0)
func (m *ZebraFotaArtifact) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPatchVersion sets the patchVersion property value. Artifact patch version (e.g.: U00)
func (m *ZebraFotaArtifact) SetPatchVersion(value *string)() {
    err := m.GetBackingStore().Set("patchVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetReleaseNotesUrl sets the releaseNotesUrl property value. Artifact release notes URL (e.g.: https://www.zebra.com/<filename.pdf>)
func (m *ZebraFotaArtifact) SetReleaseNotesUrl(value *string)() {
    err := m.GetBackingStore().Set("releaseNotesUrl", value)
    if err != nil {
        panic(err)
    }
}
// ZebraFotaArtifactable 
type ZebraFotaArtifactable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBoardSupportPackageVersion()(*string)
    GetDescription()(*string)
    GetDeviceModel()(*string)
    GetOsVersion()(*string)
    GetPatchVersion()(*string)
    GetReleaseNotesUrl()(*string)
    SetBoardSupportPackageVersion(value *string)()
    SetDescription(value *string)()
    SetDeviceModel(value *string)()
    SetOsVersion(value *string)()
    SetPatchVersion(value *string)()
    SetReleaseNotesUrl(value *string)()
}
