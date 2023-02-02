package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsRequestBuilder provides operations to manage the reportRoot singleton.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderGetQueryParameters get reports
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
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummary()(*ApplicationSignInDetailedSummaryRequestBuilder) {
    return NewApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApplicationSignInDetailedSummaryById provides operations to manage the applicationSignInDetailedSummary property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*ApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AuthenticationMethods provides operations to manage the authenticationMethods property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) AuthenticationMethods()(*AuthenticationMethodsRequestBuilder) {
    return NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports{?%24select,%24expand}";
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
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetails()(*CredentialUserRegistrationDetailsRequestBuilder) {
    return NewCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CredentialUserRegistrationDetailsById provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*CredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DailyPrintUsage provides operations to manage the dailyPrintUsage property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsage()(*DailyPrintUsageRequestBuilder) {
    return NewDailyPrintUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageById provides operations to manage the dailyPrintUsage property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageById(id string)(*DailyPrintUsagePrintUsageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDailyPrintUsagePrintUsageItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DailyPrintUsageByPrinter provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*DailyPrintUsageByPrinterRequestBuilder) {
    return NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageByPrinterById provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*DailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DailyPrintUsageByUser provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*DailyPrintUsageByUserRequestBuilder) {
    return NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageByUserById provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*DailyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDailyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DailyPrintUsageSummariesByPrinter provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*DailyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageSummariesByPrinterById provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*DailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DailyPrintUsageSummariesByUser provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*DailyPrintUsageSummariesByUserRequestBuilder) {
    return NewDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DailyPrintUsageSummariesByUserById provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*DailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Get get reports
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
func (m *ReportsRequestBuilder) MicrosoftGraphDeviceConfigurationDeviceActivity()(*MicrosoftGraphDeviceConfigurationDeviceActivityDeviceConfigurationDeviceActivityRequestBuilder) {
    return NewMicrosoftGraphDeviceConfigurationDeviceActivityDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphDeviceConfigurationUserActivity provides operations to call the deviceConfigurationUserActivity method.
func (m *ReportsRequestBuilder) MicrosoftGraphDeviceConfigurationUserActivity()(*MicrosoftGraphDeviceConfigurationUserActivityDeviceConfigurationUserActivityRequestBuilder) {
    return NewMicrosoftGraphDeviceConfigurationUserActivityDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationRepeatOffenders provides operations to call the getAttackSimulationRepeatOffenders method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationRepeatOffenders()(*MicrosoftGraphGetAttackSimulationRepeatOffendersGetAttackSimulationRepeatOffendersRequestBuilder) {
    return NewMicrosoftGraphGetAttackSimulationRepeatOffendersGetAttackSimulationRepeatOffendersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationSimulationUserCoverage provides operations to call the getAttackSimulationSimulationUserCoverage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationSimulationUserCoverage()(*MicrosoftGraphGetAttackSimulationSimulationUserCoverageGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return NewMicrosoftGraphGetAttackSimulationSimulationUserCoverageGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAttackSimulationTrainingUserCoverage provides operations to call the getAttackSimulationTrainingUserCoverage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAttackSimulationTrainingUserCoverage()(*MicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return NewMicrosoftGraphGetAttackSimulationTrainingUserCoverageGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriod provides operations to call the getAzureADApplicationSignInSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriod(period *string)(*MicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetAzureADApplicationSignInSummaryWithPeriodGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserDistributionUserCountsWithPeriod provides operations to call the getBrowserDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserDistributionUserCountsWithPeriod(period *string)(*MicrosoftGraphGetBrowserDistributionUserCountsWithPeriodGetBrowserDistributionUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetBrowserDistributionUserCountsWithPeriodGetBrowserDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserUserCountsWithPeriod provides operations to call the getBrowserUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserUserCountsWithPeriod(period *string)(*MicrosoftGraphGetBrowserUserCountsWithPeriodGetBrowserUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetBrowserUserCountsWithPeriodGetBrowserUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetBrowserUserDetailWithPeriod provides operations to call the getBrowserUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetBrowserUserDetailWithPeriod(period *string)(*MicrosoftGraphGetBrowserUserDetailWithPeriodGetBrowserUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetBrowserUserDetailWithPeriodGetBrowserUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetCredentialUsageSummaryWithPeriod provides operations to call the getCredentialUsageSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCredentialUsageSummaryWithPeriod(period *string)(*MicrosoftGraphGetCredentialUsageSummaryWithPeriodGetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetCredentialUsageSummaryWithPeriodGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetCredentialUserRegistrationCount provides operations to call the getCredentialUserRegistrationCount method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetCredentialUserRegistrationCount()(*MicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilder) {
    return NewMicrosoftGraphGetCredentialUserRegistrationCountGetCredentialUserRegistrationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetEmailActivityCountsWithPeriod provides operations to call the getEmailActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailActivityCountsWithPeriodGetEmailActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailActivityCountsWithPeriodGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailActivityUserCountsWithPeriod provides operations to call the getEmailActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailActivityUserCountsWithPeriodGetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailActivityUserCountsWithPeriodGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailActivityUserDetailWithDate provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetEmailActivityUserDetailWithDateGetEmailActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetEmailActivityUserDetailWithDateGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetEmailActivityUserDetailWithPeriod provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetEmailActivityUserDetailWithPeriodGetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailActivityUserDetailWithPeriodGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriod provides operations to call the getEmailAppUsageAppsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriodGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageAppsUserCountsWithPeriodGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageUserCountsWithPeriod provides operations to call the getEmailAppUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailAppUsageUserCountsWithPeriodGetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageUserCountsWithPeriodGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageUserDetailWithDate provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetEmailAppUsageUserDetailWithDateGetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageUserDetailWithDateGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetEmailAppUsageUserDetailWithPeriod provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageUserDetailWithPeriod(period *string)(*MicrosoftGraphGetEmailAppUsageUserDetailWithPeriodGetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageUserDetailWithPeriodGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriod provides operations to call the getEmailAppUsageVersionsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*MicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriodGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetEmailAppUsageVersionsUserCountsWithPeriodGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityCountsWithPeriod provides operations to call the getFormsUserActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetFormsUserActivityCountsWithPeriodGetFormsUserActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetFormsUserActivityCountsWithPeriodGetFormsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityUserCountsWithPeriod provides operations to call the getFormsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetFormsUserActivityUserCountsWithPeriodGetFormsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodGetFormsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetFormsUserActivityUserDetailWithDate provides operations to call the getFormsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetFormsUserActivityUserDetailWithDateGetFormsUserActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetFormsUserActivityUserDetailWithDateGetFormsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetFormsUserActivityUserDetailWithPeriod provides operations to call the getFormsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetFormsUserActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetFormsUserActivityUserDetailWithPeriodGetFormsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetFormsUserActivityUserDetailWithPeriodGetFormsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime provides operations to call the getGroupArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*MicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewMicrosoftGraphGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, groupId, startDateTime)
}
// MicrosoftGraphGetM365AppPlatformUserCountsWithPeriod provides operations to call the getM365AppPlatformUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppPlatformUserCountsWithPeriod(period *string)(*MicrosoftGraphGetM365AppPlatformUserCountsWithPeriodGetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetM365AppPlatformUserCountsWithPeriodGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetM365AppUserCountsWithPeriod provides operations to call the getM365AppUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserCountsWithPeriod(period *string)(*MicrosoftGraphGetM365AppUserCountsWithPeriodGetM365AppUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetM365AppUserCountsWithPeriodGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetM365AppUserDetailWithDate provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetM365AppUserDetailWithDateGetM365AppUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetM365AppUserDetailWithDateGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetM365AppUserDetailWithPeriod provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetM365AppUserDetailWithPeriod(period *string)(*MicrosoftGraphGetM365AppUserDetailWithPeriodGetM365AppUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetM365AppUserDetailWithPeriodGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageDetailWithPeriod provides operations to call the getMailboxUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageDetailWithPeriod(period *string)(*MicrosoftGraphGetMailboxUsageDetailWithPeriodGetMailboxUsageDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetMailboxUsageDetailWithPeriodGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriod provides operations to call the getMailboxUsageMailboxCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriod(period *string)(*MicrosoftGraphGetMailboxUsageMailboxCountsWithPeriodGetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetMailboxUsageMailboxCountsWithPeriodGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriod provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*MicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetMailboxUsageQuotaStatusMailboxCountsWithPeriodGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetMailboxUsageStorageWithPeriod provides operations to call the getMailboxUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetMailboxUsageStorageWithPeriod(period *string)(*MicrosoftGraphGetMailboxUsageStorageWithPeriodGetMailboxUsageStorageWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetMailboxUsageStorageWithPeriodGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ActivationCounts provides operations to call the getOffice365ActivationCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationCounts()(*MicrosoftGraphGetOffice365ActivationCountsGetOffice365ActivationCountsRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActivationCountsGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActivationsUserCounts provides operations to call the getOffice365ActivationsUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationsUserCounts()(*MicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActivationsUserCountsGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActivationsUserDetail provides operations to call the getOffice365ActivationsUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActivationsUserDetail()(*MicrosoftGraphGetOffice365ActivationsUserDetailGetOffice365ActivationsUserDetailRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActivationsUserDetailGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOffice365ActiveUserCountsWithPeriod provides operations to call the getOffice365ActiveUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365ActiveUserCountsWithPeriodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActiveUserCountsWithPeriodGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ActiveUserDetailWithDate provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetOffice365ActiveUserDetailWithDateGetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActiveUserDetailWithDateGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOffice365ActiveUserDetailWithPeriod provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ActiveUserDetailWithPeriod(period *string)(*MicrosoftGraphGetOffice365ActiveUserDetailWithPeriodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ActiveUserDetailWithPeriodGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriod provides operations to call the getOffice365GroupsActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityCountsWithPeriodGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityCountsWithPeriodGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityDetailWithDate provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetOffice365GroupsActivityDetailWithDateGetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityDetailWithDateGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriod provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityDetailWithPeriodGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriod provides operations to call the getOffice365GroupsActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityFileCountsWithPeriodGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriod provides operations to call the getOffice365GroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityGroupCountsWithPeriodGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriod provides operations to call the getOffice365GroupsActivityStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriod(period *string)(*MicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365GroupsActivityStorageWithPeriodGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOffice365ServicesUserCountsWithPeriod provides operations to call the getOffice365ServicesUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOffice365ServicesUserCountsWithPeriod(period *string)(*MicrosoftGraphGetOffice365ServicesUserCountsWithPeriodGetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOffice365ServicesUserCountsWithPeriodGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityFileCountsWithPeriod provides operations to call the getOneDriveActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityFileCountsWithPeriod(period *string)(*MicrosoftGraphGetOneDriveActivityFileCountsWithPeriodGetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveActivityFileCountsWithPeriodGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityUserCountsWithPeriod provides operations to call the getOneDriveActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetOneDriveActivityUserCountsWithPeriodGetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveActivityUserCountsWithPeriodGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveActivityUserDetailWithDate provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetOneDriveActivityUserDetailWithDateGetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveActivityUserDetailWithDateGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOneDriveActivityUserDetailWithPeriod provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetOneDriveActivityUserDetailWithPeriodGetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveActivityUserDetailWithPeriodGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriod provides operations to call the getOneDriveUsageAccountCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriod(period *string)(*MicrosoftGraphGetOneDriveUsageAccountCountsWithPeriodGetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageAccountCountsWithPeriodGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageAccountDetailWithDate provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetOneDriveUsageAccountDetailWithDateGetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageAccountDetailWithDateGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriod provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriod(period *string)(*MicrosoftGraphGetOneDriveUsageAccountDetailWithPeriodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageAccountDetailWithPeriodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageFileCountsWithPeriod provides operations to call the getOneDriveUsageFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageFileCountsWithPeriod(period *string)(*MicrosoftGraphGetOneDriveUsageFileCountsWithPeriodGetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageFileCountsWithPeriodGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetOneDriveUsageStorageWithPeriod provides operations to call the getOneDriveUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetOneDriveUsageStorageWithPeriod(period *string)(*MicrosoftGraphGetOneDriveUsageStorageWithPeriodGetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetOneDriveUsageStorageWithPeriodGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime provides operations to call the getPrinterArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*MicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewMicrosoftGraphGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, printerId, startDateTime)
}
// MicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriod provides operations to call the getRelyingPartyDetailedSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriod(period *string)(*MicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriodGetRelyingPartyDetailedSummaryWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetRelyingPartyDetailedSummaryWithPeriodGetRelyingPartyDetailedSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityFileCountsWithPeriod provides operations to call the getSharePointActivityFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityFileCountsWithPeriod(period *string)(*MicrosoftGraphGetSharePointActivityFileCountsWithPeriodGetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityFileCountsWithPeriodGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityPagesWithPeriod provides operations to call the getSharePointActivityPages method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityPagesWithPeriod(period *string)(*MicrosoftGraphGetSharePointActivityPagesWithPeriodGetSharePointActivityPagesWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityPagesWithPeriodGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityUserCountsWithPeriod provides operations to call the getSharePointActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSharePointActivityUserCountsWithPeriodGetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityUserCountsWithPeriodGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointActivityUserDetailWithDate provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetSharePointActivityUserDetailWithDateGetSharePointActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityUserDetailWithDateGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSharePointActivityUserDetailWithPeriod provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetSharePointActivityUserDetailWithPeriodGetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointActivityUserDetailWithPeriodGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageDetailWithDate provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetSharePointSiteUsageDetailWithDateGetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageDetailWithDateGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSharePointSiteUsageDetailWithPeriod provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageDetailWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsageDetailWithPeriodGetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageDetailWithPeriodGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriod provides operations to call the getSharePointSiteUsageFileCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriodGetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageFileCountsWithPeriodGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsagePagesWithPeriod provides operations to call the getSharePointSiteUsagePages method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsagePagesWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsagePagesWithPeriodGetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsagePagesWithPeriodGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriod provides operations to call the getSharePointSiteUsageSiteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageSiteCountsWithPeriodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSharePointSiteUsageStorageWithPeriod provides operations to call the getSharePointSiteUsageStorage method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSharePointSiteUsageStorageWithPeriod(period *string)(*MicrosoftGraphGetSharePointSiteUsageStorageWithPeriodGetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSharePointSiteUsageStorageWithPeriodGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriod provides operations to call the getSkypeForBusinessActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessActivityCountsWithPeriodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriodGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessActivityUserCountsWithPeriodGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDate provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithDateGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriod provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriodGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessActivityUserDetailWithPeriodGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriodGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageUserCountsWithPeriodGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDate provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithDateGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriodGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageUserDetailWithPeriodGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessOrganizerActivityCountsWithPeriodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriodGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessParticipantActivityUserCountsWithPeriodGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageDistributionUserCountsWithPeriodGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriodGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageTotalUserCountsWithPeriodGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriodGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageUserCountsWithPeriodGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDate provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetTeamsDeviceUsageUserDetailWithDateGetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageUserDetailWithDateGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriod provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*MicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsDeviceUsageUserDetailWithPeriodGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityCountsWithPeriod provides operations to call the getTeamsTeamActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsTeamActivityCountsWithPeriodGetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsTeamActivityCountsWithPeriodGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityDetailWithDate provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetTeamsTeamActivityDetailWithDateGetTeamsTeamActivityDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetTeamsTeamActivityDetailWithDateGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsTeamActivityDetailWithPeriod provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDetailWithPeriod(period *string)(*MicrosoftGraphGetTeamsTeamActivityDetailWithPeriodGetTeamsTeamActivityDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsTeamActivityDetailWithPeriodGetTeamsTeamActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriod provides operations to call the getTeamsTeamActivityDistributionCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriodGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsTeamActivityDistributionCountsWithPeriodGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsTeamCountsWithPeriod provides operations to call the getTeamsTeamCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsTeamCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsTeamCountsWithPeriodGetTeamsTeamCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsTeamCountsWithPeriodGetTeamsTeamCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityCountsWithPeriod provides operations to call the getTeamsUserActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityCountsWithPeriodGetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityCountsWithPeriodGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityDistributionTotalUserCountsWithPeriodGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriodGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityDistributionUserCountsWithPeriodGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriod provides operations to call the getTeamsUserActivityTotalCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriodGetTeamsUserActivityTotalCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityTotalCountsWithPeriodGetTeamsUserActivityTotalCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriod provides operations to call the getTeamsUserActivityTotalDistributionCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriodGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityTotalDistributionCountsWithPeriodGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityTotalUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriodGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityTotalUserCountsWithPeriodGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriod provides operations to call the getTeamsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityUserCountsWithPeriodGetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityUserCountsWithPeriodGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetTeamsUserActivityUserDetailWithDate provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetTeamsUserActivityUserDetailWithDateGetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityUserDetailWithDateGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriod provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetTeamsUserActivityUserDetailWithPeriodGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetTeamsUserActivityUserDetailWithPeriodGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime provides operations to call the getUserArchivedPrintJobs method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*MicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewMicrosoftGraphGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime, userId)
}
// MicrosoftGraphGetYammerActivityCountsWithPeriod provides operations to call the getYammerActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerActivityCountsWithPeriodGetYammerActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerActivityCountsWithPeriodGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerActivityUserCountsWithPeriod provides operations to call the getYammerActivityUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerActivityUserCountsWithPeriodGetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerActivityUserCountsWithPeriodGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerActivityUserDetailWithDate provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetYammerActivityUserDetailWithDateGetYammerActivityUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetYammerActivityUserDetailWithDateGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerActivityUserDetailWithPeriod provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerActivityUserDetailWithPeriod(period *string)(*MicrosoftGraphGetYammerActivityUserDetailWithPeriodGetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerActivityUserDetailWithPeriodGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriod provides operations to call the getYammerDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerDeviceUsageUserCountsWithPeriodGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerDeviceUsageUserDetailWithDate provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetYammerDeviceUsageUserDetailWithDateGetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetYammerDeviceUsageUserDetailWithDateGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriod provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriod(period *string)(*MicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriodGetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerDeviceUsageUserDetailWithPeriodGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityCountsWithPeriod provides operations to call the getYammerGroupsActivityCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerGroupsActivityCountsWithPeriodGetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerGroupsActivityCountsWithPeriodGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityDetailWithDate provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*MicrosoftGraphGetYammerGroupsActivityDetailWithDateGetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return NewMicrosoftGraphGetYammerGroupsActivityDetailWithDateGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date)
}
// MicrosoftGraphGetYammerGroupsActivityDetailWithPeriod provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityDetailWithPeriod(period *string)(*MicrosoftGraphGetYammerGroupsActivityDetailWithPeriodGetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerGroupsActivityDetailWithPeriodGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriod provides operations to call the getYammerGroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*MicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriodGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewMicrosoftGraphGetYammerGroupsActivityGroupCountsWithPeriodGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*MicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentSummary method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*MicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureDetails provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureDetails()(*MicrosoftGraphManagedDeviceEnrollmentFailureDetailsManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentFailureDetailsManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*MicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top)
}
// MicrosoftGraphManagedDeviceEnrollmentFailureTrends provides operations to call the managedDeviceEnrollmentFailureTrends method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentFailureTrends()(*MicrosoftGraphManagedDeviceEnrollmentFailureTrendsManagedDeviceEnrollmentFailureTrendsRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentFailureTrendsManagedDeviceEnrollmentFailureTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentTopFailures provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentTopFailures()(*MicrosoftGraphManagedDeviceEnrollmentTopFailuresManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentTopFailuresManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriod provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*MicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return NewMicrosoftGraphManagedDeviceEnrollmentTopFailuresWithPeriodManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period)
}
// MonthlyPrintUsageByPrinter provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*MonthlyPrintUsageByPrinterRequestBuilder) {
    return NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageByPrinterById provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*MonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MonthlyPrintUsageByUser provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*MonthlyPrintUsageByUserRequestBuilder) {
    return NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageByUserById provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*MonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MonthlyPrintUsageSummariesByPrinter provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*MonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageSummariesByPrinterById provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*MonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MonthlyPrintUsageSummariesByUser provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*MonthlyPrintUsageSummariesByUserRequestBuilder) {
    return NewMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MonthlyPrintUsageSummariesByUserById provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Patch update reports
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
func (m *ReportsRequestBuilder) Security()(*SecurityRequestBuilder) {
    return NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation get reports
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
// ToPatchRequestInformation update reports
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
func (m *ReportsRequestBuilder) UserCredentialUsageDetails()(*UserCredentialUsageDetailsRequestBuilder) {
    return NewUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserCredentialUsageDetailsById provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*UserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
