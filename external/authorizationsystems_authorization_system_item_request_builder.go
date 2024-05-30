package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthorizationsystemsAuthorizationSystemItemRequestBuilder provides operations to manage the authorizationSystems property of the microsoft.graph.externalConnectors.external entity.
type AuthorizationsystemsAuthorizationSystemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthorizationsystemsAuthorizationSystemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationsystemsAuthorizationSystemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthorizationsystemsAuthorizationSystemItemRequestBuilderGetQueryParameters represents an onboarded AWS account, Azure subscription, or GCP project that Microsoft Entra Permissions Management will collect and analyze permissions and actions on.
type AuthorizationsystemsAuthorizationSystemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthorizationsystemsAuthorizationSystemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationsystemsAuthorizationSystemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthorizationsystemsAuthorizationSystemItemRequestBuilderGetQueryParameters
}
// AuthorizationsystemsAuthorizationSystemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationsystemsAuthorizationSystemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthorizationsystemsAuthorizationSystemItemRequestBuilderInternal instantiates a new AuthorizationsystemsAuthorizationSystemItemRequestBuilder and sets the default values.
func NewAuthorizationsystemsAuthorizationSystemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthorizationsystemsAuthorizationSystemItemRequestBuilder) {
    m := &AuthorizationsystemsAuthorizationSystemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/authorizationSystems/{authorizationSystem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthorizationsystemsAuthorizationSystemItemRequestBuilder instantiates a new AuthorizationsystemsAuthorizationSystemItemRequestBuilder and sets the default values.
func NewAuthorizationsystemsAuthorizationSystemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthorizationsystemsAuthorizationSystemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthorizationsystemsAuthorizationSystemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataCollectionInfo provides operations to manage the dataCollectionInfo property of the microsoft.graph.authorizationSystem entity.
// returns a *AuthorizationsystemsItemDatacollectioninfoDataCollectionInfoRequestBuilder when successful
func (m *AuthorizationsystemsAuthorizationSystemItemRequestBuilder) DataCollectionInfo()(*AuthorizationsystemsItemDatacollectioninfoDataCollectionInfoRequestBuilder) {
    return NewAuthorizationsystemsItemDatacollectioninfoDataCollectionInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property authorizationSystems for external
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationsystemsAuthorizationSystemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthorizationsystemsAuthorizationSystemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get represents an onboarded AWS account, Azure subscription, or GCP project that Microsoft Entra Permissions Management will collect and analyze permissions and actions on.
// returns a AuthorizationSystemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationsystemsAuthorizationSystemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthorizationsystemsAuthorizationSystemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthorizationSystemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthorizationSystemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthorizationSystemable), nil
}
// Patch update the navigation property authorizationSystems in external
// returns a AuthorizationSystemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationsystemsAuthorizationSystemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthorizationSystemable, requestConfiguration *AuthorizationsystemsAuthorizationSystemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthorizationSystemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthorizationSystemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthorizationSystemable), nil
}
// ToDeleteRequestInformation delete navigation property authorizationSystems for external
// returns a *RequestInformation when successful
func (m *AuthorizationsystemsAuthorizationSystemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthorizationsystemsAuthorizationSystemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents an onboarded AWS account, Azure subscription, or GCP project that Microsoft Entra Permissions Management will collect and analyze permissions and actions on.
// returns a *RequestInformation when successful
func (m *AuthorizationsystemsAuthorizationSystemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthorizationsystemsAuthorizationSystemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property authorizationSystems in external
// returns a *RequestInformation when successful
func (m *AuthorizationsystemsAuthorizationSystemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthorizationSystemable, requestConfiguration *AuthorizationsystemsAuthorizationSystemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthorizationsystemsAuthorizationSystemItemRequestBuilder when successful
func (m *AuthorizationsystemsAuthorizationSystemItemRequestBuilder) WithUrl(rawUrl string)(*AuthorizationsystemsAuthorizationSystemItemRequestBuilder) {
    return NewAuthorizationsystemsAuthorizationSystemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
