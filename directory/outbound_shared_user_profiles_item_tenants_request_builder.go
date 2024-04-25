package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OutboundSharedUserProfilesItemTenantsRequestBuilder provides operations to manage the tenants property of the microsoft.graph.outboundSharedUserProfile entity.
type OutboundSharedUserProfilesItemTenantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OutboundSharedUserProfilesItemTenantsRequestBuilderGetQueryParameters the collection of external Microsoft Entra tenants that the user has shared profile data with. Read-only.
type OutboundSharedUserProfilesItemTenantsRequestBuilderGetQueryParameters struct {
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
// OutboundSharedUserProfilesItemTenantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundSharedUserProfilesItemTenantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OutboundSharedUserProfilesItemTenantsRequestBuilderGetQueryParameters
}
// OutboundSharedUserProfilesItemTenantsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundSharedUserProfilesItemTenantsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTenantReferenceTenantId provides operations to manage the tenants property of the microsoft.graph.outboundSharedUserProfile entity.
// returns a *OutboundSharedUserProfilesItemTenantsTenantReferenceTenantItemRequestBuilder when successful
func (m *OutboundSharedUserProfilesItemTenantsRequestBuilder) ByTenantReferenceTenantId(tenantReferenceTenantId string)(*OutboundSharedUserProfilesItemTenantsTenantReferenceTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tenantReferenceTenantId != "" {
        urlTplParams["tenantReference%2DtenantId"] = tenantReferenceTenantId
    }
    return NewOutboundSharedUserProfilesItemTenantsTenantReferenceTenantItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewOutboundSharedUserProfilesItemTenantsRequestBuilderInternal instantiates a new OutboundSharedUserProfilesItemTenantsRequestBuilder and sets the default values.
func NewOutboundSharedUserProfilesItemTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundSharedUserProfilesItemTenantsRequestBuilder) {
    m := &OutboundSharedUserProfilesItemTenantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/outboundSharedUserProfiles/{outboundSharedUserProfile%2DuserId}/tenants{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOutboundSharedUserProfilesItemTenantsRequestBuilder instantiates a new OutboundSharedUserProfilesItemTenantsRequestBuilder and sets the default values.
func NewOutboundSharedUserProfilesItemTenantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundSharedUserProfilesItemTenantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutboundSharedUserProfilesItemTenantsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *OutboundSharedUserProfilesItemTenantsCountRequestBuilder when successful
func (m *OutboundSharedUserProfilesItemTenantsRequestBuilder) Count()(*OutboundSharedUserProfilesItemTenantsCountRequestBuilder) {
    return NewOutboundSharedUserProfilesItemTenantsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of external Microsoft Entra tenants that the user has shared profile data with. Read-only.
// returns a TenantReferenceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OutboundSharedUserProfilesItemTenantsRequestBuilder) Get(ctx context.Context, requestConfiguration *OutboundSharedUserProfilesItemTenantsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantReferenceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceCollectionResponseable), nil
}
// Post create new navigation property to tenants for directory
// returns a TenantReferenceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OutboundSharedUserProfilesItemTenantsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceable, requestConfiguration *OutboundSharedUserProfilesItemTenantsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantReferenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceable), nil
}
// ToGetRequestInformation the collection of external Microsoft Entra tenants that the user has shared profile data with. Read-only.
// returns a *RequestInformation when successful
func (m *OutboundSharedUserProfilesItemTenantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OutboundSharedUserProfilesItemTenantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to tenants for directory
// returns a *RequestInformation when successful
func (m *OutboundSharedUserProfilesItemTenantsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceable, requestConfiguration *OutboundSharedUserProfilesItemTenantsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OutboundSharedUserProfilesItemTenantsRequestBuilder when successful
func (m *OutboundSharedUserProfilesItemTenantsRequestBuilder) WithUrl(rawUrl string)(*OutboundSharedUserProfilesItemTenantsRequestBuilder) {
    return NewOutboundSharedUserProfilesItemTenantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
