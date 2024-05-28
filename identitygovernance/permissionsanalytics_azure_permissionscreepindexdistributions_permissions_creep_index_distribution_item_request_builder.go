package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder provides operations to manage the permissionsCreepIndexDistributions property of the microsoft.graph.permissionsAnalytics entity.
type PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
type PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters
}
// PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthorizationSystem provides operations to manage the authorizationSystem property of the microsoft.graph.permissionsCreepIndexDistribution entity.
// returns a *PermissionsanalyticsAzurePermissionscreepindexdistributionsItemAuthorizationsystemAuthorizationSystemRequestBuilder when successful
func (m *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) AuthorizationSystem()(*PermissionsanalyticsAzurePermissionscreepindexdistributionsItemAuthorizationsystemAuthorizationSystemRequestBuilder) {
    return NewPermissionsanalyticsAzurePermissionscreepindexdistributionsItemAuthorizationsystemAuthorizationSystemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal instantiates a new PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder and sets the default values.
func NewPermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    m := &PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/{permissionsCreepIndexDistribution%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder instantiates a new PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder and sets the default values.
func NewPermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property permissionsCreepIndexDistributions for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
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
func (m *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
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
func (m *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder when successful
func (m *PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) WithUrl(rawUrl string)(*PermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    return NewPermissionsanalyticsAzurePermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
