package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConditionalAccessAuthenticationContextClassReferencesRequestBuilder provides operations to manage the authenticationContextClassReferences property of the microsoft.graph.conditionalAccessRoot entity.
type ConditionalAccessAuthenticationContextClassReferencesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalAccessAuthenticationContextClassReferencesRequestBuilderGetQueryParameters retrieve a list of authenticationContextClassReference objects.
type ConditionalAccessAuthenticationContextClassReferencesRequestBuilderGetQueryParameters struct {
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
// ConditionalAccessAuthenticationContextClassReferencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationContextClassReferencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalAccessAuthenticationContextClassReferencesRequestBuilderGetQueryParameters
}
// ConditionalAccessAuthenticationContextClassReferencesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalAccessAuthenticationContextClassReferencesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationContextClassReferenceId provides operations to manage the authenticationContextClassReferences property of the microsoft.graph.conditionalAccessRoot entity.
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) ByAuthenticationContextClassReferenceId(authenticationContextClassReferenceId string)(*ConditionalAccessAuthenticationContextClassReferencesAuthenticationContextClassReferenceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationContextClassReferenceId != "" {
        urlTplParams["authenticationContextClassReference%2Did"] = authenticationContextClassReferenceId
    }
    return NewConditionalAccessAuthenticationContextClassReferencesAuthenticationContextClassReferenceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalAccessAuthenticationContextClassReferencesRequestBuilderInternal instantiates a new AuthenticationContextClassReferencesRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationContextClassReferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) {
    m := &ConditionalAccessAuthenticationContextClassReferencesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationContextClassReferences{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewConditionalAccessAuthenticationContextClassReferencesRequestBuilder instantiates a new AuthenticationContextClassReferencesRequestBuilder and sets the default values.
func NewConditionalAccessAuthenticationContextClassReferencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalAccessAuthenticationContextClassReferencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) Count()(*ConditionalAccessAuthenticationContextClassReferencesCountRequestBuilder) {
    return NewConditionalAccessAuthenticationContextClassReferencesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of authenticationContextClassReference objects.
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conditionalaccessroot-list-authenticationcontextclassreferences?view=graph-rest-1.0
func (m *ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationContextClassReferencesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationContextClassReferenceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceCollectionResponseable), nil
}
// Post create a new authenticationContextClassReference.
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conditionalaccessroot-post-authenticationcontextclassreferences?view=graph-rest-1.0
func (m *ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable, requestConfiguration *ConditionalAccessAuthenticationContextClassReferencesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAuthenticationContextClassReferenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable), nil
}
// ToGetRequestInformation retrieve a list of authenticationContextClassReference objects.
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalAccessAuthenticationContextClassReferencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new authenticationContextClassReference.
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AuthenticationContextClassReferenceable, requestConfiguration *ConditionalAccessAuthenticationContextClassReferencesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2023-08/PrivatePreview:changeManagement on 2023-08-23 and will be removed 2023-08-23
func (m *ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) WithUrl(rawUrl string)(*ConditionalAccessAuthenticationContextClassReferencesRequestBuilder) {
    return NewConditionalAccessAuthenticationContextClassReferencesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
