package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsEnrollmentStatusScreenSettings struct {
    additionalData map[string]interface{};
    allowDeviceUseBeforeProfileAndAppInstallComplete *bool;
    allowDeviceUseOnInstallFailure *bool;
    allowLogCollectionOnInstallFailure *bool;
    blockDeviceSetupRetryByUser *bool;
    customErrorMessage *string;
    hideInstallationProgress *bool;
    installProgressTimeoutInMinutes *int32;
}
func NewWindowsEnrollmentStatusScreenSettings()(*WindowsEnrollmentStatusScreenSettings) {
    m := &WindowsEnrollmentStatusScreenSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WindowsEnrollmentStatusScreenSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseBeforeProfileAndAppInstallComplete()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeviceUseBeforeProfileAndAppInstallComplete
    }
}
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseOnInstallFailure()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeviceUseOnInstallFailure
    }
}
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowLogCollectionOnInstallFailure()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowLogCollectionOnInstallFailure
    }
}
func (m *WindowsEnrollmentStatusScreenSettings) GetBlockDeviceSetupRetryByUser()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockDeviceSetupRetryByUser
    }
}
func (m *WindowsEnrollmentStatusScreenSettings) GetCustomErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customErrorMessage
    }
}
func (m *WindowsEnrollmentStatusScreenSettings) GetHideInstallationProgress()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideInstallationProgress
    }
}
func (m *WindowsEnrollmentStatusScreenSettings) GetInstallProgressTimeoutInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installProgressTimeoutInMinutes
    }
}
func (m *WindowsEnrollmentStatusScreenSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowDeviceUseBeforeProfileAndAppInstallComplete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowDeviceUseBeforeProfileAndAppInstallComplete(val)
        return nil
    }
    res["allowDeviceUseOnInstallFailure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowDeviceUseOnInstallFailure(val)
        return nil
    }
    res["allowLogCollectionOnInstallFailure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowLogCollectionOnInstallFailure(val)
        return nil
    }
    res["blockDeviceSetupRetryByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetBlockDeviceSetupRetryByUser(val)
        return nil
    }
    res["customErrorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomErrorMessage(val)
        return nil
    }
    res["hideInstallationProgress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHideInstallationProgress(val)
        return nil
    }
    res["installProgressTimeoutInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInstallProgressTimeoutInMinutes(val)
        return nil
    }
    return res
}
func (m *WindowsEnrollmentStatusScreenSettings) IsNil()(bool) {
    return m == nil
}
func (m *WindowsEnrollmentStatusScreenSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowDeviceUseBeforeProfileAndAppInstallComplete", m.GetAllowDeviceUseBeforeProfileAndAppInstallComplete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeviceUseOnInstallFailure", m.GetAllowDeviceUseOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowLogCollectionOnInstallFailure", m.GetAllowLogCollectionOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockDeviceSetupRetryByUser", m.GetBlockDeviceSetupRetryByUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customErrorMessage", m.GetCustomErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideInstallationProgress", m.GetHideInstallationProgress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("installProgressTimeoutInMinutes", m.GetInstallProgressTimeoutInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsEnrollmentStatusScreenSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseBeforeProfileAndAppInstallComplete(value *bool)() {
    m.allowDeviceUseBeforeProfileAndAppInstallComplete = value
}
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseOnInstallFailure(value *bool)() {
    m.allowDeviceUseOnInstallFailure = value
}
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowLogCollectionOnInstallFailure(value *bool)() {
    m.allowLogCollectionOnInstallFailure = value
}
func (m *WindowsEnrollmentStatusScreenSettings) SetBlockDeviceSetupRetryByUser(value *bool)() {
    m.blockDeviceSetupRetryByUser = value
}
func (m *WindowsEnrollmentStatusScreenSettings) SetCustomErrorMessage(value *string)() {
    m.customErrorMessage = value
}
func (m *WindowsEnrollmentStatusScreenSettings) SetHideInstallationProgress(value *bool)() {
    m.hideInstallationProgress = value
}
func (m *WindowsEnrollmentStatusScreenSettings) SetInstallProgressTimeoutInMinutes(value *int32)() {
    m.installProgressTimeoutInMinutes = value
}
