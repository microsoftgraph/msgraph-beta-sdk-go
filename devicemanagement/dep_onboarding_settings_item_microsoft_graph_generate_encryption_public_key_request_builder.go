package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder provides operations to call the generateEncryptionPublicKey method.
type DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilderInternal instantiates a new MicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder) {
    m := &DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/microsoft.graph.generateEncryptionPublicKey";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder instantiates a new MicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder and sets the default values.
func NewDepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post generate a public key to use to encrypt the Apple device enrollment program token
func (m *DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder) Post(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration)(DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyGenerateEncryptionPublicKeyResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateDepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyGenerateEncryptionPublicKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyGenerateEncryptionPublicKeyResponseable), nil
}
// ToPostRequestInformation generate a public key to use to encrypt the Apple device enrollment program token
func (m *DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DepOnboardingSettingsItemMicrosoftGraphGenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
