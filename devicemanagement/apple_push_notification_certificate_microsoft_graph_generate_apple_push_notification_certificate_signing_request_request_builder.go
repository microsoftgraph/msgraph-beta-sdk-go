package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder provides operations to call the generateApplePushNotificationCertificateSigningRequest method.
type ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal instantiates a new MicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder) {
    m := &ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/applePushNotificationCertificate/microsoft.graph.generateApplePushNotificationCertificateSigningRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder instantiates a new MicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Post download Apple push notification certificate signing request
func (m *ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder) Post(ctx context.Context, requestConfiguration *ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration)(ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestGenerateApplePushNotificationCertificateSigningRequestResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestGenerateApplePushNotificationCertificateSigningRequestResponseable), nil
}
// ToPostRequestInformation download Apple push notification certificate signing request
func (m *ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ApplePushNotificationCertificateMicrosoftGraphGenerateApplePushNotificationCertificateSigningRequestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
