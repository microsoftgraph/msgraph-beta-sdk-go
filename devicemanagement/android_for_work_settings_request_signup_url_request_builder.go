package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidForWorkSettingsRequestSignupUrlRequestBuilder provides operations to call the requestSignupUrl method.
type AndroidForWorkSettingsRequestSignupUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidForWorkSettingsRequestSignupUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidForWorkSettingsRequestSignupUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidForWorkSettingsRequestSignupUrlRequestBuilderInternal instantiates a new RequestSignupUrlRequestBuilder and sets the default values.
func NewAndroidForWorkSettingsRequestSignupUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkSettingsRequestSignupUrlRequestBuilder) {
    m := &AndroidForWorkSettingsRequestSignupUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidForWorkSettings/requestSignupUrl", pathParameters),
    }
    return m
}
// NewAndroidForWorkSettingsRequestSignupUrlRequestBuilder instantiates a new RequestSignupUrlRequestBuilder and sets the default values.
func NewAndroidForWorkSettingsRequestSignupUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidForWorkSettingsRequestSignupUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidForWorkSettingsRequestSignupUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action requestSignupUrl
// Deprecated: This method is obsolete. Use PostAsRequestSignupUrlPostResponse instead.
func (m *AndroidForWorkSettingsRequestSignupUrlRequestBuilder) Post(ctx context.Context, body AndroidForWorkSettingsRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidForWorkSettingsRequestSignupUrlRequestBuilderPostRequestConfiguration)(AndroidForWorkSettingsRequestSignupUrlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidForWorkSettingsRequestSignupUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidForWorkSettingsRequestSignupUrlResponseable), nil
}
// PostAsRequestSignupUrlPostResponse invoke action requestSignupUrl
func (m *AndroidForWorkSettingsRequestSignupUrlRequestBuilder) PostAsRequestSignupUrlPostResponse(ctx context.Context, body AndroidForWorkSettingsRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidForWorkSettingsRequestSignupUrlRequestBuilderPostRequestConfiguration)(AndroidForWorkSettingsRequestSignupUrlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidForWorkSettingsRequestSignupUrlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidForWorkSettingsRequestSignupUrlPostResponseable), nil
}
// ToPostRequestInformation invoke action requestSignupUrl
func (m *AndroidForWorkSettingsRequestSignupUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidForWorkSettingsRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidForWorkSettingsRequestSignupUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AndroidForWorkSettingsRequestSignupUrlRequestBuilder) WithUrl(rawUrl string)(*AndroidForWorkSettingsRequestSignupUrlRequestBuilder) {
    return NewAndroidForWorkSettingsRequestSignupUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
