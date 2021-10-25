package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsBaseline struct {
    Entity
    appHealthMetrics *UserExperienceAnalyticsCategory;
    bestPracticesMetrics *UserExperienceAnalyticsCategory;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    deviceBootPerformanceMetrics *UserExperienceAnalyticsCategory;
    displayName *string;
    isBuiltIn *bool;
    rebootAnalyticsMetrics *UserExperienceAnalyticsCategory;
    resourcePerformanceMetrics *UserExperienceAnalyticsCategory;
    workFromAnywhereMetrics *UserExperienceAnalyticsCategory;
}
func NewUserExperienceAnalyticsBaseline()(*UserExperienceAnalyticsBaseline) {
    m := &UserExperienceAnalyticsBaseline{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsBaseline) GetAppHealthMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.appHealthMetrics
    }
}
func (m *UserExperienceAnalyticsBaseline) GetBestPracticesMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.bestPracticesMetrics
    }
}
func (m *UserExperienceAnalyticsBaseline) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *UserExperienceAnalyticsBaseline) GetDeviceBootPerformanceMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceBootPerformanceMetrics
    }
}
func (m *UserExperienceAnalyticsBaseline) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *UserExperienceAnalyticsBaseline) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
func (m *UserExperienceAnalyticsBaseline) GetRebootAnalyticsMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.rebootAnalyticsMetrics
    }
}
func (m *UserExperienceAnalyticsBaseline) GetResourcePerformanceMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.resourcePerformanceMetrics
    }
}
func (m *UserExperienceAnalyticsBaseline) GetWorkFromAnywhereMetrics()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.workFromAnywhereMetrics
    }
}
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
func (m *UserExperienceAnalyticsBaseline) SetAppHealthMetrics(value *UserExperienceAnalyticsCategory)() {
    m.appHealthMetrics = value
}
func (m *UserExperienceAnalyticsBaseline) SetBestPracticesMetrics(value *UserExperienceAnalyticsCategory)() {
    m.bestPracticesMetrics = value
}
func (m *UserExperienceAnalyticsBaseline) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *UserExperienceAnalyticsBaseline) SetDeviceBootPerformanceMetrics(value *UserExperienceAnalyticsCategory)() {
    m.deviceBootPerformanceMetrics = value
}
func (m *UserExperienceAnalyticsBaseline) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *UserExperienceAnalyticsBaseline) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
func (m *UserExperienceAnalyticsBaseline) SetRebootAnalyticsMetrics(value *UserExperienceAnalyticsCategory)() {
    m.rebootAnalyticsMetrics = value
}
func (m *UserExperienceAnalyticsBaseline) SetResourcePerformanceMetrics(value *UserExperienceAnalyticsCategory)() {
    m.resourcePerformanceMetrics = value
}
func (m *UserExperienceAnalyticsBaseline) SetWorkFromAnywhereMetrics(value *UserExperienceAnalyticsCategory)() {
    m.workFromAnywhereMetrics = value
}
