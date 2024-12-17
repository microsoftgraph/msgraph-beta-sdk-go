package certificateauthorities

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder provides operations to manage the trustedCertificateAuthorities property of the microsoft.graph.trustedCertificateAuthorityAsEntityBase entity.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderGetQueryParameters collection of trusted certificate authorities.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderGetQueryParameters
}
// CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCertificateAuthorityAsEntityId provides operations to manage the trustedCertificateAuthorities property of the microsoft.graph.trustedCertificateAuthorityAsEntityBase entity.
// returns a *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) ByCertificateAuthorityAsEntityId(certificateAuthorityAsEntityId string)(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if certificateAuthorityAsEntityId != "" {
        urlTplParams["certificateAuthorityAsEntity%2Did"] = certificateAuthorityAsEntityId
    }
    return NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderInternal instantiates a new CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder and sets the default values.
func NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) {
    m := &CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/certificateAuthorities/certificateBasedApplicationConfigurations/{certificateBasedApplicationConfiguration%2Did}/trustedCertificateAuthorities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder instantiates a new CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder and sets the default values.
func NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCountRequestBuilder when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) Count()(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCountRequestBuilder) {
    return NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of trusted certificate authorities.
// returns a CertificateAuthorityAsEntityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateAuthorityAsEntityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityCollectionResponseable), nil
}
// Post create new navigation property to trustedCertificateAuthorities for certificateAuthorities
// returns a CertificateAuthorityAsEntityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation collection of trusted certificate authorities.
// returns a *RequestInformation when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to trustedCertificateAuthorities for certificateAuthorities
// returns a *RequestInformation when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, requestConfiguration *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder when successful
func (m *CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) WithUrl(rawUrl string)(*CertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder) {
    return NewCertificateBasedApplicationConfigurationsItemTrustedCertificateAuthoritiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
