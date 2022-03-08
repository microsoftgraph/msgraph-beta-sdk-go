package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsOverview provides operations to manage the deviceManagement singleton.
type UserExperienceAnalyticsOverview struct {
    Entity
    // The user experience analytics insights.
    insights []UserExperienceAnalyticsInsightable;
}
// NewUserExperienceAnalyticsOverview instantiates a new userExperienceAnalyticsOverview and sets the default values.
func NewUserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverview) {
    m := &UserExperienceAnalyticsOverview{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserExperienceAnalyticsOverview(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsInsightFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsInsightable, len(val))
            for i, v := range val {
                res[i] = v.(UserExperienceAnalyticsInsightable)
            }
            m.SetInsights(res)
        }
        return nil
    }
    return res
}
// GetInsights gets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) GetInsights()([]UserExperienceAnalyticsInsightable) {
    if m == nil {
        return nil
    } else {
        return m.insights
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("insights", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInsights sets the insights property value. The user experience analytics insights.
func (m *UserExperienceAnalyticsOverview) SetInsights(value []UserExperienceAnalyticsInsightable)() {
    if m != nil {
        m.insights = value
    }
}
