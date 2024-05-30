package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder provides operations to manage the postFederationSignup property of the microsoft.graph.userFlowApiConnectorConfiguration entity.
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderGetQueryParameters get postFederationSignup from identity
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderGetQueryParameters
}
// B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderInternal instantiates a new B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder and sets the default values.
func NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) {
    m := &B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/apiConnectorConfiguration/postFederationSignup{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder instantiates a new B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder and sets the default values.
func NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property postFederationSignup for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get postFederationSignup from identity
// returns a IdentityApiConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityApiConnectorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property postFederationSignup in identity
// returns a IdentityApiConnectorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityApiConnectorable, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityApiConnectorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Ref provides operations to manage the collection of identityContainer entities.
// returns a *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) Ref()(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property postFederationSignup for identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get postFederationSignup from identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property postFederationSignup in identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityApiConnectorable, requestConfiguration *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// UploadClientCertificate provides operations to call the uploadClientCertificate method.
// returns a *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupUploadclientcertificateUploadClientCertificateRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) UploadClientCertificate()(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupUploadclientcertificateUploadClientCertificateRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupUploadclientcertificateUploadClientCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder when successful
func (m *B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationPostfederationsignupPostFederationSignupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
