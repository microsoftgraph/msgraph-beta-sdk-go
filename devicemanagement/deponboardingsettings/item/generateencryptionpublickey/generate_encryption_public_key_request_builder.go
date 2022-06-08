package generateencryptionpublickey

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GenerateEncryptionPublicKeyRequestBuilder provides operations to call the generateEncryptionPublicKey method.
type GenerateEncryptionPublicKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGenerateEncryptionPublicKeyRequestBuilderInternal instantiates a new GenerateEncryptionPublicKeyRequestBuilder and sets the default values.
func NewGenerateEncryptionPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GenerateEncryptionPublicKeyRequestBuilder) {
    m := &GenerateEncryptionPublicKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/microsoft.graph.generateEncryptionPublicKey";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGenerateEncryptionPublicKeyRequestBuilder instantiates a new GenerateEncryptionPublicKeyRequestBuilder and sets the default values.
func NewGenerateEncryptionPublicKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GenerateEncryptionPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGenerateEncryptionPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation generate a public key to use to encrypt the Apple device enrollment program token
func (m *GenerateEncryptionPublicKeyRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration generate a public key to use to encrypt the Apple device enrollment program token
func (m *GenerateEncryptionPublicKeyRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *GenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post generate a public key to use to encrypt the Apple device enrollment program token
func (m *GenerateEncryptionPublicKeyRequestBuilder) Post()(GenerateEncryptionPublicKeyResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler generate a public key to use to encrypt the Apple device enrollment program token
func (m *GenerateEncryptionPublicKeyRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *GenerateEncryptionPublicKeyRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GenerateEncryptionPublicKeyResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGenerateEncryptionPublicKeyResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GenerateEncryptionPublicKeyResponseable), nil
}
