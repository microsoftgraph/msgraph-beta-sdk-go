package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder provides operations to call the uploadClientCertificate method.
type B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilderInternal instantiates a new UploadClientCertificateRequestBuilder and sets the default values.
func NewB2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder) {
    m := &B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/apiConnectorConfiguration/postFederationSignup/uploadClientCertificate", pathParameters),
    }
    return m
}
// NewB2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder instantiates a new UploadClientCertificateRequestBuilder and sets the default values.
func NewB2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a PKCS 12 format key (.pfx) to an API connector's authentication configuration. The input is a base-64 encoded value of the PKCS 12 certificate contents. This method returns an apiConnector. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityapiconnector-uploadclientcertificate?view=graph-rest-1.0
func (m *B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder) Post(ctx context.Context, body B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificatePostRequestBodyable, requestConfiguration *B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityApiConnectorable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityApiConnectorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityApiConnectorable), nil
}
// ToPostRequestInformation upload a PKCS 12 format key (.pfx) to an API connector's authentication configuration. The input is a base-64 encoded value of the PKCS 12 certificate contents. This method returns an apiConnector. This API is available in the following national cloud deployments.
func (m *B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder) ToPostRequestInformation(ctx context.Context, body B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificatePostRequestBodyable, requestConfiguration *B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder) WithUrl(rawUrl string)(*B2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder) {
    return NewB2xUserFlowsItemApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
