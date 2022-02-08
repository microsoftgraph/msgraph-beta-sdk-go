package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RestrictedAppsViolation 
type RestrictedAppsViolation struct {
    Entity
    // Device configuration profile unique identifier, must be Guid
    deviceConfigurationId *string;
    // Device configuration profile name
    deviceConfigurationName *string;
    // Device name
    deviceName *string;
    // Managed device unique identifier, must be Guid
    managedDeviceId *string;
    // Platform type. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
    platformType *PolicyPlatformType;
    // List of violated restricted apps
    restrictedApps []ManagedDeviceReportedApp;
    // Restricted apps state. Possible values are: prohibitedApps, notApprovedApps.
    restrictedAppsState *RestrictedAppsState;
    // User unique identifier, must be Guid
    userId *string;
    // User name
    userName *string;
}
// NewRestrictedAppsViolation instantiates a new restrictedAppsViolation and sets the default values.
func NewRestrictedAppsViolation()(*RestrictedAppsViolation) {
    m := &RestrictedAppsViolation{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeviceConfigurationId gets the deviceConfigurationId property value. Device configuration profile unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetDeviceConfigurationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationId
    }
}
// GetDeviceConfigurationName gets the deviceConfigurationName property value. Device configuration profile name
func (m *RestrictedAppsViolation) GetDeviceConfigurationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationName
    }
}
// GetDeviceName gets the deviceName property value. Device name
func (m *RestrictedAppsViolation) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetManagedDeviceId gets the managedDeviceId property value. Managed device unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// GetPlatformType gets the platformType property value. Platform type. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
func (m *RestrictedAppsViolation) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// GetRestrictedApps gets the restrictedApps property value. List of violated restricted apps
func (m *RestrictedAppsViolation) GetRestrictedApps()([]ManagedDeviceReportedApp) {
    if m == nil {
        return nil
    } else {
        return m.restrictedApps
    }
}
// GetRestrictedAppsState gets the restrictedAppsState property value. Restricted apps state. Possible values are: prohibitedApps, notApprovedApps.
func (m *RestrictedAppsViolation) GetRestrictedAppsState()(*RestrictedAppsState) {
    if m == nil {
        return nil
    } else {
        return m.restrictedAppsState
    }
}
// GetUserId gets the userId property value. User unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserName gets the userName property value. User name
func (m *RestrictedAppsViolation) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RestrictedAppsViolation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfigurationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationId(val)
        }
        return nil
    }
    res["deviceConfigurationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationName(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    res["restrictedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceReportedApp() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceReportedApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDeviceReportedApp))
            }
            m.SetRestrictedApps(res)
        }
        return nil
    }
    res["restrictedAppsState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRestrictedAppsState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictedAppsState(val.(*RestrictedAppsState))
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
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
    return res
}
func (m *RestrictedAppsViolation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRestrictedApps() != nil {
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
        cast := (*m.GetRestrictedAppsState()).String()
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
// SetDeviceConfigurationId sets the deviceConfigurationId property value. Device configuration profile unique identifier, must be Guid
func (m *RestrictedAppsViolation) SetDeviceConfigurationId(value *string)() {
    if m != nil {
        m.deviceConfigurationId = value
    }
}
// SetDeviceConfigurationName sets the deviceConfigurationName property value. Device configuration profile name
func (m *RestrictedAppsViolation) SetDeviceConfigurationName(value *string)() {
    if m != nil {
        m.deviceConfigurationName = value
    }
}
// SetDeviceName sets the deviceName property value. Device name
func (m *RestrictedAppsViolation) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Managed device unique identifier, must be Guid
func (m *RestrictedAppsViolation) SetManagedDeviceId(value *string)() {
    if m != nil {
        m.managedDeviceId = value
    }
}
// SetPlatformType sets the platformType property value. Platform type. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, androidAOSP, all.
func (m *RestrictedAppsViolation) SetPlatformType(value *PolicyPlatformType)() {
    if m != nil {
        m.platformType = value
    }
}
// SetRestrictedApps sets the restrictedApps property value. List of violated restricted apps
func (m *RestrictedAppsViolation) SetRestrictedApps(value []ManagedDeviceReportedApp)() {
    if m != nil {
        m.restrictedApps = value
    }
}
// SetRestrictedAppsState sets the restrictedAppsState property value. Restricted apps state. Possible values are: prohibitedApps, notApprovedApps.
func (m *RestrictedAppsViolation) SetRestrictedAppsState(value *RestrictedAppsState)() {
    if m != nil {
        m.restrictedAppsState = value
    }
}
// SetUserId sets the userId property value. User unique identifier, must be Guid
func (m *RestrictedAppsViolation) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserName sets the userName property value. User name
func (m *RestrictedAppsViolation) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
