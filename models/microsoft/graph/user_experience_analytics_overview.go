package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsOverview struct {
    Entity
    // The user experience analytics insights.
    insights []UserExperienceAnalyticsInsight;
}
// Instantiates a new userExperienceAnalyticsOverview and sets the default values.
func NewUserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverview) {
    m := &UserExperienceAnalyticsOverview{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) GetInsights()([]UserExperienceAnalyticsInsight) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsInsight() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsInsight, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsInsight))
        }
        m.SetInsights(res)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsOverview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
// Sets the insights property value. The user experience analytics insights.
// Parameters:
//  - value : Value to set for the insights property.
func (m *UserExperienceAnalyticsOverview) SetInsights(value []UserExperienceAnalyticsInsight)() {
    m.insights = value
}
