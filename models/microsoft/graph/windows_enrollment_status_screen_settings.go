package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new windowsEnrollmentStatusScreenSettings and sets the default values.
func NewWindowsEnrollmentStatusScreenSettings()(*WindowsEnrollmentStatusScreenSettings) {
    m := &WindowsEnrollmentStatusScreenSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsEnrollmentStatusScreenSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowDeviceUseBeforeProfileAndAppInstallComplete property value. Allow or block user to use device before profile and app installation complete
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseBeforeProfileAndAppInstallComplete()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeviceUseBeforeProfileAndAppInstallComplete
    }
}
// Gets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowDeviceUseOnInstallFailure()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeviceUseOnInstallFailure
    }
}
// Gets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetAllowLogCollectionOnInstallFailure()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowLogCollectionOnInstallFailure
    }
}
// Gets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetBlockDeviceSetupRetryByUser()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockDeviceSetupRetryByUser
    }
}
// Gets the customErrorMessage property value. Set custom error message to show upon installation failure
func (m *WindowsEnrollmentStatusScreenSettings) GetCustomErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customErrorMessage
    }
}
// Gets the hideInstallationProgress property value. Show or hide installation progress to user
func (m *WindowsEnrollmentStatusScreenSettings) GetHideInstallationProgress()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideInstallationProgress
    }
}
// Gets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
func (m *WindowsEnrollmentStatusScreenSettings) GetInstallProgressTimeoutInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installProgressTimeoutInMinutes
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WindowsEnrollmentStatusScreenSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowDeviceUseBeforeProfileAndAppInstallComplete property value. Allow or block user to use device before profile and app installation complete
// Parameters:
//  - value : Value to set for the allowDeviceUseBeforeProfileAndAppInstallComplete property.
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseBeforeProfileAndAppInstallComplete(value *bool)() {
    m.allowDeviceUseBeforeProfileAndAppInstallComplete = value
}
// Sets the allowDeviceUseOnInstallFailure property value. Allow the user to continue using the device on installation failure
// Parameters:
//  - value : Value to set for the allowDeviceUseOnInstallFailure property.
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowDeviceUseOnInstallFailure(value *bool)() {
    m.allowDeviceUseOnInstallFailure = value
}
// Sets the allowLogCollectionOnInstallFailure property value. Allow or block log collection on installation failure
// Parameters:
//  - value : Value to set for the allowLogCollectionOnInstallFailure property.
func (m *WindowsEnrollmentStatusScreenSettings) SetAllowLogCollectionOnInstallFailure(value *bool)() {
    m.allowLogCollectionOnInstallFailure = value
}
// Sets the blockDeviceSetupRetryByUser property value. Allow the user to retry the setup on installation failure
// Parameters:
//  - value : Value to set for the blockDeviceSetupRetryByUser property.
func (m *WindowsEnrollmentStatusScreenSettings) SetBlockDeviceSetupRetryByUser(value *bool)() {
    m.blockDeviceSetupRetryByUser = value
}
// Sets the customErrorMessage property value. Set custom error message to show upon installation failure
// Parameters:
//  - value : Value to set for the customErrorMessage property.
func (m *WindowsEnrollmentStatusScreenSettings) SetCustomErrorMessage(value *string)() {
    m.customErrorMessage = value
}
// Sets the hideInstallationProgress property value. Show or hide installation progress to user
// Parameters:
//  - value : Value to set for the hideInstallationProgress property.
func (m *WindowsEnrollmentStatusScreenSettings) SetHideInstallationProgress(value *bool)() {
    m.hideInstallationProgress = value
}
// Sets the installProgressTimeoutInMinutes property value. Set installation progress timeout in minutes
// Parameters:
//  - value : Value to set for the installProgressTimeoutInMinutes property.
func (m *WindowsEnrollmentStatusScreenSettings) SetInstallProgressTimeoutInMinutes(value *int32)() {
    m.installProgressTimeoutInMinutes = value
}
