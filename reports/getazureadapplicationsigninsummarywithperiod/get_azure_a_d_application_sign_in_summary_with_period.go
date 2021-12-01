package getazureadapplicationsigninsummarywithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetAzureADApplicationSignInSummaryWithPeriod 
type GetAzureADApplicationSignInSummaryWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // Name of the application that the user signed in to.
    appDisplayName *string;
    // Count of failed sign-ins made by the application.
    failedSignInCount *int64;
    // Count of successful sign-ins made by the application.
    successfulSignInCount *int64;
    // Percentage of successful sign-ins made by the application.
    successPercentage *float64;
}
// NewGetAzureADApplicationSignInSummaryWithPeriod instantiates a new getAzureADApplicationSignInSummaryWithPeriod and sets the default values.
func NewGetAzureADApplicationSignInSummaryWithPeriod()(*GetAzureADApplicationSignInSummaryWithPeriod) {
    m := &GetAzureADApplicationSignInSummaryWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAppDisplayName gets the appDisplayName property value. Name of the application that the user signed in to.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetFailedSignInCount gets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetFailedSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failedSignInCount
    }
}
// GetSuccessfulSignInCount gets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetSuccessfulSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successfulSignInCount
    }
}
// GetSuccessPercentage gets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetSuccessPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.successPercentage
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetAzureADApplicationSignInSummaryWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetAzureADApplicationSignInSummaryWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetAppDisplayName sets the appDisplayName property value. Name of the application that the user signed in to.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetFailedSignInCount sets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetFailedSignInCount(value *int64)() {
    if m != nil {
        m.failedSignInCount = value
    }
}
// SetSuccessfulSignInCount sets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetSuccessfulSignInCount(value *int64)() {
    if m != nil {
        m.successfulSignInCount = value
    }
}
// SetSuccessPercentage sets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetSuccessPercentage(value *float64)() {
    if m != nil {
        m.successPercentage = value
    }
}
