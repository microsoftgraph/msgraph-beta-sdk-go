package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder provides operations to manage the permissionsRequestChanges property of the microsoft.graph.permissionsManagement entity.
type PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderGetQueryParameters read the properties and relationships of a permissionsRequestChange object.
type PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderGetQueryParameters
}
// PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderInternal instantiates a new PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder and sets the default values.
func NewPermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) {
    m := &PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsManagement/permissionsRequestChanges/{permissionsRequestChange%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder instantiates a new PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder and sets the default values.
func NewPermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property permissionsRequestChanges for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a permissionsRequestChange object.
// returns a PermissionsRequestChangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permissionsrequestchange-get?view=graph-rest-1.0
func (m *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property permissionsRequestChanges in identityGovernance
// returns a PermissionsRequestChangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeable, requestConfiguration *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property permissionsRequestChanges for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a permissionsRequestChange object.
// returns a *RequestInformation when successful
func (m *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property permissionsRequestChanges in identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsRequestChangeable, requestConfiguration *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder when successful
func (m *PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) WithUrl(rawUrl string)(*PermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder) {
    return NewPermissionsManagementPermissionsRequestChangesPermissionsRequestChangeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
