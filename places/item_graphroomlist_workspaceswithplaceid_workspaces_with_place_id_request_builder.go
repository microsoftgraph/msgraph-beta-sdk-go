package places

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder provides operations to manage the workspaces property of the microsoft.graph.roomList entity.
type ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderGetQueryParameters get workspaces from places
type ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderGetQueryParameters
}
// ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderInternal instantiates a new ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder and sets the default values.
func NewItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, placeId *string)(*ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) {
    m := &ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/places/{place%2Did}/graph.roomList/workspaces(placeId='{placeId}'){?%24expand,%24select}", pathParameters),
    }
    if placeId != nil {
        m.BaseRequestBuilder.PathParameters["placeId"] = *placeId
    }
    return m
}
// NewItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder instantiates a new ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder and sets the default values.
func NewItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property workspaces for places
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get workspaces from places
// returns a Workspaceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workspaceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkspaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workspaceable), nil
}
// Patch update the navigation property workspaces in places
// returns a Workspaceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workspaceable, requestConfiguration *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workspaceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkspaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workspaceable), nil
}
// ToDeleteRequestInformation delete navigation property workspaces for places
// returns a *RequestInformation when successful
func (m *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get workspaces from places
// returns a *RequestInformation when successful
func (m *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property workspaces in places
// returns a *RequestInformation when successful
func (m *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Workspaceable, requestConfiguration *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder when successful
func (m *ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) WithUrl(rawUrl string)(*ItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder) {
    return NewItemGraphroomlistWorkspaceswithplaceidWorkspacesWithPlaceIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
