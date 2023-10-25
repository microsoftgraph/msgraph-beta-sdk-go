package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder provides operations to call the downloadApplePushNotificationCertificateSigningRequest method.
type ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal instantiates a new DownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    m := &ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/applePushNotificationCertificate/downloadApplePushNotificationCertificateSigningRequest()", pathParameters),
    }
    return m
}
// NewApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder instantiates a new DownloadApplePushNotificationCertificateSigningRequestRequestBuilder and sets the default values.
func NewApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Get download Apple push notification certificate signing request
// Deprecated: This method is obsolete. Use GetAsDownloadApplePushNotificationCertificateSigningRequestGetResponse instead.
func (m *ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) Get(ctx context.Context, requestConfiguration *ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration)(ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestResponseable), nil
}
// GetAsDownloadApplePushNotificationCertificateSigningRequestGetResponse download Apple push notification certificate signing request
func (m *ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) GetAsDownloadApplePushNotificationCertificateSigningRequestGetResponse(ctx context.Context, requestConfiguration *ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration)(ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestGetResponseable), nil
}
// ToGetRequestInformation download Apple push notification certificate signing request
func (m *ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) WithUrl(rawUrl string)(*ApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder) {
    return NewApplePushNotificationCertificateDownloadApplePushNotificationCertificateSigningRequestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
