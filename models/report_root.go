package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportRoot 
type ReportRoot struct {
    Entity
}
// NewReportRoot instantiates a new reportRoot and sets the default values.
func NewReportRoot()(*ReportRoot) {
    m := &ReportRoot{
        Entity: *NewEntity(),
    }
    return m
}
// CreateReportRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReportRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportRoot(), nil
}
// GetAppCredentialSignInActivities gets the appCredentialSignInActivities property value. Represents a collection of sign-in activities of application credentials.
func (m *ReportRoot) GetAppCredentialSignInActivities()([]AppCredentialSignInActivityable) {
    val, err := m.GetBackingStore().Get("appCredentialSignInActivities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppCredentialSignInActivityable)
    }
    return nil
}
// GetApplicationSignInDetailedSummary gets the applicationSignInDetailedSummary property value. Represents a detailed summary of an application sign-in.
func (m *ReportRoot) GetApplicationSignInDetailedSummary()([]ApplicationSignInDetailedSummaryable) {
    val, err := m.GetBackingStore().Get("applicationSignInDetailedSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApplicationSignInDetailedSummaryable)
    }
    return nil
}
// GetAuthenticationMethods gets the authenticationMethods property value. Container for navigation properties for Microsoft Entra authentication methods resources.
func (m *ReportRoot) GetAuthenticationMethods()(AuthenticationMethodsRootable) {
    val, err := m.GetBackingStore().Get("authenticationMethods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationMethodsRootable)
    }
    return nil
}
// GetCredentialUserRegistrationDetails gets the credentialUserRegistrationDetails property value. Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
func (m *ReportRoot) GetCredentialUserRegistrationDetails()([]CredentialUserRegistrationDetailsable) {
    val, err := m.GetBackingStore().Get("credentialUserRegistrationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CredentialUserRegistrationDetailsable)
    }
    return nil
}
// GetDailyPrintUsage gets the dailyPrintUsage property value. The dailyPrintUsage property
func (m *ReportRoot) GetDailyPrintUsage()([]PrintUsageable) {
    val, err := m.GetBackingStore().Get("dailyPrintUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageable)
    }
    return nil
}
// GetDailyPrintUsageByPrinter gets the dailyPrintUsageByPrinter property value. Retrieve a list of daily print usage summaries, grouped by printer.
func (m *ReportRoot) GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable) {
    val, err := m.GetBackingStore().Get("dailyPrintUsageByPrinter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageByPrinterable)
    }
    return nil
}
// GetDailyPrintUsageByUser gets the dailyPrintUsageByUser property value. Retrieve a list of daily print usage summaries, grouped by user.
func (m *ReportRoot) GetDailyPrintUsageByUser()([]PrintUsageByUserable) {
    val, err := m.GetBackingStore().Get("dailyPrintUsageByUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageByUserable)
    }
    return nil
}
// GetDailyPrintUsageSummariesByPrinter gets the dailyPrintUsageSummariesByPrinter property value. The dailyPrintUsageSummariesByPrinter property
func (m *ReportRoot) GetDailyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable) {
    val, err := m.GetBackingStore().Get("dailyPrintUsageSummariesByPrinter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageByPrinterable)
    }
    return nil
}
// GetDailyPrintUsageSummariesByUser gets the dailyPrintUsageSummariesByUser property value. The dailyPrintUsageSummariesByUser property
func (m *ReportRoot) GetDailyPrintUsageSummariesByUser()([]PrintUsageByUserable) {
    val, err := m.GetBackingStore().Get("dailyPrintUsageSummariesByUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageByUserable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReportRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appCredentialSignInActivities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppCredentialSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppCredentialSignInActivityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppCredentialSignInActivityable)
                }
            }
            m.SetAppCredentialSignInActivities(res)
        }
        return nil
    }
    res["applicationSignInDetailedSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicationSignInDetailedSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApplicationSignInDetailedSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApplicationSignInDetailedSummaryable)
                }
            }
            m.SetApplicationSignInDetailedSummary(res)
        }
        return nil
    }
    res["authenticationMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethods(val.(AuthenticationMethodsRootable))
        }
        return nil
    }
    res["credentialUserRegistrationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCredentialUserRegistrationDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CredentialUserRegistrationDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CredentialUserRegistrationDetailsable)
                }
            }
            m.SetCredentialUserRegistrationDetails(res)
        }
        return nil
    }
    res["dailyPrintUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageable)
                }
            }
            m.SetDailyPrintUsage(res)
        }
        return nil
    }
    res["dailyPrintUsageByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageByPrinterable)
                }
            }
            m.SetDailyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["dailyPrintUsageByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageByUserable)
                }
            }
            m.SetDailyPrintUsageByUser(res)
        }
        return nil
    }
    res["dailyPrintUsageSummariesByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageByPrinterable)
                }
            }
            m.SetDailyPrintUsageSummariesByPrinter(res)
        }
        return nil
    }
    res["dailyPrintUsageSummariesByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageByUserable)
                }
            }
            m.SetDailyPrintUsageSummariesByUser(res)
        }
        return nil
    }
    res["monthlyPrintUsageByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageByPrinterable)
                }
            }
            m.SetMonthlyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["monthlyPrintUsageByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageByUserable)
                }
            }
            m.SetMonthlyPrintUsageByUser(res)
        }
        return nil
    }
    res["monthlyPrintUsageSummariesByPrinter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByPrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageByPrinterable)
                }
            }
            m.SetMonthlyPrintUsageSummariesByPrinter(res)
        }
        return nil
    }
    res["monthlyPrintUsageSummariesByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintUsageByUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrintUsageByUserable)
                }
            }
            m.SetMonthlyPrintUsageSummariesByUser(res)
        }
        return nil
    }
    res["security"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityReportsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurity(val.(SecurityReportsRootable))
        }
        return nil
    }
    res["servicePrincipalSignInActivities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalSignInActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalSignInActivityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServicePrincipalSignInActivityable)
                }
            }
            m.SetServicePrincipalSignInActivities(res)
        }
        return nil
    }
    res["sla"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceLevelAgreementRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSla(val.(ServiceLevelAgreementRootable))
        }
        return nil
    }
    res["userCredentialUsageDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserCredentialUsageDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserCredentialUsageDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserCredentialUsageDetailsable)
                }
            }
            m.SetUserCredentialUsageDetails(res)
        }
        return nil
    }
    res["userInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserInsightsRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserInsights(val.(UserInsightsRootable))
        }
        return nil
    }
    return res
}
// GetMonthlyPrintUsageByPrinter gets the monthlyPrintUsageByPrinter property value. Retrieve a list of monthly print usage summaries, grouped by printer.
func (m *ReportRoot) GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable) {
    val, err := m.GetBackingStore().Get("monthlyPrintUsageByPrinter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageByPrinterable)
    }
    return nil
}
// GetMonthlyPrintUsageByUser gets the monthlyPrintUsageByUser property value. Retrieve a list of monthly print usage summaries, grouped by user.
func (m *ReportRoot) GetMonthlyPrintUsageByUser()([]PrintUsageByUserable) {
    val, err := m.GetBackingStore().Get("monthlyPrintUsageByUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageByUserable)
    }
    return nil
}
// GetMonthlyPrintUsageSummariesByPrinter gets the monthlyPrintUsageSummariesByPrinter property value. The monthlyPrintUsageSummariesByPrinter property
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable) {
    val, err := m.GetBackingStore().Get("monthlyPrintUsageSummariesByPrinter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageByPrinterable)
    }
    return nil
}
// GetMonthlyPrintUsageSummariesByUser gets the monthlyPrintUsageSummariesByUser property value. The monthlyPrintUsageSummariesByUser property
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByUser()([]PrintUsageByUserable) {
    val, err := m.GetBackingStore().Get("monthlyPrintUsageSummariesByUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrintUsageByUserable)
    }
    return nil
}
// GetSecurity gets the security property value. Provides the ability to launch a realistically simulated phishing attack that organizations can learn from.
func (m *ReportRoot) GetSecurity()(SecurityReportsRootable) {
    val, err := m.GetBackingStore().Get("security")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SecurityReportsRootable)
    }
    return nil
}
// GetServicePrincipalSignInActivities gets the servicePrincipalSignInActivities property value. Represents a collection of sign-in activities of service principals.
func (m *ReportRoot) GetServicePrincipalSignInActivities()([]ServicePrincipalSignInActivityable) {
    val, err := m.GetBackingStore().Get("servicePrincipalSignInActivities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServicePrincipalSignInActivityable)
    }
    return nil
}
// GetSla gets the sla property value. A placeholder to allow for the desired URL path for SLA.
func (m *ReportRoot) GetSla()(ServiceLevelAgreementRootable) {
    val, err := m.GetBackingStore().Get("sla")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ServiceLevelAgreementRootable)
    }
    return nil
}
// GetUserCredentialUsageDetails gets the userCredentialUsageDetails property value. Represents the self-service password reset (SSPR) usage for a given tenant.
func (m *ReportRoot) GetUserCredentialUsageDetails()([]UserCredentialUsageDetailsable) {
    val, err := m.GetBackingStore().Get("userCredentialUsageDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserCredentialUsageDetailsable)
    }
    return nil
}
// GetUserInsights gets the userInsights property value. The userInsights property
func (m *ReportRoot) GetUserInsights()(UserInsightsRootable) {
    val, err := m.GetBackingStore().Get("userInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserInsightsRootable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ReportRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppCredentialSignInActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppCredentialSignInActivities()))
        for i, v := range m.GetAppCredentialSignInActivities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appCredentialSignInActivities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicationSignInDetailedSummary() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApplicationSignInDetailedSummary()))
        for i, v := range m.GetApplicationSignInDetailedSummary() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCredentialUserRegistrationDetails()))
        for i, v := range m.GetCredentialUserRegistrationDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("credentialUserRegistrationDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsage() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsage()))
        for i, v := range m.GetDailyPrintUsage() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsage", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageByPrinter()))
        for i, v := range m.GetDailyPrintUsageByPrinter() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageByUser()))
        for i, v := range m.GetDailyPrintUsageByUser() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageSummariesByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageSummariesByPrinter()))
        for i, v := range m.GetDailyPrintUsageSummariesByPrinter() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageSummariesByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDailyPrintUsageSummariesByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDailyPrintUsageSummariesByUser()))
        for i, v := range m.GetDailyPrintUsageSummariesByUser() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageSummariesByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageByPrinter()))
        for i, v := range m.GetMonthlyPrintUsageByPrinter() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageByUser()))
        for i, v := range m.GetMonthlyPrintUsageByUser() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageSummariesByPrinter() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageSummariesByPrinter()))
        for i, v := range m.GetMonthlyPrintUsageSummariesByPrinter() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageSummariesByPrinter", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMonthlyPrintUsageSummariesByUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonthlyPrintUsageSummariesByUser()))
        for i, v := range m.GetMonthlyPrintUsageSummariesByUser() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    if m.GetServicePrincipalSignInActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePrincipalSignInActivities()))
        for i, v := range m.GetServicePrincipalSignInActivities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("servicePrincipalSignInActivities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sla", m.GetSla())
        if err != nil {
            return err
        }
    }
    if m.GetUserCredentialUsageDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserCredentialUsageDetails()))
        for i, v := range m.GetUserCredentialUsageDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userCredentialUsageDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userInsights", m.GetUserInsights())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppCredentialSignInActivities sets the appCredentialSignInActivities property value. Represents a collection of sign-in activities of application credentials.
func (m *ReportRoot) SetAppCredentialSignInActivities(value []AppCredentialSignInActivityable)() {
    err := m.GetBackingStore().Set("appCredentialSignInActivities", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationSignInDetailedSummary sets the applicationSignInDetailedSummary property value. Represents a detailed summary of an application sign-in.
func (m *ReportRoot) SetApplicationSignInDetailedSummary(value []ApplicationSignInDetailedSummaryable)() {
    err := m.GetBackingStore().Set("applicationSignInDetailedSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthenticationMethods sets the authenticationMethods property value. Container for navigation properties for Microsoft Entra authentication methods resources.
func (m *ReportRoot) SetAuthenticationMethods(value AuthenticationMethodsRootable)() {
    err := m.GetBackingStore().Set("authenticationMethods", value)
    if err != nil {
        panic(err)
    }
}
// SetCredentialUserRegistrationDetails sets the credentialUserRegistrationDetails property value. Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
func (m *ReportRoot) SetCredentialUserRegistrationDetails(value []CredentialUserRegistrationDetailsable)() {
    err := m.GetBackingStore().Set("credentialUserRegistrationDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetDailyPrintUsage sets the dailyPrintUsage property value. The dailyPrintUsage property
func (m *ReportRoot) SetDailyPrintUsage(value []PrintUsageable)() {
    err := m.GetBackingStore().Set("dailyPrintUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetDailyPrintUsageByPrinter sets the dailyPrintUsageByPrinter property value. Retrieve a list of daily print usage summaries, grouped by printer.
func (m *ReportRoot) SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)() {
    err := m.GetBackingStore().Set("dailyPrintUsageByPrinter", value)
    if err != nil {
        panic(err)
    }
}
// SetDailyPrintUsageByUser sets the dailyPrintUsageByUser property value. Retrieve a list of daily print usage summaries, grouped by user.
func (m *ReportRoot) SetDailyPrintUsageByUser(value []PrintUsageByUserable)() {
    err := m.GetBackingStore().Set("dailyPrintUsageByUser", value)
    if err != nil {
        panic(err)
    }
}
// SetDailyPrintUsageSummariesByPrinter sets the dailyPrintUsageSummariesByPrinter property value. The dailyPrintUsageSummariesByPrinter property
func (m *ReportRoot) SetDailyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)() {
    err := m.GetBackingStore().Set("dailyPrintUsageSummariesByPrinter", value)
    if err != nil {
        panic(err)
    }
}
// SetDailyPrintUsageSummariesByUser sets the dailyPrintUsageSummariesByUser property value. The dailyPrintUsageSummariesByUser property
func (m *ReportRoot) SetDailyPrintUsageSummariesByUser(value []PrintUsageByUserable)() {
    err := m.GetBackingStore().Set("dailyPrintUsageSummariesByUser", value)
    if err != nil {
        panic(err)
    }
}
// SetMonthlyPrintUsageByPrinter sets the monthlyPrintUsageByPrinter property value. Retrieve a list of monthly print usage summaries, grouped by printer.
func (m *ReportRoot) SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)() {
    err := m.GetBackingStore().Set("monthlyPrintUsageByPrinter", value)
    if err != nil {
        panic(err)
    }
}
// SetMonthlyPrintUsageByUser sets the monthlyPrintUsageByUser property value. Retrieve a list of monthly print usage summaries, grouped by user.
func (m *ReportRoot) SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)() {
    err := m.GetBackingStore().Set("monthlyPrintUsageByUser", value)
    if err != nil {
        panic(err)
    }
}
// SetMonthlyPrintUsageSummariesByPrinter sets the monthlyPrintUsageSummariesByPrinter property value. The monthlyPrintUsageSummariesByPrinter property
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)() {
    err := m.GetBackingStore().Set("monthlyPrintUsageSummariesByPrinter", value)
    if err != nil {
        panic(err)
    }
}
// SetMonthlyPrintUsageSummariesByUser sets the monthlyPrintUsageSummariesByUser property value. The monthlyPrintUsageSummariesByUser property
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByUser(value []PrintUsageByUserable)() {
    err := m.GetBackingStore().Set("monthlyPrintUsageSummariesByUser", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurity sets the security property value. Provides the ability to launch a realistically simulated phishing attack that organizations can learn from.
func (m *ReportRoot) SetSecurity(value SecurityReportsRootable)() {
    err := m.GetBackingStore().Set("security", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipalSignInActivities sets the servicePrincipalSignInActivities property value. Represents a collection of sign-in activities of service principals.
func (m *ReportRoot) SetServicePrincipalSignInActivities(value []ServicePrincipalSignInActivityable)() {
    err := m.GetBackingStore().Set("servicePrincipalSignInActivities", value)
    if err != nil {
        panic(err)
    }
}
// SetSla sets the sla property value. A placeholder to allow for the desired URL path for SLA.
func (m *ReportRoot) SetSla(value ServiceLevelAgreementRootable)() {
    err := m.GetBackingStore().Set("sla", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCredentialUsageDetails sets the userCredentialUsageDetails property value. Represents the self-service password reset (SSPR) usage for a given tenant.
func (m *ReportRoot) SetUserCredentialUsageDetails(value []UserCredentialUsageDetailsable)() {
    err := m.GetBackingStore().Set("userCredentialUsageDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetUserInsights sets the userInsights property value. The userInsights property
func (m *ReportRoot) SetUserInsights(value UserInsightsRootable)() {
    err := m.GetBackingStore().Set("userInsights", value)
    if err != nil {
        panic(err)
    }
}
// ReportRootable 
type ReportRootable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppCredentialSignInActivities()([]AppCredentialSignInActivityable)
    GetApplicationSignInDetailedSummary()([]ApplicationSignInDetailedSummaryable)
    GetAuthenticationMethods()(AuthenticationMethodsRootable)
    GetCredentialUserRegistrationDetails()([]CredentialUserRegistrationDetailsable)
    GetDailyPrintUsage()([]PrintUsageable)
    GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetDailyPrintUsageByUser()([]PrintUsageByUserable)
    GetDailyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable)
    GetDailyPrintUsageSummariesByUser()([]PrintUsageByUserable)
    GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetMonthlyPrintUsageByUser()([]PrintUsageByUserable)
    GetMonthlyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable)
    GetMonthlyPrintUsageSummariesByUser()([]PrintUsageByUserable)
    GetSecurity()(SecurityReportsRootable)
    GetServicePrincipalSignInActivities()([]ServicePrincipalSignInActivityable)
    GetSla()(ServiceLevelAgreementRootable)
    GetUserCredentialUsageDetails()([]UserCredentialUsageDetailsable)
    GetUserInsights()(UserInsightsRootable)
    SetAppCredentialSignInActivities(value []AppCredentialSignInActivityable)()
    SetApplicationSignInDetailedSummary(value []ApplicationSignInDetailedSummaryable)()
    SetAuthenticationMethods(value AuthenticationMethodsRootable)()
    SetCredentialUserRegistrationDetails(value []CredentialUserRegistrationDetailsable)()
    SetDailyPrintUsage(value []PrintUsageable)()
    SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetDailyPrintUsageByUser(value []PrintUsageByUserable)()
    SetDailyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)()
    SetDailyPrintUsageSummariesByUser(value []PrintUsageByUserable)()
    SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)()
    SetMonthlyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)()
    SetMonthlyPrintUsageSummariesByUser(value []PrintUsageByUserable)()
    SetSecurity(value SecurityReportsRootable)()
    SetServicePrincipalSignInActivities(value []ServicePrincipalSignInActivityable)()
    SetSla(value ServiceLevelAgreementRootable)()
    SetUserCredentialUsageDetails(value []UserCredentialUsageDetailsable)()
    SetUserInsights(value UserInsightsRootable)()
}
