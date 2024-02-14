package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsRequestBuilder provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
type ItemOnlineMeetingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlineMeetingsRequestBuilderGetQueryParameters retrieve the properties and relationships of an onlineMeeting object. For example, you can: Teams live event attendee report (deprecated) and Teams live event recordings (deprecated) are online meeting artifacts. For details, see Online meeting artifacts and permissions. 
type ItemOnlineMeetingsRequestBuilderGetQueryParameters struct {
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
// ItemOnlineMeetingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlineMeetingsRequestBuilderGetQueryParameters
}
// ItemOnlineMeetingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByOnlineMeetingId provides operations to manage the onlineMeetings property of the microsoft.graph.user entity.
// returns a *ItemOnlineMeetingsOnlineMeetingItemRequestBuilder when successful
func (m *ItemOnlineMeetingsRequestBuilder) ByOnlineMeetingId(onlineMeetingId string)(*ItemOnlineMeetingsOnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if onlineMeetingId != "" {
        urlTplParams["onlineMeeting%2Did"] = onlineMeetingId
    }
    return NewItemOnlineMeetingsOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOnlineMeetingsRequestBuilderInternal instantiates a new ItemOnlineMeetingsRequestBuilder and sets the default values.
func NewItemOnlineMeetingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsRequestBuilder) {
    m := &ItemOnlineMeetingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOnlineMeetingsRequestBuilder instantiates a new ItemOnlineMeetingsRequestBuilder and sets the default values.
func NewItemOnlineMeetingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemOnlineMeetingsCountRequestBuilder when successful
func (m *ItemOnlineMeetingsRequestBuilder) Count()(*ItemOnlineMeetingsCountRequestBuilder) {
    return NewItemOnlineMeetingsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateOrGet provides operations to call the createOrGet method.
// returns a *ItemOnlineMeetingsCreateOrGetRequestBuilder when successful
func (m *ItemOnlineMeetingsRequestBuilder) CreateOrGet()(*ItemOnlineMeetingsCreateOrGetRequestBuilder) {
    return NewItemOnlineMeetingsCreateOrGetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the properties and relationships of an onlineMeeting object. For example, you can: Teams live event attendee report (deprecated) and Teams live event recordings (deprecated) are online meeting artifacts. For details, see Online meeting artifacts and permissions. 
// returns a OnlineMeetingCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onlinemeeting-get?view=graph-rest-1.0
func (m *ItemOnlineMeetingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlineMeetingsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingCollectionResponseable), nil
}
// GetAllRecordings provides operations to call the getAllRecordings method.
// returns a *ItemOnlineMeetingsGetAllRecordingsRequestBuilder when successful
func (m *ItemOnlineMeetingsRequestBuilder) GetAllRecordings()(*ItemOnlineMeetingsGetAllRecordingsRequestBuilder) {
    return NewItemOnlineMeetingsGetAllRecordingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetAllTranscripts provides operations to call the getAllTranscripts method.
// returns a *ItemOnlineMeetingsGetAllTranscriptsRequestBuilder when successful
func (m *ItemOnlineMeetingsRequestBuilder) GetAllTranscripts()(*ItemOnlineMeetingsGetAllTranscriptsRequestBuilder) {
    return NewItemOnlineMeetingsGetAllTranscriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create an online meeting on behalf of a user.
// returns a OnlineMeetingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/application-post-onlinemeetings?view=graph-rest-1.0
func (m *ItemOnlineMeetingsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *ItemOnlineMeetingsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnlineMeetingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable), nil
}
// ToGetRequestInformation retrieve the properties and relationships of an onlineMeeting object. For example, you can: Teams live event attendee report (deprecated) and Teams live event recordings (deprecated) are online meeting artifacts. For details, see Online meeting artifacts and permissions. 
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create an online meeting on behalf of a user.
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnlineMeetingable, requestConfiguration *ItemOnlineMeetingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/users/{user%2Did}/onlineMeetings", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemOnlineMeetingsRequestBuilder when successful
func (m *ItemOnlineMeetingsRequestBuilder) WithUrl(rawUrl string)(*ItemOnlineMeetingsRequestBuilder) {
    return NewItemOnlineMeetingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
