package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder provides operations to manage the permissionsCreepIndexDistributions property of the microsoft.graph.permissionsAnalytics entity.
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderGetQueryParameters represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderGetQueryParameters
}
// PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPermissionsCreepIndexDistributionId provides operations to manage the permissionsCreepIndexDistributions property of the microsoft.graph.permissionsAnalytics entity.
// returns a *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) ByPermissionsCreepIndexDistributionId(permissionsCreepIndexDistributionId string)(*PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if permissionsCreepIndexDistributionId != "" {
        urlTplParams["permissionsCreepIndexDistribution%2Did"] = permissionsCreepIndexDistributionId
    }
    return NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderInternal instantiates a new PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder and sets the default values.
func NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) {
    m := &PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder instantiates a new PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder and sets the default values.
func NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PermissionsanalyticsGcpPermissionscreepindexdistributionsCountRequestBuilder when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) Count()(*PermissionsanalyticsGcpPermissionscreepindexdistributionsCountRequestBuilder) {
    return NewPermissionsanalyticsGcpPermissionscreepindexdistributionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
// returns a PermissionsCreepIndexDistributionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionCollectionResponseable), nil
}
// Post create new navigation property to permissionsCreepIndexDistributions for identityGovernance
// returns a PermissionsCreepIndexDistributionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to permissionsCreepIndexDistributions for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder when successful
func (m *PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) WithUrl(rawUrl string)(*PermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder) {
    return NewPermissionsanalyticsGcpPermissionscreepindexdistributionsPermissionsCreepIndexDistributionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
