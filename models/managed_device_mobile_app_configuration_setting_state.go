package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceMobileAppConfigurationSettingState managed Device Mobile App Configuration Setting State for a given device.
type ManagedDeviceMobileAppConfigurationSettingState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Current value of setting on device
    currentValue *string
    // Error code for the setting
    errorCode *int64
    // Error description
    errorDescription *string
    // Name of setting instance that is being reported.
    instanceDisplayName *string
    // The OdataType property
    odataType *string
    // The setting that is being reported
    setting *string
    // SettingInstanceId
    settingInstanceId *string
    // Localized/user friendly setting name that is being reported
    settingName *string
    // Contributing policies
    sources []SettingSourceable
    // The state property
    state *ComplianceStatus
    // UserEmail
    userEmail *string
    // UserId
    userId *string
    // UserName
    userName *string
    // UserPrincipalName.
    userPrincipalName *string
}
// NewManagedDeviceMobileAppConfigurationSettingState instantiates a new managedDeviceMobileAppConfigurationSettingState and sets the default values.
func NewManagedDeviceMobileAppConfigurationSettingState()(*ManagedDeviceMobileAppConfigurationSettingState) {
    m := &ManagedDeviceMobileAppConfigurationSettingState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagedDeviceMobileAppConfigurationSettingStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceMobileAppConfigurationSettingStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceMobileAppConfigurationSettingState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCurrentValue gets the currentValue property value. Current value of setting on device
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetCurrentValue()(*string) {
    return m.currentValue
}
// GetErrorCode gets the errorCode property value. Error code for the setting
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetErrorCode()(*int64) {
    return m.errorCode
}
// GetErrorDescription gets the errorDescription property value. Error description
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetErrorDescription()(*string) {
    return m.errorDescription
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCurrentValue)
    res["errorCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetErrorCode)
    res["errorDescription"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetErrorDescription)
    res["instanceDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInstanceDisplayName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["setting"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSetting)
    res["settingInstanceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSettingInstanceId)
    res["settingName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSettingName)
    res["sources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSettingSourceFromDiscriminatorValue , m.SetSources)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseComplianceStatus , m.SetState)
    res["userEmail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserEmail)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserName)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetInstanceDisplayName gets the instanceDisplayName property value. Name of setting instance that is being reported.
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetInstanceDisplayName()(*string) {
    return m.instanceDisplayName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetOdataType()(*string) {
    return m.odataType
}
// GetSetting gets the setting property value. The setting that is being reported
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetSetting()(*string) {
    return m.setting
}
// GetSettingInstanceId gets the settingInstanceId property value. SettingInstanceId
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetSettingInstanceId()(*string) {
    return m.settingInstanceId
}
// GetSettingName gets the settingName property value. Localized/user friendly setting name that is being reported
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetSettingName()(*string) {
    return m.settingName
}
// GetSources gets the sources property value. Contributing policies
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetSources()([]SettingSourceable) {
    return m.sources
}
// GetState gets the state property value. The state property
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetState()(*ComplianceStatus) {
    return m.state
}
// GetUserEmail gets the userEmail property value. UserEmail
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetUserEmail()(*string) {
    return m.userEmail
}
// GetUserId gets the userId property value. UserId
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetUserId()(*string) {
    return m.userId
}
// GetUserName gets the userName property value. UserName
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. UserPrincipalName.
func (m *ManagedDeviceMobileAppConfigurationSettingState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *ManagedDeviceMobileAppConfigurationSettingState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("currentValue", m.GetCurrentValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorDescription", m.GetErrorDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("instanceDisplayName", m.GetInstanceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingInstanceId", m.GetSettingInstanceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    if m.GetSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSources())
        err := writer.WriteCollectionOfObjectValues("sources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userEmail", m.GetUserEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCurrentValue sets the currentValue property value. Current value of setting on device
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetCurrentValue(value *string)() {
    m.currentValue = value
}
// SetErrorCode sets the errorCode property value. Error code for the setting
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetErrorCode(value *int64)() {
    m.errorCode = value
}
// SetErrorDescription sets the errorDescription property value. Error description
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetErrorDescription(value *string)() {
    m.errorDescription = value
}
// SetInstanceDisplayName sets the instanceDisplayName property value. Name of setting instance that is being reported.
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetInstanceDisplayName(value *string)() {
    m.instanceDisplayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSetting sets the setting property value. The setting that is being reported
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetSetting(value *string)() {
    m.setting = value
}
// SetSettingInstanceId sets the settingInstanceId property value. SettingInstanceId
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetSettingInstanceId(value *string)() {
    m.settingInstanceId = value
}
// SetSettingName sets the settingName property value. Localized/user friendly setting name that is being reported
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetSettingName(value *string)() {
    m.settingName = value
}
// SetSources sets the sources property value. Contributing policies
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetSources(value []SettingSourceable)() {
    m.sources = value
}
// SetState sets the state property value. The state property
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// SetUserEmail sets the userEmail property value. UserEmail
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetUserEmail(value *string)() {
    m.userEmail = value
}
// SetUserId sets the userId property value. UserId
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetUserId(value *string)() {
    m.userId = value
}
// SetUserName sets the userName property value. UserName
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. UserPrincipalName.
func (m *ManagedDeviceMobileAppConfigurationSettingState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
