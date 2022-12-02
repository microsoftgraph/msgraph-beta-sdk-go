package print

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrintReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.print entity.
type PrintReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrintReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrintReportsRequestBuilderGetQueryParameters get reports from print
type PrintReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrintReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrintReportsRequestBuilderGetQueryParameters
}
// PrintReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrintReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationSignInDetailedSummary provides operations to manage the applicationSignInDetailedSummary property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) ApplicationSignInDetailedSummary()(*PrintReportsApplicationSignInDetailedSummaryRequestBuilder) {
    return NewPrintReportsApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationSignInDetailedSummaryById provides operations to manage the applicationSignInDetailedSummary property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*PrintReportsApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationSignInDetailedSummary%2Did"] = id
    }
    return NewPrintReportsApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationMethods provides operations to manage the authenticationMethods property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) AuthenticationMethods()(*PrintReportsAuthenticationMethodsRequestBuilder) {
    return NewPrintReportsAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrintReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewPrintReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintReportsRequestBuilder) {
    m := &PrintReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewPrintReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrintReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reports for print
func (m *PrintReportsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PrintReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get reports from print
func (m *PrintReportsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PrintReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property reports in print
func (m *PrintReportsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *PrintReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CredentialUserRegistrationDetails provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) CredentialUserRegistrationDetails()(*PrintReportsCredentialUserRegistrationDetailsRequestBuilder) {
    return NewPrintReportsCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationDetailsById provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*PrintReportsCredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationDetails%2Did"] = id
    }
    return NewPrintReportsCredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsage provides operations to manage the dailyPrintUsage property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsage()(*PrintReportsDailyPrintUsageRequestBuilder) {
    return NewPrintReportsDailyPrintUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageById provides operations to manage the dailyPrintUsage property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageById(id string)(*PrintReportsDailyPrintUsagePrintUsageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsage%2Did"] = id
    }
    return NewPrintReportsDailyPrintUsagePrintUsageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByPrinter provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageByPrinter()(*PrintReportsDailyPrintUsageByPrinterRequestBuilder) {
    return NewPrintReportsDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByPrinterById provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*PrintReportsDailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewPrintReportsDailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByUser provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageByUser()(*PrintReportsDailyPrintUsageByUserRequestBuilder) {
    return NewPrintReportsDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByUserById provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*PrintReportsDailyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewPrintReportsDailyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinter provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*PrintReportsDailyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewPrintReportsDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinterById provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*PrintReportsDailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewPrintReportsDailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByUser provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*PrintReportsDailyPrintUsageSummariesByUserRequestBuilder) {
    return NewPrintReportsDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByUserById provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*PrintReportsDailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewPrintReportsDailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property reports for print
func (m *PrintReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrintReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceConfigurationDeviceActivity provides operations to call the deviceConfigurationDeviceActivity method.
func (m *PrintReportsRequestBuilder) DeviceConfigurationDeviceActivity()(*PrintReportsDeviceConfigurationDeviceActivityRequestBuilder) {
    return NewPrintReportsDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationUserActivity provides operations to call the deviceConfigurationUserActivity method.
func (m *PrintReportsRequestBuilder) DeviceConfigurationUserActivity()(*PrintReportsDeviceConfigurationUserActivityRequestBuilder) {
    return NewPrintReportsDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get reports from print
func (m *PrintReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrintReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// GetAttackSimulationRepeatOffenders provides operations to call the getAttackSimulationRepeatOffenders method.
func (m *PrintReportsRequestBuilder) GetAttackSimulationRepeatOffenders()(*PrintReportsGetAttackSimulationRepeatOffendersRequestBuilder) {
    return NewPrintReportsGetAttackSimulationRepeatOffendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationSimulationUserCoverage provides operations to call the getAttackSimulationSimulationUserCoverage method.
func (m *PrintReportsRequestBuilder) GetAttackSimulationSimulationUserCoverage()(*PrintReportsGetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return NewPrintReportsGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationTrainingUserCoverage provides operations to call the getAttackSimulationTrainingUserCoverage method.
func (m *PrintReportsRequestBuilder) GetAttackSimulationTrainingUserCoverage()(*PrintReportsGetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return NewPrintReportsGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAzureADApplicationSignInSummaryWithPeriod provides operations to call the getAzureADApplicationSignInSummary method.
func (m *PrintReportsRequestBuilder) GetAzureADApplicationSignInSummaryWithPeriod(period *string)(*PrintReportsGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return NewPrintReportsGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserDistributionUserCountsWithPeriod provides operations to call the getBrowserDistributionUserCounts method.
func (m *PrintReportsRequestBuilder) GetBrowserDistributionUserCountsWithPeriod(period *string)(*PrintReportsGetBrowserDistributionUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetBrowserDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserUserCountsWithPeriod provides operations to call the getBrowserUserCounts method.
func (m *PrintReportsRequestBuilder) GetBrowserUserCountsWithPeriod(period *string)(*PrintReportsGetBrowserUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetBrowserUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserUserDetailWithPeriod provides operations to call the getBrowserUserDetail method.
func (m *PrintReportsRequestBuilder) GetBrowserUserDetailWithPeriod(period *string)(*PrintReportsGetBrowserUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetBrowserUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUsageSummaryWithPeriod provides operations to call the getCredentialUsageSummary method.
func (m *PrintReportsRequestBuilder) GetCredentialUsageSummaryWithPeriod(period *string)(*PrintReportsGetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return NewPrintReportsGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUserRegistrationCount provides operations to call the getCredentialUserRegistrationCount method.
func (m *PrintReportsRequestBuilder) GetCredentialUserRegistrationCount()(*PrintReportsGetCredentialUserRegistrationCountRequestBuilder) {
    return NewPrintReportsGetCredentialUserRegistrationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEmailActivityCountsWithPeriod provides operations to call the getEmailActivityCounts method.
func (m *PrintReportsRequestBuilder) GetEmailActivityCountsWithPeriod(period *string)(*PrintReportsGetEmailActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserCountsWithPeriod provides operations to call the getEmailActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetEmailActivityUserCountsWithPeriod(period *string)(*PrintReportsGetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserDetailWithDate provides operations to call the getEmailActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetEmailActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetEmailActivityUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailActivityUserDetailWithPeriod provides operations to call the getEmailActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetEmailActivityUserDetailWithPeriod(period *string)(*PrintReportsGetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageAppsUserCountsWithPeriod provides operations to call the getEmailAppUsageAppsUserCounts method.
func (m *PrintReportsRequestBuilder) GetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*PrintReportsGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserCountsWithPeriod provides operations to call the getEmailAppUsageUserCounts method.
func (m *PrintReportsRequestBuilder) GetEmailAppUsageUserCountsWithPeriod(period *string)(*PrintReportsGetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserDetailWithDate provides operations to call the getEmailAppUsageUserDetail method.
func (m *PrintReportsRequestBuilder) GetEmailAppUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailAppUsageUserDetailWithPeriod provides operations to call the getEmailAppUsageUserDetail method.
func (m *PrintReportsRequestBuilder) GetEmailAppUsageUserDetailWithPeriod(period *string)(*PrintReportsGetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageVersionsUserCountsWithPeriod provides operations to call the getEmailAppUsageVersionsUserCounts method.
func (m *PrintReportsRequestBuilder) GetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*PrintReportsGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime provides operations to call the getGroupArchivedPrintJobs method.
func (m *PrintReportsRequestBuilder) GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*PrintReportsGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewPrintReportsGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, groupId, startDateTime);
}
// GetM365AppPlatformUserCountsWithPeriod provides operations to call the getM365AppPlatformUserCounts method.
func (m *PrintReportsRequestBuilder) GetM365AppPlatformUserCountsWithPeriod(period *string)(*PrintReportsGetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserCountsWithPeriod provides operations to call the getM365AppUserCounts method.
func (m *PrintReportsRequestBuilder) GetM365AppUserCountsWithPeriod(period *string)(*PrintReportsGetM365AppUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserDetailWithDate provides operations to call the getM365AppUserDetail method.
func (m *PrintReportsRequestBuilder) GetM365AppUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetM365AppUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetM365AppUserDetailWithPeriod provides operations to call the getM365AppUserDetail method.
func (m *PrintReportsRequestBuilder) GetM365AppUserDetailWithPeriod(period *string)(*PrintReportsGetM365AppUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageDetailWithPeriod provides operations to call the getMailboxUsageDetail method.
func (m *PrintReportsRequestBuilder) GetMailboxUsageDetailWithPeriod(period *string)(*PrintReportsGetMailboxUsageDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageMailboxCountsWithPeriod provides operations to call the getMailboxUsageMailboxCounts method.
func (m *PrintReportsRequestBuilder) GetMailboxUsageMailboxCountsWithPeriod(period *string)(*PrintReportsGetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageQuotaStatusMailboxCountsWithPeriod provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
func (m *PrintReportsRequestBuilder) GetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*PrintReportsGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageStorageWithPeriod provides operations to call the getMailboxUsageStorage method.
func (m *PrintReportsRequestBuilder) GetMailboxUsageStorageWithPeriod(period *string)(*PrintReportsGetMailboxUsageStorageWithPeriodRequestBuilder) {
    return NewPrintReportsGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActivationCounts provides operations to call the getOffice365ActivationCounts method.
func (m *PrintReportsRequestBuilder) GetOffice365ActivationCounts()(*PrintReportsGetOffice365ActivationCountsRequestBuilder) {
    return NewPrintReportsGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserCounts provides operations to call the getOffice365ActivationsUserCounts method.
func (m *PrintReportsRequestBuilder) GetOffice365ActivationsUserCounts()(*PrintReportsGetOffice365ActivationsUserCountsRequestBuilder) {
    return NewPrintReportsGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserDetail provides operations to call the getOffice365ActivationsUserDetail method.
func (m *PrintReportsRequestBuilder) GetOffice365ActivationsUserDetail()(*PrintReportsGetOffice365ActivationsUserDetailRequestBuilder) {
    return NewPrintReportsGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActiveUserCountsWithPeriod provides operations to call the getOffice365ActiveUserCounts method.
func (m *PrintReportsRequestBuilder) GetOffice365ActiveUserCountsWithPeriod(period *string)(*PrintReportsGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActiveUserDetailWithDate provides operations to call the getOffice365ActiveUserDetail method.
func (m *PrintReportsRequestBuilder) GetOffice365ActiveUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365ActiveUserDetailWithPeriod provides operations to call the getOffice365ActiveUserDetail method.
func (m *PrintReportsRequestBuilder) GetOffice365ActiveUserDetailWithPeriod(period *string)(*PrintReportsGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityCountsWithPeriod provides operations to call the getOffice365GroupsActivityCounts method.
func (m *PrintReportsRequestBuilder) GetOffice365GroupsActivityCountsWithPeriod(period *string)(*PrintReportsGetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityDetailWithDate provides operations to call the getOffice365GroupsActivityDetail method.
func (m *PrintReportsRequestBuilder) GetOffice365GroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return NewPrintReportsGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365GroupsActivityDetailWithPeriod provides operations to call the getOffice365GroupsActivityDetail method.
func (m *PrintReportsRequestBuilder) GetOffice365GroupsActivityDetailWithPeriod(period *string)(*PrintReportsGetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityFileCountsWithPeriod provides operations to call the getOffice365GroupsActivityFileCounts method.
func (m *PrintReportsRequestBuilder) GetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*PrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityGroupCountsWithPeriod provides operations to call the getOffice365GroupsActivityGroupCounts method.
func (m *PrintReportsRequestBuilder) GetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*PrintReportsGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityStorageWithPeriod provides operations to call the getOffice365GroupsActivityStorage method.
func (m *PrintReportsRequestBuilder) GetOffice365GroupsActivityStorageWithPeriod(period *string)(*PrintReportsGetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return NewPrintReportsGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ServicesUserCountsWithPeriod provides operations to call the getOffice365ServicesUserCounts method.
func (m *PrintReportsRequestBuilder) GetOffice365ServicesUserCountsWithPeriod(period *string)(*PrintReportsGetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityFileCountsWithPeriod provides operations to call the getOneDriveActivityFileCounts method.
func (m *PrintReportsRequestBuilder) GetOneDriveActivityFileCountsWithPeriod(period *string)(*PrintReportsGetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserCountsWithPeriod provides operations to call the getOneDriveActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetOneDriveActivityUserCountsWithPeriod(period *string)(*PrintReportsGetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserDetailWithDate provides operations to call the getOneDriveActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetOneDriveActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveActivityUserDetailWithPeriod provides operations to call the getOneDriveActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetOneDriveActivityUserDetailWithPeriod(period *string)(*PrintReportsGetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountCountsWithPeriod provides operations to call the getOneDriveUsageAccountCounts method.
func (m *PrintReportsRequestBuilder) GetOneDriveUsageAccountCountsWithPeriod(period *string)(*PrintReportsGetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountDetailWithDate provides operations to call the getOneDriveUsageAccountDetail method.
func (m *PrintReportsRequestBuilder) GetOneDriveUsageAccountDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return NewPrintReportsGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveUsageAccountDetailWithPeriod provides operations to call the getOneDriveUsageAccountDetail method.
func (m *PrintReportsRequestBuilder) GetOneDriveUsageAccountDetailWithPeriod(period *string)(*PrintReportsGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageFileCountsWithPeriod provides operations to call the getOneDriveUsageFileCounts method.
func (m *PrintReportsRequestBuilder) GetOneDriveUsageFileCountsWithPeriod(period *string)(*PrintReportsGetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageStorageWithPeriod provides operations to call the getOneDriveUsageStorage method.
func (m *PrintReportsRequestBuilder) GetOneDriveUsageStorageWithPeriod(period *string)(*PrintReportsGetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return NewPrintReportsGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime provides operations to call the getPrinterArchivedPrintJobs method.
func (m *PrintReportsRequestBuilder) GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*PrintReportsGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewPrintReportsGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, printerId, startDateTime);
}
// GetRelyingPartyDetailedSummaryWithPeriod provides operations to call the getRelyingPartyDetailedSummary method.
func (m *PrintReportsRequestBuilder) GetRelyingPartyDetailedSummaryWithPeriod(period *string)(*PrintReportsGetRelyingPartyDetailedSummaryWithPeriodRequestBuilder) {
    return NewPrintReportsGetRelyingPartyDetailedSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityFileCountsWithPeriod provides operations to call the getSharePointActivityFileCounts method.
func (m *PrintReportsRequestBuilder) GetSharePointActivityFileCountsWithPeriod(period *string)(*PrintReportsGetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityPagesWithPeriod provides operations to call the getSharePointActivityPages method.
func (m *PrintReportsRequestBuilder) GetSharePointActivityPagesWithPeriod(period *string)(*PrintReportsGetSharePointActivityPagesWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserCountsWithPeriod provides operations to call the getSharePointActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetSharePointActivityUserCountsWithPeriod(period *string)(*PrintReportsGetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserDetailWithDate provides operations to call the getSharePointActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetSharePointActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetSharePointActivityUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointActivityUserDetailWithPeriod provides operations to call the getSharePointActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetSharePointActivityUserDetailWithPeriod(period *string)(*PrintReportsGetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageDetailWithDate provides operations to call the getSharePointSiteUsageDetail method.
func (m *PrintReportsRequestBuilder) GetSharePointSiteUsageDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return NewPrintReportsGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointSiteUsageDetailWithPeriod provides operations to call the getSharePointSiteUsageDetail method.
func (m *PrintReportsRequestBuilder) GetSharePointSiteUsageDetailWithPeriod(period *string)(*PrintReportsGetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageFileCountsWithPeriod provides operations to call the getSharePointSiteUsageFileCounts method.
func (m *PrintReportsRequestBuilder) GetSharePointSiteUsageFileCountsWithPeriod(period *string)(*PrintReportsGetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsagePagesWithPeriod provides operations to call the getSharePointSiteUsagePages method.
func (m *PrintReportsRequestBuilder) GetSharePointSiteUsagePagesWithPeriod(period *string)(*PrintReportsGetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageSiteCountsWithPeriod provides operations to call the getSharePointSiteUsageSiteCounts method.
func (m *PrintReportsRequestBuilder) GetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*PrintReportsGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageStorageWithPeriod provides operations to call the getSharePointSiteUsageStorage method.
func (m *PrintReportsRequestBuilder) GetSharePointSiteUsageStorageWithPeriod(period *string)(*PrintReportsGetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return NewPrintReportsGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityCountsWithPeriod provides operations to call the getSkypeForBusinessActivityCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessActivityCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserDetailWithDate provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessActivityUserDetailWithPeriod provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithDate provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityMinuteCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityMinuteCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityMinuteCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*PrintReportsGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionTotalUserCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod(period *string)(*PrintReportsGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*PrintReportsGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageTotalUserCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsDeviceUsageTotalUserCountsWithPeriod(period *string)(*PrintReportsGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageUserCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*PrintReportsGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserDetailWithDate provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *PrintReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsDeviceUsageUserDetailWithPeriod provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *PrintReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*PrintReportsGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityCountsWithPeriod provides operations to call the getTeamsTeamActivityCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsTeamActivityCountsWithPeriod(period *string)(*PrintReportsGetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityDetailWithDate provides operations to call the getTeamsTeamActivityDetail method.
func (m *PrintReportsRequestBuilder) GetTeamsTeamActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetTeamsTeamActivityDetailWithDateRequestBuilder) {
    return NewPrintReportsGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsTeamActivityDetailWithPeriod provides operations to call the getTeamsTeamActivityDetail method.
func (m *PrintReportsRequestBuilder) GetTeamsTeamActivityDetailWithPeriod(period *string)(*PrintReportsGetTeamsTeamActivityDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsTeamActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityDistributionCountsWithPeriod provides operations to call the getTeamsTeamActivityDistributionCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsTeamActivityDistributionCountsWithPeriod(period *string)(*PrintReportsGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityCountsWithPeriod provides operations to call the getTeamsUserActivityCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityCountsWithPeriod(period *string)(*PrintReportsGetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionTotalUserCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityDistributionTotalUserCountsWithPeriod(period *string)(*PrintReportsGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionUserCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityDistributionUserCountsWithPeriod(period *string)(*PrintReportsGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalCountsWithPeriod provides operations to call the getTeamsUserActivityTotalCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityTotalCountsWithPeriod(period *string)(*PrintReportsGetTeamsUserActivityTotalCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityTotalCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalDistributionCountsWithPeriod provides operations to call the getTeamsUserActivityTotalDistributionCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityTotalDistributionCountsWithPeriod(period *string)(*PrintReportsGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityTotalUserCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityTotalUserCountsWithPeriod(period *string)(*PrintReportsGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserCountsWithPeriod provides operations to call the getTeamsUserActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityUserCountsWithPeriod(period *string)(*PrintReportsGetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserDetailWithDate provides operations to call the getTeamsUserActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsUserActivityUserDetailWithPeriod provides operations to call the getTeamsUserActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetTeamsUserActivityUserDetailWithPeriod(period *string)(*PrintReportsGetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime provides operations to call the getUserArchivedPrintJobs method.
func (m *PrintReportsRequestBuilder) GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*PrintReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewPrintReportsGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime, userId);
}
// GetYammerActivityCountsWithPeriod provides operations to call the getYammerActivityCounts method.
func (m *PrintReportsRequestBuilder) GetYammerActivityCountsWithPeriod(period *string)(*PrintReportsGetYammerActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserCountsWithPeriod provides operations to call the getYammerActivityUserCounts method.
func (m *PrintReportsRequestBuilder) GetYammerActivityUserCountsWithPeriod(period *string)(*PrintReportsGetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserDetailWithDate provides operations to call the getYammerActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetYammerActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetYammerActivityUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerActivityUserDetailWithPeriod provides operations to call the getYammerActivityUserDetail method.
func (m *PrintReportsRequestBuilder) GetYammerActivityUserDetailWithPeriod(period *string)(*PrintReportsGetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
func (m *PrintReportsRequestBuilder) GetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*PrintReportsGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserCountsWithPeriod provides operations to call the getYammerDeviceUsageUserCounts method.
func (m *PrintReportsRequestBuilder) GetYammerDeviceUsageUserCountsWithPeriod(period *string)(*PrintReportsGetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserDetailWithDate provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *PrintReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewPrintReportsGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerDeviceUsageUserDetailWithPeriod provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *PrintReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithPeriod(period *string)(*PrintReportsGetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityCountsWithPeriod provides operations to call the getYammerGroupsActivityCounts method.
func (m *PrintReportsRequestBuilder) GetYammerGroupsActivityCountsWithPeriod(period *string)(*PrintReportsGetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityDetailWithDate provides operations to call the getYammerGroupsActivityDetail method.
func (m *PrintReportsRequestBuilder) GetYammerGroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*PrintReportsGetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return NewPrintReportsGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerGroupsActivityDetailWithPeriod provides operations to call the getYammerGroupsActivityDetail method.
func (m *PrintReportsRequestBuilder) GetYammerGroupsActivityDetailWithPeriod(period *string)(*PrintReportsGetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityGroupCountsWithPeriod provides operations to call the getYammerGroupsActivityGroupCounts method.
func (m *PrintReportsRequestBuilder) GetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*PrintReportsGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewPrintReportsGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentDetails method.
func (m *PrintReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*PrintReportsManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewPrintReportsManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentSummary method.
func (m *PrintReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*PrintReportsManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewPrintReportsManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentFailureDetails provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *PrintReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetails()(*PrintReportsManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return NewPrintReportsManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *PrintReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*PrintReportsManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewPrintReportsManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentFailureTrends provides operations to call the managedDeviceEnrollmentFailureTrends method.
func (m *PrintReportsRequestBuilder) ManagedDeviceEnrollmentFailureTrends()(*PrintReportsManagedDeviceEnrollmentFailureTrendsRequestBuilder) {
    return NewPrintReportsManagedDeviceEnrollmentFailureTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailures provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *PrintReportsRequestBuilder) ManagedDeviceEnrollmentTopFailures()(*PrintReportsManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return NewPrintReportsManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailuresWithPeriod provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *PrintReportsRequestBuilder) ManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*PrintReportsManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return NewPrintReportsManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// MonthlyPrintUsageByPrinter provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*PrintReportsMonthlyPrintUsageByPrinterRequestBuilder) {
    return NewPrintReportsMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByPrinterById provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*PrintReportsMonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewPrintReportsMonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageByUser provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) MonthlyPrintUsageByUser()(*PrintReportsMonthlyPrintUsageByUserRequestBuilder) {
    return NewPrintReportsMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByUserById provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*PrintReportsMonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewPrintReportsMonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinter provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*PrintReportsMonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewPrintReportsMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinterById provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*PrintReportsMonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewPrintReportsMonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUser provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*PrintReportsMonthlyPrintUsageSummariesByUserRequestBuilder) {
    return NewPrintReportsMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUserById provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*PrintReportsMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewPrintReportsMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property reports in print
func (m *PrintReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *PrintReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// Security provides operations to manage the security property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) Security()(*PrintReportsSecurityRequestBuilder) {
    return NewPrintReportsSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetails provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) UserCredentialUsageDetails()(*PrintReportsUserCredentialUsageDetailsRequestBuilder) {
    return NewPrintReportsUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetailsById provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
func (m *PrintReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*PrintReportsUserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userCredentialUsageDetails%2Did"] = id
    }
    return NewPrintReportsUserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
