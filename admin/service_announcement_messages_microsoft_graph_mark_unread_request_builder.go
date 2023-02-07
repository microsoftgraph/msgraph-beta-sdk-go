package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder provides operations to call the markUnread method.
type ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilderInternal instantiates a new MicrosoftGraphMarkUnreadRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder) {
    m := &ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages/microsoft.graph.markUnread";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder instantiates a new MicrosoftGraphMarkUnreadRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark a list of serviceUpdateMessages as **unread** for the signed in user.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceupdatemessage-markunread?view=graph-rest-1.0
func (m *ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder) Post(ctx context.Context, body ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadMarkUnreadPostRequestBodyable, requestConfiguration *ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilderPostRequestConfiguration)(ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadMarkUnreadResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateServiceAnnouncementMessagesMicrosoftGraphMarkUnreadMarkUnreadResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadMarkUnreadResponseable), nil
}
// ToPostRequestInformation mark a list of serviceUpdateMessages as **unread** for the signed in user.
func (m *ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadMarkUnreadPostRequestBodyable, requestConfiguration *ServiceAnnouncementMessagesMicrosoftGraphMarkUnreadRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
