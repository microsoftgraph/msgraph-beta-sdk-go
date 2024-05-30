package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder provides operations to manage the outboundSharedUserProfiles property of the microsoft.graph.directory entity.
type OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderGetQueryParameters get the details of an outboundSharedUserProfile.
type OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderGetQueryParameters
}
// OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderInternal instantiates a new OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder and sets the default values.
func NewOutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) {
    m := &OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/outboundSharedUserProfiles/{outboundSharedUserProfile%2DuserId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder instantiates a new OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder and sets the default values.
func NewOutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property outboundSharedUserProfiles for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the details of an outboundSharedUserProfile.
// returns a OutboundSharedUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/outboundshareduserprofile-get?view=graph-rest-beta
func (m *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutboundSharedUserProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable), nil
}
// Patch update the navigation property outboundSharedUserProfiles in directory
// returns a OutboundSharedUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutboundSharedUserProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable), nil
}
// Tenants provides operations to manage the tenants property of the microsoft.graph.outboundSharedUserProfile entity.
// returns a *OutboundshareduserprofilesItemTenantsRequestBuilder when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) Tenants()(*OutboundshareduserprofilesItemTenantsRequestBuilder) {
    return NewOutboundshareduserprofilesItemTenantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property outboundSharedUserProfiles for directory
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the details of an outboundSharedUserProfile.
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property outboundSharedUserProfiles in directory
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) WithUrl(rawUrl string)(*OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) {
    return NewOutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
