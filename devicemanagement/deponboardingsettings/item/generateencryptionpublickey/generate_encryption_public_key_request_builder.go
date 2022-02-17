package generateencryptionpublickey

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GenerateEncryptionPublicKeyRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\microsoft.graph.generateEncryptionPublicKey
type GenerateEncryptionPublicKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GenerateEncryptionPublicKeyRequestBuilderPostOptions options for Post
type GenerateEncryptionPublicKeyRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGenerateEncryptionPublicKeyRequestBuilderInternal instantiates a new GenerateEncryptionPublicKeyRequestBuilder and sets the default values.
func NewGenerateEncryptionPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GenerateEncryptionPublicKeyRequestBuilder) {
    m := &GenerateEncryptionPublicKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}/microsoft.graph.generateEncryptionPublicKey";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGenerateEncryptionPublicKeyRequestBuilder instantiates a new GenerateEncryptionPublicKeyRequestBuilder and sets the default values.
func NewGenerateEncryptionPublicKeyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GenerateEncryptionPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGenerateEncryptionPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation generate a public key to use to encrypt the Apple device enrollment program token
func (m *GenerateEncryptionPublicKeyRequestBuilder) CreatePostRequestInformation(options *GenerateEncryptionPublicKeyRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post generate a public key to use to encrypt the Apple device enrollment program token
func (m *GenerateEncryptionPublicKeyRequestBuilder) Post(options *GenerateEncryptionPublicKeyRequestBuilderPostOptions)(*string, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "string", nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*string), nil
}
