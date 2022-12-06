package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcOrganizationSettings 
type CloudPcOrganizationSettings struct {
    Entity
    // Specifies whether new Cloud PCs will be automatically enrolled in Microsoft Endpoint Manager(MEM). The default value is false.
    enableMEMAutoEnroll *bool
    // The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11, unknownFutureValue.
    osVersion *CloudPcOperatingSystem
    // The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator, unknownFutureValue.
    userAccountType *CloudPcUserAccountType
    // Represents the Cloud PC organization settings for a tenant. A tenant has only one cloudPcOrganizationSettings object. The default language value en-US.
    windowsSettings CloudPcWindowsSettingsable
}
// NewCloudPcOrganizationSettings instantiates a new CloudPcOrganizationSettings and sets the default values.
func NewCloudPcOrganizationSettings()(*CloudPcOrganizationSettings) {
    m := &CloudPcOrganizationSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcOrganizationSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcOrganizationSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOrganizationSettings(), nil
}
// GetEnableMEMAutoEnroll gets the enableMEMAutoEnroll property value. Specifies whether new Cloud PCs will be automatically enrolled in Microsoft Endpoint Manager(MEM). The default value is false.
func (m *CloudPcOrganizationSettings) GetEnableMEMAutoEnroll()(*bool) {
    return m.enableMEMAutoEnroll
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOrganizationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enableMEMAutoEnroll"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableMEMAutoEnroll)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcOperatingSystem , m.SetOsVersion)
    res["userAccountType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcUserAccountType , m.SetUserAccountType)
    res["windowsSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCloudPcWindowsSettingsFromDiscriminatorValue , m.SetWindowsSettings)
    return res
}
// GetOsVersion gets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11, unknownFutureValue.
func (m *CloudPcOrganizationSettings) GetOsVersion()(*CloudPcOperatingSystem) {
    return m.osVersion
}
// GetUserAccountType gets the userAccountType property value. The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator, unknownFutureValue.
func (m *CloudPcOrganizationSettings) GetUserAccountType()(*CloudPcUserAccountType) {
    return m.userAccountType
}
// GetWindowsSettings gets the windowsSettings property value. Represents the Cloud PC organization settings for a tenant. A tenant has only one cloudPcOrganizationSettings object. The default language value en-US.
func (m *CloudPcOrganizationSettings) GetWindowsSettings()(CloudPcWindowsSettingsable) {
    return m.windowsSettings
}
// Serialize serializes information the current object
func (m *CloudPcOrganizationSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enableMEMAutoEnroll", m.GetEnableMEMAutoEnroll())
        if err != nil {
            return err
        }
    }
    if m.GetOsVersion() != nil {
        cast := (*m.GetOsVersion()).String()
        err = writer.WriteStringValue("osVersion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountType() != nil {
        cast := (*m.GetUserAccountType()).String()
        err = writer.WriteStringValue("userAccountType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsSettings", m.GetWindowsSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnableMEMAutoEnroll sets the enableMEMAutoEnroll property value. Specifies whether new Cloud PCs will be automatically enrolled in Microsoft Endpoint Manager(MEM). The default value is false.
func (m *CloudPcOrganizationSettings) SetEnableMEMAutoEnroll(value *bool)() {
    m.enableMEMAutoEnroll = value
}
// SetOsVersion sets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11, unknownFutureValue.
func (m *CloudPcOrganizationSettings) SetOsVersion(value *CloudPcOperatingSystem)() {
    m.osVersion = value
}
// SetUserAccountType sets the userAccountType property value. The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator, unknownFutureValue.
func (m *CloudPcOrganizationSettings) SetUserAccountType(value *CloudPcUserAccountType)() {
    m.userAccountType = value
}
// SetWindowsSettings sets the windowsSettings property value. Represents the Cloud PC organization settings for a tenant. A tenant has only one cloudPcOrganizationSettings object. The default language value en-US.
func (m *CloudPcOrganizationSettings) SetWindowsSettings(value CloudPcWindowsSettingsable)() {
    m.windowsSettings = value
}
