package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new depMacOSEnrollmentProfile and sets the default values.
func NewDepMacOSEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    m := &DepMacOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    return m
}
// Gets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
func (m *DepMacOSEnrollmentProfile) GetAccessibilityScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accessibilityScreenDisabled
    }
}
// Gets the chooseYourLockScreenDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) GetChooseYourLockScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.chooseYourLockScreenDisabled
    }
}
// Gets the fileVaultDisabled property value. Indicates if file vault is disabled
func (m *DepMacOSEnrollmentProfile) GetFileVaultDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fileVaultDisabled
    }
}
// Gets the iCloudDiagnosticsDisabled property value. Indicates if iCloud Analytics screen is disabled
func (m *DepMacOSEnrollmentProfile) GetICloudDiagnosticsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudDiagnosticsDisabled
    }
}
// Gets the iCloudStorageDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) GetICloudStorageDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudStorageDisabled
    }
}
// Gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepMacOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passCodeDisabled
    }
}
// Gets the registrationDisabled property value. Indicates if registration is disabled
func (m *DepMacOSEnrollmentProfile) GetRegistrationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.registrationDisabled
    }
}
// Gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepMacOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.zoomDisabled
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
// Parameters:
//  - value : Value to set for the accessibilityScreenDisabled property.
func (m *DepMacOSEnrollmentProfile) SetAccessibilityScreenDisabled(value *bool)() {
    m.accessibilityScreenDisabled = value
}
// Sets the chooseYourLockScreenDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
// Parameters:
//  - value : Value to set for the chooseYourLockScreenDisabled property.
func (m *DepMacOSEnrollmentProfile) SetChooseYourLockScreenDisabled(value *bool)() {
    m.chooseYourLockScreenDisabled = value
}
// Sets the fileVaultDisabled property value. Indicates if file vault is disabled
// Parameters:
//  - value : Value to set for the fileVaultDisabled property.
func (m *DepMacOSEnrollmentProfile) SetFileVaultDisabled(value *bool)() {
    m.fileVaultDisabled = value
}
// Sets the iCloudDiagnosticsDisabled property value. Indicates if iCloud Analytics screen is disabled
// Parameters:
//  - value : Value to set for the iCloudDiagnosticsDisabled property.
func (m *DepMacOSEnrollmentProfile) SetICloudDiagnosticsDisabled(value *bool)() {
    m.iCloudDiagnosticsDisabled = value
}
// Sets the iCloudStorageDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
// Parameters:
//  - value : Value to set for the iCloudStorageDisabled property.
func (m *DepMacOSEnrollmentProfile) SetICloudStorageDisabled(value *bool)() {
    m.iCloudStorageDisabled = value
}
// Sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
// Parameters:
//  - value : Value to set for the passCodeDisabled property.
func (m *DepMacOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    m.passCodeDisabled = value
}
// Sets the registrationDisabled property value. Indicates if registration is disabled
// Parameters:
//  - value : Value to set for the registrationDisabled property.
func (m *DepMacOSEnrollmentProfile) SetRegistrationDisabled(value *bool)() {
    m.registrationDisabled = value
}
// Sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
// Parameters:
//  - value : Value to set for the zoomDisabled property.
func (m *DepMacOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    m.zoomDisabled = value
}
