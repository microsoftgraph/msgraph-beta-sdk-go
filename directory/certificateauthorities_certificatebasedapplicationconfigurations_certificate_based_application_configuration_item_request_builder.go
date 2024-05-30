package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder provides operations to manage the certificateBasedApplicationConfigurations property of the microsoft.graph.certificateAuthorityPath entity.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetQueryParameters read the properties and relationships of a certificateBasedApplicationConfiguration object.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetQueryParameters
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderInternal instantiates a new CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder and sets the default values.
func NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) {
    m := &CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/certificateAuthorities/certificateBasedApplicationConfigurations/{certificateBasedApplicationConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder instantiates a new CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder and sets the default values.
func NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete the properties and relationships of a certificateBasedApplicationConfiguration object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificatebasedapplicationconfiguration-delete?view=graph-rest-beta
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a certificateBasedApplicationConfiguration object.
// returns a CertificateBasedApplicationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificatebasedapplicationconfiguration-get?view=graph-rest-beta
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, error) {
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
// Patch update the properties of a certificateBasedApplicationConfiguration object. To update the trustedCertificateAuthorities within a certificateBasedApplicationConfiguration object, use the Update certificateAuthorityAsEntity operation.
// returns a CertificateBasedApplicationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificatebasedapplicationconfiguration-update?view=graph-rest-beta
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, error) {
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
// ToDeleteRequestInformation delete the properties and relationships of a certificateBasedApplicationConfiguration object.
// returns a *RequestInformation when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a certificateBasedApplicationConfiguration object.
// returns a *RequestInformation when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a certificateBasedApplicationConfiguration object. To update the trustedCertificateAuthorities within a certificateBasedApplicationConfiguration object, use the Update certificateAuthorityAsEntity operation.
// returns a *RequestInformation when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesTrustedCertificateAuthoritiesRequestBuilder when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) TrustedCertificateAuthorities()(*CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesTrustedCertificateAuthoritiesRequestBuilder) {
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesTrustedCertificateAuthoritiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) {
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
