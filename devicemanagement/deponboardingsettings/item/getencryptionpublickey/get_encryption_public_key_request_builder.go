package getencryptionpublickey

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetEncryptionPublicKeyRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\{depOnboardingSetting-id}\microsoft.graph.getEncryptionPublicKey()
type GetEncryptionPublicKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetEncryptionPublicKeyRequestBuilderGetOptions options for Get
type GetEncryptionPublicKeyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetEncryptionPublicKeyRequestBuilderInternal instantiates a new GetEncryptionPublicKeyRequestBuilder and sets the default values.
func NewGetEncryptionPublicKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetEncryptionPublicKeyRequestBuilder) {
    m := &GetEncryptionPublicKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting_id}/microsoft.graph.getEncryptionPublicKey()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetEncryptionPublicKeyRequestBuilder instantiates a new GetEncryptionPublicKeyRequestBuilder and sets the default values.
func NewGetEncryptionPublicKeyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetEncryptionPublicKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetEncryptionPublicKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get a public key to use to encrypt the Apple device enrollment program token
func (m *GetEncryptionPublicKeyRequestBuilder) CreateGetRequestInformation(options *GetEncryptionPublicKeyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get get a public key to use to encrypt the Apple device enrollment program token
func (m *GetEncryptionPublicKeyRequestBuilder) Get(options *GetEncryptionPublicKeyRequestBuilderGetOptions)(*string, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "string", nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*string), nil
}
