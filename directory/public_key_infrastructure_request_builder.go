package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PublicKeyInfrastructureRequestBuilder provides operations to manage the publicKeyInfrastructure property of the microsoft.graph.directory entity.
type PublicKeyInfrastructureRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PublicKeyInfrastructureRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PublicKeyInfrastructureRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PublicKeyInfrastructureRequestBuilderGetQueryParameters get publicKeyInfrastructure from directory
type PublicKeyInfrastructureRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PublicKeyInfrastructureRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PublicKeyInfrastructureRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PublicKeyInfrastructureRequestBuilderGetQueryParameters
}
// PublicKeyInfrastructureRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PublicKeyInfrastructureRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CertificateBasedAuthConfigurations provides operations to manage the certificateBasedAuthConfigurations property of the microsoft.graph.publicKeyInfrastructureRoot entity.
// returns a *PublicKeyInfrastructureCertificateBasedAuthConfigurationsRequestBuilder when successful
func (m *PublicKeyInfrastructureRequestBuilder) CertificateBasedAuthConfigurations()(*PublicKeyInfrastructureCertificateBasedAuthConfigurationsRequestBuilder) {
    return NewPublicKeyInfrastructureCertificateBasedAuthConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPublicKeyInfrastructureRequestBuilderInternal instantiates a new PublicKeyInfrastructureRequestBuilder and sets the default values.
func NewPublicKeyInfrastructureRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PublicKeyInfrastructureRequestBuilder) {
    m := &PublicKeyInfrastructureRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/publicKeyInfrastructure{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPublicKeyInfrastructureRequestBuilder instantiates a new PublicKeyInfrastructureRequestBuilder and sets the default values.
func NewPublicKeyInfrastructureRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PublicKeyInfrastructureRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPublicKeyInfrastructureRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property publicKeyInfrastructure for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PublicKeyInfrastructureRequestBuilder) Delete(ctx context.Context, requestConfiguration *PublicKeyInfrastructureRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get publicKeyInfrastructure from directory
// returns a PublicKeyInfrastructureRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PublicKeyInfrastructureRequestBuilder) Get(ctx context.Context, requestConfiguration *PublicKeyInfrastructureRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicKeyInfrastructureRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublicKeyInfrastructureRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicKeyInfrastructureRootable), nil
}
// Patch update the navigation property publicKeyInfrastructure in directory
// returns a PublicKeyInfrastructureRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PublicKeyInfrastructureRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicKeyInfrastructureRootable, requestConfiguration *PublicKeyInfrastructureRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicKeyInfrastructureRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublicKeyInfrastructureRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicKeyInfrastructureRootable), nil
}
// ToDeleteRequestInformation delete navigation property publicKeyInfrastructure for directory
// returns a *RequestInformation when successful
func (m *PublicKeyInfrastructureRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PublicKeyInfrastructureRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get publicKeyInfrastructure from directory
// returns a *RequestInformation when successful
func (m *PublicKeyInfrastructureRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PublicKeyInfrastructureRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property publicKeyInfrastructure in directory
// returns a *RequestInformation when successful
func (m *PublicKeyInfrastructureRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicKeyInfrastructureRootable, requestConfiguration *PublicKeyInfrastructureRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PublicKeyInfrastructureRequestBuilder when successful
func (m *PublicKeyInfrastructureRequestBuilder) WithUrl(rawUrl string)(*PublicKeyInfrastructureRequestBuilder) {
    return NewPublicKeyInfrastructureRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
