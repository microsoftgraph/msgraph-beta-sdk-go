package certificateauthorities

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder provides operations to manage the mutualTlsOauthConfigurations property of the microsoft.graph.certificateAuthorityPath entity.
type MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetQueryParameters get mutualTlsOauthConfigurations from certificateAuthorities
type MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetQueryParameters
}
// MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderInternal instantiates a new MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder and sets the default values.
func NewMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) {
    m := &MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/certificateAuthorities/mutualTlsOauthConfigurations/{mutualTlsOauthConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder instantiates a new MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder and sets the default values.
func NewMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mutualTlsOauthConfigurations for certificateAuthorities
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get mutualTlsOauthConfigurations from certificateAuthorities
// returns a MutualTlsOauthConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable, error) {
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
// Patch update the navigation property mutualTlsOauthConfigurations in certificateAuthorities
// returns a MutualTlsOauthConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable, requestConfiguration *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable, error) {
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
// ToDeleteRequestInformation delete navigation property mutualTlsOauthConfigurations for certificateAuthorities
// returns a *RequestInformation when successful
func (m *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get mutualTlsOauthConfigurations from certificateAuthorities
// returns a *RequestInformation when successful
func (m *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mutualTlsOauthConfigurations in certificateAuthorities
// returns a *RequestInformation when successful
func (m *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MutualTlsOauthConfigurationable, requestConfiguration *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder when successful
func (m *MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*MutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder) {
    return NewMutualTlsOauthConfigurationsMutualTlsOauthConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
