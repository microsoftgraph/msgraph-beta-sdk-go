package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder provides operations to call the generateEncryptionPublicKey method.
type DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderInternal instantiates a new DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder) {
    m := &DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/generateEncryptionPublicKey", pathParameters),
    }
    return m
}
// NewDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder instantiates a new DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder and sets the default values.
func NewDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generate a public key to use to encrypt the Apple device enrollment program token
// Deprecated: This method is obsolete. Use PostAsGenerateEncryptionPublicKeyPostResponse instead.
// returns a DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder) Post(ctx context.Context, requestConfiguration *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration)(DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyResponseable), nil
}
// PostAsGenerateEncryptionPublicKeyPostResponse generate a public key to use to encrypt the Apple device enrollment program token
// returns a DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder) PostAsGenerateEncryptionPublicKeyPostResponse(ctx context.Context, requestConfiguration *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration)(DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyPostResponseable), nil
}
// ToPostRequestInformation generate a public key to use to encrypt the Apple device enrollment program token
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder when successful
func (m *DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder) {
    return NewDeponboardingsettingsItemGenerateencryptionpublickeyGenerateEncryptionPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
