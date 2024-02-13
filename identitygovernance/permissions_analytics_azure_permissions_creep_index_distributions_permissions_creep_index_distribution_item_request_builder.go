package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder provides operations to manage the permissionsCreepIndexDistributions property of the microsoft.graph.permissionsAnalytics entity.
type PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
type PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters
}
// PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthorizationSystem provides operations to manage the authorizationSystem property of the microsoft.graph.permissionsCreepIndexDistribution entity.
// returns a *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder when successful
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) AuthorizationSystem()(*PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder) {
    return NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal instantiates a new PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder and sets the default values.
func NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    m := &PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/{permissionsCreepIndexDistribution%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder instantiates a new PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder and sets the default values.
func NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property permissionsCreepIndexDistributions for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
// returns a PermissionsCreepIndexDistributionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable), nil
}
// Patch update the navigation property permissionsCreepIndexDistributions in identityGovernance
// returns a PermissionsCreepIndexDistributionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable), nil
}
// ToDeleteRequestInformation delete navigation property permissionsCreepIndexDistributions for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/{permissionsCreepIndexDistribution%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
// returns a *RequestInformation when successful
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property permissionsCreepIndexDistributions in identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/{permissionsCreepIndexDistribution%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder when successful
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) WithUrl(rawUrl string)(*PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    return NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
