package teamtemplatedefinition

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
type TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters a collection of permissions granted to apps to access the team.
type TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters
}
// TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberGroups()(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemCheckMemberGroupsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberObjects()(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    m := &TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamTemplateDefinition/{teamTemplateDefinition%2Did}/teamDefinition/permissionGrants/{resourceSpecificPermissionGrant%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property permissionGrants for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of permissions granted to apps to access the team.
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property permissionGrants in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property permissionGrants for teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of permissions granted to apps to access the team.
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) GetMemberGroups()(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemGetMemberGroupsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) GetMemberObjects()(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemGetMemberObjectsRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property permissionGrants in teamTemplateDefinition
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable), nil
}
// Restore provides operations to call the restore method.
func (m *TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) Restore()(*TeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemRestoreRequestBuilder) {
    return NewTeamTemplateDefinitionItemTeamDefinitionPermissionGrantsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
