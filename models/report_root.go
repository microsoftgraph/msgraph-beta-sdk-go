package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportRoot 
type ReportRoot struct {
    Entity
    // Represents a detailed summary of an application sign-in.
    applicationSignInDetailedSummary []ApplicationSignInDetailedSummaryable
    // Container for navigation properties for Azure AD authentication methods resources.
    authenticationMethods AuthenticationMethodsRootable
    // Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
    credentialUserRegistrationDetails []CredentialUserRegistrationDetailsable
    // The dailyPrintUsageByPrinter property
    dailyPrintUsageByPrinter []PrintUsageByPrinterable
    // The dailyPrintUsageByUser property
    dailyPrintUsageByUser []PrintUsageByUserable
    // The dailyPrintUsageSummariesByPrinter property
    dailyPrintUsageSummariesByPrinter []PrintUsageByPrinterable
    // The dailyPrintUsageSummariesByUser property
    dailyPrintUsageSummariesByUser []PrintUsageByUserable
    // The monthlyPrintUsageByPrinter property
    monthlyPrintUsageByPrinter []PrintUsageByPrinterable
    // The monthlyPrintUsageByUser property
    monthlyPrintUsageByUser []PrintUsageByUserable
    // The monthlyPrintUsageSummariesByPrinter property
    monthlyPrintUsageSummariesByPrinter []PrintUsageByPrinterable
    // The monthlyPrintUsageSummariesByUser property
    monthlyPrintUsageSummariesByUser []PrintUsageByUserable
    // Provides the ability to launch a realistic simulated phishing attack that organizations can learn from.
    security SecurityReportsRootable
    // Represents the self-service password reset (SSPR) usage for a given tenant.
    userCredentialUsageDetails []UserCredentialUsageDetailsable
}
// NewReportRoot instantiates a new ReportRoot and sets the default values.
func NewReportRoot()(*ReportRoot) {
    m := &ReportRoot{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.reportRoot";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateReportRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportRoot(), nil
}
// GetApplicationSignInDetailedSummary gets the applicationSignInDetailedSummary property value. Represents a detailed summary of an application sign-in.
func (m *ReportRoot) GetApplicationSignInDetailedSummary()([]ApplicationSignInDetailedSummaryable) {
    return m.applicationSignInDetailedSummary
}
// GetAuthenticationMethods gets the authenticationMethods property value. Container for navigation properties for Azure AD authentication methods resources.
func (m *ReportRoot) GetAuthenticationMethods()(AuthenticationMethodsRootable) {
    return m.authenticationMethods
}
// GetCredentialUserRegistrationDetails gets the credentialUserRegistrationDetails property value. Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
func (m *ReportRoot) GetCredentialUserRegistrationDetails()([]CredentialUserRegistrationDetailsable) {
    return m.credentialUserRegistrationDetails
}
// GetDailyPrintUsageByPrinter gets the dailyPrintUsageByPrinter property value. The dailyPrintUsageByPrinter property
func (m *ReportRoot) GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable) {
    return m.dailyPrintUsageByPrinter
}
// GetDailyPrintUsageByUser gets the dailyPrintUsageByUser property value. The dailyPrintUsageByUser property
func (m *ReportRoot) GetDailyPrintUsageByUser()([]PrintUsageByUserable) {
    return m.dailyPrintUsageByUser
}
// GetDailyPrintUsageSummariesByPrinter gets the dailyPrintUsageSummariesByPrinter property value. The dailyPrintUsageSummariesByPrinter property
func (m *ReportRoot) GetDailyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable) {
    return m.dailyPrintUsageSummariesByPrinter
}
// GetDailyPrintUsageSummariesByUser gets the dailyPrintUsageSummariesByUser property value. The dailyPrintUsageSummariesByUser property
func (m *ReportRoot) GetDailyPrintUsageSummariesByUser()([]PrintUsageByUserable) {
    return m.dailyPrintUsageSummariesByUser
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationSignInDetailedSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateApplicationSignInDetailedSummaryFromDiscriminatorValue , m.SetApplicationSignInDetailedSummary)
    res["authenticationMethods"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAuthenticationMethodsRootFromDiscriminatorValue , m.SetAuthenticationMethods)
    res["credentialUserRegistrationDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCredentialUserRegistrationDetailsFromDiscriminatorValue , m.SetCredentialUserRegistrationDetails)
    res["dailyPrintUsageByPrinter"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue , m.SetDailyPrintUsageByPrinter)
    res["dailyPrintUsageByUser"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue , m.SetDailyPrintUsageByUser)
    res["dailyPrintUsageSummariesByPrinter"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue , m.SetDailyPrintUsageSummariesByPrinter)
    res["dailyPrintUsageSummariesByUser"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue , m.SetDailyPrintUsageSummariesByUser)
    res["monthlyPrintUsageByPrinter"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue , m.SetMonthlyPrintUsageByPrinter)
    res["monthlyPrintUsageByUser"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue , m.SetMonthlyPrintUsageByUser)
    res["monthlyPrintUsageSummariesByPrinter"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue , m.SetMonthlyPrintUsageSummariesByPrinter)
    res["monthlyPrintUsageSummariesByUser"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue , m.SetMonthlyPrintUsageSummariesByUser)
    res["security"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSecurityReportsRootFromDiscriminatorValue , m.SetSecurity)
    res["userCredentialUsageDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserCredentialUsageDetailsFromDiscriminatorValue , m.SetUserCredentialUsageDetails)
    return res
}
// GetMonthlyPrintUsageByPrinter gets the monthlyPrintUsageByPrinter property value. The monthlyPrintUsageByPrinter property
func (m *ReportRoot) GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable) {
    return m.monthlyPrintUsageByPrinter
}
// GetMonthlyPrintUsageByUser gets the monthlyPrintUsageByUser property value. The monthlyPrintUsageByUser property
func (m *ReportRoot) GetMonthlyPrintUsageByUser()([]PrintUsageByUserable) {
    return m.monthlyPrintUsageByUser
}
// GetMonthlyPrintUsageSummariesByPrinter gets the monthlyPrintUsageSummariesByPrinter property value. The monthlyPrintUsageSummariesByPrinter property
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable) {
    return m.monthlyPrintUsageSummariesByPrinter
}
// GetMonthlyPrintUsageSummariesByUser gets the monthlyPrintUsageSummariesByUser property value. The monthlyPrintUsageSummariesByUser property
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByUser()([]PrintUsageByUserable) {
    return m.monthlyPrintUsageSummariesByUser
}
// GetSecurity gets the security property value. Provides the ability to launch a realistic simulated phishing attack that organizations can learn from.
func (m *ReportRoot) GetSecurity()(SecurityReportsRootable) {
    return m.security
}
// GetUserCredentialUsageDetails gets the userCredentialUsageDetails property value. Represents the self-service password reset (SSPR) usage for a given tenant.
func (m *ReportRoot) GetUserCredentialUsageDetails()([]UserCredentialUsageDetailsable) {
    return m.userCredentialUsageDetails
}
// Serialize serializes information the current object
func (m *ReportRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicationSignInDetailedSummary() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetApplicationSignInDetailedSummary())
        err = writer.WriteCollectionOfObjectValues("applicationSignInDetailedSummary", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationMethods", m.GetAuthenticationMethods())
        if err != nil {
            return err
        }
    }
    if m.GetCredentialUserRegistrationDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCredentialUserRegistrationDetails())
        err = writer.WriteCollectionOfObjectValues("credentialUserRegistrationDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageByPrinter() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDailyPrintUsageByPrinter())
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageByUser() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDailyPrintUsageByUser())
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageSummariesByPrinter() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDailyPrintUsageSummariesByPrinter())
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageSummariesByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageSummariesByUser() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDailyPrintUsageSummariesByUser())
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageSummariesByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageByPrinter() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMonthlyPrintUsageByPrinter())
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageByUser() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMonthlyPrintUsageByUser())
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageSummariesByPrinter() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMonthlyPrintUsageSummariesByPrinter())
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageSummariesByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageSummariesByUser() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMonthlyPrintUsageSummariesByUser())
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageSummariesByUser", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("security", m.GetSecurity())
        if err != nil {
            return err
        }
    }
    if m.GetUserCredentialUsageDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserCredentialUsageDetails())
        err = writer.WriteCollectionOfObjectValues("userCredentialUsageDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationSignInDetailedSummary sets the applicationSignInDetailedSummary property value. Represents a detailed summary of an application sign-in.
func (m *ReportRoot) SetApplicationSignInDetailedSummary(value []ApplicationSignInDetailedSummaryable)() {
    m.applicationSignInDetailedSummary = value
}
// SetAuthenticationMethods sets the authenticationMethods property value. Container for navigation properties for Azure AD authentication methods resources.
func (m *ReportRoot) SetAuthenticationMethods(value AuthenticationMethodsRootable)() {
    m.authenticationMethods = value
}
// SetCredentialUserRegistrationDetails sets the credentialUserRegistrationDetails property value. Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
func (m *ReportRoot) SetCredentialUserRegistrationDetails(value []CredentialUserRegistrationDetailsable)() {
    m.credentialUserRegistrationDetails = value
}
// SetDailyPrintUsageByPrinter sets the dailyPrintUsageByPrinter property value. The dailyPrintUsageByPrinter property
func (m *ReportRoot) SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)() {
    m.dailyPrintUsageByPrinter = value
}
// SetDailyPrintUsageByUser sets the dailyPrintUsageByUser property value. The dailyPrintUsageByUser property
func (m *ReportRoot) SetDailyPrintUsageByUser(value []PrintUsageByUserable)() {
    m.dailyPrintUsageByUser = value
}
// SetDailyPrintUsageSummariesByPrinter sets the dailyPrintUsageSummariesByPrinter property value. The dailyPrintUsageSummariesByPrinter property
func (m *ReportRoot) SetDailyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)() {
    m.dailyPrintUsageSummariesByPrinter = value
}
// SetDailyPrintUsageSummariesByUser sets the dailyPrintUsageSummariesByUser property value. The dailyPrintUsageSummariesByUser property
func (m *ReportRoot) SetDailyPrintUsageSummariesByUser(value []PrintUsageByUserable)() {
    m.dailyPrintUsageSummariesByUser = value
}
// SetMonthlyPrintUsageByPrinter sets the monthlyPrintUsageByPrinter property value. The monthlyPrintUsageByPrinter property
func (m *ReportRoot) SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)() {
    m.monthlyPrintUsageByPrinter = value
}
// SetMonthlyPrintUsageByUser sets the monthlyPrintUsageByUser property value. The monthlyPrintUsageByUser property
func (m *ReportRoot) SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)() {
    m.monthlyPrintUsageByUser = value
}
// SetMonthlyPrintUsageSummariesByPrinter sets the monthlyPrintUsageSummariesByPrinter property value. The monthlyPrintUsageSummariesByPrinter property
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)() {
    m.monthlyPrintUsageSummariesByPrinter = value
}
// SetMonthlyPrintUsageSummariesByUser sets the monthlyPrintUsageSummariesByUser property value. The monthlyPrintUsageSummariesByUser property
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByUser(value []PrintUsageByUserable)() {
    m.monthlyPrintUsageSummariesByUser = value
}
// SetSecurity sets the security property value. Provides the ability to launch a realistic simulated phishing attack that organizations can learn from.
func (m *ReportRoot) SetSecurity(value SecurityReportsRootable)() {
    m.security = value
}
// SetUserCredentialUsageDetails sets the userCredentialUsageDetails property value. Represents the self-service password reset (SSPR) usage for a given tenant.
func (m *ReportRoot) SetUserCredentialUsageDetails(value []UserCredentialUsageDetailsable)() {
    m.userCredentialUsageDetails = value
}
