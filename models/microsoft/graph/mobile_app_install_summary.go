package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileAppInstallSummary struct {
    Entity
    // Number of Devices that have failed to install this app.
    failedDeviceCount *int32;
    // Number of Users that have 1 or more device that failed to install this app.
    failedUserCount *int32;
    // Number of Devices that have successfully installed this app.
    installedDeviceCount *int32;
    // Number of Users whose devices have all succeeded to install this app.
    installedUserCount *int32;
    // Number of Devices that are not applicable for this app.
    notApplicableDeviceCount *int32;
    // Number of Users whose devices were all not applicable for this app.
    notApplicableUserCount *int32;
    // Number of Devices that does not have this app installed.
    notInstalledDeviceCount *int32;
    // Number of Users that have 1 or more devices that did not install this app.
    notInstalledUserCount *int32;
    // Number of Devices that have been notified to install this app.
    pendingInstallDeviceCount *int32;
    // Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
    pendingInstallUserCount *int32;
}
// Instantiates a new mobileAppInstallSummary and sets the default values.
func NewMobileAppInstallSummary()(*MobileAppInstallSummary) {
    m := &MobileAppInstallSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the failedDeviceCount property value. Number of Devices that have failed to install this app.
func (m *MobileAppInstallSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// Gets the failedUserCount property value. Number of Users that have 1 or more device that failed to install this app.
func (m *MobileAppInstallSummary) GetFailedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUserCount
    }
}
// Gets the installedDeviceCount property value. Number of Devices that have successfully installed this app.
func (m *MobileAppInstallSummary) GetInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedDeviceCount
    }
}
// Gets the installedUserCount property value. Number of Users whose devices have all succeeded to install this app.
func (m *MobileAppInstallSummary) GetInstalledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedUserCount
    }
}
// Gets the notApplicableDeviceCount property value. Number of Devices that are not applicable for this app.
func (m *MobileAppInstallSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// Gets the notApplicableUserCount property value. Number of Users whose devices were all not applicable for this app.
func (m *MobileAppInstallSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
// Gets the notInstalledDeviceCount property value. Number of Devices that does not have this app installed.
func (m *MobileAppInstallSummary) GetNotInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledDeviceCount
    }
}
// Gets the notInstalledUserCount property value. Number of Users that have 1 or more devices that did not install this app.
func (m *MobileAppInstallSummary) GetNotInstalledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledUserCount
    }
}
// Gets the pendingInstallDeviceCount property value. Number of Devices that have been notified to install this app.
func (m *MobileAppInstallSummary) GetPendingInstallDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingInstallDeviceCount
    }
}
// Gets the pendingInstallUserCount property value. Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
func (m *MobileAppInstallSummary) GetPendingInstallUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingInstallUserCount
    }
}
// The deserialization information for the current model
func (m *MobileAppInstallSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    res["failedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUserCount(val)
        }
        return nil
    }
    res["installedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledDeviceCount(val)
        }
        return nil
    }
    res["installedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledUserCount(val)
        }
        return nil
    }
    res["notApplicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["notApplicableUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableUserCount(val)
        }
        return nil
    }
    res["notInstalledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledDeviceCount(val)
        }
        return nil
    }
    res["notInstalledUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledUserCount(val)
        }
        return nil
    }
    res["pendingInstallDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingInstallDeviceCount(val)
        }
        return nil
    }
    res["pendingInstallUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingInstallUserCount(val)
        }
        return nil
    }
    return res
}
func (m *MobileAppInstallSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileAppInstallSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedUserCount", m.GetFailedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installedDeviceCount", m.GetInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installedUserCount", m.GetInstalledUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableDeviceCount", m.GetNotApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableUserCount", m.GetNotApplicableUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notInstalledDeviceCount", m.GetNotInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notInstalledUserCount", m.GetNotInstalledUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingInstallDeviceCount", m.GetPendingInstallDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingInstallUserCount", m.GetPendingInstallUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the failedDeviceCount property value. Number of Devices that have failed to install this app.
// Parameters:
//  - value : Value to set for the failedDeviceCount property.
func (m *MobileAppInstallSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
// Sets the failedUserCount property value. Number of Users that have 1 or more device that failed to install this app.
// Parameters:
//  - value : Value to set for the failedUserCount property.
func (m *MobileAppInstallSummary) SetFailedUserCount(value *int32)() {
    m.failedUserCount = value
}
// Sets the installedDeviceCount property value. Number of Devices that have successfully installed this app.
// Parameters:
//  - value : Value to set for the installedDeviceCount property.
func (m *MobileAppInstallSummary) SetInstalledDeviceCount(value *int32)() {
    m.installedDeviceCount = value
}
// Sets the installedUserCount property value. Number of Users whose devices have all succeeded to install this app.
// Parameters:
//  - value : Value to set for the installedUserCount property.
func (m *MobileAppInstallSummary) SetInstalledUserCount(value *int32)() {
    m.installedUserCount = value
}
// Sets the notApplicableDeviceCount property value. Number of Devices that are not applicable for this app.
// Parameters:
//  - value : Value to set for the notApplicableDeviceCount property.
func (m *MobileAppInstallSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// Sets the notApplicableUserCount property value. Number of Users whose devices were all not applicable for this app.
// Parameters:
//  - value : Value to set for the notApplicableUserCount property.
func (m *MobileAppInstallSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
// Sets the notInstalledDeviceCount property value. Number of Devices that does not have this app installed.
// Parameters:
//  - value : Value to set for the notInstalledDeviceCount property.
func (m *MobileAppInstallSummary) SetNotInstalledDeviceCount(value *int32)() {
    m.notInstalledDeviceCount = value
}
// Sets the notInstalledUserCount property value. Number of Users that have 1 or more devices that did not install this app.
// Parameters:
//  - value : Value to set for the notInstalledUserCount property.
func (m *MobileAppInstallSummary) SetNotInstalledUserCount(value *int32)() {
    m.notInstalledUserCount = value
}
// Sets the pendingInstallDeviceCount property value. Number of Devices that have been notified to install this app.
// Parameters:
//  - value : Value to set for the pendingInstallDeviceCount property.
func (m *MobileAppInstallSummary) SetPendingInstallDeviceCount(value *int32)() {
    m.pendingInstallDeviceCount = value
}
// Sets the pendingInstallUserCount property value. Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
// Parameters:
//  - value : Value to set for the pendingInstallUserCount property.
func (m *MobileAppInstallSummary) SetPendingInstallUserCount(value *int32)() {
    m.pendingInstallUserCount = value
}
