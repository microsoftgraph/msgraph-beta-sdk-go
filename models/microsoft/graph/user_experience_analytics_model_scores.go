package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsModelScores struct {
    Entity
    // The user experience analytics model app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    appReliabilityScore *float64;
    // The user experience analytics model score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    endpointAnalyticsScore *float64;
    // The health state of the user experience analytics model. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
    healthStatus *UserExperienceAnalyticsHealthState;
    // A unique identifier of the user experience analytics model scores: device manufacturer.
    manufacturer *string;
    // A unique identifier of the user experience analytics model scores: device model.
    model *string;
    // The user experience analytics model device count. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
    modelDeviceCount *int64;
    // The user experience analytics model startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    startupPerformanceScore *float64;
}
// Instantiates a new userExperienceAnalyticsModelScores and sets the default values.
func NewUserExperienceAnalyticsModelScores()(*UserExperienceAnalyticsModelScores) {
    m := &UserExperienceAnalyticsModelScores{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appReliabilityScore property value. The user experience analytics model app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) GetAppReliabilityScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.appReliabilityScore
    }
}
// Gets the endpointAnalyticsScore property value. The user experience analytics model score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) GetEndpointAnalyticsScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.endpointAnalyticsScore
    }
}
// Gets the healthStatus property value. The health state of the user experience analytics model. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
func (m *UserExperienceAnalyticsModelScores) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
// Gets the manufacturer property value. A unique identifier of the user experience analytics model scores: device manufacturer.
func (m *UserExperienceAnalyticsModelScores) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. A unique identifier of the user experience analytics model scores: device model.
func (m *UserExperienceAnalyticsModelScores) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the modelDeviceCount property value. The user experience analytics model device count. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
func (m *UserExperienceAnalyticsModelScores) GetModelDeviceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.modelDeviceCount
    }
}
// Gets the startupPerformanceScore property value. The user experience analytics model startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) GetStartupPerformanceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.startupPerformanceScore
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsModelScores) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appReliabilityScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAppReliabilityScore(val)
        return nil
    }
    res["endpointAnalyticsScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetEndpointAnalyticsScore(val)
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
    res["modelDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetModelDeviceCount(val)
        return nil
    }
    res["startupPerformanceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetStartupPerformanceScore(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsModelScores) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsModelScores) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("appReliabilityScore", m.GetAppReliabilityScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("endpointAnalyticsScore", m.GetEndpointAnalyticsScore())
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
        err = writer.WriteInt64Value("modelDeviceCount", m.GetModelDeviceCount())
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
// Sets the appReliabilityScore property value. The user experience analytics model app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the appReliabilityScore property.
func (m *UserExperienceAnalyticsModelScores) SetAppReliabilityScore(value *float64)() {
    m.appReliabilityScore = value
}
// Sets the endpointAnalyticsScore property value. The user experience analytics model score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the endpointAnalyticsScore property.
func (m *UserExperienceAnalyticsModelScores) SetEndpointAnalyticsScore(value *float64)() {
    m.endpointAnalyticsScore = value
}
// Sets the healthStatus property value. The health state of the user experience analytics model. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
// Parameters:
//  - value : Value to set for the healthStatus property.
func (m *UserExperienceAnalyticsModelScores) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// Sets the manufacturer property value. A unique identifier of the user experience analytics model scores: device manufacturer.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *UserExperienceAnalyticsModelScores) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. A unique identifier of the user experience analytics model scores: device model.
// Parameters:
//  - value : Value to set for the model property.
func (m *UserExperienceAnalyticsModelScores) SetModel(value *string)() {
    m.model = value
}
// Sets the modelDeviceCount property value. The user experience analytics model device count. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
// Parameters:
//  - value : Value to set for the modelDeviceCount property.
func (m *UserExperienceAnalyticsModelScores) SetModelDeviceCount(value *int64)() {
    m.modelDeviceCount = value
}
// Sets the startupPerformanceScore property value. The user experience analytics model startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the startupPerformanceScore property.
func (m *UserExperienceAnalyticsModelScores) SetStartupPerformanceScore(value *float64)() {
    m.startupPerformanceScore = value
}
