package getrelyingpartydetailedsummarywithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetRelyingPartyDetailedSummaryWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    failedSignInCount *int64;
    migrationStatus *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MigrationStatus;
    migrationValidationDetails []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair;
    relyingPartyId *string;
    relyingPartyName *string;
    replyUrls []string;
    serviceId *string;
    signInSuccessRate *float64;
    successfulSignInCount *int64;
    totalSignInCount *int64;
    uniqueUserCount *int64;
}
func NewGetRelyingPartyDetailedSummaryWithPeriod()(*GetRelyingPartyDetailedSummaryWithPeriod) {
    m := &GetRelyingPartyDetailedSummaryWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetFailedSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failedSignInCount
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetMigrationStatus()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MigrationStatus) {
    if m == nil {
        return nil
    } else {
        return m.migrationStatus
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetMigrationValidationDetails()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.migrationValidationDetails
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetRelyingPartyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relyingPartyId
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetRelyingPartyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relyingPartyName
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetReplyUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.replyUrls
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetServiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceId
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetSignInSuccessRate()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.signInSuccessRate
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetSuccessfulSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successfulSignInCount
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetTotalSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalSignInCount
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetUniqueUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.uniqueUserCount
    }
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedSignInCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetFailedSignInCount(val)
        return nil
    }
    res["migrationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseMigrationStatus)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MigrationStatus)
        m.SetMigrationStatus(&cast)
        return nil
    }
    res["migrationValidationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair))
        }
        m.SetMigrationValidationDetails(res)
        return nil
    }
    res["relyingPartyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRelyingPartyId(val)
        return nil
    }
    res["relyingPartyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRelyingPartyName(val)
        return nil
    }
    res["replyUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetReplyUrls(res)
        return nil
    }
    res["serviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServiceId(val)
        return nil
    }
    res["signInSuccessRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetSignInSuccessRate(val)
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
    res["totalSignInCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalSignInCount(val)
        return nil
    }
    res["uniqueUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetUniqueUserCount(val)
        return nil
    }
    return res
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("failedSignInCount", m.GetFailedSignInCount())
        if err != nil {
            return err
        }
    }
    if m.GetMigrationStatus() != nil {
        cast := m.GetMigrationStatus().String()
        err = writer.WriteStringValue("migrationStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMigrationValidationDetails()))
        for i, v := range m.GetMigrationValidationDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("migrationValidationDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("relyingPartyId", m.GetRelyingPartyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("relyingPartyName", m.GetRelyingPartyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("replyUrls", m.GetReplyUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("signInSuccessRate", m.GetSignInSuccessRate())
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
        err = writer.WriteInt64Value("totalSignInCount", m.GetTotalSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("uniqueUserCount", m.GetUniqueUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetFailedSignInCount(value *int64)() {
    m.failedSignInCount = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetMigrationStatus(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MigrationStatus)() {
    m.migrationStatus = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetMigrationValidationDetails(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair)() {
    m.migrationValidationDetails = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetRelyingPartyId(value *string)() {
    m.relyingPartyId = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetRelyingPartyName(value *string)() {
    m.relyingPartyName = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetReplyUrls(value []string)() {
    m.replyUrls = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetServiceId(value *string)() {
    m.serviceId = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetSignInSuccessRate(value *float64)() {
    m.signInSuccessRate = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetSuccessfulSignInCount(value *int64)() {
    m.successfulSignInCount = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetTotalSignInCount(value *int64)() {
    m.totalSignInCount = value
}
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetUniqueUserCount(value *int64)() {
    m.uniqueUserCount = value
}
