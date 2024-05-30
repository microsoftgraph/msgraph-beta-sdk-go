package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder provides operations to call the requestSignupUrl method.
type AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderInternal instantiates a new AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder and sets the default values.
func NewAndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder) {
    m := &AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidForWorkSettings/requestSignupUrl", pathParameters),
    }
    return m
}
// NewAndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder instantiates a new AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder and sets the default values.
func NewAndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action requestSignupUrl
// Deprecated: This method is obsolete. Use PostAsRequestSignupUrlPostResponse instead.
// returns a AndroidforworksettingsRequestsignupurlRequestSignupUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder) Post(ctx context.Context, body AndroidforworksettingsRequestsignupurlRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration)(AndroidforworksettingsRequestsignupurlRequestSignupUrlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidforworksettingsRequestsignupurlRequestSignupUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidforworksettingsRequestsignupurlRequestSignupUrlResponseable), nil
}
// PostAsRequestSignupUrlPostResponse invoke action requestSignupUrl
// returns a AndroidforworksettingsRequestsignupurlRequestSignupUrlPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder) PostAsRequestSignupUrlPostResponse(ctx context.Context, body AndroidforworksettingsRequestsignupurlRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration)(AndroidforworksettingsRequestsignupurlRequestSignupUrlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidforworksettingsRequestsignupurlRequestSignupUrlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidforworksettingsRequestsignupurlRequestSignupUrlPostResponseable), nil
}
// ToPostRequestInformation invoke action requestSignupUrl
// returns a *RequestInformation when successful
func (m *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidforworksettingsRequestsignupurlRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder when successful
func (m *AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder) WithUrl(rawUrl string)(*AndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder) {
    return NewAndroidforworksettingsRequestsignupurlRequestSignupUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
