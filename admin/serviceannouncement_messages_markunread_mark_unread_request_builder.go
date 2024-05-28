package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder provides operations to call the markUnread method.
type ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderInternal instantiates a new ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder and sets the default values.
func NewServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder) {
    m := &ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/markUnread", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder instantiates a new ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder and sets the default values.
func NewServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark a list of serviceUpdateMessages as unread for the signed in user.
// Deprecated: This method is obsolete. Use PostAsMarkUnreadPostResponse instead.
// returns a ServiceannouncementMessagesMarkunreadMarkUnreadResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-markunread?view=graph-rest-beta
func (m *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder) Post(ctx context.Context, body ServiceannouncementMessagesMarkunreadMarkUnreadPostRequestBodyable, requestConfiguration *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesMarkunreadMarkUnreadResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesMarkunreadMarkUnreadResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesMarkunreadMarkUnreadResponseable), nil
}
// PostAsMarkUnreadPostResponse mark a list of serviceUpdateMessages as unread for the signed in user.
// returns a ServiceannouncementMessagesMarkunreadMarkUnreadPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-markunread?view=graph-rest-beta
func (m *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder) PostAsMarkUnreadPostResponse(ctx context.Context, body ServiceannouncementMessagesMarkunreadMarkUnreadPostRequestBodyable, requestConfiguration *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesMarkunreadMarkUnreadPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesMarkunreadMarkUnreadPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesMarkunreadMarkUnreadPostResponseable), nil
}
// ToPostRequestInformation mark a list of serviceUpdateMessages as unread for the signed in user.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceannouncementMessagesMarkunreadMarkUnreadPostRequestBodyable, requestConfiguration *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder when successful
func (m *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder) {
    return NewServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
