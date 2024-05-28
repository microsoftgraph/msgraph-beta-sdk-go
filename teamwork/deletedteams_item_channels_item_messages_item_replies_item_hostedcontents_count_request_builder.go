package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder provides operations to count the resources in the collection.
type DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderGetQueryParameters get the number of the resource
type DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderGetQueryParameters
}
// NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderInternal instantiates a new DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder) {
    m := &DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/hostedContents/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder instantiates a new DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder when successful
func (m *DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder) {
    return NewDeletedteamsItemChannelsItemMessagesItemRepliesItemHostedcontentsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
