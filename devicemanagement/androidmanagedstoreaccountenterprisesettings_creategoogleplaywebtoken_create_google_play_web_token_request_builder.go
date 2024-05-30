package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder provides operations to call the createGooglePlayWebToken method.
type AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderInternal instantiates a new AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder) {
    m := &AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/createGooglePlayWebToken", pathParameters),
    }
    return m
}
// NewAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder instantiates a new AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generates a web token that is used in an embeddable component.
// Deprecated: This method is obsolete. Use PostAsCreateGooglePlayWebTokenPostResponse instead.
// returns a AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder) Post(ctx context.Context, body AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenResponseable), nil
}
// PostAsCreateGooglePlayWebTokenPostResponse generates a web token that is used in an embeddable component.
// returns a AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder) PostAsCreateGooglePlayWebTokenPostResponse(ctx context.Context, body AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenPostResponseable), nil
}
// ToPostRequestInformation generates a web token that is used in an embeddable component.
// returns a *RequestInformation when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder when successful
func (m *AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder) WithUrl(rawUrl string)(*AndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsCreategoogleplaywebtokenCreateGooglePlayWebTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
