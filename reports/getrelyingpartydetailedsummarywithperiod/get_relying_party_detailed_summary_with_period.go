package getrelyingpartydetailedsummarywithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetRelyingPartyDetailedSummaryWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // Number of failed sign in on Active Directory Federation Service in the period specified.
    failedSignInCount *int64;
    // Indication of whether the application can be moved to Azure AD or require more investigation. Possible values are: ready, needsReview, additionalStepsRequired, unknownFutureValue.
    migrationStatus *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MigrationStatus;
    // Specifies all the validations check done on applications configuration details to evaluate if the application is ready to be moved to Azure AD.
    migrationValidationDetails []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair;
    // This identifier is used to identify the relying party to this Federation Service. It is used when issuing claims to the relying party.
    relyingPartyId *string;
    // Name of application or other entity on the internet that uses an identity provider to authenticate a user who wants to log in.
    relyingPartyName *string;
    // Specifies where the relying party expects to receive the token.
    replyUrls []string;
    // Uniquely identifies the Active Directory forest.
    serviceId *string;
    // Number of successful / (number of successful + number of failed sign ins) on Active Directory Federation Service in the period specified.
    signInSuccessRate *float64;
    // Number of successful sign ins on Active Directory Federation Service.
    successfulSignInCount *int64;
    // Number of successful + failed sign ins failed sign ins on Active Directory Federation Service in the period specified.
    totalSignInCount *int64;
    // Number of unique users that have signed into the application.
    uniqueUserCount *int64;
}
// Instantiates a new getRelyingPartyDetailedSummaryWithPeriod and sets the default values.
func NewGetRelyingPartyDetailedSummaryWithPeriod()(*GetRelyingPartyDetailedSummaryWithPeriod) {
    m := &GetRelyingPartyDetailedSummaryWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the failedSignInCount property value. Number of failed sign in on Active Directory Federation Service in the period specified.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetFailedSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.failedSignInCount
    }
}
// Gets the migrationStatus property value. Indication of whether the application can be moved to Azure AD or require more investigation. Possible values are: ready, needsReview, additionalStepsRequired, unknownFutureValue.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetMigrationStatus()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MigrationStatus) {
    if m == nil {
        return nil
    } else {
        return m.migrationStatus
    }
}
// Gets the migrationValidationDetails property value. Specifies all the validations check done on applications configuration details to evaluate if the application is ready to be moved to Azure AD.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetMigrationValidationDetails()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.migrationValidationDetails
    }
}
// Gets the relyingPartyId property value. This identifier is used to identify the relying party to this Federation Service. It is used when issuing claims to the relying party.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetRelyingPartyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relyingPartyId
    }
}
// Gets the relyingPartyName property value. Name of application or other entity on the internet that uses an identity provider to authenticate a user who wants to log in.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetRelyingPartyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relyingPartyName
    }
}
// Gets the replyUrls property value. Specifies where the relying party expects to receive the token.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetReplyUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.replyUrls
    }
}
// Gets the serviceId property value. Uniquely identifies the Active Directory forest.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetServiceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceId
    }
}
// Gets the signInSuccessRate property value. Number of successful / (number of successful + number of failed sign ins) on Active Directory Federation Service in the period specified.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetSignInSuccessRate()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.signInSuccessRate
    }
}
// Gets the successfulSignInCount property value. Number of successful sign ins on Active Directory Federation Service.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetSuccessfulSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.successfulSignInCount
    }
}
// Gets the totalSignInCount property value. Number of successful + failed sign ins failed sign ins on Active Directory Federation Service in the period specified.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetTotalSignInCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalSignInCount
    }
}
// Gets the uniqueUserCount property value. Number of unique users that have signed into the application.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) GetUniqueUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.uniqueUserCount
    }
}
// The deserialization information for the current model
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
            res[i] = *(v.(*string))
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the failedSignInCount property value. Number of failed sign in on Active Directory Federation Service in the period specified.
// Parameters:
//  - value : Value to set for the failedSignInCount property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetFailedSignInCount(value *int64)() {
    m.failedSignInCount = value
}
// Sets the migrationStatus property value. Indication of whether the application can be moved to Azure AD or require more investigation. Possible values are: ready, needsReview, additionalStepsRequired, unknownFutureValue.
// Parameters:
//  - value : Value to set for the migrationStatus property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetMigrationStatus(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MigrationStatus)() {
    m.migrationStatus = value
}
// Sets the migrationValidationDetails property value. Specifies all the validations check done on applications configuration details to evaluate if the application is ready to be moved to Azure AD.
// Parameters:
//  - value : Value to set for the migrationValidationDetails property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetMigrationValidationDetails(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair)() {
    m.migrationValidationDetails = value
}
// Sets the relyingPartyId property value. This identifier is used to identify the relying party to this Federation Service. It is used when issuing claims to the relying party.
// Parameters:
//  - value : Value to set for the relyingPartyId property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetRelyingPartyId(value *string)() {
    m.relyingPartyId = value
}
// Sets the relyingPartyName property value. Name of application or other entity on the internet that uses an identity provider to authenticate a user who wants to log in.
// Parameters:
//  - value : Value to set for the relyingPartyName property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetRelyingPartyName(value *string)() {
    m.relyingPartyName = value
}
// Sets the replyUrls property value. Specifies where the relying party expects to receive the token.
// Parameters:
//  - value : Value to set for the replyUrls property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetReplyUrls(value []string)() {
    m.replyUrls = value
}
// Sets the serviceId property value. Uniquely identifies the Active Directory forest.
// Parameters:
//  - value : Value to set for the serviceId property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetServiceId(value *string)() {
    m.serviceId = value
}
// Sets the signInSuccessRate property value. Number of successful / (number of successful + number of failed sign ins) on Active Directory Federation Service in the period specified.
// Parameters:
//  - value : Value to set for the signInSuccessRate property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetSignInSuccessRate(value *float64)() {
    m.signInSuccessRate = value
}
// Sets the successfulSignInCount property value. Number of successful sign ins on Active Directory Federation Service.
// Parameters:
//  - value : Value to set for the successfulSignInCount property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetSuccessfulSignInCount(value *int64)() {
    m.successfulSignInCount = value
}
// Sets the totalSignInCount property value. Number of successful + failed sign ins failed sign ins on Active Directory Federation Service in the period specified.
// Parameters:
//  - value : Value to set for the totalSignInCount property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetTotalSignInCount(value *int64)() {
    m.totalSignInCount = value
}
// Sets the uniqueUserCount property value. Number of unique users that have signed into the application.
// Parameters:
//  - value : Value to set for the uniqueUserCount property.
func (m *GetRelyingPartyDetailedSummaryWithPeriod) SetUniqueUserCount(value *int64)() {
    m.uniqueUserCount = value
}
