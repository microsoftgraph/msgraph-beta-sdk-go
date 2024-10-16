package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder provides operations to manage the certificateAuthorities property of the microsoft.graph.certificateBasedAuthPki entity.
type PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderGetQueryParameters get certificateAuthorities from directory
type PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderGetQueryParameters
}
// PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderInternal instantiates a new PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder and sets the default values.
func NewPublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) {
    m := &PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/{certificateBasedAuthPki%2Did}/certificateAuthorities/{certificateAuthorityDetail%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder instantiates a new PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder and sets the default values.
func NewPublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property certificateAuthorities for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get certificateAuthorities from directory
// returns a CertificateAuthorityDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateAuthorityDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityDetailable), nil
}
// Patch update the navigation property certificateAuthorities in directory
// returns a CertificateAuthorityDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityDetailable, requestConfiguration *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityDetailable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateAuthorityDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityDetailable), nil
}
// ToDeleteRequestInformation delete navigation property certificateAuthorities for directory
// returns a *RequestInformation when successful
func (m *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get certificateAuthorities from directory
// returns a *RequestInformation when successful
func (m *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property certificateAuthorities in directory
// returns a *RequestInformation when successful
func (m *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityDetailable, requestConfiguration *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder when successful
func (m *PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) WithUrl(rawUrl string)(*PublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder) {
    return NewPublicKeyInfrastructureCertificateBasedAuthConfigurationsItemCertificateAuthoritiesCertificateAuthorityDetailItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
