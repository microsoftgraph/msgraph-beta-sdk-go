package certificateauthorities

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder provides operations to manage the certificateBasedApplicationConfigurations property of the microsoft.graph.certificateAuthorityPath entity.
type CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetQueryParameters defines the trusted certificate authorities for certificates that can be added to apps and service principals in the tenant.
type CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetQueryParameters
}
// CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderInternal instantiates a new CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder and sets the default values.
func NewCertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) {
    m := &CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/certificateAuthorities/certificateBasedApplicationConfigurations/{certificateBasedApplicationConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder instantiates a new CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder and sets the default values.
func NewCertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property certificateBasedApplicationConfigurations for certificateAuthorities
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get defines the trusted certificate authorities for certificates that can be added to apps and service principals in the tenant.
// returns a CertificateBasedApplicationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateBasedApplicationConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable), nil
}
// Patch update the navigation property certificateBasedApplicationConfigurations in certificateAuthorities
// returns a CertificateBasedApplicationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, requestConfiguration *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateBasedApplicationConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property certificateBasedApplicationConfigurations for certificateAuthorities
// returns a *RequestInformation when successful
func (m *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation defines the trusted certificate authorities for certificates that can be added to apps and service principals in the tenant.
// returns a *RequestInformation when successful
func (m *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property certificateBasedApplicationConfigurations in certificateAuthorities
// returns a *RequestInformation when successful
func (m *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, requestConfiguration *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TrustedCertificateAuthorities provides operations to manage the trustedCertificateAuthorities property of the microsoft.graph.trustedCertificateAuthorityAsEntityBase entity.
// returns a *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder when successful
func (m *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) TrustedCertificateAuthorities()(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) {
    return NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder when successful
func (m *CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*CertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) {
    return NewCertificateBasedApplicationConfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
