package certificateauthorities

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder provides operations to manage the trustedCertificateAuthorities property of the microsoft.graph.trustedCertificateAuthorityAsEntityBase entity.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters collection of trusted certificate authorities.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters
}
// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal instantiates a new CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder and sets the default values.
func NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    m := &CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/certificateAuthorities/certificateBasedApplicationConfigurations/{certificateBasedApplicationConfiguration%2Did}/trustedCertificateAuthorities/{certificateAuthorityAsEntity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder instantiates a new CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder and sets the default values.
func NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property trustedCertificateAuthorities for certificateAuthorities
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of trusted certificate authorities.
// returns a CertificateAuthorityAsEntityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateAuthorityAsEntityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable), nil
}
// Patch update the navigation property trustedCertificateAuthorities in certificateAuthorities
// returns a CertificateAuthorityAsEntityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateAuthorityAsEntityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable), nil
}
// ToDeleteRequestInformation delete navigation property trustedCertificateAuthorities for certificateAuthorities
// returns a *RequestInformation when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of trusted certificate authorities.
// returns a *RequestInformation when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property trustedCertificateAuthorities in certificateAuthorities
// returns a *RequestInformation when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) WithUrl(rawUrl string)(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    return NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
