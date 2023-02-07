package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PresenceMicrosoftGraphSetStatusMessageRequestBuilder provides operations to call the setStatusMessage method.
type PresenceMicrosoftGraphSetStatusMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PresenceMicrosoftGraphSetStatusMessageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceMicrosoftGraphSetStatusMessageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresenceMicrosoftGraphSetStatusMessageRequestBuilderInternal instantiates a new MicrosoftGraphSetStatusMessageRequestBuilder and sets the default values.
func NewPresenceMicrosoftGraphSetStatusMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceMicrosoftGraphSetStatusMessageRequestBuilder) {
    m := &PresenceMicrosoftGraphSetStatusMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/presence/microsoft.graph.setStatusMessage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPresenceMicrosoftGraphSetStatusMessageRequestBuilder instantiates a new MicrosoftGraphSetStatusMessageRequestBuilder and sets the default values.
func NewPresenceMicrosoftGraphSetStatusMessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceMicrosoftGraphSetStatusMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresenceMicrosoftGraphSetStatusMessageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set a presence status message for a user. An optional expiration date and time can be supplied.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/presence-setstatusmessage?view=graph-rest-1.0
func (m *PresenceMicrosoftGraphSetStatusMessageRequestBuilder) Post(ctx context.Context, body PresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBodyable, requestConfiguration *PresenceMicrosoftGraphSetStatusMessageRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation set a presence status message for a user. An optional expiration date and time can be supplied.
func (m *PresenceMicrosoftGraphSetStatusMessageRequestBuilder) ToPostRequestInformation(ctx context.Context, body PresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBodyable, requestConfiguration *PresenceMicrosoftGraphSetStatusMessageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
