package identitygovernance

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the summary method.
type LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new MicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/runs/{run%2Did}/userProcessingResults/microsoft.graph.identityGovernance.summary(startDateTime={startDateTime},endDateTime={endDateTime})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if endDateTime != nil {
        urlTplParams["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        urlTplParams["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new MicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function summary
func (m *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateUserSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.UserSummaryable), nil
}
// ToGetRequestInformation invoke function summary
func (m *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemRunsItemUserProcessingResultsMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
