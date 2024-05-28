package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder provides operations to call the requestSignupUrl method.
type AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderInternal instantiates a new AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder) {
    m := &AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/requestSignupUrl", pathParameters),
    }
    return m
}
// NewAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder instantiates a new AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action requestSignupUrl
// Deprecated: This method is obsolete. Use PostAsRequestSignupUrlPostResponse instead.
// returns a AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder) Post(ctx context.Context, body AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration)(AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlResponseable), nil
}
// PostAsRequestSignupUrlPostResponse invoke action requestSignupUrl
// returns a AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder) PostAsRequestSignupUrlPostResponse(ctx context.Context, body AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration)(AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlPostResponseable), nil
}
// ToPostRequestInformation invoke action requestSignupUrl
// returns a *RequestInformation when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlPostRequestBodyable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder) WithUrl(rawUrl string)(*AndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsRequestsignupurlRequestSignupUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
