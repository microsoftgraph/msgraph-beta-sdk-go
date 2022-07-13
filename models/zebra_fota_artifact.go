package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaArtifact 
type ZebraFotaArtifact struct {
    Entity
    // The version of the Board Support Package.
    boardSupportPackageVersion *string
    // Artifact device model.
    deviceModel *string
    // Artifact OS version.
    osVersion *string
    // Artifact patch version.
    patchVersion *string
    // Artifact release notes URL.
    releaseNotesUrl *string
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
// GetBoardSupportPackageVersion gets the boardSupportPackageVersion property value. The version of the Board Support Package.
func (m *ZebraFotaArtifact) GetBoardSupportPackageVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.boardSupportPackageVersion
    }
}
// GetDeviceModel gets the deviceModel property value. Artifact device model.
func (m *ZebraFotaArtifact) GetDeviceModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceModel
    }
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
// GetOsVersion gets the osVersion property value. Artifact OS version.
func (m *ZebraFotaArtifact) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetPatchVersion gets the patchVersion property value. Artifact patch version.
func (m *ZebraFotaArtifact) GetPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.patchVersion
    }
}
// GetReleaseNotesUrl gets the releaseNotesUrl property value. Artifact release notes URL.
func (m *ZebraFotaArtifact) GetReleaseNotesUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.releaseNotesUrl
    }
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
// SetBoardSupportPackageVersion sets the boardSupportPackageVersion property value. The version of the Board Support Package.
func (m *ZebraFotaArtifact) SetBoardSupportPackageVersion(value *string)() {
    if m != nil {
        m.boardSupportPackageVersion = value
    }
}
// SetDeviceModel sets the deviceModel property value. Artifact device model.
func (m *ZebraFotaArtifact) SetDeviceModel(value *string)() {
    if m != nil {
        m.deviceModel = value
    }
}
// SetOsVersion sets the osVersion property value. Artifact OS version.
func (m *ZebraFotaArtifact) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetPatchVersion sets the patchVersion property value. Artifact patch version.
func (m *ZebraFotaArtifact) SetPatchVersion(value *string)() {
    if m != nil {
        m.patchVersion = value
    }
}
// SetReleaseNotesUrl sets the releaseNotesUrl property value. Artifact release notes URL.
func (m *ZebraFotaArtifact) SetReleaseNotesUrl(value *string)() {
    if m != nil {
        m.releaseNotesUrl = value
    }
}
