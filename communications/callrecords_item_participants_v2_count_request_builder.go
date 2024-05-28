package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallrecordsItemParticipants_v2CountRequestBuilder provides operations to count the resources in the collection.
type CallrecordsItemParticipants_v2CountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallrecordsItemParticipants_v2CountRequestBuilderGetQueryParameters get the number of the resource
type CallrecordsItemParticipants_v2CountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// CallrecordsItemParticipants_v2CountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemParticipants_v2CountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallrecordsItemParticipants_v2CountRequestBuilderGetQueryParameters
}
// NewCallrecordsItemParticipants_v2CountRequestBuilderInternal instantiates a new CallrecordsItemParticipants_v2CountRequestBuilder and sets the default values.
func NewCallrecordsItemParticipants_v2CountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsItemParticipants_v2CountRequestBuilder) {
    m := &CallrecordsItemParticipants_v2CountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/{callRecord%2Did}/participants_v2/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewCallrecordsItemParticipants_v2CountRequestBuilder instantiates a new CallrecordsItemParticipants_v2CountRequestBuilder and sets the default values.
func NewCallrecordsItemParticipants_v2CountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsItemParticipants_v2CountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallrecordsItemParticipants_v2CountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemParticipants_v2CountRequestBuilder) Get(ctx context.Context, requestConfiguration *CallrecordsItemParticipants_v2CountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *CallrecordsItemParticipants_v2CountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallrecordsItemParticipants_v2CountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CallrecordsItemParticipants_v2CountRequestBuilder when successful
func (m *CallrecordsItemParticipants_v2CountRequestBuilder) WithUrl(rawUrl string)(*CallrecordsItemParticipants_v2CountRequestBuilder) {
    return NewCallrecordsItemParticipants_v2CountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
