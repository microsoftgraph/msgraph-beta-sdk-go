package identitygovernance

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsInsightsRequestBuilder provides operations to manage the insights property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type LifecycleWorkflowsInsightsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsInsightsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsInsightsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LifecycleWorkflowsInsightsRequestBuilderGetQueryParameters the insight container holding workflow insight summaries for a tenant.
type LifecycleWorkflowsInsightsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsInsightsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsInsightsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsInsightsRequestBuilderGetQueryParameters
}
// LifecycleWorkflowsInsightsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsInsightsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleWorkflowsInsightsRequestBuilderInternal instantiates a new LifecycleWorkflowsInsightsRequestBuilder and sets the default values.
func NewLifecycleWorkflowsInsightsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsInsightsRequestBuilder) {
    m := &LifecycleWorkflowsInsightsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/insights{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsInsightsRequestBuilder instantiates a new LifecycleWorkflowsInsightsRequestBuilder and sets the default values.
func NewLifecycleWorkflowsInsightsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsInsightsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsInsightsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property insights for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleWorkflowsInsightsRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleWorkflowsInsightsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the insight container holding workflow insight summaries for a tenant.
// returns a Insightsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleWorkflowsInsightsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsInsightsRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Insightsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateInsightsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Insightsable), nil
}
// MicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTime provides operations to call the topTasksProcessedSummary method.
// returns a *LifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *LifecycleWorkflowsInsightsRequestBuilder) MicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*LifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewLifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceTopTasksProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, startDateTime)
}
// MicrosoftGraphIdentityGovernanceTopWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTime provides operations to call the topWorkflowsProcessedSummary method.
// returns a *LifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceTopWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *LifecycleWorkflowsInsightsRequestBuilder) MicrosoftGraphIdentityGovernanceTopWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*LifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceTopWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewLifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceTopWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, startDateTime)
}
// MicrosoftGraphIdentityGovernanceWorkflowsProcessedByCategoryWithStartDateTimeWithEndDateTime provides operations to call the workflowsProcessedByCategory method.
// returns a *LifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceWorkflowsProcessedByCategoryWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *LifecycleWorkflowsInsightsRequestBuilder) MicrosoftGraphIdentityGovernanceWorkflowsProcessedByCategoryWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*LifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceWorkflowsProcessedByCategoryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewLifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceWorkflowsProcessedByCategoryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, startDateTime)
}
// MicrosoftGraphIdentityGovernanceWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTime provides operations to call the workflowsProcessedSummary method.
// returns a *LifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *LifecycleWorkflowsInsightsRequestBuilder) MicrosoftGraphIdentityGovernanceWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*LifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewLifecycleWorkflowsInsightsMicrosoftGraphIdentityGovernanceWorkflowsProcessedSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, startDateTime)
}
// Patch update the navigation property insights in identityGovernance
// returns a Insightsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleWorkflowsInsightsRequestBuilder) Patch(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Insightsable, requestConfiguration *LifecycleWorkflowsInsightsRequestBuilderPatchRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Insightsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateInsightsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Insightsable), nil
}
// ToDeleteRequestInformation delete navigation property insights for identityGovernance
// returns a *RequestInformation when successful
func (m *LifecycleWorkflowsInsightsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsInsightsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the insight container holding workflow insight summaries for a tenant.
// returns a *RequestInformation when successful
func (m *LifecycleWorkflowsInsightsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsInsightsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property insights in identityGovernance
// returns a *RequestInformation when successful
func (m *LifecycleWorkflowsInsightsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Insightsable, requestConfiguration *LifecycleWorkflowsInsightsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleWorkflowsInsightsRequestBuilder when successful
func (m *LifecycleWorkflowsInsightsRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsInsightsRequestBuilder) {
    return NewLifecycleWorkflowsInsightsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
