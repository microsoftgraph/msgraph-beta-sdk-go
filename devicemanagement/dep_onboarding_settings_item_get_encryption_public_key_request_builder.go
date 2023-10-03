package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder provides operations to call the getEncryptionPublicKey method.
type DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderInternal instantiates a new GetEncryptionPublicKeyRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) {
    m := &DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/getEncryptionPublicKey()", pathParameters),
    }
    return m
}
// NewDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder instantiates a new GetEncryptionPublicKeyRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a public key to use to encrypt the Apple device enrollment program token
// Deprecated: This method is obsolete. Use GetAsGetEncryptionPublicKeyGetResponse instead.
func (m *DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration)(DepOnboardingSettingsItemGetEncryptionPublicKeyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsItemGetEncryptionPublicKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsItemGetEncryptionPublicKeyResponseable), nil
}
// GetAsGetEncryptionPublicKeyGetResponse get a public key to use to encrypt the Apple device enrollment program token
func (m *DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) GetAsGetEncryptionPublicKeyGetResponse(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration)(DepOnboardingSettingsItemGetEncryptionPublicKeyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsItemGetEncryptionPublicKeyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsItemGetEncryptionPublicKeyGetResponseable), nil
}
// ToGetRequestInformation get a public key to use to encrypt the Apple device enrollment program token
func (m *DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) WithUrl(rawUrl string)(*DepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder) {
    return NewDepOnboardingSettingsItemGetEncryptionPublicKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
