package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RestrictedAppsViolation struct {
    Entity
    deviceConfigurationId *string;
    deviceConfigurationName *string;
    deviceName *string;
    managedDeviceId *string;
    platformType *PolicyPlatformType;
    restrictedApps []ManagedDeviceReportedApp;
    restrictedAppsState *RestrictedAppsState;
    userId *string;
    userName *string;
}
func NewRestrictedAppsViolation()(*RestrictedAppsViolation) {
    m := &RestrictedAppsViolation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *RestrictedAppsViolation) GetDeviceConfigurationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationId
    }
}
func (m *RestrictedAppsViolation) GetDeviceConfigurationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationName
    }
}
func (m *RestrictedAppsViolation) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *RestrictedAppsViolation) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
func (m *RestrictedAppsViolation) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
func (m *RestrictedAppsViolation) GetRestrictedApps()([]ManagedDeviceReportedApp) {
    if m == nil {
        return nil
    } else {
        return m.restrictedApps
    }
}
func (m *RestrictedAppsViolation) GetRestrictedAppsState()(*RestrictedAppsState) {
    if m == nil {
        return nil
    } else {
        return m.restrictedAppsState
    }
}
func (m *RestrictedAppsViolation) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *RestrictedAppsViolation) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
func (m *RestrictedAppsViolation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfigurationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceConfigurationId(val)
        return nil
    }
    res["deviceConfigurationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceConfigurationName(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceId(val)
        return nil
    }
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        cast := val.(PolicyPlatformType)
        m.SetPlatformType(&cast)
        return nil
    }
    res["restrictedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceReportedApp() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceReportedApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceReportedApp))
        }
        m.SetRestrictedApps(res)
        return nil
    }
    res["restrictedAppsState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRestrictedAppsState)
        if err != nil {
            return err
        }
        cast := val.(RestrictedAppsState)
        m.SetRestrictedAppsState(&cast)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
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
    return res
}
func (m *RestrictedAppsViolation) IsNil()(bool) {
    return m == nil
}
func (m *RestrictedAppsViolation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceConfigurationId", m.GetDeviceConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceConfigurationName", m.GetDeviceConfigurationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := m.GetPlatformType().String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRestrictedApps()))
        for i, v := range m.GetRestrictedApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("restrictedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRestrictedAppsState() != nil {
        cast := m.GetRestrictedAppsState().String()
        err = writer.WriteStringValue("restrictedAppsState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
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
    return nil
}
func (m *RestrictedAppsViolation) SetDeviceConfigurationId(value *string)() {
    m.deviceConfigurationId = value
}
func (m *RestrictedAppsViolation) SetDeviceConfigurationName(value *string)() {
    m.deviceConfigurationName = value
}
func (m *RestrictedAppsViolation) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *RestrictedAppsViolation) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
func (m *RestrictedAppsViolation) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
func (m *RestrictedAppsViolation) SetRestrictedApps(value []ManagedDeviceReportedApp)() {
    m.restrictedApps = value
}
func (m *RestrictedAppsViolation) SetRestrictedAppsState(value *RestrictedAppsState)() {
    m.restrictedAppsState = value
}
func (m *RestrictedAppsViolation) SetUserId(value *string)() {
    m.userId = value
}
func (m *RestrictedAppsViolation) SetUserName(value *string)() {
    m.userName = value
}
