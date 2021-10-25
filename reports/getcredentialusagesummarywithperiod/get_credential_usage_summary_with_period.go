package getcredentialusagesummarywithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetCredentialUsageSummaryWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    authMethod *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UsageAuthMethod;
    failureActivityCount *int64;
    feature *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FeatureType;
    successfulActivityCount *int64;
}
func NewGetCredentialUsageSummaryWithPeriod()(*GetCredentialUsageSummaryWithPeriod) {
    m := &GetCredentialUsageSummaryWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetCredentialUsageSummaryWithPeriod) GetAuthMethod()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UsageAuthMethod) {
    if m == nil {
        return nil
    } else {
        return m.authMethod
    }
}
func (m *GetCredentialUsageSummaryWithPeriod) GetFailureActivityCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failureActivityCount
    }
}
func (m *GetCredentialUsageSummaryWithPeriod) GetFeature()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FeatureType) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
func (m *GetCredentialUsageSummaryWithPeriod) GetSuccessfulActivityCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successfulActivityCount
    }
}
func (m *GetCredentialUsageSummaryWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseUsageAuthMethod)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UsageAuthMethod)
        m.SetAuthMethod(&cast)
        return nil
    }
    res["failureActivityCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetFailureActivityCount(val)
        return nil
    }
    res["feature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseFeatureType)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FeatureType)
        m.SetFeature(&cast)
        return nil
    }
    res["successfulActivityCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSuccessfulActivityCount(val)
        return nil
    }
    return res
}
func (m *GetCredentialUsageSummaryWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetCredentialUsageSummaryWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthMethod() != nil {
        cast := m.GetAuthMethod().String()
        err = writer.WriteStringValue("authMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("failureActivityCount", m.GetFailureActivityCount())
        if err != nil {
            return err
        }
    }
    if m.GetFeature() != nil {
        cast := m.GetFeature().String()
        err = writer.WriteStringValue("feature", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("successfulActivityCount", m.GetSuccessfulActivityCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetCredentialUsageSummaryWithPeriod) SetAuthMethod(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UsageAuthMethod)() {
    m.authMethod = value
}
func (m *GetCredentialUsageSummaryWithPeriod) SetFailureActivityCount(value *int64)() {
    m.failureActivityCount = value
}
func (m *GetCredentialUsageSummaryWithPeriod) SetFeature(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.FeatureType)() {
    m.feature = value
}
func (m *GetCredentialUsageSummaryWithPeriod) SetSuccessfulActivityCount(value *int64)() {
    m.successfulActivityCount = value
}
