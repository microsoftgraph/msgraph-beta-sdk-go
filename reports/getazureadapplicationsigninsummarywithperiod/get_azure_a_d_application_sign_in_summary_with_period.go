package getazureadapplicationsigninsummarywithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetAzureADApplicationSignInSummaryWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    appDisplayName *string;
    failedSignInCount *int64;
    successfulSignInCount *int64;
    successPercentage *float64;
}
func NewGetAzureADApplicationSignInSummaryWithPeriod()(*GetAzureADApplicationSignInSummaryWithPeriod) {
    m := &GetAzureADApplicationSignInSummaryWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetFailedSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failedSignInCount
    }
}
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetSuccessfulSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successfulSignInCount
    }
}
func (m *GetAzureADApplicationSignInSummaryWithPeriod) GetSuccessPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.successPercentage
    }
}
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
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetFailedSignInCount(value *int64)() {
    m.failedSignInCount = value
}
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetSuccessfulSignInCount(value *int64)() {
    m.successfulSignInCount = value
}
func (m *GetAzureADApplicationSignInSummaryWithPeriod) SetSuccessPercentage(value *float64)() {
    m.successPercentage = value
}
