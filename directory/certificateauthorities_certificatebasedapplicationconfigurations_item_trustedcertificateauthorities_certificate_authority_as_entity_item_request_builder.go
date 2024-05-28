package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder provides operations to manage the trustedCertificateAuthorities property of the microsoft.graph.trustedCertificateAuthorityAsEntityBase entity.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters read the properties and relationships of a certificateAuthorityAsEntity object.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetQueryParameters
}
// CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal instantiates a new CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder and sets the default values.
func NewCertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    m := &CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/certificateAuthorities/certificateBasedApplicationConfigurations/{certificateBasedApplicationConfiguration%2Did}/trustedCertificateAuthorities/{certificateAuthorityAsEntity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder instantiates a new CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder and sets the default values.
func NewCertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a certificateAuthorityAsEntity object. You can't delete all items in the collection because this collection requires at least one object that is a root authority to always persist.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificateauthorityasentity-delete?view=graph-rest-beta
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a certificateAuthorityAsEntity object.
// returns a CertificateAuthorityAsEntityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificateauthorityasentity-get?view=graph-rest-beta
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, error) {
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
// Patch update the properties of a certificateAuthorityAsEntity object.
// returns a CertificateAuthorityAsEntityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/certificateauthorityasentity-update?view=graph-rest-beta
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, error) {
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
// ToDeleteRequestInformation delete a certificateAuthorityAsEntity object. You can't delete all items in the collection because this collection requires at least one object that is a root authority to always persist.
// returns a *RequestInformation when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a certificateAuthorityAsEntity object.
// returns a *RequestInformation when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a certificateAuthorityAsEntity object.
// returns a *RequestInformation when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CertificateAuthorityAsEntityable, requestConfiguration *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder when successful
func (m *CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) WithUrl(rawUrl string)(*CertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder) {
    return NewCertificateauthoritiesCertificatebasedapplicationconfigurationsItemTrustedcertificateauthoritiesCertificateAuthorityAsEntityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
