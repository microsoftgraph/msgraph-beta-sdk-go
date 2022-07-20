package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepMacOSEnrollmentProfile 
type DepMacOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
    // Indicates if Accessibility screen is disabled
    accessibilityScreenDisabled *bool
    // Indicates if UnlockWithWatch screen is disabled
    autoUnlockWithWatchDisabled *bool
    // Indicates if iCloud Documents and Desktop screen is disabled
    chooseYourLockScreenDisabled *bool
    // Indicates if file vault is disabled
    fileVaultDisabled *bool
    // Indicates if iCloud Analytics screen is disabled
    iCloudDiagnosticsDisabled *bool
    // Indicates if iCloud Documents and Desktop screen is disabled
    iCloudStorageDisabled *bool
    // Indicates if Passcode setup pane is disabled
    passCodeDisabled *bool
    // Indicates if registration is disabled
    registrationDisabled *bool
    // Indicates if zoom setup pane is disabled
    zoomDisabled *bool
}
// NewDepMacOSEnrollmentProfile instantiates a new DepMacOSEnrollmentProfile and sets the default values.
func NewDepMacOSEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    m := &DepMacOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    odataTypeValue := "#microsoft.graph.depMacOSEnrollmentProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDepMacOSEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepMacOSEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepMacOSEnrollmentProfile(), nil
}
// GetAccessibilityScreenDisabled gets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
func (m *DepMacOSEnrollmentProfile) GetAccessibilityScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accessibilityScreenDisabled
    }
}
// GetAutoUnlockWithWatchDisabled gets the autoUnlockWithWatchDisabled property value. Indicates if UnlockWithWatch screen is disabled
func (m *DepMacOSEnrollmentProfile) GetAutoUnlockWithWatchDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoUnlockWithWatchDisabled
    }
}
// GetChooseYourLockScreenDisabled gets the chooseYourLockScreenDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) GetChooseYourLockScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.chooseYourLockScreenDisabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepMacOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["accessibilityScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibilityScreenDisabled(val)
        }
        return nil
    }
    res["autoUnlockWithWatchDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoUnlockWithWatchDisabled(val)
        }
        return nil
    }
    res["chooseYourLockScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChooseYourLockScreenDisabled(val)
        }
        return nil
    }
    res["fileVaultDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultDisabled(val)
        }
        return nil
    }
    res["iCloudDiagnosticsDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudDiagnosticsDisabled(val)
        }
        return nil
    }
    res["iCloudStorageDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudStorageDisabled(val)
        }
        return nil
    }
    res["passCodeDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassCodeDisabled(val)
        }
        return nil
    }
    res["registrationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDisabled(val)
        }
        return nil
    }
    res["zoomDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoomDisabled(val)
        }
        return nil
    }
    return res
}
// GetFileVaultDisabled gets the fileVaultDisabled property value. Indicates if file vault is disabled
func (m *DepMacOSEnrollmentProfile) GetFileVaultDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fileVaultDisabled
    }
}
// GetICloudDiagnosticsDisabled gets the iCloudDiagnosticsDisabled property value. Indicates if iCloud Analytics screen is disabled
func (m *DepMacOSEnrollmentProfile) GetICloudDiagnosticsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudDiagnosticsDisabled
    }
}
// GetICloudStorageDisabled gets the iCloudStorageDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) GetICloudStorageDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudStorageDisabled
    }
}
// GetPassCodeDisabled gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepMacOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passCodeDisabled
    }
}
// GetRegistrationDisabled gets the registrationDisabled property value. Indicates if registration is disabled
func (m *DepMacOSEnrollmentProfile) GetRegistrationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.registrationDisabled
    }
}
// GetZoomDisabled gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepMacOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.zoomDisabled
    }
}
// Serialize serializes information the current object
func (m *DepMacOSEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DepEnrollmentBaseProfile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accessibilityScreenDisabled", m.GetAccessibilityScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoUnlockWithWatchDisabled", m.GetAutoUnlockWithWatchDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("chooseYourLockScreenDisabled", m.GetChooseYourLockScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fileVaultDisabled", m.GetFileVaultDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudDiagnosticsDisabled", m.GetICloudDiagnosticsDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudStorageDisabled", m.GetICloudStorageDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passCodeDisabled", m.GetPassCodeDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("registrationDisabled", m.GetRegistrationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("zoomDisabled", m.GetZoomDisabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibilityScreenDisabled sets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
func (m *DepMacOSEnrollmentProfile) SetAccessibilityScreenDisabled(value *bool)() {
    if m != nil {
        m.accessibilityScreenDisabled = value
    }
}
// SetAutoUnlockWithWatchDisabled sets the autoUnlockWithWatchDisabled property value. Indicates if UnlockWithWatch screen is disabled
func (m *DepMacOSEnrollmentProfile) SetAutoUnlockWithWatchDisabled(value *bool)() {
    if m != nil {
        m.autoUnlockWithWatchDisabled = value
    }
}
// SetChooseYourLockScreenDisabled sets the chooseYourLockScreenDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) SetChooseYourLockScreenDisabled(value *bool)() {
    if m != nil {
        m.chooseYourLockScreenDisabled = value
    }
}
// SetFileVaultDisabled sets the fileVaultDisabled property value. Indicates if file vault is disabled
func (m *DepMacOSEnrollmentProfile) SetFileVaultDisabled(value *bool)() {
    if m != nil {
        m.fileVaultDisabled = value
    }
}
// SetICloudDiagnosticsDisabled sets the iCloudDiagnosticsDisabled property value. Indicates if iCloud Analytics screen is disabled
func (m *DepMacOSEnrollmentProfile) SetICloudDiagnosticsDisabled(value *bool)() {
    if m != nil {
        m.iCloudDiagnosticsDisabled = value
    }
}
// SetICloudStorageDisabled sets the iCloudStorageDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) SetICloudStorageDisabled(value *bool)() {
    if m != nil {
        m.iCloudStorageDisabled = value
    }
}
// SetPassCodeDisabled sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepMacOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    if m != nil {
        m.passCodeDisabled = value
    }
}
// SetRegistrationDisabled sets the registrationDisabled property value. Indicates if registration is disabled
func (m *DepMacOSEnrollmentProfile) SetRegistrationDisabled(value *bool)() {
    if m != nil {
        m.registrationDisabled = value
    }
}
// SetZoomDisabled sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepMacOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    if m != nil {
        m.zoomDisabled = value
    }
}
