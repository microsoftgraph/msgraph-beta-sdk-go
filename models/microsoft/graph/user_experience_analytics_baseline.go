package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsBaseline struct {
    Entity
    // The user experience analytics app health metrics.
    appHealthMetrics *UserExperienceAnalyticsCategory;
    // The user experience analytics battery health metrics.
    batteryHealthMetrics *UserExperienceAnalyticsCategory;
    // The user experience analytics best practices metrics.
    bestPracticesMetrics *UserExperienceAnalyticsCategory;
    // The date the custom baseline was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The user experience analytics device boot performance metrics.
    deviceBootPerformanceMetrics *UserExperienceAnalyticsCategory;
    // The name of the user experience analytics baseline.
    displayName *string;
    // Signifies if the current baseline is the commercial median baseline or a custom baseline.
    isBuiltIn *bool;
    // The user experience analytics reboot analytics metrics.
    rebootAnalyticsMetrics *UserExperienceAnalyticsCategory;
    // The user experience analytics resource performance metrics.
    resourcePerformanceMetrics *UserExperienceAnalyticsCategory;
    // The user experience analytics work from anywhere metrics.
    workFromAnywhereMetrics *UserExperienceAnalyticsCategory;
}
// Instantiates a new userExperienceAnalyticsBaseline and sets the default values.
func NewUserExperienceAnalyticsBaseline()(*UserExperienceAnalyticsBaseline) {
    m := &UserExperienceAnalyticsBaseline{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appHealthMetrics property value. The user experience analytics app health metrics.
func (m *UserExperienceAnalyticsBaseline) GetAppHealthMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.appHealthMetrics
    }
}
// Gets the batteryHealthMetrics property value. The user experience analytics battery health metrics.
func (m *UserExperienceAnalyticsBaseline) GetBatteryHealthMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.batteryHealthMetrics
    }
}
// Gets the bestPracticesMetrics property value. The user experience analytics best practices metrics.
func (m *UserExperienceAnalyticsBaseline) GetBestPracticesMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.bestPracticesMetrics
    }
}
// Gets the createdDateTime property value. The date the custom baseline was created.
func (m *UserExperienceAnalyticsBaseline) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deviceBootPerformanceMetrics property value. The user experience analytics device boot performance metrics.
func (m *UserExperienceAnalyticsBaseline) GetDeviceBootPerformanceMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceBootPerformanceMetrics
    }
}
// Gets the displayName property value. The name of the user experience analytics baseline.
func (m *UserExperienceAnalyticsBaseline) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isBuiltIn property value. Signifies if the current baseline is the commercial median baseline or a custom baseline.
func (m *UserExperienceAnalyticsBaseline) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// Gets the rebootAnalyticsMetrics property value. The user experience analytics reboot analytics metrics.
func (m *UserExperienceAnalyticsBaseline) GetRebootAnalyticsMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.rebootAnalyticsMetrics
    }
}
// Gets the resourcePerformanceMetrics property value. The user experience analytics resource performance metrics.
func (m *UserExperienceAnalyticsBaseline) GetResourcePerformanceMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.resourcePerformanceMetrics
    }
}
// Gets the workFromAnywhereMetrics property value. The user experience analytics work from anywhere metrics.
func (m *UserExperienceAnalyticsBaseline) GetWorkFromAnywhereMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.workFromAnywhereMetrics
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsBaseline) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appHealthMetrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        m.SetAppHealthMetrics(val.(*UserExperienceAnalyticsCategory))
        return nil
    }
    res["batteryHealthMetrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        m.SetBatteryHealthMetrics(val.(*UserExperienceAnalyticsCategory))
        return nil
    }
    res["bestPracticesMetrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        m.SetBestPracticesMetrics(val.(*UserExperienceAnalyticsCategory))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["deviceBootPerformanceMetrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        m.SetDeviceBootPerformanceMetrics(val.(*UserExperienceAnalyticsCategory))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsBuiltIn(val)
        return nil
    }
    res["rebootAnalyticsMetrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        m.SetRebootAnalyticsMetrics(val.(*UserExperienceAnalyticsCategory))
        return nil
    }
    res["resourcePerformanceMetrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        m.SetResourcePerformanceMetrics(val.(*UserExperienceAnalyticsCategory))
        return nil
    }
    res["workFromAnywhereMetrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        m.SetWorkFromAnywhereMetrics(val.(*UserExperienceAnalyticsCategory))
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBaseline) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsBaseline) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appHealthMetrics", m.GetAppHealthMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("batteryHealthMetrics", m.GetBatteryHealthMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bestPracticesMetrics", m.GetBestPracticesMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceBootPerformanceMetrics", m.GetDeviceBootPerformanceMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rebootAnalyticsMetrics", m.GetRebootAnalyticsMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourcePerformanceMetrics", m.GetResourcePerformanceMetrics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("workFromAnywhereMetrics", m.GetWorkFromAnywhereMetrics())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appHealthMetrics property value. The user experience analytics app health metrics.
// Parameters:
//  - value : Value to set for the appHealthMetrics property.
func (m *UserExperienceAnalyticsBaseline) SetAppHealthMetrics(value *UserExperienceAnalyticsCategory)() {
    m.appHealthMetrics = value
}
// Sets the batteryHealthMetrics property value. The user experience analytics battery health metrics.
// Parameters:
//  - value : Value to set for the batteryHealthMetrics property.
func (m *UserExperienceAnalyticsBaseline) SetBatteryHealthMetrics(value *UserExperienceAnalyticsCategory)() {
    m.batteryHealthMetrics = value
}
// Sets the bestPracticesMetrics property value. The user experience analytics best practices metrics.
// Parameters:
//  - value : Value to set for the bestPracticesMetrics property.
func (m *UserExperienceAnalyticsBaseline) SetBestPracticesMetrics(value *UserExperienceAnalyticsCategory)() {
    m.bestPracticesMetrics = value
}
// Sets the createdDateTime property value. The date the custom baseline was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *UserExperienceAnalyticsBaseline) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deviceBootPerformanceMetrics property value. The user experience analytics device boot performance metrics.
// Parameters:
//  - value : Value to set for the deviceBootPerformanceMetrics property.
func (m *UserExperienceAnalyticsBaseline) SetDeviceBootPerformanceMetrics(value *UserExperienceAnalyticsCategory)() {
    m.deviceBootPerformanceMetrics = value
}
// Sets the displayName property value. The name of the user experience analytics baseline.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *UserExperienceAnalyticsBaseline) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isBuiltIn property value. Signifies if the current baseline is the commercial median baseline or a custom baseline.
// Parameters:
//  - value : Value to set for the isBuiltIn property.
func (m *UserExperienceAnalyticsBaseline) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// Sets the rebootAnalyticsMetrics property value. The user experience analytics reboot analytics metrics.
// Parameters:
//  - value : Value to set for the rebootAnalyticsMetrics property.
func (m *UserExperienceAnalyticsBaseline) SetRebootAnalyticsMetrics(value *UserExperienceAnalyticsCategory)() {
    m.rebootAnalyticsMetrics = value
}
// Sets the resourcePerformanceMetrics property value. The user experience analytics resource performance metrics.
// Parameters:
//  - value : Value to set for the resourcePerformanceMetrics property.
func (m *UserExperienceAnalyticsBaseline) SetResourcePerformanceMetrics(value *UserExperienceAnalyticsCategory)() {
    m.resourcePerformanceMetrics = value
}
// Sets the workFromAnywhereMetrics property value. The user experience analytics work from anywhere metrics.
// Parameters:
//  - value : Value to set for the workFromAnywhereMetrics property.
func (m *UserExperienceAnalyticsBaseline) SetWorkFromAnywhereMetrics(value *UserExperienceAnalyticsCategory)() {
    m.workFromAnywhereMetrics = value
}
