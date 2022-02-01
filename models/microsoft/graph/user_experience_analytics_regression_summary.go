package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsRegressionSummary 
type UserExperienceAnalyticsRegressionSummary struct {
    Entity
    // The metric values for the user experience analytics Manufacturer regression.
    manufacturerRegression []UserExperienceAnalyticsMetric;
    // The metric values for the user experience analytics model regression.
    modelRegression []UserExperienceAnalyticsMetric;
    // The metric values for the user experience analytics operating system regression.
    operatingSystemRegression []UserExperienceAnalyticsMetric;
}
// NewUserExperienceAnalyticsRegressionSummary instantiates a new userExperienceAnalyticsRegressionSummary and sets the default values.
func NewUserExperienceAnalyticsRegressionSummary()(*UserExperienceAnalyticsRegressionSummary) {
    m := &UserExperienceAnalyticsRegressionSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetManufacturerRegression gets the manufacturerRegression property value. The metric values for the user experience analytics Manufacturer regression.
func (m *UserExperienceAnalyticsRegressionSummary) GetManufacturerRegression()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.manufacturerRegression
    }
}
// GetModelRegression gets the modelRegression property value. The metric values for the user experience analytics model regression.
func (m *UserExperienceAnalyticsRegressionSummary) GetModelRegression()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.modelRegression
    }
}
// GetOperatingSystemRegression gets the operatingSystemRegression property value. The metric values for the user experience analytics operating system regression.
func (m *UserExperienceAnalyticsRegressionSummary) GetOperatingSystemRegression()([]UserExperienceAnalyticsMetric) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemRegression
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsRegressionSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["manufacturerRegression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetric, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserExperienceAnalyticsMetric))
            }
            m.SetManufacturerRegression(res)
        }
        return nil
    }
    res["modelRegression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetric, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserExperienceAnalyticsMetric))
            }
            m.SetModelRegression(res)
        }
        return nil
    }
    res["operatingSystemRegression"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetric() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetric, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserExperienceAnalyticsMetric))
            }
            m.SetOperatingSystemRegression(res)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsRegressionSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsRegressionSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetManufacturerRegression() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManufacturerRegression()))
        for i, v := range m.GetManufacturerRegression() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("manufacturerRegression", cast)
        if err != nil {
            return err
        }
    }
    if m.GetModelRegression() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetModelRegression()))
        for i, v := range m.GetModelRegression() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("modelRegression", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperatingSystemRegression() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperatingSystemRegression()))
        for i, v := range m.GetOperatingSystemRegression() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operatingSystemRegression", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManufacturerRegression sets the manufacturerRegression property value. The metric values for the user experience analytics Manufacturer regression.
func (m *UserExperienceAnalyticsRegressionSummary) SetManufacturerRegression(value []UserExperienceAnalyticsMetric)() {
    if m != nil {
        m.manufacturerRegression = value
    }
}
// SetModelRegression sets the modelRegression property value. The metric values for the user experience analytics model regression.
func (m *UserExperienceAnalyticsRegressionSummary) SetModelRegression(value []UserExperienceAnalyticsMetric)() {
    if m != nil {
        m.modelRegression = value
    }
}
// SetOperatingSystemRegression sets the operatingSystemRegression property value. The metric values for the user experience analytics operating system regression.
func (m *UserExperienceAnalyticsRegressionSummary) SetOperatingSystemRegression(value []UserExperienceAnalyticsMetric)() {
    if m != nil {
        m.operatingSystemRegression = value
    }
}
