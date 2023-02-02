package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder provides operations to call the unfavorite method.
type ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilderInternal instantiates a new UnfavoriteRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder) {
    m := &ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages/microsoft.graph.unfavorite";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder instantiates a new UnfavoriteRequestBuilder and sets the default values.
func NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove the favorite status of serviceUpdateMessages for the signed in user.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/serviceupdatemessage-unfavorite?view=graph-rest-1.0
func (m *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder) Post(ctx context.Context, body ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBodyable, requestConfiguration *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilderPostRequestConfiguration)(ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteResponseable), nil
}
// ToPostRequestInformation remove the favorite status of serviceUpdateMessages for the signed in user.
func (m *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilder) ToPostRequestInformation(ctx context.Context, body ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBodyable, requestConfiguration *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoriteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
