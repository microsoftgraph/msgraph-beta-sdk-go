package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RelyingPartyDetailedSummary 
type RelyingPartyDetailedSummary struct {
    Entity
    // Number of failed sign in on Active Directory Federation Service in the period specified.
    failedSignInCount *int64
    // The migrationStatus property
    migrationStatus *MigrationStatus
    // Specifies all the validations check done on applications configuration details to evaluate if the application is ready to be moved to Azure AD.
    migrationValidationDetails []KeyValuePairable
    // This identifier is used to identify the relying party to this Federation Service. It is used when issuing claims to the relying party.
    relyingPartyId *string
    // Name of application or other entity on the internet that uses an identity provider to authenticate a user who wants to log in.
    relyingPartyName *string
    // Specifies where the relying party expects to receive the token.
    replyUrls []string
    // Uniquely identifies the Active Directory forest.
    serviceId *string
    // Number of successful / (number of successful + number of failed sign ins) on Active Directory Federation Service in the period specified.
    signInSuccessRate *float64
    // Number of successful sign ins on Active Directory Federation Service.
    successfulSignInCount *int64
    // Number of successful + failed sign ins failed sign ins on Active Directory Federation Service in the period specified.
    totalSignInCount *int64
    // Number of unique users that have signed into the application.
    uniqueUserCount *int64
}
// NewRelyingPartyDetailedSummary instantiates a new RelyingPartyDetailedSummary and sets the default values.
func NewRelyingPartyDetailedSummary()(*RelyingPartyDetailedSummary) {
    m := &RelyingPartyDetailedSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRelyingPartyDetailedSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRelyingPartyDetailedSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelyingPartyDetailedSummary(), nil
}
// GetFailedSignInCount gets the failedSignInCount property value. Number of failed sign in on Active Directory Federation Service in the period specified.
func (m *RelyingPartyDetailedSummary) GetFailedSignInCount()(*int64) {
    return m.failedSignInCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RelyingPartyDetailedSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedSignInCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetFailedSignInCount)
    res["migrationStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMigrationStatus , m.SetMigrationStatus)
    res["migrationValidationDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetMigrationValidationDetails)
    res["relyingPartyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRelyingPartyId)
    res["relyingPartyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRelyingPartyName)
    res["replyUrls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetReplyUrls)
    res["serviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetServiceId)
    res["signInSuccessRate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetSignInSuccessRate)
    res["successfulSignInCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetSuccessfulSignInCount)
    res["totalSignInCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetTotalSignInCount)
    res["uniqueUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetUniqueUserCount)
    return res
}
// GetMigrationStatus gets the migrationStatus property value. The migrationStatus property
func (m *RelyingPartyDetailedSummary) GetMigrationStatus()(*MigrationStatus) {
    return m.migrationStatus
}
// GetMigrationValidationDetails gets the migrationValidationDetails property value. Specifies all the validations check done on applications configuration details to evaluate if the application is ready to be moved to Azure AD.
func (m *RelyingPartyDetailedSummary) GetMigrationValidationDetails()([]KeyValuePairable) {
    return m.migrationValidationDetails
}
// GetRelyingPartyId gets the relyingPartyId property value. This identifier is used to identify the relying party to this Federation Service. It is used when issuing claims to the relying party.
func (m *RelyingPartyDetailedSummary) GetRelyingPartyId()(*string) {
    return m.relyingPartyId
}
// GetRelyingPartyName gets the relyingPartyName property value. Name of application or other entity on the internet that uses an identity provider to authenticate a user who wants to log in.
func (m *RelyingPartyDetailedSummary) GetRelyingPartyName()(*string) {
    return m.relyingPartyName
}
// GetReplyUrls gets the replyUrls property value. Specifies where the relying party expects to receive the token.
func (m *RelyingPartyDetailedSummary) GetReplyUrls()([]string) {
    return m.replyUrls
}
// GetServiceId gets the serviceId property value. Uniquely identifies the Active Directory forest.
func (m *RelyingPartyDetailedSummary) GetServiceId()(*string) {
    return m.serviceId
}
// GetSignInSuccessRate gets the signInSuccessRate property value. Number of successful / (number of successful + number of failed sign ins) on Active Directory Federation Service in the period specified.
func (m *RelyingPartyDetailedSummary) GetSignInSuccessRate()(*float64) {
    return m.signInSuccessRate
}
// GetSuccessfulSignInCount gets the successfulSignInCount property value. Number of successful sign ins on Active Directory Federation Service.
func (m *RelyingPartyDetailedSummary) GetSuccessfulSignInCount()(*int64) {
    return m.successfulSignInCount
}
// GetTotalSignInCount gets the totalSignInCount property value. Number of successful + failed sign ins failed sign ins on Active Directory Federation Service in the period specified.
func (m *RelyingPartyDetailedSummary) GetTotalSignInCount()(*int64) {
    return m.totalSignInCount
}
// GetUniqueUserCount gets the uniqueUserCount property value. Number of unique users that have signed into the application.
func (m *RelyingPartyDetailedSummary) GetUniqueUserCount()(*int64) {
    return m.uniqueUserCount
}
// Serialize serializes information the current object
func (m *RelyingPartyDetailedSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetMigrationStatus()).String()
        err = writer.WriteStringValue("migrationStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMigrationValidationDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMigrationValidationDetails())
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
    if m.GetReplyUrls() != nil {
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
// SetFailedSignInCount sets the failedSignInCount property value. Number of failed sign in on Active Directory Federation Service in the period specified.
func (m *RelyingPartyDetailedSummary) SetFailedSignInCount(value *int64)() {
    m.failedSignInCount = value
}
// SetMigrationStatus sets the migrationStatus property value. The migrationStatus property
func (m *RelyingPartyDetailedSummary) SetMigrationStatus(value *MigrationStatus)() {
    m.migrationStatus = value
}
// SetMigrationValidationDetails sets the migrationValidationDetails property value. Specifies all the validations check done on applications configuration details to evaluate if the application is ready to be moved to Azure AD.
func (m *RelyingPartyDetailedSummary) SetMigrationValidationDetails(value []KeyValuePairable)() {
    m.migrationValidationDetails = value
}
// SetRelyingPartyId sets the relyingPartyId property value. This identifier is used to identify the relying party to this Federation Service. It is used when issuing claims to the relying party.
func (m *RelyingPartyDetailedSummary) SetRelyingPartyId(value *string)() {
    m.relyingPartyId = value
}
// SetRelyingPartyName sets the relyingPartyName property value. Name of application or other entity on the internet that uses an identity provider to authenticate a user who wants to log in.
func (m *RelyingPartyDetailedSummary) SetRelyingPartyName(value *string)() {
    m.relyingPartyName = value
}
// SetReplyUrls sets the replyUrls property value. Specifies where the relying party expects to receive the token.
func (m *RelyingPartyDetailedSummary) SetReplyUrls(value []string)() {
    m.replyUrls = value
}
// SetServiceId sets the serviceId property value. Uniquely identifies the Active Directory forest.
func (m *RelyingPartyDetailedSummary) SetServiceId(value *string)() {
    m.serviceId = value
}
// SetSignInSuccessRate sets the signInSuccessRate property value. Number of successful / (number of successful + number of failed sign ins) on Active Directory Federation Service in the period specified.
func (m *RelyingPartyDetailedSummary) SetSignInSuccessRate(value *float64)() {
    m.signInSuccessRate = value
}
// SetSuccessfulSignInCount sets the successfulSignInCount property value. Number of successful sign ins on Active Directory Federation Service.
func (m *RelyingPartyDetailedSummary) SetSuccessfulSignInCount(value *int64)() {
    m.successfulSignInCount = value
}
// SetTotalSignInCount sets the totalSignInCount property value. Number of successful + failed sign ins failed sign ins on Active Directory Federation Service in the period specified.
func (m *RelyingPartyDetailedSummary) SetTotalSignInCount(value *int64)() {
    m.totalSignInCount = value
}
// SetUniqueUserCount sets the uniqueUserCount property value. Number of unique users that have signed into the application.
func (m *RelyingPartyDetailedSummary) SetUniqueUserCount(value *int64)() {
    m.uniqueUserCount = value
}
