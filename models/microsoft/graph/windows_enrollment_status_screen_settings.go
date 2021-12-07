package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsEnrollmentStatusScreenSettings 
type WindowsEnrollmentStatusScreenSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Allow or block user to use device before profile and app installation complete
    allowDeviceUseBeforeProfileAndAppInstallComplete *bool;
    // Allow the user to continue using the device on installation failure
    allowDeviceUseOnInstallFailure *bool;
    // Allow or block log collection on installation failure
    allowLogCollectionOnInstallFailure *bool;
    // Allow the user to retry the setup on installation failure
    blockDeviceSetupRetryByUser *bool;
    // Set custom error message to show upon installation failure
    customErrorMessage *string;
    // Show or hide installation progress to user
    hideInstallationProgress *bool;
    // Set installation progress timeout in minutes
    installProgressTimeoutInMinutes *int32;
}
// NewWindowsEnrollmentStatusScreenSettings instantiates a new windowsEnrollmentStatusScreenSettings and sets the default values.
func NewWindowsEnrollmentStatusScreenSettings()(*WindowsEnrollmentStatusScreenSettings) {
    m := &WindowsEnrollmentStatusScreenSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsEnrollmentStatusScreenSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowDeviceUseBeforeProfileAndAppInstallComplete gets the allowDeviceUseBeforeProfileAndAppInstallComplete property value. Allow or block user to use device before profile and app installation complete
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseBeforeProfileAndAppInstallComplete()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeviceUseBeforeProfileAndAppInstallComplete
    }
}
// GetAllowDeviceUseOnInstallFailure gets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseOnInstallFailure()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeviceUseOnInstallFailure
    }
}
// GetAllowLogCollectionOnInstallFailure gets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowLogCollectionOnInstallFailure()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowLogCollectionOnInstallFailure
    }
}
// GetBlockDeviceSetupRetryByUser gets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetBlockDeviceSetupRetryByUser()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockDeviceSetupRetryByUser
    }
}
// GetCustomErrorMessage gets the customErrorMessage property value. Set custom error message to show upon installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetCustomErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customErrorMessage
    }
}
// GetHideInstallationProgress gets the hideInstallationProgress property value. Show or hide installation progress to user
func (m *WindowsEnrollmentStatusScreenSettings) GetHideInstallationProgress()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideInstallationProgress
    }
}
// GetInstallProgressTimeoutInMinutes gets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
func (m *WindowsEnrollmentStatusScreenSettings) GetInstallProgressTimeoutInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installProgressTimeoutInMinutes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsEnrollmentStatusScreenSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowDeviceUseBeforeProfileAndAppInstallComplete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceUseBeforeProfileAndAppInstallComplete(val)
        }
        return nil
    }
    res["allowDeviceUseOnInstallFailure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceUseOnInstallFailure(val)
        }
        return nil
    }
    res["allowLogCollectionOnInstallFailure"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLogCollectionOnInstallFailure(val)
        }
        return nil
    }
    res["blockDeviceSetupRetryByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDeviceSetupRetryByUser(val)
        }
        return nil
    }
    res["customErrorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomErrorMessage(val)
        }
        return nil
    }
    res["hideInstallationProgress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideInstallationProgress(val)
        }
        return nil
    }
    res["installProgressTimeoutInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallProgressTimeoutInMinutes(val)
        }
        return nil
    }
    return res
}
func (m *WindowsEnrollmentStatusScreenSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsEnrollmentStatusScreenSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowDeviceUseBeforeProfileAndAppInstallComplete sets the allowDeviceUseBeforeProfileAndAppInstallComplete property value. Allow or block user to use device before profile and app installation complete
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseBeforeProfileAndAppInstallComplete(value *bool)() {
    if m != nil {
        m.allowDeviceUseBeforeProfileAndAppInstallComplete = value
    }
}
// SetAllowDeviceUseOnInstallFailure sets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseOnInstallFailure(value *bool)() {
    if m != nil {
        m.allowDeviceUseOnInstallFailure = value
    }
}
// SetAllowLogCollectionOnInstallFailure sets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowLogCollectionOnInstallFailure(value *bool)() {
    if m != nil {
        m.allowLogCollectionOnInstallFailure = value
    }
}
// SetBlockDeviceSetupRetryByUser sets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetBlockDeviceSetupRetryByUser(value *bool)() {
    if m != nil {
        m.blockDeviceSetupRetryByUser = value
    }
}
// SetCustomErrorMessage sets the customErrorMessage property value. Set custom error message to show upon installation failure
func (m *WindowsEnrollmentStatusScreenSettings) SetCustomErrorMessage(value *string)() {
    if m != nil {
        m.customErrorMessage = value
    }
}
// SetHideInstallationProgress sets the hideInstallationProgress property value. Show or hide installation progress to user
func (m *WindowsEnrollmentStatusScreenSettings) SetHideInstallationProgress(value *bool)() {
    if m != nil {
        m.hideInstallationProgress = value
    }
}
// SetInstallProgressTimeoutInMinutes sets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
func (m *WindowsEnrollmentStatusScreenSettings) SetInstallProgressTimeoutInMinutes(value *int32)() {
    if m != nil {
        m.installProgressTimeoutInMinutes = value
    }
}
