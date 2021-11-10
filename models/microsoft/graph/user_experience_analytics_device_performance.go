package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsDevicePerformance struct {
    Entity
    // Average (mean) number of Blue Screens per device in the last 14 days. Valid values 0 to 9999999
    averageBlueScreens *float64;
    // Average (mean) number of Restarts per device in the last 14 days. Valid values 0 to 9999999
    averageRestarts *float64;
    // Number of Blue Screens in the last 14 days. Valid values 0 to 9999999
    blueScreenCount *int32;
    // The user experience analytics device boot score.
    bootScore *int32;
    // The user experience analytics device core boot time in milliseconds.
    coreBootTimeInMs *int32;
    // The user experience analytics device core login time in milliseconds.
    coreLoginTimeInMs *int32;
    // User experience analytics summarized device count.
    deviceCount *int64;
    // The user experience analytics device name.
    deviceName *string;
    // The user experience analytics device disk type. Possible values are: unkown, hdd, ssd.
    diskType *DiskType;
    // The user experience analytics device group policy boot time in milliseconds.
    groupPolicyBootTimeInMs *int32;
    // The user experience analytics device group policy login time in milliseconds.
    groupPolicyLoginTimeInMs *int32;
    // The health state of the user experience analytics device. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
    healthStatus *UserExperienceAnalyticsHealthState;
    // The user experience analytics device login score.
    loginScore *int32;
    // The user experience analytics device manufacturer.
    manufacturer *string;
    // The user experience analytics device model.
    model *string;
    // The user experience analytics model level startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    modelStartupPerformanceScore *float64;
    // The user experience analytics device Operating System version.
    operatingSystemVersion *string;
    // The user experience analytics responsive desktop time in milliseconds.
    responsiveDesktopTimeInMs *int32;
    // Number of Restarts in the last 14 days. Valid values 0 to 9999999
    restartCount *int32;
    // The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    startupPerformanceScore *float64;
}
// Instantiates a new userExperienceAnalyticsDevicePerformance and sets the default values.
func NewUserExperienceAnalyticsDevicePerformance()(*UserExperienceAnalyticsDevicePerformance) {
    m := &UserExperienceAnalyticsDevicePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the averageBlueScreens property value. Average (mean) number of Blue Screens per device in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) GetAverageBlueScreens()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageBlueScreens
    }
}
// Gets the averageRestarts property value. Average (mean) number of Restarts per device in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) GetAverageRestarts()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageRestarts
    }
}
// Gets the blueScreenCount property value. Number of Blue Screens in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) GetBlueScreenCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.blueScreenCount
    }
}
// Gets the bootScore property value. The user experience analytics device boot score.
func (m *UserExperienceAnalyticsDevicePerformance) GetBootScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bootScore
    }
}
// Gets the coreBootTimeInMs property value. The user experience analytics device core boot time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetCoreBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coreBootTimeInMs
    }
}
// Gets the coreLoginTimeInMs property value. The user experience analytics device core login time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetCoreLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coreLoginTimeInMs
    }
}
// Gets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsDevicePerformance) GetDeviceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// Gets the deviceName property value. The user experience analytics device name.
func (m *UserExperienceAnalyticsDevicePerformance) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the diskType property value. The user experience analytics device disk type. Possible values are: unkown, hdd, ssd.
func (m *UserExperienceAnalyticsDevicePerformance) GetDiskType()(*DiskType) {
    if m == nil {
        return nil
    } else {
        return m.diskType
    }
}
// Gets the groupPolicyBootTimeInMs property value. The user experience analytics device group policy boot time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetGroupPolicyBootTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyBootTimeInMs
    }
}
// Gets the groupPolicyLoginTimeInMs property value. The user experience analytics device group policy login time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetGroupPolicyLoginTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyLoginTimeInMs
    }
}
// Gets the healthStatus property value. The health state of the user experience analytics device. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
func (m *UserExperienceAnalyticsDevicePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
// Gets the loginScore property value. The user experience analytics device login score.
func (m *UserExperienceAnalyticsDevicePerformance) GetLoginScore()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.loginScore
    }
}
// Gets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsDevicePerformance) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsDevicePerformance) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the modelStartupPerformanceScore property value. The user experience analytics model level startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDevicePerformance) GetModelStartupPerformanceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.modelStartupPerformanceScore
    }
}
// Gets the operatingSystemVersion property value. The user experience analytics device Operating System version.
func (m *UserExperienceAnalyticsDevicePerformance) GetOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemVersion
    }
}
// Gets the responsiveDesktopTimeInMs property value. The user experience analytics responsive desktop time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) GetResponsiveDesktopTimeInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.responsiveDesktopTimeInMs
    }
}
// Gets the restartCount property value. Number of Restarts in the last 14 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) GetRestartCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.restartCount
    }
}
// Gets the startupPerformanceScore property value. The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDevicePerformance) GetStartupPerformanceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.startupPerformanceScore
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsDevicePerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["averageBlueScreens"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageBlueScreens(val)
        }
        return nil
    }
    res["averageRestarts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageRestarts(val)
        }
        return nil
    }
    res["blueScreenCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueScreenCount(val)
        }
        return nil
    }
    res["bootScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootScore(val)
        }
        return nil
    }
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
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
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
    res["diskType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiskType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DiskType)
            m.SetDiskType(&cast)
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
    res["healthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(UserExperienceAnalyticsHealthState)
            m.SetHealthStatus(&cast)
        }
        return nil
    }
    res["loginScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginScore(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["modelStartupPerformanceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelStartupPerformanceScore(val)
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
    res["restartCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartCount(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsDevicePerformance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err = writer.WriteFloat64Value("modelStartupPerformanceScore", m.GetModelStartupPerformanceScore())
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
    {
        err = writer.WriteFloat64Value("startupPerformanceScore", m.GetStartupPerformanceScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the averageBlueScreens property value. Average (mean) number of Blue Screens per device in the last 14 days. Valid values 0 to 9999999
// Parameters:
//  - value : Value to set for the averageBlueScreens property.
func (m *UserExperienceAnalyticsDevicePerformance) SetAverageBlueScreens(value *float64)() {
    m.averageBlueScreens = value
}
// Sets the averageRestarts property value. Average (mean) number of Restarts per device in the last 14 days. Valid values 0 to 9999999
// Parameters:
//  - value : Value to set for the averageRestarts property.
func (m *UserExperienceAnalyticsDevicePerformance) SetAverageRestarts(value *float64)() {
    m.averageRestarts = value
}
// Sets the blueScreenCount property value. Number of Blue Screens in the last 14 days. Valid values 0 to 9999999
// Parameters:
//  - value : Value to set for the blueScreenCount property.
func (m *UserExperienceAnalyticsDevicePerformance) SetBlueScreenCount(value *int32)() {
    m.blueScreenCount = value
}
// Sets the bootScore property value. The user experience analytics device boot score.
// Parameters:
//  - value : Value to set for the bootScore property.
func (m *UserExperienceAnalyticsDevicePerformance) SetBootScore(value *int32)() {
    m.bootScore = value
}
// Sets the coreBootTimeInMs property value. The user experience analytics device core boot time in milliseconds.
// Parameters:
//  - value : Value to set for the coreBootTimeInMs property.
func (m *UserExperienceAnalyticsDevicePerformance) SetCoreBootTimeInMs(value *int32)() {
    m.coreBootTimeInMs = value
}
// Sets the coreLoginTimeInMs property value. The user experience analytics device core login time in milliseconds.
// Parameters:
//  - value : Value to set for the coreLoginTimeInMs property.
func (m *UserExperienceAnalyticsDevicePerformance) SetCoreLoginTimeInMs(value *int32)() {
    m.coreLoginTimeInMs = value
}
// Sets the deviceCount property value. User experience analytics summarized device count.
// Parameters:
//  - value : Value to set for the deviceCount property.
func (m *UserExperienceAnalyticsDevicePerformance) SetDeviceCount(value *int64)() {
    m.deviceCount = value
}
// Sets the deviceName property value. The user experience analytics device name.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *UserExperienceAnalyticsDevicePerformance) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the diskType property value. The user experience analytics device disk type. Possible values are: unkown, hdd, ssd.
// Parameters:
//  - value : Value to set for the diskType property.
func (m *UserExperienceAnalyticsDevicePerformance) SetDiskType(value *DiskType)() {
    m.diskType = value
}
// Sets the groupPolicyBootTimeInMs property value. The user experience analytics device group policy boot time in milliseconds.
// Parameters:
//  - value : Value to set for the groupPolicyBootTimeInMs property.
func (m *UserExperienceAnalyticsDevicePerformance) SetGroupPolicyBootTimeInMs(value *int32)() {
    m.groupPolicyBootTimeInMs = value
}
// Sets the groupPolicyLoginTimeInMs property value. The user experience analytics device group policy login time in milliseconds.
// Parameters:
//  - value : Value to set for the groupPolicyLoginTimeInMs property.
func (m *UserExperienceAnalyticsDevicePerformance) SetGroupPolicyLoginTimeInMs(value *int32)() {
    m.groupPolicyLoginTimeInMs = value
}
// Sets the healthStatus property value. The health state of the user experience analytics device. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
// Parameters:
//  - value : Value to set for the healthStatus property.
func (m *UserExperienceAnalyticsDevicePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// Sets the loginScore property value. The user experience analytics device login score.
// Parameters:
//  - value : Value to set for the loginScore property.
func (m *UserExperienceAnalyticsDevicePerformance) SetLoginScore(value *int32)() {
    m.loginScore = value
}
// Sets the manufacturer property value. The user experience analytics device manufacturer.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *UserExperienceAnalyticsDevicePerformance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. The user experience analytics device model.
// Parameters:
//  - value : Value to set for the model property.
func (m *UserExperienceAnalyticsDevicePerformance) SetModel(value *string)() {
    m.model = value
}
// Sets the modelStartupPerformanceScore property value. The user experience analytics model level startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the modelStartupPerformanceScore property.
func (m *UserExperienceAnalyticsDevicePerformance) SetModelStartupPerformanceScore(value *float64)() {
    m.modelStartupPerformanceScore = value
}
// Sets the operatingSystemVersion property value. The user experience analytics device Operating System version.
// Parameters:
//  - value : Value to set for the operatingSystemVersion property.
func (m *UserExperienceAnalyticsDevicePerformance) SetOperatingSystemVersion(value *string)() {
    m.operatingSystemVersion = value
}
// Sets the responsiveDesktopTimeInMs property value. The user experience analytics responsive desktop time in milliseconds.
// Parameters:
//  - value : Value to set for the responsiveDesktopTimeInMs property.
func (m *UserExperienceAnalyticsDevicePerformance) SetResponsiveDesktopTimeInMs(value *int32)() {
    m.responsiveDesktopTimeInMs = value
}
// Sets the restartCount property value. Number of Restarts in the last 14 days. Valid values 0 to 9999999
// Parameters:
//  - value : Value to set for the restartCount property.
func (m *UserExperienceAnalyticsDevicePerformance) SetRestartCount(value *int32)() {
    m.restartCount = value
}
// Sets the startupPerformanceScore property value. The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the startupPerformanceScore property.
func (m *UserExperienceAnalyticsDevicePerformance) SetStartupPerformanceScore(value *float64)() {
    m.startupPerformanceScore = value
}
