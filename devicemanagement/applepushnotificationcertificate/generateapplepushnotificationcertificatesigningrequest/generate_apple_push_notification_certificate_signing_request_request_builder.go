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
// GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
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
// CreatePostRequestInformationWithRequestConfiguration download Apple push notification certificate signing request
func (m *GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration download Apple push notification certificate signing request
func (m *GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler download Apple push notification certificate signing request
func (m *GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) PostWithResponseHandler(requestConfiguration *GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration)(GenerateApplePushNotificationCertificateSigningRequestResponseable, error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler download Apple push notification certificate signing request
func (m *GenerateApplePushNotificationCertificateSigningRequestRequestBuilder) PostWithResponseHandler(requestConfiguration *GenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GenerateApplePushNotificationCertificateSigningRequestResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GenerateApplePushNotificationCertificateSigningRequestResponseable), nil
}
