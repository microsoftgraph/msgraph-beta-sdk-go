package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserAppInstallStatus struct {
    Entity
    // The navigation link to the mobile app.
    app *MobileApp;
    // The install state of the app on devices.
    deviceStatuses []MobileAppInstallStatus;
    // Failed Device Count.
    failedDeviceCount *int32;
    // Installed Device Count.
    installedDeviceCount *int32;
    // Not installed device count.
    notInstalledDeviceCount *int32;
    // User name.
    userName *string;
    // User Principal Name.
    userPrincipalName *string;
}
// Instantiates a new userAppInstallStatus and sets the default values.
func NewUserAppInstallStatus()(*UserAppInstallStatus) {
    m := &UserAppInstallStatus{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the app property value. The navigation link to the mobile app.
func (m *UserAppInstallStatus) GetApp()(*MobileApp) {
    if m == nil {
        return nil
    } else {
        return m.app
    }
}
// Gets the deviceStatuses property value. The install state of the app on devices.
func (m *UserAppInstallStatus) GetDeviceStatuses()([]MobileAppInstallStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// Gets the failedDeviceCount property value. Failed Device Count.
func (m *UserAppInstallStatus) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// Gets the installedDeviceCount property value. Installed Device Count.
func (m *UserAppInstallStatus) GetInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedDeviceCount
    }
}
// Gets the notInstalledDeviceCount property value. Not installed device count.
func (m *UserAppInstallStatus) GetNotInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledDeviceCount
    }
}
// Gets the userName property value. User name.
func (m *UserAppInstallStatus) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// Gets the userPrincipalName property value. User Principal Name.
func (m *UserAppInstallStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *UserAppInstallStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["app"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileApp() })
        if err != nil {
            return err
        }
        m.SetApp(val.(*MobileApp))
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppInstallStatus() })
        if err != nil {
            return err
        }
        res := make([]MobileAppInstallStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileAppInstallStatus))
        }
        m.SetDeviceStatuses(res)
        return nil
    }
    res["failedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedDeviceCount(val)
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
    res["notInstalledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotInstalledDeviceCount(val)
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
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
func (m *UserAppInstallStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserAppInstallStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("app", m.GetApp())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
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
        err = writer.WriteInt32Value("notInstalledDeviceCount", m.GetNotInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
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
// Sets the app property value. The navigation link to the mobile app.
// Parameters:
//  - value : Value to set for the app property.
func (m *UserAppInstallStatus) SetApp(value *MobileApp)() {
    m.app = value
}
// Sets the deviceStatuses property value. The install state of the app on devices.
// Parameters:
//  - value : Value to set for the deviceStatuses property.
func (m *UserAppInstallStatus) SetDeviceStatuses(value []MobileAppInstallStatus)() {
    m.deviceStatuses = value
}
// Sets the failedDeviceCount property value. Failed Device Count.
// Parameters:
//  - value : Value to set for the failedDeviceCount property.
func (m *UserAppInstallStatus) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
// Sets the installedDeviceCount property value. Installed Device Count.
// Parameters:
//  - value : Value to set for the installedDeviceCount property.
func (m *UserAppInstallStatus) SetInstalledDeviceCount(value *int32)() {
    m.installedDeviceCount = value
}
// Sets the notInstalledDeviceCount property value. Not installed device count.
// Parameters:
//  - value : Value to set for the notInstalledDeviceCount property.
func (m *UserAppInstallStatus) SetNotInstalledDeviceCount(value *int32)() {
    m.notInstalledDeviceCount = value
}
// Sets the userName property value. User name.
// Parameters:
//  - value : Value to set for the userName property.
func (m *UserAppInstallStatus) SetUserName(value *string)() {
    m.userName = value
}
// Sets the userPrincipalName property value. User Principal Name.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *UserAppInstallStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
