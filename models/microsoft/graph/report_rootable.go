package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ReportRootable 
type ReportRootable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationSignInDetailedSummary()([]ApplicationSignInDetailedSummaryable)
    GetAuthenticationMethods()(AuthenticationMethodsRootable)
    GetCredentialUserRegistrationDetails()([]CredentialUserRegistrationDetailsable)
    GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetDailyPrintUsageByUser()([]PrintUsageByUserable)
    GetDailyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable)
    GetDailyPrintUsageSummariesByUser()([]PrintUsageByUserable)
    GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetMonthlyPrintUsageByUser()([]PrintUsageByUserable)
    GetMonthlyPrintUsageSummariesByPrinter()([]PrintUsageByPrinterable)
    GetMonthlyPrintUsageSummariesByUser()([]PrintUsageByUserable)
    GetUserCredentialUsageDetails()([]UserCredentialUsageDetailsable)
    SetApplicationSignInDetailedSummary(value []ApplicationSignInDetailedSummaryable)()
    SetAuthenticationMethods(value AuthenticationMethodsRootable)()
    SetCredentialUserRegistrationDetails(value []CredentialUserRegistrationDetailsable)()
    SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetDailyPrintUsageByUser(value []PrintUsageByUserable)()
    SetDailyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)()
    SetDailyPrintUsageSummariesByUser(value []PrintUsageByUserable)()
    SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)()
    SetMonthlyPrintUsageSummariesByPrinter(value []PrintUsageByPrinterable)()
    SetMonthlyPrintUsageSummariesByUser(value []PrintUsageByUserable)()
    SetUserCredentialUsageDetails(value []UserCredentialUsageDetailsable)()
}
