package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DepMacOSEnrollmentProfile 
type DepMacOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
    // Indicates if Accessibility screen is disabled
    accessibilityScreenDisabled *bool;
    // Indicates if iCloud Documents and Desktop screen is disabled
    chooseYourLockScreenDisabled *bool;
    // Indicates if file vault is disabled
    fileVaultDisabled *bool;
    // Indicates if iCloud Analytics screen is disabled
    iCloudDiagnosticsDisabled *bool;
    // Indicates if iCloud Documents and Desktop screen is disabled
    iCloudStorageDisabled *bool;
    // Indicates if Passcode setup pane is disabled
    passCodeDisabled *bool;
    // Indicates if registration is disabled
    registrationDisabled *bool;
    // Indicates if zoom setup pane is disabled
    zoomDisabled *bool;
}
// NewDepMacOSEnrollmentProfile instantiates a new depMacOSEnrollmentProfile and sets the default values.
func NewDepMacOSEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    m := &DepMacOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    return m
}
// GetAccessibilityScreenDisabled gets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
func (m *DepMacOSEnrollmentProfile) GetAccessibilityScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accessibilityScreenDisabled
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
// GetFieldDeserializers the deserialization information for the current model
func (m *DepMacOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["accessibilityScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibilityScreenDisabled(val)
        }
        return nil
    }
    res["chooseYourLockScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChooseYourLockScreenDisabled(val)
        }
        return nil
    }
    res["fileVaultDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultDisabled(val)
        }
        return nil
    }
    res["iCloudDiagnosticsDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudDiagnosticsDisabled(val)
        }
        return nil
    }
    res["iCloudStorageDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudStorageDisabled(val)
        }
        return nil
    }
    res["passCodeDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassCodeDisabled(val)
        }
        return nil
    }
    res["registrationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDisabled(val)
        }
        return nil
    }
    res["zoomDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *DepMacOSEnrollmentProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DepMacOSEnrollmentProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
