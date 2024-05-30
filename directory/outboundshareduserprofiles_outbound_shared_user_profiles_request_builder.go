package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder provides operations to manage the outboundSharedUserProfiles property of the microsoft.graph.directory entity.
type OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderGetQueryParameters retrieve the properties of all outboundSharedUserProfiles.
type OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderGetQueryParameters struct {
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
// OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderGetQueryParameters
}
// OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOutboundSharedUserProfileUserId provides operations to manage the outboundSharedUserProfiles property of the microsoft.graph.directory entity.
// returns a *OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) ByOutboundSharedUserProfileUserId(outboundSharedUserProfileUserId string)(*OutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if outboundSharedUserProfileUserId != "" {
        urlTplParams["outboundSharedUserProfile%2DuserId"] = outboundSharedUserProfileUserId
    }
    return NewOutboundshareduserprofilesOutboundSharedUserProfileUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewOutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderInternal instantiates a new OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder and sets the default values.
func NewOutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) {
    m := &OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/outboundSharedUserProfiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewOutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder instantiates a new OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder and sets the default values.
func NewOutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *OutboundshareduserprofilesCountRequestBuilder when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) Count()(*OutboundshareduserprofilesCountRequestBuilder) {
    return NewOutboundshareduserprofilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties of all outboundSharedUserProfiles.
// returns a OutboundSharedUserProfileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-list-outboundshareduserprofiles?view=graph-rest-beta
func (m *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOutboundSharedUserProfileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileCollectionResponseable), nil
}
// Post create new navigation property to outboundSharedUserProfiles for directory
// returns a OutboundSharedUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation retrieve the properties of all outboundSharedUserProfiles.
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to outboundSharedUserProfiles for directory
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OutboundSharedUserProfileable, requestConfiguration *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder when successful
func (m *OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) WithUrl(rawUrl string)(*OutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder) {
    return NewOutboundshareduserprofilesOutboundSharedUserProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
