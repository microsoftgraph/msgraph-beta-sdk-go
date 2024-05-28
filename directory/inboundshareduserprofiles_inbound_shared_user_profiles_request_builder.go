package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder provides operations to manage the inboundSharedUserProfiles property of the microsoft.graph.directory entity.
type InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderGetQueryParameters retrieve the properties of all inboundSharedUserProfiles.
type InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderGetQueryParameters struct {
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
// InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderGetQueryParameters
}
// InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByInboundSharedUserProfileUserId provides operations to manage the inboundSharedUserProfiles property of the microsoft.graph.directory entity.
// returns a *InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder when successful
func (m *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) ByInboundSharedUserProfileUserId(inboundSharedUserProfileUserId string)(*InboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if inboundSharedUserProfileUserId != "" {
        urlTplParams["inboundSharedUserProfile%2DuserId"] = inboundSharedUserProfileUserId
    }
    return NewInboundshareduserprofilesInboundSharedUserProfileUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewInboundshareduserprofilesInboundSharedUserProfilesRequestBuilderInternal instantiates a new InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder and sets the default values.
func NewInboundshareduserprofilesInboundSharedUserProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) {
    m := &InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/inboundSharedUserProfiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewInboundshareduserprofilesInboundSharedUserProfilesRequestBuilder instantiates a new InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder and sets the default values.
func NewInboundshareduserprofilesInboundSharedUserProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInboundshareduserprofilesInboundSharedUserProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *InboundshareduserprofilesCountRequestBuilder when successful
func (m *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) Count()(*InboundshareduserprofilesCountRequestBuilder) {
    return NewInboundshareduserprofilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties of all inboundSharedUserProfiles.
// returns a InboundSharedUserProfileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directory-list-inboundshareduserprofiles?view=graph-rest-beta
func (m *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInboundSharedUserProfileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileCollectionResponseable), nil
}
// Post create new navigation property to inboundSharedUserProfiles for directory
// returns a InboundSharedUserProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInboundSharedUserProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable), nil
}
// ToGetRequestInformation retrieve the properties of all inboundSharedUserProfiles.
// returns a *RequestInformation when successful
func (m *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to inboundSharedUserProfiles for directory
// returns a *RequestInformation when successful
func (m *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InboundSharedUserProfileable, requestConfiguration *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder when successful
func (m *InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) WithUrl(rawUrl string)(*InboundshareduserprofilesInboundSharedUserProfilesRequestBuilder) {
    return NewInboundshareduserprofilesInboundSharedUserProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
