package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesUnarchiveRequestBuilder provides operations to call the unarchive method.
type ServiceannouncementMessagesUnarchiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesUnarchiveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesUnarchiveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementMessagesUnarchiveRequestBuilderInternal instantiates a new ServiceannouncementMessagesUnarchiveRequestBuilder and sets the default values.
func NewServiceannouncementMessagesUnarchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesUnarchiveRequestBuilder) {
    m := &ServiceannouncementMessagesUnarchiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/unarchive", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesUnarchiveRequestBuilder instantiates a new ServiceannouncementMessagesUnarchiveRequestBuilder and sets the default values.
func NewServiceannouncementMessagesUnarchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesUnarchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesUnarchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unarchive a list of serviceUpdateMessages for the signed in user.
// Deprecated: This method is obsolete. Use PostAsUnarchivePostResponse instead.
// returns a ServiceannouncementMessagesUnarchiveResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-unarchive?view=graph-rest-beta
func (m *ServiceannouncementMessagesUnarchiveRequestBuilder) Post(ctx context.Context, body ServiceannouncementMessagesUnarchivePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesUnarchiveRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesUnarchiveResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesUnarchiveResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesUnarchiveResponseable), nil
}
// PostAsUnarchivePostResponse unarchive a list of serviceUpdateMessages for the signed in user.
// returns a ServiceannouncementMessagesUnarchivePostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-unarchive?view=graph-rest-beta
func (m *ServiceannouncementMessagesUnarchiveRequestBuilder) PostAsUnarchivePostResponse(ctx context.Context, body ServiceannouncementMessagesUnarchivePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesUnarchiveRequestBuilderPostRequestConfiguration)(ServiceannouncementMessagesUnarchivePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceannouncementMessagesUnarchivePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceannouncementMessagesUnarchivePostResponseable), nil
}
// ToPostRequestInformation unarchive a list of serviceUpdateMessages for the signed in user.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesUnarchiveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceannouncementMessagesUnarchivePostRequestBodyable, requestConfiguration *ServiceannouncementMessagesUnarchiveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementMessagesUnarchiveRequestBuilder when successful
func (m *ServiceannouncementMessagesUnarchiveRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesUnarchiveRequestBuilder) {
    return NewServiceannouncementMessagesUnarchiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
