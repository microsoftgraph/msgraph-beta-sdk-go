package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsDeviceStartupHistory 
type UserExperienceAnalyticsDeviceStartupHistory struct {
    Entity
    // The user experience analytics device core boot time in milliseconds.
    coreBootTimeInMs *int32;
    // The user experience analytics device core login time in milliseconds.
    coreLoginTimeInMs *int32;
    // The user experience analytics device id.
    deviceId *string;
    // The user experience analytics device feature update time in milliseconds.
    featureUpdateBootTimeInMs *int32;
    // The User experience analytics Device group policy boot time in milliseconds.
    groupPolicyBootTimeInMs *int32;
    // The User experience analytics Device group policy login time in milliseconds.
    groupPolicyLoginTimeInMs *int32;
    // The user experience analytics device boot record is a feature update.
    isFeatureUpdate *bool;
    // The user experience analytics device first login.
    isFirstLogin *bool;
    // The user experience analytics device boot record's operating system version.
    operatingSystemVersion *string;
    // The user experience analytics responsive desktop time in milliseconds.
    responsiveDesktopTimeInMs *int32;
    // OS restart category. Possible values are: unknown, restartWithUpdate, restartWithoutUpdate, blueScreen, shutdownWithUpdate, shutdownWithoutUpdate, longPowerButtonPress, bootError, update.
    restartCategory *UserExperienceAnalyticsOperatingSystemRestartCategory;
    // OS restart fault bucket. The fault bucket is used to find additional information about a system crash.
    restartFaultBucket *string;
    // OS restart stop code. This shows the bug check code which can be used to look up the blue screen reason.
    restartStopCode *string;
    // The user experience analytics device boot start time.
    startTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The user experience analytics device total boot time in milliseconds.
    totalBootTimeInMs *int32;
    // The user experience analytics device total login time in milliseconds.
    totalLoginTimeInMs *int32;
}
// NewUserExperienceAnalyticsDeviceStartupHistory instantiates a new userExperienceAnalyticsDeviceStartupHistory and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupHistory()(*UserExperienceAnalyticsDeviceStartupHistory) {
    m := &UserExperienceAnalyticsDeviceStartupHistory{
        Entity: *NewEntity(),
    }
    return m
}
// GetCoreBootTimeInMs gets the coreBootTimeInMs property value. The user experience analytics device core boot time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetCoreBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coreBootTimeInMs
    }
}
// GetCoreLoginTimeInMs gets the coreLoginTimeInMs property value. The user experience analytics device core login time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetCoreLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coreLoginTimeInMs
    }
}
// GetDeviceId gets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetFeatureUpdateBootTimeInMs gets the featureUpdateBootTimeInMs property value. The user experience analytics device feature update time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetFeatureUpdateBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdateBootTimeInMs
    }
}
// GetGroupPolicyBootTimeInMs gets the groupPolicyBootTimeInMs property value. The User experience analytics Device group policy boot time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetGroupPolicyBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyBootTimeInMs
    }
}
// GetGroupPolicyLoginTimeInMs gets the groupPolicyLoginTimeInMs property value. The User experience analytics Device group policy login time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetGroupPolicyLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyLoginTimeInMs
    }
}
// GetIsFeatureUpdate gets the isFeatureUpdate property value. The user experience analytics device boot record is a feature update.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetIsFeatureUpdate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFeatureUpdate
    }
}
// GetIsFirstLogin gets the isFirstLogin property value. The user experience analytics device first login.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetIsFirstLogin()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFirstLogin
    }
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. The user experience analytics device boot record's operating system version.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemVersion
    }
}
// GetResponsiveDesktopTimeInMs gets the responsiveDesktopTimeInMs property value. The user experience analytics responsive desktop time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetResponsiveDesktopTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.responsiveDesktopTimeInMs
    }
}
// GetRestartCategory gets the restartCategory property value. OS restart category. Possible values are: unknown, restartWithUpdate, restartWithoutUpdate, blueScreen, shutdownWithUpdate, shutdownWithoutUpdate, longPowerButtonPress, bootError, update.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetRestartCategory()(*UserExperienceAnalyticsOperatingSystemRestartCategory) {
    if m == nil {
        return nil
    } else {
        return m.restartCategory
    }
}
// GetRestartFaultBucket gets the restartFaultBucket property value. OS restart fault bucket. The fault bucket is used to find additional information about a system crash.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetRestartFaultBucket()(*string) {
    if m == nil {
        return nil
    } else {
        return m.restartFaultBucket
    }
}
// GetRestartStopCode gets the restartStopCode property value. OS restart stop code. This shows the bug check code which can be used to look up the blue screen reason.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetRestartStopCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.restartStopCode
    }
}
// GetStartTime gets the startTime property value. The user experience analytics device boot start time.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetStartTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
// GetTotalBootTimeInMs gets the totalBootTimeInMs property value. The user experience analytics device total boot time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetTotalBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalBootTimeInMs
    }
}
// GetTotalLoginTimeInMs gets the totalLoginTimeInMs property value. The user experience analytics device total login time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetTotalLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalLoginTimeInMs
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceStartupHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["coreBootTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreBootTimeInMs(val)
        }
        return nil
    }
    res["coreLoginTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreLoginTimeInMs(val)
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["featureUpdateBootTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdateBootTimeInMs(val)
        }
        return nil
    }
    res["groupPolicyBootTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyBootTimeInMs(val)
        }
        return nil
    }
    res["groupPolicyLoginTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyLoginTimeInMs(val)
        }
        return nil
    }
    res["isFeatureUpdate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFeatureUpdate(val)
        }
        return nil
    }
    res["isFirstLogin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFirstLogin(val)
        }
        return nil
    }
    res["operatingSystemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemVersion(val)
        }
        return nil
    }
    res["responsiveDesktopTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsiveDesktopTimeInMs(val)
        }
        return nil
    }
    res["restartCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsOperatingSystemRestartCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartCategory(val.(*UserExperienceAnalyticsOperatingSystemRestartCategory))
        }
        return nil
    }
    res["restartFaultBucket"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartFaultBucket(val)
        }
        return nil
    }
    res["restartStopCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartStopCode(val)
        }
        return nil
    }
    res["startTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    res["totalBootTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBootTimeInMs(val)
        }
        return nil
    }
    res["totalLoginTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLoginTimeInMs(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsDeviceStartupHistory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceStartupHistory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("coreBootTimeInMs", m.GetCoreBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("coreLoginTimeInMs", m.GetCoreLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("featureUpdateBootTimeInMs", m.GetFeatureUpdateBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupPolicyBootTimeInMs", m.GetGroupPolicyBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupPolicyLoginTimeInMs", m.GetGroupPolicyLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFeatureUpdate", m.GetIsFeatureUpdate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFirstLogin", m.GetIsFirstLogin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("responsiveDesktopTimeInMs", m.GetResponsiveDesktopTimeInMs())
        if err != nil {
            return err
        }
    }
    if m.GetRestartCategory() != nil {
        cast := (*m.GetRestartCategory()).String()
        err = writer.WriteStringValue("restartCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("restartFaultBucket", m.GetRestartFaultBucket())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("restartStopCode", m.GetRestartStopCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startTime", m.GetStartTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalBootTimeInMs", m.GetTotalBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLoginTimeInMs", m.GetTotalLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCoreBootTimeInMs sets the coreBootTimeInMs property value. The user experience analytics device core boot time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetCoreBootTimeInMs(value *int32)() {
    if m != nil {
        m.coreBootTimeInMs = value
    }
}
// SetCoreLoginTimeInMs sets the coreLoginTimeInMs property value. The user experience analytics device core login time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetCoreLoginTimeInMs(value *int32)() {
    if m != nil {
        m.coreLoginTimeInMs = value
    }
}
// SetDeviceId sets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetFeatureUpdateBootTimeInMs sets the featureUpdateBootTimeInMs property value. The user experience analytics device feature update time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetFeatureUpdateBootTimeInMs(value *int32)() {
    if m != nil {
        m.featureUpdateBootTimeInMs = value
    }
}
// SetGroupPolicyBootTimeInMs sets the groupPolicyBootTimeInMs property value. The User experience analytics Device group policy boot time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetGroupPolicyBootTimeInMs(value *int32)() {
    if m != nil {
        m.groupPolicyBootTimeInMs = value
    }
}
// SetGroupPolicyLoginTimeInMs sets the groupPolicyLoginTimeInMs property value. The User experience analytics Device group policy login time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetGroupPolicyLoginTimeInMs(value *int32)() {
    if m != nil {
        m.groupPolicyLoginTimeInMs = value
    }
}
// SetIsFeatureUpdate sets the isFeatureUpdate property value. The user experience analytics device boot record is a feature update.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetIsFeatureUpdate(value *bool)() {
    if m != nil {
        m.isFeatureUpdate = value
    }
}
// SetIsFirstLogin sets the isFirstLogin property value. The user experience analytics device first login.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetIsFirstLogin(value *bool)() {
    if m != nil {
        m.isFirstLogin = value
    }
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. The user experience analytics device boot record's operating system version.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetOperatingSystemVersion(value *string)() {
    if m != nil {
        m.operatingSystemVersion = value
    }
}
// SetResponsiveDesktopTimeInMs sets the responsiveDesktopTimeInMs property value. The user experience analytics responsive desktop time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetResponsiveDesktopTimeInMs(value *int32)() {
    if m != nil {
        m.responsiveDesktopTimeInMs = value
    }
}
// SetRestartCategory sets the restartCategory property value. OS restart category. Possible values are: unknown, restartWithUpdate, restartWithoutUpdate, blueScreen, shutdownWithUpdate, shutdownWithoutUpdate, longPowerButtonPress, bootError, update.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetRestartCategory(value *UserExperienceAnalyticsOperatingSystemRestartCategory)() {
    if m != nil {
        m.restartCategory = value
    }
}
// SetRestartFaultBucket sets the restartFaultBucket property value. OS restart fault bucket. The fault bucket is used to find additional information about a system crash.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetRestartFaultBucket(value *string)() {
    if m != nil {
        m.restartFaultBucket = value
    }
}
// SetRestartStopCode sets the restartStopCode property value. OS restart stop code. This shows the bug check code which can be used to look up the blue screen reason.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetRestartStopCode(value *string)() {
    if m != nil {
        m.restartStopCode = value
    }
}
// SetStartTime sets the startTime property value. The user experience analytics device boot start time.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetStartTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startTime = value
    }
}
// SetTotalBootTimeInMs sets the totalBootTimeInMs property value. The user experience analytics device total boot time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetTotalBootTimeInMs(value *int32)() {
    if m != nil {
        m.totalBootTimeInMs = value
    }
}
// SetTotalLoginTimeInMs sets the totalLoginTimeInMs property value. The user experience analytics device total login time in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupHistory) SetTotalLoginTimeInMs(value *int32)() {
    if m != nil {
        m.totalLoginTimeInMs = value
    }
}
