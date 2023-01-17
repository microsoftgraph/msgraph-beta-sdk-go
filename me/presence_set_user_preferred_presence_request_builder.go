package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PresenceSetUserPreferredPresenceRequestBuilder provides operations to call the setUserPreferredPresence method.
type PresenceSetUserPreferredPresenceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PresenceSetUserPreferredPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceSetUserPreferredPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresenceSetUserPreferredPresenceRequestBuilderInternal instantiates a new SetUserPreferredPresenceRequestBuilder and sets the default values.
func NewPresenceSetUserPreferredPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceSetUserPreferredPresenceRequestBuilder) {
    m := &PresenceSetUserPreferredPresenceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/presence/microsoft.graph.setUserPreferredPresence";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPresenceSetUserPreferredPresenceRequestBuilder instantiates a new SetUserPreferredPresenceRequestBuilder and sets the default values.
func NewPresenceSetUserPreferredPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceSetUserPreferredPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresenceSetUserPreferredPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the preferred availability and activity status for a user. If the preferred presence of a user is set, the user's presence is the preferred presence. Preferred presence takes effect only when there is at least one presence session of the user. Otherwise, the user's presence stays as Offline. A presence session can be created as a result of a successful setPresence operation, or if the user is signed in on a Teams client.  Read more about presence sessions and their time-out and expiration. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/presence-setuserpreferredpresence?view=graph-rest-1.0
func (m *PresenceSetUserPreferredPresenceRequestBuilder) Post(ctx context.Context, body PresenceSetUserPreferredPresencePostRequestBodyable, requestConfiguration *PresenceSetUserPreferredPresenceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation set the preferred availability and activity status for a user. If the preferred presence of a user is set, the user's presence is the preferred presence. Preferred presence takes effect only when there is at least one presence session of the user. Otherwise, the user's presence stays as Offline. A presence session can be created as a result of a successful setPresence operation, or if the user is signed in on a Teams client.  Read more about presence sessions and their time-out and expiration. 
func (m *PresenceSetUserPreferredPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body PresenceSetUserPreferredPresencePostRequestBodyable, requestConfiguration *PresenceSetUserPreferredPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
