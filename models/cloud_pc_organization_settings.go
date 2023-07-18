package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcOrganizationSettings 
type CloudPcOrganizationSettings struct {
    Entity
}
// NewCloudPcOrganizationSettings instantiates a new cloudPcOrganizationSettings and sets the default values.
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
// GetEnableMEMAutoEnroll gets the enableMEMAutoEnroll property value. Specifies whether new Cloud PCs will be automatically enrolled in Microsoft Endpoint Manager (MEM). The default value is false.
func (m *CloudPcOrganizationSettings) GetEnableMEMAutoEnroll()(*bool) {
    val, err := m.GetBackingStore().Get("enableMEMAutoEnroll")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnableSingleSignOn gets the enableSingleSignOn property value. True if the provisioned Cloud PC can be accessed by single sign-on. False indicates that the provisioned Cloud PC doesn't support this feature. Default value is false. Windows 365 users can use single sign-on to authenticate to Azure Active Directory (Azure AD) with passwordless options (for example, FIDO keys) to access their Cloud PC. Optional.
func (m *CloudPcOrganizationSettings) GetEnableSingleSignOn()(*bool) {
    val, err := m.GetBackingStore().Get("enableSingleSignOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOrganizationSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enableMEMAutoEnroll"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableMEMAutoEnroll(val)
        }
        return nil
    }
    res["enableSingleSignOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleSignOn(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOperatingSystem)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val.(*CloudPcOperatingSystem))
        }
        return nil
    }
    res["userAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcUserAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*CloudPcUserAccountType))
        }
        return nil
    }
    res["windowsSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcWindowsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSettings(val.(CloudPcWindowsSettingsable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcOrganizationSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11, unknownFutureValue.
func (m *CloudPcOrganizationSettings) GetOsVersion()(*CloudPcOperatingSystem) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcOperatingSystem)
    }
    return nil
}
// GetUserAccountType gets the userAccountType property value. The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator, unknownFutureValue.
func (m *CloudPcOrganizationSettings) GetUserAccountType()(*CloudPcUserAccountType) {
    val, err := m.GetBackingStore().Get("userAccountType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcUserAccountType)
    }
    return nil
}
// GetWindowsSettings gets the windowsSettings property value. Represents the Cloud PC organization settings for a tenant. A tenant has only one cloudPcOrganizationSettings object. The default language value en-US.
func (m *CloudPcOrganizationSettings) GetWindowsSettings()(CloudPcWindowsSettingsable) {
    val, err := m.GetBackingStore().Get("windowsSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcWindowsSettingsable)
    }
    return nil
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
    {
        err = writer.WriteBoolValue("enableSingleSignOn", m.GetEnableSingleSignOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetEnableMEMAutoEnroll sets the enableMEMAutoEnroll property value. Specifies whether new Cloud PCs will be automatically enrolled in Microsoft Endpoint Manager (MEM). The default value is false.
func (m *CloudPcOrganizationSettings) SetEnableMEMAutoEnroll(value *bool)() {
    err := m.GetBackingStore().Set("enableMEMAutoEnroll", value)
    if err != nil {
        panic(err)
    }
}
// SetEnableSingleSignOn sets the enableSingleSignOn property value. True if the provisioned Cloud PC can be accessed by single sign-on. False indicates that the provisioned Cloud PC doesn't support this feature. Default value is false. Windows 365 users can use single sign-on to authenticate to Azure Active Directory (Azure AD) with passwordless options (for example, FIDO keys) to access their Cloud PC. Optional.
func (m *CloudPcOrganizationSettings) SetEnableSingleSignOn(value *bool)() {
    err := m.GetBackingStore().Set("enableSingleSignOn", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcOrganizationSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. The possible values are: windows10, windows11, unknownFutureValue.
func (m *CloudPcOrganizationSettings) SetOsVersion(value *CloudPcOperatingSystem)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAccountType sets the userAccountType property value. The account type of the user on provisioned Cloud PCs. The possible values are: standardUser, administrator, unknownFutureValue.
func (m *CloudPcOrganizationSettings) SetUserAccountType(value *CloudPcUserAccountType)() {
    err := m.GetBackingStore().Set("userAccountType", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsSettings sets the windowsSettings property value. Represents the Cloud PC organization settings for a tenant. A tenant has only one cloudPcOrganizationSettings object. The default language value en-US.
func (m *CloudPcOrganizationSettings) SetWindowsSettings(value CloudPcWindowsSettingsable)() {
    err := m.GetBackingStore().Set("windowsSettings", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcOrganizationSettingsable 
type CloudPcOrganizationSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnableMEMAutoEnroll()(*bool)
    GetEnableSingleSignOn()(*bool)
    GetOdataType()(*string)
    GetOsVersion()(*CloudPcOperatingSystem)
    GetUserAccountType()(*CloudPcUserAccountType)
    GetWindowsSettings()(CloudPcWindowsSettingsable)
    SetEnableMEMAutoEnroll(value *bool)()
    SetEnableSingleSignOn(value *bool)()
    SetOdataType(value *string)()
    SetOsVersion(value *CloudPcOperatingSystem)()
    SetUserAccountType(value *CloudPcUserAccountType)()
    SetWindowsSettings(value CloudPcWindowsSettingsable)()
}
