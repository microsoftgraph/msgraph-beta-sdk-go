package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DepMacOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
    accessibilityScreenDisabled *bool;
    chooseYourLockScreenDisabled *bool;
    fileVaultDisabled *bool;
    iCloudDiagnosticsDisabled *bool;
    iCloudStorageDisabled *bool;
    passCodeDisabled *bool;
    registrationDisabled *bool;
    zoomDisabled *bool;
}
func NewDepMacOSEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    m := &DepMacOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    return m
}
func (m *DepMacOSEnrollmentProfile) GetAccessibilityScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accessibilityScreenDisabled
    }
}
func (m *DepMacOSEnrollmentProfile) GetChooseYourLockScreenDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.chooseYourLockScreenDisabled
    }
}
func (m *DepMacOSEnrollmentProfile) GetFileVaultDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fileVaultDisabled
    }
}
func (m *DepMacOSEnrollmentProfile) GetICloudDiagnosticsDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudDiagnosticsDisabled
    }
}
func (m *DepMacOSEnrollmentProfile) GetICloudStorageDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudStorageDisabled
    }
}
func (m *DepMacOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passCodeDisabled
    }
}
func (m *DepMacOSEnrollmentProfile) GetRegistrationDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.registrationDisabled
    }
}
func (m *DepMacOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.zoomDisabled
    }
}
func (m *DepMacOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["accessibilityScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAccessibilityScreenDisabled(val)
        return nil
    }
    res["chooseYourLockScreenDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetChooseYourLockScreenDisabled(val)
        return nil
    }
    res["fileVaultDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFileVaultDisabled(val)
        return nil
    }
    res["iCloudDiagnosticsDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetICloudDiagnosticsDisabled(val)
        return nil
    }
    res["iCloudStorageDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetICloudStorageDisabled(val)
        return nil
    }
    res["passCodeDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPassCodeDisabled(val)
        return nil
    }
    res["registrationDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRegistrationDisabled(val)
        return nil
    }
    res["zoomDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetZoomDisabled(val)
        return nil
    }
    return res
}
func (m *DepMacOSEnrollmentProfile) IsNil()(bool) {
    return m == nil
}
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
func (m *DepMacOSEnrollmentProfile) SetAccessibilityScreenDisabled(value *bool)() {
    m.accessibilityScreenDisabled = value
}
func (m *DepMacOSEnrollmentProfile) SetChooseYourLockScreenDisabled(value *bool)() {
    m.chooseYourLockScreenDisabled = value
}
func (m *DepMacOSEnrollmentProfile) SetFileVaultDisabled(value *bool)() {
    m.fileVaultDisabled = value
}
func (m *DepMacOSEnrollmentProfile) SetICloudDiagnosticsDisabled(value *bool)() {
    m.iCloudDiagnosticsDisabled = value
}
func (m *DepMacOSEnrollmentProfile) SetICloudStorageDisabled(value *bool)() {
    m.iCloudStorageDisabled = value
}
func (m *DepMacOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    m.passCodeDisabled = value
}
func (m *DepMacOSEnrollmentProfile) SetRegistrationDisabled(value *bool)() {
    m.registrationDisabled = value
}
func (m *DepMacOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    m.zoomDisabled = value
}
