package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PresencesItemSetpresenceSetPresenceRequestBuilder provides operations to call the setPresence method.
type PresencesItemSetpresenceSetPresenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PresencesItemSetpresenceSetPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresencesItemSetpresenceSetPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresencesItemSetpresenceSetPresenceRequestBuilderInternal instantiates a new PresencesItemSetpresenceSetPresenceRequestBuilder and sets the default values.
func NewPresencesItemSetpresenceSetPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemSetpresenceSetPresenceRequestBuilder) {
    m := &PresencesItemSetpresenceSetPresenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/presences/{presence%2Did}/setPresence", pathParameters),
    }
    return m
}
// NewPresencesItemSetpresenceSetPresenceRequestBuilder instantiates a new PresencesItemSetpresenceSetPresenceRequestBuilder and sets the default values.
func NewPresencesItemSetpresenceSetPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemSetpresenceSetPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresencesItemSetpresenceSetPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the availability and activity status in a presence session of an application for a user.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-setpresence?view=graph-rest-beta
func (m *PresencesItemSetpresenceSetPresenceRequestBuilder) Post(ctx context.Context, body PresencesItemSetpresenceSetPresencePostRequestBodyable, requestConfiguration *PresencesItemSetpresenceSetPresenceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation set the availability and activity status in a presence session of an application for a user.
// returns a *RequestInformation when successful
func (m *PresencesItemSetpresenceSetPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body PresencesItemSetpresenceSetPresencePostRequestBodyable, requestConfiguration *PresencesItemSetpresenceSetPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PresencesItemSetpresenceSetPresenceRequestBuilder when successful
func (m *PresencesItemSetpresenceSetPresenceRequestBuilder) WithUrl(rawUrl string)(*PresencesItemSetpresenceSetPresenceRequestBuilder) {
    return NewPresencesItemSetpresenceSetPresenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
