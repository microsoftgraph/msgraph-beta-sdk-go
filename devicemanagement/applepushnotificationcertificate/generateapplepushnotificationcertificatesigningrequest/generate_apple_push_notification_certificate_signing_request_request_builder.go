package generateapplepushnotificationcertificatesigningrequest

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GenerateApplePushNotificationCertificateSigningRequestRequestBuilder provides operations to call the generateApplePushNotificationCertificateSigningRequest method.
type GenerateApplePushNotificationCertificateSigningRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostOptions options for Post
type GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal instantiates a new GenerateApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) {
    m := &GenerateApplePushNotificationCertificateSigningRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/applePushNotificationCertificate/microsoft.graph.generateApplePushNotificationCertificateSigningRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilder instantiates a new GenerateApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation download Apple push notification certificate signing request
func (m *GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) CreatePostRequestInformation(options *GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post download Apple push notification certificate signing request
func (m *GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) Post(options *GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostOptions)(GenerateApplePushNotificationCertificateSigningRequestResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GenerateApplePushNotificationCertificateSigningRequestResponseable), nil
}
