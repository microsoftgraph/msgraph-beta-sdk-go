package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder provides operations to manage the platformCredentialMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderGetQueryParameters read the properties and relationships of a platformCredentialAuthenticationMethod object.
type ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// NewItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderInternal instantiates a new PlatformCredentialAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) {
    m := &ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/platformCredentialMethods/{platformCredentialAuthenticationMethod%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder instantiates a new PlatformCredentialAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a platformCredentialAuthenticationMethod object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/platformcredentialauthenticationmethod-delete?view=graph-rest-1.0
func (m *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Device provides operations to manage the device property of the microsoft.graph.platformCredentialAuthenticationMethod entity.
func (m *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) Device()(*ItemAuthenticationPlatformCredentialMethodsItemDeviceRequestBuilder) {
    return NewItemAuthenticationPlatformCredentialMethodsItemDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a platformCredentialAuthenticationMethod object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/platformcredentialauthenticationmethod-get?view=graph-rest-1.0
func (m *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlatformCredentialAuthenticationMethodable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlatformCredentialAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlatformCredentialAuthenticationMethodable), nil
}
// ToDeleteRequestInformation delete a platformCredentialAuthenticationMethod object.
func (m *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a platformCredentialAuthenticationMethod object.
func (m *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder) {
    return NewItemAuthenticationPlatformCredentialMethodsPlatformCredentialAuthenticationMethodItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
