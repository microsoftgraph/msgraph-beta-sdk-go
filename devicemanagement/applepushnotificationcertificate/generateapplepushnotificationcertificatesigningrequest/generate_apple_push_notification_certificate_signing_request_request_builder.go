package generateapplepushnotificationcertificatesigningrequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GenerateApplePushNotificationCertificateSigningRequestRequestBuilder builds and executes requests for operations under \deviceManagement\applePushNotificationCertificate\microsoft.graph.generateApplePushNotificationCertificateSigningRequest
type GenerateApplePushNotificationCertificateSigningRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostOptions options for Post
type GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal instantiates a new GenerateApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) {
    m := &GenerateApplePushNotificationCertificateSigningRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/applePushNotificationCertificate/microsoft.graph.generateApplePushNotificationCertificateSigningRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilder instantiates a new GenerateApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation download Apple push notification certificate signing request
func (m *GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) CreatePostRequestInformation(options *GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post download Apple push notification certificate signing request
func (m *GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) Post(options *GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostOptions)(*string, error) {
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
