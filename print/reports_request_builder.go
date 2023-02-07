package print

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.print entity.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReportsRequestBuilderGetQueryParameters get reports from print
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsRequestBuilderGetQueryParameters
}
// ReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationSignInDetailedSummary provides operations to manage the applicationSignInDetailedSummary property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummary()(*ReportsApplicationSignInDetailedSummaryRequestBuilder) {
    return NewReportsApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplicationSignInDetailedSummaryById provides operations to manage the applicationSignInDetailedSummary property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*ReportsApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationSignInDetailedSummary%2Did"] = id
    }
    return NewReportsApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AuthenticationMethods provides operations to manage the authenticationMethods property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) AuthenticationMethods()(*ReportsAuthenticationMethodsRequestBuilder) {
    return NewReportsAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CredentialUserRegistrationDetails provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetails()(*ReportsCredentialUserRegistrationDetailsRequestBuilder) {
    return NewReportsCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CredentialUserRegistrationDetailsById provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*ReportsCredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationDetails%2Did"] = id
    }
    return NewReportsCredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DailyPrintUsage provides operations to manage the dailyPrintUsage property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsage()(*ReportsDailyPrintUsageRequestBuilder) {
    return NewReportsDailyPrintUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageById provides operations to manage the dailyPrintUsage property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageById(id string)(*ReportsDailyPrintUsagePrintUsageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsage%2Did"] = id
    }
    return NewReportsDailyPrintUsagePrintUsageItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DailyPrintUsageByPrinter provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*ReportsDailyPrintUsageByPrinterRequestBuilder) {
    return NewReportsDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageByPrinterById provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*ReportsDailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewReportsDailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DailyPrintUsageByUser provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*ReportsDailyPrintUsageByUserRequestBuilder) {
    return NewReportsDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageByUserById provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*ReportsDailyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewReportsDailyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DailyPrintUsageSummariesByPrinter provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*ReportsDailyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewReportsDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageSummariesByPrinterById provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*ReportsDailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewReportsDailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DailyPrintUsageSummariesByUser provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*ReportsDailyPrintUsageSummariesByUserRequestBuilder) {
    return NewReportsDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageSummariesByUserById provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*ReportsDailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewReportsDailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete navigation property reports for print
func (m *ReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get reports from print
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// MicrosoftGraphDeviceConfigurationDeviceActivity provides operations to call the deviceConfigurationDeviceActivity method.
func (m *ReportsRequestBuilder) MicrosoftGraphDeviceConfigurationDeviceActivity()(*ReportsMicrosoftGraphDeviceConfigurationDeviceActivityRequestBuilder) {
    return NewReportsMicrosoftGraphDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeviceConfigurationUserActivity provides operations to call the deviceConfigurationUserActivity method.
func (m *ReportsRequestBuilder) MicrosoftGraphDeviceConfigurationUserActivity()(*ReportsMicrosoftGraphDeviceConfigurationUserActivityRequestBuilder) {
    return NewReportsMicrosoftGraphDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationRepeatOffenders provides operations to call the getAttackSimulationRepeatOffenders method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationRepeatOffenders()(*ReportsMicrosoftGraphGetAttackSimulationRepeatOffendersRequestBuilder) {
    return NewReportsMicrosoftGraphGetAttackSimulationRepeatOffendersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationSimulationUserCoverage provides operations to call the getAttackSimulationSimulationUserCoverage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationSimulationUserCoverage()(*ReportsMicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return NewReportsMicrosoftGraphGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationTrainingUserCoverage provides operations to call the getAttackSimulationTrainingUserCoverage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationTrainingUserCoverage()(*ReportsMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return NewReportsMicrosoftGraphGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriod provides operations to call the getAzureADApplicationSignInSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriod(period *string)(*ReportsMicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserDistributionUserCountsWithPeriod provides operations to call the getBrowserDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetBrowserDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetBrowserDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserUserCountsWithPeriod provides operations to call the getBrowserUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetBrowserUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetBrowserUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserUserDetailWithPeriod provides operations to call the getBrowserUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetBrowserUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetBrowserUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetCredentialUsageSummaryWithPeriod provides operations to call the getCredentialUsageSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCredentialUsageSummaryWithPeriod(period *string)(*ReportsMicrosoftGraphGetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetCredentialUserRegistrationCount provides operations to call the getCredentialUserRegistrationCount method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCredentialUserRegistrationCount()(*ReportsMicrosoftGraphGetCredentialUserRegistrationCountRequestBuilder) {
    return NewReportsMicrosoftGraphGetCredentialUserRegistrationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetEmailActivityCountsWithPeriod provides operations to call the getEmailActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailActivityUserCountsWithPeriod provides operations to call the getEmailActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailActivityUserDetailWithDate provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetEmailActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetEmailActivityUserDetailWithPeriod provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriod provides operations to call the getEmailAppUsageAppsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageUserCountsWithPeriod provides operations to call the getEmailAppUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageUserDetailWithDate provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetEmailAppUsageUserDetailWithPeriod provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriod provides operations to call the getEmailAppUsageVersionsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityCountsWithPeriod provides operations to call the getFormsUserActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetFormsUserActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetFormsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityUserCountsWithPeriod provides operations to call the getFormsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityUserDetailWithDate provides operations to call the getFormsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetFormsUserActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetFormsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetFormsUserActivityUserDetailWithPeriod provides operations to call the getFormsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetFormsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetFormsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime provides operations to call the getGroupArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewReportsMicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, groupId, startDateTime)
}
// MicrosoftGraphGetM365AppPlatformUserCountsWithPeriod provides operations to call the getM365AppPlatformUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppPlatformUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetM365AppUserCountsWithPeriod provides operations to call the getM365AppUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetM365AppUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetM365AppUserDetailWithDate provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetM365AppUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetM365AppUserDetailWithPeriod provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetM365AppUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageDetailWithPeriod provides operations to call the getMailboxUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetMailboxUsageDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriod provides operations to call the getMailboxUsageMailboxCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriod provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageStorageWithPeriod provides operations to call the getMailboxUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageStorageWithPeriod(period *string)(*ReportsMicrosoftGraphGetMailboxUsageStorageWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ActivationCounts provides operations to call the getOffice365ActivationCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationCounts()(*ReportsMicrosoftGraphGetOffice365ActivationCountsRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActivationsUserCounts provides operations to call the getOffice365ActivationsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationsUserCounts()(*ReportsMicrosoftGraphGetOffice365ActivationsUserCountsRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActivationsUserDetail provides operations to call the getOffice365ActivationsUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationsUserDetail()(*ReportsMicrosoftGraphGetOffice365ActivationsUserDetailRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActiveUserCountsWithPeriod provides operations to call the getOffice365ActiveUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ActiveUserDetailWithDate provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOffice365ActiveUserDetailWithPeriod provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriod provides operations to call the getOffice365GroupsActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityDetailWithDate provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriod provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriod provides operations to call the getOffice365GroupsActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriod provides operations to call the getOffice365GroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriod provides operations to call the getOffice365GroupsActivityStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ServicesUserCountsWithPeriod provides operations to call the getOffice365ServicesUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ServicesUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityFileCountsWithPeriod provides operations to call the getOneDriveActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityUserCountsWithPeriod provides operations to call the getOneDriveActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityUserDetailWithDate provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOneDriveActivityUserDetailWithPeriod provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriod provides operations to call the getOneDriveUsageAccountCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageAccountDetailWithDate provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriod provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageFileCountsWithPeriod provides operations to call the getOneDriveUsageFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageStorageWithPeriod provides operations to call the getOneDriveUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageStorageWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime provides operations to call the getPrinterArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewReportsMicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, printerId, startDateTime)
}
// MicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriod provides operations to call the getRelyingPartyDetailedSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriod(period *string)(*ReportsMicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityFileCountsWithPeriod provides operations to call the getSharePointActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityPagesWithPeriod provides operations to call the getSharePointActivityPages method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityPagesWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointActivityPagesWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityUserCountsWithPeriod provides operations to call the getSharePointActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityUserDetailWithDate provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetSharePointActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSharePointActivityUserDetailWithPeriod provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageDetailWithDate provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSharePointSiteUsageDetailWithPeriod provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriod provides operations to call the getSharePointSiteUsageFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsagePagesWithPeriod provides operations to call the getSharePointSiteUsagePages method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsagePagesWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriod provides operations to call the getSharePointSiteUsageSiteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageStorageWithPeriod provides operations to call the getSharePointSiteUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageStorageWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriod provides operations to call the getSkypeForBusinessActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDate provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriod provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDate provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDate provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriod provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityCountsWithPeriod provides operations to call the getTeamsTeamActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityDetailWithDate provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetTeamsTeamActivityDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsTeamActivityDetailWithPeriod provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsTeamActivityDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriod provides operations to call the getTeamsTeamActivityDistributionCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamCountsWithPeriod provides operations to call the getTeamsTeamCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsTeamCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityCountsWithPeriod provides operations to call the getTeamsUserActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriod provides operations to call the getTeamsUserActivityTotalCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriod provides operations to call the getTeamsUserActivityTotalDistributionCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriod provides operations to call the getTeamsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityUserDetailWithDate provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriod provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime provides operations to call the getUserArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*ReportsMicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewReportsMicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime, userId)
}
// MicrosoftGraphGetYammerActivityCountsWithPeriod provides operations to call the getYammerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerActivityUserCountsWithPeriod provides operations to call the getYammerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerActivityUserDetailWithDate provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetYammerActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerActivityUserDetailWithPeriod provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriod provides operations to call the getYammerDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageUserDetailWithDate provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriod provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityCountsWithPeriod provides operations to call the getYammerGroupsActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityDetailWithDate provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerGroupsActivityDetailWithPeriod provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriod provides operations to call the getYammerGroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ReportsMicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ReportsMicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureDetails provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureDetails()(*ReportsMicrosoftGraphManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ReportsMicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureTrends provides operations to call the managedDeviceEnrollmentFailureTrends method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureTrends()(*ReportsMicrosoftGraphManagedDeviceEnrollmentFailureTrendsRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentFailureTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentTopFailures provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentTopFailures()(*ReportsMicrosoftGraphManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriod provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*ReportsMicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MonthlyPrintUsageByPrinter provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*ReportsMonthlyPrintUsageByPrinterRequestBuilder) {
    return NewReportsMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageByPrinterById provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*ReportsMonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewReportsMonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MonthlyPrintUsageByUser provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*ReportsMonthlyPrintUsageByUserRequestBuilder) {
    return NewReportsMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageByUserById provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*ReportsMonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewReportsMonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MonthlyPrintUsageSummariesByPrinter provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*ReportsMonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewReportsMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageSummariesByPrinterById provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*ReportsMonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewReportsMonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// MonthlyPrintUsageSummariesByUser provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*ReportsMonthlyPrintUsageSummariesByUserRequestBuilder) {
    return NewReportsMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageSummariesByUserById provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*ReportsMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewReportsMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property reports in print
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// Security provides operations to manage the security property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) Security()(*ReportsSecurityRequestBuilder) {
    return NewReportsSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property reports for print
func (m *ReportsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get reports from print
func (m *ReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property reports in print
func (m *ReportsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UserCredentialUsageDetails provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) UserCredentialUsageDetails()(*ReportsUserCredentialUsageDetailsRequestBuilder) {
    return NewReportsUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserCredentialUsageDetailsById provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*ReportsUserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userCredentialUsageDetails%2Did"] = id
    }
    return NewReportsUserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
