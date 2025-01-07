package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder provides operations to manage the mutualTlsOauthConfigurations property of the microsoft.graph.certificateAuthorityPath entity.
type CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetQueryParameters get the properties and relationships of the specified mutualTlsOauthConfiguration resource.
type CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetQueryParameters
}
// CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderInternal instantiates a new CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder and sets the default values.
func NewCertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) {
    m := &CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/certificateAuthorities/mutualTlsOauthConfigurations/{mutualTlsOauthConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder instantiates a new CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder and sets the default values.
func NewCertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mutualTlsOauthConfigurations for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties and relationships of the specified mutualTlsOauthConfiguration resource.
// returns a MutualTlsOauthConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mutualtlsoauthconfiguration-get?view=graph-rest-beta
func (m *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMutualTlsOauthConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable), nil
}
// Patch update the specified mutualTlsOauthConfiguration resource. You can only update the following two properties: displayName, certificateAuthority. To update a subset of objects in the certificateAuthorities collection, first get the complete list, make your modifications, and then repost the entire contents of the certificateAuthorities attribute list in the request body. Excluding a subset of objects removes them from the collection.
// returns a MutualTlsOauthConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mutualtlsoauthconfiguration-update?view=graph-rest-beta
func (m *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable, requestConfiguration *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMutualTlsOauthConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property mutualTlsOauthConfigurations for directory
// returns a *RequestInformation when successful
func (m *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of the specified mutualTlsOauthConfiguration resource.
// returns a *RequestInformation when successful
func (m *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the specified mutualTlsOauthConfiguration resource. You can only update the following two properties: displayName, certificateAuthority. To update a subset of objects in the certificateAuthorities collection, first get the complete list, make your modifications, and then repost the entire contents of the certificateAuthorities attribute list in the request body. Excluding a subset of objects removes them from the collection.
// returns a *RequestInformation when successful
func (m *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable, requestConfiguration *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder when successful
func (m *CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*CertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) {
    return NewCertificateAuthoritiesMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
