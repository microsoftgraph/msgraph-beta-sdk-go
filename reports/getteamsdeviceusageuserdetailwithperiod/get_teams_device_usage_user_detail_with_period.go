package getteamsdeviceusageuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetTeamsDeviceUsageUserDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
    deletedDate *string;
    // Whether this user has been deleted or soft deleted.
    isDeleted *bool;
    // Whether the user has been assigned a Teams license.
    isLicensed *bool;
    // The last date that the user participated in a Microsoft Teams activity.
    lastActivityDate *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // Whether the user was active on the Teams mobile client for Android.
    usedAndroidPhone *bool;
    // Whether the user was active in the Teams desktop client on a ChromeOS computer.
    usedChromeOS *bool;
    // Whether the user was active on the Teams mobile client for iOS.
    usediOS *bool;
    // Whether the user was active in the Teams desktop client on a Linux computer.
    usedLinux *bool;
    // Whether the user was active in the Teams desktop client on a macOS computer.
    usedMac *bool;
    // Whether the user was active in the Teams web client on devices.
    usedWeb *bool;
    // Whether the user was active in the Teams desktop client on a Windows-based computer.
    usedWindows *bool;
    // Whether the user was active on the Teams mobile client for Windows phone.
    usedWindowsPhone *bool;
    // The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
    userPrincipalName *string;
}
// Instantiates a new getTeamsDeviceUsageUserDetailWithPeriod and sets the default values.
func NewGetTeamsDeviceUsageUserDetailWithPeriod()(*GetTeamsDeviceUsageUserDetailWithPeriod) {
    m := &GetTeamsDeviceUsageUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// Gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the isLicensed property value. Whether the user has been assigned a Teams license.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetIsLicensed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLicensed
    }
}
// Gets the lastActivityDate property value. The last date that the user participated in a Microsoft Teams activity.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the usedAndroidPhone property value. Whether the user was active on the Teams mobile client for Android.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUsedAndroidPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedAndroidPhone
    }
}
// Gets the usedChromeOS property value. Whether the user was active in the Teams desktop client on a ChromeOS computer.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUsedChromeOS()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedChromeOS
    }
}
// Gets the usediOS property value. Whether the user was active on the Teams mobile client for iOS.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUsediOS()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediOS
    }
}
// Gets the usedLinux property value. Whether the user was active in the Teams desktop client on a Linux computer.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUsedLinux()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedLinux
    }
}
// Gets the usedMac property value. Whether the user was active in the Teams desktop client on a macOS computer.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUsedMac()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedMac
    }
}
// Gets the usedWeb property value. Whether the user was active in the Teams web client on devices.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUsedWeb()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWeb
    }
}
// Gets the usedWindows property value. Whether the user was active in the Teams desktop client on a Windows-based computer.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUsedWindows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindows
    }
}
// Gets the usedWindowsPhone property value. Whether the user was active on the Teams mobile client for Windows phone.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUsedWindowsPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindowsPhone
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeletedDate(val)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["isLicensed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsLicensed(val)
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActivityDate(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["usedAndroidPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedAndroidPhone(val)
        return nil
    }
    res["usedChromeOS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedChromeOS(val)
        return nil
    }
    res["usediOS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsediOS(val)
        return nil
    }
    res["usedLinux"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedLinux(val)
        return nil
    }
    res["usedMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedMac(val)
        return nil
    }
    res["usedWeb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedWeb(val)
        return nil
    }
    res["usedWindows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedWindows(val)
        return nil
    }
    res["usedWindowsPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedWindowsPhone(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deletedDate", m.GetDeletedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLicensed", m.GetIsLicensed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedAndroidPhone", m.GetUsedAndroidPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedChromeOS", m.GetUsedChromeOS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usediOS", m.GetUsediOS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedLinux", m.GetUsedLinux())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedMac", m.GetUsedMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWeb", m.GetUsedWeb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWindows", m.GetUsedWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWindowsPhone", m.GetUsedWindowsPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
// Parameters:
//  - value : Value to set for the deletedDate property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
// Sets the isDeleted property value. Whether this user has been deleted or soft deleted.
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the isLicensed property value. Whether the user has been assigned a Teams license.
// Parameters:
//  - value : Value to set for the isLicensed property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetIsLicensed(value *bool)() {
    m.isLicensed = value
}
// Sets the lastActivityDate property value. The last date that the user participated in a Microsoft Teams activity.
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the usedAndroidPhone property value. Whether the user was active on the Teams mobile client for Android.
// Parameters:
//  - value : Value to set for the usedAndroidPhone property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUsedAndroidPhone(value *bool)() {
    m.usedAndroidPhone = value
}
// Sets the usedChromeOS property value. Whether the user was active in the Teams desktop client on a ChromeOS computer.
// Parameters:
//  - value : Value to set for the usedChromeOS property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUsedChromeOS(value *bool)() {
    m.usedChromeOS = value
}
// Sets the usediOS property value. Whether the user was active on the Teams mobile client for iOS.
// Parameters:
//  - value : Value to set for the usediOS property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUsediOS(value *bool)() {
    m.usediOS = value
}
// Sets the usedLinux property value. Whether the user was active in the Teams desktop client on a Linux computer.
// Parameters:
//  - value : Value to set for the usedLinux property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUsedLinux(value *bool)() {
    m.usedLinux = value
}
// Sets the usedMac property value. Whether the user was active in the Teams desktop client on a macOS computer.
// Parameters:
//  - value : Value to set for the usedMac property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUsedMac(value *bool)() {
    m.usedMac = value
}
// Sets the usedWeb property value. Whether the user was active in the Teams web client on devices.
// Parameters:
//  - value : Value to set for the usedWeb property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUsedWeb(value *bool)() {
    m.usedWeb = value
}
// Sets the usedWindows property value. Whether the user was active in the Teams desktop client on a Windows-based computer.
// Parameters:
//  - value : Value to set for the usedWindows property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUsedWindows(value *bool)() {
    m.usedWindows = value
}
// Sets the usedWindowsPhone property value. Whether the user was active on the Teams mobile client for Windows phone.
// Parameters:
//  - value : Value to set for the usedWindowsPhone property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUsedWindowsPhone(value *bool)() {
    m.usedWindowsPhone = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetTeamsDeviceUsageUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
