package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder provides operations to call the assign method.
type EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedSIMActivationCodePoolsItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolsItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEmbeddedSIMActivationCodePoolsItemAssignRequestBuilderInternal instantiates a new AssignRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolsItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) {
    m := &EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}/assign", pathParameters),
    }
    return m
}
// NewEmbeddedSIMActivationCodePoolsItemAssignRequestBuilder instantiates a new AssignRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolsItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedSIMActivationCodePoolsItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assign
// Deprecated: This method is obsolete. Use PostAsAssignPostResponse instead.
func (m *EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) Post(ctx context.Context, body EmbeddedSIMActivationCodePoolsItemAssignPostRequestBodyable, requestConfiguration *EmbeddedSIMActivationCodePoolsItemAssignRequestBuilderPostRequestConfiguration)(EmbeddedSIMActivationCodePoolsItemAssignResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmbeddedSIMActivationCodePoolsItemAssignResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EmbeddedSIMActivationCodePoolsItemAssignResponseable), nil
}
// PostAsAssignPostResponse invoke action assign
func (m *EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) PostAsAssignPostResponse(ctx context.Context, body EmbeddedSIMActivationCodePoolsItemAssignPostRequestBodyable, requestConfiguration *EmbeddedSIMActivationCodePoolsItemAssignRequestBuilderPostRequestConfiguration)(EmbeddedSIMActivationCodePoolsItemAssignPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEmbeddedSIMActivationCodePoolsItemAssignPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EmbeddedSIMActivationCodePoolsItemAssignPostResponseable), nil
}
// ToPostRequestInformation invoke action assign
func (m *EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmbeddedSIMActivationCodePoolsItemAssignPostRequestBodyable, requestConfiguration *EmbeddedSIMActivationCodePoolsItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) WithUrl(rawUrl string)(*EmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) {
    return NewEmbeddedSIMActivationCodePoolsItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
