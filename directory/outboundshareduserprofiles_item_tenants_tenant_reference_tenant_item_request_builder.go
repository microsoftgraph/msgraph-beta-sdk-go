package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder provides operations to manage the tenants property of the microsoft.graph.outboundSharedUserProfile entity.
type OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderGetQueryParameters the collection of external Microsoft Entra tenants that the user shared profile data with. Read-only.
type OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderGetQueryParameters
}
// OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderInternal instantiates a new OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder and sets the default values.
func NewOutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) {
    m := &OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/outboundSharedUserProfiles/{outboundSharedUserProfile%2DuserId}/tenants/{tenantReference%2DtenantId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder instantiates a new OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder and sets the default values.
func NewOutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tenants for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of external Microsoft Entra tenants that the user shared profile data with. Read-only.
// returns a TenantReferenceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property tenants in directory
// returns a TenantReferenceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceable, requestConfiguration *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// RemovePersonalData provides operations to call the removePersonalData method.
// returns a *OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder when successful
func (m *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) RemovePersonalData()(*OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder) {
    return NewOutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property tenants for directory
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of external Microsoft Entra tenants that the user shared profile data with. Read-only.
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property tenants in directory
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantReferenceable, requestConfiguration *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder when successful
func (m *OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) WithUrl(rawUrl string)(*OutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder) {
    return NewOutboundshareduserprofilesItemTenantsTenantReferenceTenantItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
