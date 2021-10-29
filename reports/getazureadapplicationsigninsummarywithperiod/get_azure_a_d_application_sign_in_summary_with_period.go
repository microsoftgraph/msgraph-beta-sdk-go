package getazureadapplicationsigninsummarywithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
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
// Instantiates a new getAzureADApplicationSignInSummaryWithPeriod and sets the default values.
func NewGetAzureADApplicationSignInSummaryWithPeriod()(*GetAzureADApplicationSignInSummaryWithPeriod) {
    m := &GetAzureADApplicationSignInSummaryWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the appDisplayName property value. Name of the application that the user signed in to.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the failedSignInCount property value. Count of failed sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetFailedSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failedSignInCount
    }
}
// Gets the successfulSignInCount property value. Count of successful sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetSuccessfulSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successfulSignInCount
    }
}
// Gets the successPercentage property value. Percentage of successful sign-ins made by the application.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetSuccessPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.successPercentage
    }
}
// The deserialization information for the current model
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["failedSignInCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetFailedSignInCount(val)
        return nil
    }
    res["successfulSignInCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSuccessfulSignInCount(val)
        return nil
    }
    res["successPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetSuccessPercentage(val)
        return nil
    }
    return res
}
func (m *GetAzureADApplicationSignInSummaryWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the appDisplayName property value. Name of the application that the user signed in to.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the failedSignInCount property value. Count of failed sign-ins made by the application.
// Parameters:
//  - value : Value to set for the failedSignInCount property.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetFailedSignInCount(value *int64)() {
    m.failedSignInCount = value
}
// Sets the successfulSignInCount property value. Count of successful sign-ins made by the application.
// Parameters:
//  - value : Value to set for the successfulSignInCount property.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetSuccessfulSignInCount(value *int64)() {
    m.successfulSignInCount = value
}
// Sets the successPercentage property value. Percentage of successful sign-ins made by the application.
// Parameters:
//  - value : Value to set for the successPercentage property.
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetSuccessPercentage(value *float64)() {
    m.successPercentage = value
}
