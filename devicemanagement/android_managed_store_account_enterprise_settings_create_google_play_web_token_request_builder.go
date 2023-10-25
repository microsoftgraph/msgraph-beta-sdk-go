package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder provides operations to call the createGooglePlayWebToken method.
type AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderInternal instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) {
    m := &AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/createGooglePlayWebToken", pathParameters),
    }
    return m
}
// NewAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generates a web token that is used in an embeddable component.
// Deprecated: This method is obsolete. Use PostAsCreateGooglePlayWebTokenPostResponse instead.
func (m *AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) Post(ctx context.Context, body AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenResponseable), nil
}
// PostAsCreateGooglePlayWebTokenPostResponse generates a web token that is used in an embeddable component.
func (m *AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) PostAsCreateGooglePlayWebTokenPostResponse(ctx context.Context, body AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenPostResponseable), nil
}
// ToPostRequestInformation generates a web token that is used in an embeddable component.
func (m *AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) WithUrl(rawUrl string)(*AndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
