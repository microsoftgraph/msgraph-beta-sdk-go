package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MobileAppInstallSummary struct {
    Entity
    failedDeviceCount *int32;
    failedUserCount *int32;
    installedDeviceCount *int32;
    installedUserCount *int32;
    notApplicableDeviceCount *int32;
    notApplicableUserCount *int32;
    notInstalledDeviceCount *int32;
    notInstalledUserCount *int32;
    pendingInstallDeviceCount *int32;
    pendingInstallUserCount *int32;
}
func NewMobileAppInstallSummary()(*MobileAppInstallSummary) {
    m := &MobileAppInstallSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MobileAppInstallSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
func (m *MobileAppInstallSummary) GetFailedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUserCount
    }
}
func (m *MobileAppInstallSummary) GetInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedDeviceCount
    }
}
func (m *MobileAppInstallSummary) GetInstalledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedUserCount
    }
}
func (m *MobileAppInstallSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
func (m *MobileAppInstallSummary) GetNotApplicableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableUserCount
    }
}
func (m *MobileAppInstallSummary) GetNotInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledDeviceCount
    }
}
func (m *MobileAppInstallSummary) GetNotInstalledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledUserCount
    }
}
func (m *MobileAppInstallSummary) GetPendingInstallDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingInstallDeviceCount
    }
}
func (m *MobileAppInstallSummary) GetPendingInstallUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingInstallUserCount
    }
}
func (m *MobileAppInstallSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedDeviceCount(val)
        return nil
    }
    res["failedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedUserCount(val)
        return nil
    }
    res["installedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInstalledDeviceCount(val)
        return nil
    }
    res["installedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInstalledUserCount(val)
        return nil
    }
    res["notApplicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotApplicableDeviceCount(val)
        return nil
    }
    res["notApplicableUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotApplicableUserCount(val)
        return nil
    }
    res["notInstalledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotInstalledDeviceCount(val)
        return nil
    }
    res["notInstalledUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotInstalledUserCount(val)
        return nil
    }
    res["pendingInstallDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingInstallDeviceCount(val)
        return nil
    }
    res["pendingInstallUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPendingInstallUserCount(val)
        return nil
    }
    return res
}
func (m *MobileAppInstallSummary) IsNil()(bool) {
    return m == nil
}
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
func (m *MobileAppInstallSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
func (m *MobileAppInstallSummary) SetFailedUserCount(value *int32)() {
    m.failedUserCount = value
}
func (m *MobileAppInstallSummary) SetInstalledDeviceCount(value *int32)() {
    m.installedDeviceCount = value
}
func (m *MobileAppInstallSummary) SetInstalledUserCount(value *int32)() {
    m.installedUserCount = value
}
func (m *MobileAppInstallSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
func (m *MobileAppInstallSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
func (m *MobileAppInstallSummary) SetNotInstalledDeviceCount(value *int32)() {
    m.notInstalledDeviceCount = value
}
func (m *MobileAppInstallSummary) SetNotInstalledUserCount(value *int32)() {
    m.notInstalledUserCount = value
}
func (m *MobileAppInstallSummary) SetPendingInstallDeviceCount(value *int32)() {
    m.pendingInstallDeviceCount = value
}
func (m *MobileAppInstallSummary) SetPendingInstallUserCount(value *int32)() {
    m.pendingInstallUserCount = value
}
