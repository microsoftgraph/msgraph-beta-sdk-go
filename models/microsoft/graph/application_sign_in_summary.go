package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplicationSignInSummary 
type ApplicationSignInSummary struct {
    Entity
    // Name of the application that the user signed into.
    appDisplayName *string;
    // Count of failed sign-ins made by the application.
    failedSignInCount *int64;
    // Count of successful sign-ins made by the application.
    successfulSignInCount *int64;
    // Percentage of successful sign-ins made by the application.
    successPercentage *float64;
}
// NewApplicationSignInSummary instantiates a new ApplicationSignInSummary and sets the default values.
func NewApplicationSignInSummary()(*ApplicationSignInSummary) {
    m := &ApplicationSignInSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppDisplayName gets the appDisplayName property value. Name of the application that the user signed into.
func (m *ApplicationSignInSummary) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetFailedSignInCount gets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *ApplicationSignInSummary) GetFailedSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failedSignInCount
    }
}
// GetSuccessfulSignInCount gets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) GetSuccessfulSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successfulSignInCount
    }
}
// GetSuccessPercentage gets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) GetSuccessPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.successPercentage
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationSignInSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["failedSignInCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedSignInCount(val)
        }
        return nil
    }
    res["successfulSignInCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulSignInCount(val)
        }
        return nil
    }
    res["successPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessPercentage(val)
        }
        return nil
    }
    return res
}
func (m *ApplicationSignInSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApplicationSignInSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("failedSignInCount", m.GetFailedSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("successfulSignInCount", m.GetSuccessfulSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("successPercentage", m.GetSuccessPercentage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. Name of the application that the user signed into.
func (m *ApplicationSignInSummary) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetFailedSignInCount sets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *ApplicationSignInSummary) SetFailedSignInCount(value *int64)() {
    if m != nil {
        m.failedSignInCount = value
    }
}
// SetSuccessfulSignInCount sets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) SetSuccessfulSignInCount(value *int64)() {
    if m != nil {
        m.successfulSignInCount = value
    }
}
// SetSuccessPercentage sets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *ApplicationSignInSummary) SetSuccessPercentage(value *float64)() {
    if m != nil {
        m.successPercentage = value
    }
}
