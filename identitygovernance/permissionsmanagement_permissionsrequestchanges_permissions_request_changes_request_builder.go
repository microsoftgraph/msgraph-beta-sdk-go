package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder provides operations to manage the permissionsRequestChanges property of the microsoft.graph.permissionsManagement entity.
type PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderGetQueryParameters list the permissionsRequestChange objects and their properties.
type PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderGetQueryParameters struct {
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
// PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderGetQueryParameters
}
// PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPermissionsRequestChangeId provides operations to manage the permissionsRequestChanges property of the microsoft.graph.permissionsManagement entity.
// returns a *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangeItemRequestBuilder when successful
func (m *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) ByPermissionsRequestChangeId(permissionsRequestChangeId string)(*PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if permissionsRequestChangeId != "" {
        urlTplParams["permissionsRequestChange%2Did"] = permissionsRequestChangeId
    }
    return NewPermissionsmanagementPermissionsrequestchangesPermissionsRequestChangeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderInternal instantiates a new PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder and sets the default values.
func NewPermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) {
    m := &PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsManagement/permissionsRequestChanges{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder instantiates a new PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder and sets the default values.
func NewPermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PermissionsmanagementPermissionsrequestchangesCountRequestBuilder when successful
func (m *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) Count()(*PermissionsmanagementPermissionsrequestchangesCountRequestBuilder) {
    return NewPermissionsmanagementPermissionsrequestchangesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the permissionsRequestChange objects and their properties.
// returns a PermissionsRequestChangeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissionsmanagement-list-permissionsrequestchanges?view=graph-rest-beta
func (m *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsRequestChangeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeCollectionResponseable), nil
}
// Post create new navigation property to permissionsRequestChanges for identityGovernance
// returns a PermissionsRequestChangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeable, requestConfiguration *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsRequestChangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeable), nil
}
// ToGetRequestInformation list the permissionsRequestChange objects and their properties.
// returns a *RequestInformation when successful
func (m *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to permissionsRequestChanges for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeable, requestConfiguration *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder when successful
func (m *PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) WithUrl(rawUrl string)(*PermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder) {
    return NewPermissionsmanagementPermissionsrequestchangesPermissionsRequestChangesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
