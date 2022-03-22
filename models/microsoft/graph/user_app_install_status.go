package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserAppInstallStatus 
type UserAppInstallStatus struct {
    Entity
    // The navigation link to the mobile app.
    app MobileAppable;
    // The install state of the app on devices.
    deviceStatuses []MobileAppInstallStatusable;
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
// NewUserAppInstallStatus instantiates a new userAppInstallStatus and sets the default values.
func NewUserAppInstallStatus()(*UserAppInstallStatus) {
    m := &UserAppInstallStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserAppInstallStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserAppInstallStatusFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserAppInstallStatus(), nil
}
// GetApp gets the app property value. The navigation link to the mobile app.
func (m *UserAppInstallStatus) GetApp()(MobileAppable) {
    if m == nil {
        return nil
    } else {
        return m.app
    }
}
// GetDeviceStatuses gets the deviceStatuses property value. The install state of the app on devices.
func (m *UserAppInstallStatus) GetDeviceStatuses()([]MobileAppInstallStatusable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Failed Device Count.
func (m *UserAppInstallStatus) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAppInstallStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["app"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApp(val.(MobileAppable))
        }
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppInstallStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppInstallStatusable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppInstallStatusable)
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
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
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetInstalledDeviceCount gets the installedDeviceCount property value. Installed Device Count.
func (m *UserAppInstallStatus) GetInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedDeviceCount
    }
}
// GetNotInstalledDeviceCount gets the notInstalledDeviceCount property value. Not installed device count.
func (m *UserAppInstallStatus) GetNotInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledDeviceCount
    }
}
// GetUserName gets the userName property value. User name.
func (m *UserAppInstallStatus) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name.
func (m *UserAppInstallStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Serialize serializes information the current object
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
    if m.GetDeviceStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetApp sets the app property value. The navigation link to the mobile app.
func (m *UserAppInstallStatus) SetApp(value MobileAppable)() {
    if m != nil {
        m.app = value
    }
}
// SetDeviceStatuses sets the deviceStatuses property value. The install state of the app on devices.
func (m *UserAppInstallStatus) SetDeviceStatuses(value []MobileAppInstallStatusable)() {
    if m != nil {
        m.deviceStatuses = value
    }
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Failed Device Count.
func (m *UserAppInstallStatus) SetFailedDeviceCount(value *int32)() {
    if m != nil {
        m.failedDeviceCount = value
    }
}
// SetInstalledDeviceCount sets the installedDeviceCount property value. Installed Device Count.
func (m *UserAppInstallStatus) SetInstalledDeviceCount(value *int32)() {
    if m != nil {
        m.installedDeviceCount = value
    }
}
// SetNotInstalledDeviceCount sets the notInstalledDeviceCount property value. Not installed device count.
func (m *UserAppInstallStatus) SetNotInstalledDeviceCount(value *int32)() {
    if m != nil {
        m.notInstalledDeviceCount = value
    }
}
// SetUserName sets the userName property value. User name.
func (m *UserAppInstallStatus) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name.
func (m *UserAppInstallStatus) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
