package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder provides operations to manage the certificateBasedApplicationConfigurations property of the microsoft.graph.certificateAuthorityPath entity.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderGetQueryParameters get a list of certificateBasedApplicationConfiguration objects.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderGetQueryParameters struct {
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
// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderGetQueryParameters
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCertificateBasedApplicationConfigurationId provides operations to manage the certificateBasedApplicationConfigurations property of the microsoft.graph.certificateAuthorityPath entity.
// returns a *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) ByCertificateBasedApplicationConfigurationId(certificateBasedApplicationConfigurationId string)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if certificateBasedApplicationConfigurationId != "" {
        urlTplParams["certificateBasedApplicationConfiguration%2Did"] = certificateBasedApplicationConfigurationId
    }
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderInternal instantiates a new CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder and sets the default values.
func NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) {
    m := &CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/certificateAuthorities/certificateBasedApplicationConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder instantiates a new CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder and sets the default values.
func NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CertificateauthoritiesCertificatebasedapplicationconfigurationsCountRequestBuilder when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) Count()(*CertificateauthoritiesCertificatebasedapplicationconfigurationsCountRequestBuilder) {
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of certificateBasedApplicationConfiguration objects.
// returns a CertificateBasedApplicationConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificateauthoritypath-list-certificatebasedapplicationconfigurations?view=graph-rest-beta
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCertificateBasedApplicationConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationCollectionResponseable), nil
}
// Post create new navigation property to certificateBasedApplicationConfigurations for directory
// returns a CertificateBasedApplicationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get a list of certificateBasedApplicationConfiguration objects.
// returns a *RequestInformation when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to certificateBasedApplicationConfigurations for directory
// returns a *RequestInformation when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateBasedApplicationConfigurationable, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) WithUrl(rawUrl string)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder) {
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsCertificateBasedApplicationConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
