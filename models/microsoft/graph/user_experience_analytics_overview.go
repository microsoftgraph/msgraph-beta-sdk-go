package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsOverview 
type UserExperienceAnalyticsOverview struct {
    Entity
    // The user experience analytics insights.
    insights []UserExperienceAnalyticsInsight;
}
// NewUserExperienceAnalyticsOverview instantiates a new userExperienceAnalyticsOverview and sets the default values.
func NewUserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverview) {
    m := &UserExperienceAnalyticsOverview{
        Entity: *NewEntity(),
    }
    return m
}
// GetInsights gets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) GetInsights()([]UserExperienceAnalyticsInsight) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsInsight() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsInsight, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserExperienceAnalyticsInsight))
            }
            m.SetInsights(res)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsOverview) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInsights() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInsights()))
        for i, v := range m.GetInsights() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("insights", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInsights sets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) SetInsights(value []UserExperienceAnalyticsInsight)() {
    if m != nil {
        m.insights = value
    }
}
