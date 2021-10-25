package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsDevicePerformance struct {
    Entity
    averageBlueScreens *float64;
    averageRestarts *float64;
    blueScreenCount *int32;
    bootScore *int32;
    coreBootTimeInMs *int32;
    coreLoginTimeInMs *int32;
    deviceCount *int64;
    deviceName *string;
    diskType *DiskType;
    groupPolicyBootTimeInMs *int32;
    groupPolicyLoginTimeInMs *int32;
    healthStatus *UserExperienceAnalyticsHealthState;
    loginScore *int32;
    manufacturer *string;
    model *string;
    operatingSystemVersion *string;
    responsiveDesktopTimeInMs *int32;
    restartCount *int32;
}
func NewUserExperienceAnalyticsDevicePerformance()(*UserExperienceAnalyticsDevicePerformance) {
    m := &UserExperienceAnalyticsDevicePerformance{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsDevicePerformance) GetAverageBlueScreens()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageBlueScreens
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetAverageRestarts()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageRestarts
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetBlueScreenCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.blueScreenCount
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetBootScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bootScore
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetCoreBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coreBootTimeInMs
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetCoreLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coreLoginTimeInMs
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetDeviceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetDiskType()(*DiskType) {
    if m == nil {
        return nil
    } else {
        return m.diskType
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetGroupPolicyBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyBootTimeInMs
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetGroupPolicyLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyLoginTimeInMs
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetLoginScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.loginScore
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemVersion
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetResponsiveDesktopTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.responsiveDesktopTimeInMs
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetRestartCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.restartCount
    }
}
func (m *UserExperienceAnalyticsDevicePerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["averageBlueScreens"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAverageBlueScreens(val)
        return nil
    }
    res["averageRestarts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAverageRestarts(val)
        return nil
    }
    res["blueScreenCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBlueScreenCount(val)
        return nil
    }
    res["bootScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBootScore(val)
        return nil
    }
    res["coreBootTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCoreBootTimeInMs(val)
        return nil
    }
    res["coreLoginTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCoreLoginTimeInMs(val)
        return nil
    }
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDeviceCount(val)
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
    res["diskType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiskType)
        if err != nil {
            return err
        }
        cast := val.(DiskType)
        m.SetDiskType(&cast)
        return nil
    }
    res["groupPolicyBootTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetGroupPolicyBootTimeInMs(val)
        return nil
    }
    res["groupPolicyLoginTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetGroupPolicyLoginTimeInMs(val)
        return nil
    }
    res["healthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        cast := val.(UserExperienceAnalyticsHealthState)
        m.SetHealthStatus(&cast)
        return nil
    }
    res["loginScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLoginScore(val)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["operatingSystemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystemVersion(val)
        return nil
    }
    res["responsiveDesktopTimeInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetResponsiveDesktopTimeInMs(val)
        return nil
    }
    res["restartCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRestartCount(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsDevicePerformance) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsDevicePerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("averageBlueScreens", m.GetAverageBlueScreens())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("averageRestarts", m.GetAverageRestarts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("blueScreenCount", m.GetBlueScreenCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("bootScore", m.GetBootScore())
        if err != nil {
            return err
        }
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
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
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
    if m.GetDiskType() != nil {
        cast := m.GetDiskType().String()
        err = writer.WriteStringValue("diskType", &cast)
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
    if m.GetHealthStatus() != nil {
        cast := m.GetHealthStatus().String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("loginScore", m.GetLoginScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
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
    {
        err = writer.WriteInt32Value("restartCount", m.GetRestartCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserExperienceAnalyticsDevicePerformance) SetAverageBlueScreens(value *float64)() {
    m.averageBlueScreens = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetAverageRestarts(value *float64)() {
    m.averageRestarts = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetBlueScreenCount(value *int32)() {
    m.blueScreenCount = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetBootScore(value *int32)() {
    m.bootScore = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetCoreBootTimeInMs(value *int32)() {
    m.coreBootTimeInMs = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetCoreLoginTimeInMs(value *int32)() {
    m.coreLoginTimeInMs = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetDeviceCount(value *int64)() {
    m.deviceCount = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetDiskType(value *DiskType)() {
    m.diskType = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetGroupPolicyBootTimeInMs(value *int32)() {
    m.groupPolicyBootTimeInMs = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetGroupPolicyLoginTimeInMs(value *int32)() {
    m.groupPolicyLoginTimeInMs = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetLoginScore(value *int32)() {
    m.loginScore = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetModel(value *string)() {
    m.model = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetOperatingSystemVersion(value *string)() {
    m.operatingSystemVersion = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetResponsiveDesktopTimeInMs(value *int32)() {
    m.responsiveDesktopTimeInMs = value
}
func (m *UserExperienceAnalyticsDevicePerformance) SetRestartCount(value *int32)() {
    m.restartCount = value
}
