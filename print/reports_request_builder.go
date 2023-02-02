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
    idPtr := &id
    return NewReportsApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsCredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsDailyPrintUsagePrintUsageItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsDailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsDailyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsDailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsDailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
func (m *ReportsRequestBuilder) MicrosoftGraphDeviceConfigurationDeviceActivity()(*ReportsMicrosoftGraphDeviceConfigurationDeviceActivityDeviceConfigurationDeviceActivityRequestBuilder) {
    return NewReportsMicrosoftGraphDeviceConfigurationDeviceActivityDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeviceConfigurationUserActivity provides operations to call the deviceConfigurationUserActivity method.
func (m *ReportsRequestBuilder) MicrosoftGraphDeviceConfigurationUserActivity()(*ReportsMicrosoftGraphDeviceConfigurationUserActivityDeviceConfigurationUserActivityRequestBuilder) {
    return NewReportsMicrosoftGraphDeviceConfigurationUserActivityDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationRepeatOffenders provides operations to call the getAttackSimulationRepeatOffenders method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationRepeatOffenders()(*ReportsMicrosoftGraphGetAttackSimulationRepeatOffendersGetAttackSimulationRepeatOffendersRequestBuilder) {
    return NewReportsMicrosoftGraphGetAttackSimulationRepeatOffendersGetAttackSimulationRepeatOffendersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationSimulationUserCoverage provides operations to call the getAttackSimulationSimulationUserCoverage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationSimulationUserCoverage()(*ReportsMicrosoftGraphGetAttackSimulationSimulationUserCoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return NewReportsMicrosoftGraphGetAttackSimulationSimulationUserCoverageGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationTrainingUserCoverage provides operations to call the getAttackSimulationTrainingUserCoverage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationTrainingUserCoverage()(*ReportsMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return NewReportsMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriod provides operations to call the getAzureADApplicationSignInSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriod(period *string)(*ReportsMicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserDistributionUserCountsWithPeriod provides operations to call the getBrowserDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetBrowserDistributionUserCountsWithPeriodGetBrowserDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetBrowserDistributionUserCountsWithPeriodGetBrowserDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserUserCountsWithPeriod provides operations to call the getBrowserUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetBrowserUserCountsWithPeriodGetBrowserUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetBrowserUserCountsWithPeriodGetBrowserUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserUserDetailWithPeriod provides operations to call the getBrowserUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetBrowserUserDetailWithPeriodGetBrowserUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetBrowserUserDetailWithPeriodGetBrowserUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetCredentialUsageSummaryWithPeriod provides operations to call the getCredentialUsageSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCredentialUsageSummaryWithPeriod(period *string)(*ReportsMicrosoftGraphGetCredentialUsageSummaryWithPeriodGetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetCredentialUsageSummaryWithPeriodGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetCredentialUserRegistrationCount provides operations to call the getCredentialUserRegistrationCount method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCredentialUserRegistrationCount()(*ReportsMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder) {
    return NewReportsMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetEmailActivityCountsWithPeriod provides operations to call the getEmailActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailActivityCountsWithPeriodGetEmailActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailActivityCountsWithPeriodGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailActivityUserCountsWithPeriod provides operations to call the getEmailActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailActivityUserCountsWithPeriodGetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailActivityUserCountsWithPeriodGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailActivityUserDetailWithDate provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetEmailActivityUserDetailWithDateGetEmailActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailActivityUserDetailWithDateGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetEmailActivityUserDetailWithPeriod provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailActivityUserDetailWithPeriodGetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailActivityUserDetailWithPeriodGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriod provides operations to call the getEmailAppUsageAppsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriodGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriodGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageUserCountsWithPeriod provides operations to call the getEmailAppUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailAppUsageUserCountsWithPeriodGetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageUserCountsWithPeriodGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageUserDetailWithDate provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetEmailAppUsageUserDetailWithDateGetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageUserDetailWithDateGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetEmailAppUsageUserDetailWithPeriod provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailAppUsageUserDetailWithPeriodGetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageUserDetailWithPeriodGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriod provides operations to call the getEmailAppUsageVersionsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriodGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriodGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityCountsWithPeriod provides operations to call the getFormsUserActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetFormsUserActivityCountsWithPeriodGetFormsUserActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetFormsUserActivityCountsWithPeriodGetFormsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityUserCountsWithPeriod provides operations to call the getFormsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodGetFormsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodGetFormsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityUserDetailWithDate provides operations to call the getFormsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetFormsUserActivityUserDetailWithDateGetFormsUserActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetFormsUserActivityUserDetailWithDateGetFormsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetFormsUserActivityUserDetailWithPeriod provides operations to call the getFormsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetFormsUserActivityUserDetailWithPeriodGetFormsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetFormsUserActivityUserDetailWithPeriodGetFormsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime provides operations to call the getGroupArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewReportsMicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, groupId, startDateTime)
}
// MicrosoftGraphGetM365AppPlatformUserCountsWithPeriod provides operations to call the getM365AppPlatformUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppPlatformUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetM365AppPlatformUserCountsWithPeriodGetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetM365AppPlatformUserCountsWithPeriodGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetM365AppUserCountsWithPeriod provides operations to call the getM365AppUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetM365AppUserCountsWithPeriodGetM365AppUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetM365AppUserCountsWithPeriodGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetM365AppUserDetailWithDate provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetM365AppUserDetailWithDateGetM365AppUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetM365AppUserDetailWithDateGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetM365AppUserDetailWithPeriod provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetM365AppUserDetailWithPeriodGetM365AppUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetM365AppUserDetailWithPeriodGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageDetailWithPeriod provides operations to call the getMailboxUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetMailboxUsageDetailWithPeriodGetMailboxUsageDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetMailboxUsageDetailWithPeriodGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriod provides operations to call the getMailboxUsageMailboxCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetMailboxUsageMailboxCountsWithPeriodGetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetMailboxUsageMailboxCountsWithPeriodGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriod provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageStorageWithPeriod provides operations to call the getMailboxUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageStorageWithPeriod(period *string)(*ReportsMicrosoftGraphGetMailboxUsageStorageWithPeriodGetMailboxUsageStorageWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetMailboxUsageStorageWithPeriodGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ActivationCounts provides operations to call the getOffice365ActivationCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationCounts()(*ReportsMicrosoftGraphGetOffice365ActivationCountsGetOffice365ActivationCountsRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActivationCountsGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActivationsUserCounts provides operations to call the getOffice365ActivationsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationsUserCounts()(*ReportsMicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActivationsUserDetail provides operations to call the getOffice365ActivationsUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationsUserDetail()(*ReportsMicrosoftGraphGetOffice365ActivationsUserDetailGetOffice365ActivationsUserDetailRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActivationsUserDetailGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActiveUserCountsWithPeriod provides operations to call the getOffice365ActiveUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365ActiveUserCountsWithPeriodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActiveUserCountsWithPeriodGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ActiveUserDetailWithDate provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetOffice365ActiveUserDetailWithDateGetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActiveUserDetailWithDateGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOffice365ActiveUserDetailWithPeriod provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365ActiveUserDetailWithPeriodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ActiveUserDetailWithPeriodGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriod provides operations to call the getOffice365GroupsActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityCountsWithPeriodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityCountsWithPeriodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityDetailWithDate provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithDateGetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithDateGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriod provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriod provides operations to call the getOffice365GroupsActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriod provides operations to call the getOffice365GroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriod provides operations to call the getOffice365GroupsActivityStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ServicesUserCountsWithPeriod provides operations to call the getOffice365ServicesUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ServicesUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOffice365ServicesUserCountsWithPeriodGetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOffice365ServicesUserCountsWithPeriodGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityFileCountsWithPeriod provides operations to call the getOneDriveActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveActivityFileCountsWithPeriodGetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveActivityFileCountsWithPeriodGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityUserCountsWithPeriod provides operations to call the getOneDriveActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveActivityUserCountsWithPeriodGetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveActivityUserCountsWithPeriodGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityUserDetailWithDate provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetOneDriveActivityUserDetailWithDateGetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveActivityUserDetailWithDateGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOneDriveActivityUserDetailWithPeriod provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveActivityUserDetailWithPeriodGetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveActivityUserDetailWithPeriodGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriod provides operations to call the getOneDriveUsageAccountCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveUsageAccountCountsWithPeriodGetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageAccountCountsWithPeriodGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageAccountDetailWithDate provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetOneDriveUsageAccountDetailWithDateGetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageAccountDetailWithDateGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriod provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveUsageAccountDetailWithPeriodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageAccountDetailWithPeriodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageFileCountsWithPeriod provides operations to call the getOneDriveUsageFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveUsageFileCountsWithPeriodGetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageFileCountsWithPeriodGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageStorageWithPeriod provides operations to call the getOneDriveUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageStorageWithPeriod(period *string)(*ReportsMicrosoftGraphGetOneDriveUsageStorageWithPeriodGetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetOneDriveUsageStorageWithPeriodGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime provides operations to call the getPrinterArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewReportsMicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, printerId, startDateTime)
}
// MicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriod provides operations to call the getRelyingPartyDetailedSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriod(period *string)(*ReportsMicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriodGetRelyingPartyDetailedSummaryWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriodGetRelyingPartyDetailedSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityFileCountsWithPeriod provides operations to call the getSharePointActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointActivityFileCountsWithPeriodGetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityFileCountsWithPeriodGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityPagesWithPeriod provides operations to call the getSharePointActivityPages method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityPagesWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointActivityPagesWithPeriodGetSharePointActivityPagesWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityPagesWithPeriodGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityUserCountsWithPeriod provides operations to call the getSharePointActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointActivityUserCountsWithPeriodGetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityUserCountsWithPeriodGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityUserDetailWithDate provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetSharePointActivityUserDetailWithDateGetSharePointActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityUserDetailWithDateGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSharePointActivityUserDetailWithPeriod provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointActivityUserDetailWithPeriodGetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointActivityUserDetailWithPeriodGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageDetailWithDate provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetSharePointSiteUsageDetailWithDateGetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageDetailWithDateGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSharePointSiteUsageDetailWithPeriod provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsageDetailWithPeriodGetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageDetailWithPeriodGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriod provides operations to call the getSharePointSiteUsageFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriodGetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriodGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsagePagesWithPeriod provides operations to call the getSharePointSiteUsagePages method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsagePagesWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsagePagesWithPeriodGetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsagePagesWithPeriodGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriod provides operations to call the getSharePointSiteUsageSiteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageStorageWithPeriod provides operations to call the getSharePointSiteUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageStorageWithPeriod(period *string)(*ReportsMicrosoftGraphGetSharePointSiteUsageStorageWithPeriodGetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSharePointSiteUsageStorageWithPeriodGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriod provides operations to call the getSkypeForBusinessActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriodGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriodGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDate provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriod provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriodGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriodGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriodGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriodGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDate provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriodGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriodGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriodGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriodGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriodGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriodGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriodGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriodGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDate provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetTeamsDeviceUsageUserDetailWithDateGetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageUserDetailWithDateGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriod provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityCountsWithPeriod provides operations to call the getTeamsTeamActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsTeamActivityCountsWithPeriodGetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamActivityCountsWithPeriodGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityDetailWithDate provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetTeamsTeamActivityDetailWithDateGetTeamsTeamActivityDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamActivityDetailWithDateGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsTeamActivityDetailWithPeriod provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsTeamActivityDetailWithPeriodGetTeamsTeamActivityDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamActivityDetailWithPeriodGetTeamsTeamActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriod provides operations to call the getTeamsTeamActivityDistributionCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriodGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriodGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamCountsWithPeriod provides operations to call the getTeamsTeamCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsTeamCountsWithPeriodGetTeamsTeamCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsTeamCountsWithPeriodGetTeamsTeamCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityCountsWithPeriod provides operations to call the getTeamsUserActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityCountsWithPeriodGetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityCountsWithPeriodGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriodGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriodGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriod provides operations to call the getTeamsUserActivityTotalCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriodGetTeamsUserActivityTotalCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriodGetTeamsUserActivityTotalCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriod provides operations to call the getTeamsUserActivityTotalDistributionCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriodGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriodGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriodGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriodGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriod provides operations to call the getTeamsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityUserCountsWithPeriodGetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityUserCountsWithPeriodGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityUserDetailWithDate provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetTeamsUserActivityUserDetailWithDateGetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityUserDetailWithDateGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriod provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetTeamsUserActivityUserDetailWithPeriodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetTeamsUserActivityUserDetailWithPeriodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime provides operations to call the getUserArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*ReportsMicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewReportsMicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime, userId)
}
// MicrosoftGraphGetYammerActivityCountsWithPeriod provides operations to call the getYammerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerActivityCountsWithPeriodGetYammerActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerActivityCountsWithPeriodGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerActivityUserCountsWithPeriod provides operations to call the getYammerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerActivityUserCountsWithPeriodGetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerActivityUserCountsWithPeriodGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerActivityUserDetailWithDate provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetYammerActivityUserDetailWithDateGetYammerActivityUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerActivityUserDetailWithDateGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerActivityUserDetailWithPeriod provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerActivityUserDetailWithPeriodGetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerActivityUserDetailWithPeriodGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriod provides operations to call the getYammerDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageUserDetailWithDate provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetYammerDeviceUsageUserDetailWithDateGetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerDeviceUsageUserDetailWithDateGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriod provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriodGetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriodGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityCountsWithPeriod provides operations to call the getYammerGroupsActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerGroupsActivityCountsWithPeriodGetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerGroupsActivityCountsWithPeriodGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityDetailWithDate provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*ReportsMicrosoftGraphGetYammerGroupsActivityDetailWithDateGetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerGroupsActivityDetailWithDateGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerGroupsActivityDetailWithPeriod provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityDetailWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerGroupsActivityDetailWithPeriodGetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerGroupsActivityDetailWithPeriodGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriod provides operations to call the getYammerGroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*ReportsMicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriodGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriodGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ReportsMicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ReportsMicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureDetails provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureDetails()(*ReportsMicrosoftGraphManagedDeviceEnrollmentFailureDetailsManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentFailureDetailsManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ReportsMicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureTrends provides operations to call the managedDeviceEnrollmentFailureTrends method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureTrends()(*ReportsMicrosoftGraphManagedDeviceEnrollmentFailureTrendsManagedDeviceEnrollmentFailureTrendsRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentFailureTrendsManagedDeviceEnrollmentFailureTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentTopFailures provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentTopFailures()(*ReportsMicrosoftGraphManagedDeviceEnrollmentTopFailuresManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentTopFailuresManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriod provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*ReportsMicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return NewReportsMicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
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
    idPtr := &id
    return NewReportsMonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsMonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsMonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    idPtr := &id
    return NewReportsUserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
