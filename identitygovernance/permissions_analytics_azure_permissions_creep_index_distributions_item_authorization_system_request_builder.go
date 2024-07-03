package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder provides operations to manage the authorizationSystem property of the microsoft.graph.permissionsCreepIndexDistribution entity.
type PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderGetQueryParameters represents an authorization system onboarded to Permissions Management.
type PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderGetQueryParameters
}
// NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderInternal instantiates a new PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder and sets the default values.
func NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder) {
    m := &PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsAnalytics/azure/permissionsCreepIndexDistributions/{permissionsCreepIndexDistribution%2Did}/authorizationSystem{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder instantiates a new PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder and sets the default values.
func NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents an authorization system onboarded to Permissions Management.
// returns a AuthorizationSystemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthorizationSystemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthorizationSystemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthorizationSystemable), nil
}
// ToGetRequestInformation represents an authorization system onboarded to Permissions Management.
// returns a *RequestInformation when successful
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder when successful
func (m *PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder) WithUrl(rawUrl string)(*PermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder) {
    return NewPermissionsAnalyticsAzurePermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
