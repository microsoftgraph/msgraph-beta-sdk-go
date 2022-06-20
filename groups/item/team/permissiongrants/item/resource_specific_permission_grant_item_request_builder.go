package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11cbe513533da44c68e3ad6043db933eff02f422bbfcfbf0811474479132e745 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants/item/checkmemberobjects"
    i1b79629625ba021d03ce17e2cac30d27db597a4734a4bfd14e2acb610bc708bd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants/item/getmembergroups"
    i2863f2f0c769cf58de6f6eee0e0d6882a3a1be318d45d475c3d63f3b1fce43c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants/item/getmemberobjects"
    i2c781a75929cae50a139655aee3679f1e52a29600e79341f24d30fb60c23eb78 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants/item/restore"
    i8096bc0b4783224e4134c4c6945f7d59e2554d8ffea1b64f1da9580df3a3ad6e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/team/permissiongrants/item/checkmembergroups"
)

// ResourceSpecificPermissionGrantItemRequestBuilder provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
type ResourceSpecificPermissionGrantItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters a collection of permissions granted to apps to access the team.
type ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters
}
// ResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups the checkMemberGroups property
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberGroups()(*i8096bc0b4783224e4134c4c6945f7d59e2554d8ffea1b64f1da9580df3a3ad6e.CheckMemberGroupsRequestBuilder) {
    return i8096bc0b4783224e4134c4c6945f7d59e2554d8ffea1b64f1da9580df3a3ad6e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberObjects()(*i11cbe513533da44c68e3ad6043db933eff02f422bbfcfbf0811474479132e745.CheckMemberObjectsRequestBuilder) {
    return i11cbe513533da44c68e3ad6043db933eff02f422bbfcfbf0811474479132e745.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewResourceSpecificPermissionGrantItemRequestBuilderInternal instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewResourceSpecificPermissionGrantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceSpecificPermissionGrantItemRequestBuilder) {
    m := &ResourceSpecificPermissionGrantItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/permissionGrants/{resourceSpecificPermissionGrant%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResourceSpecificPermissionGrantItemRequestBuilder instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewResourceSpecificPermissionGrantItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property permissionGrants for groups
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property permissionGrants for groups
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a collection of permissions granted to apps to access the team.
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property permissionGrants in groups
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property permissionGrants in groups
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *ResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property permissionGrants for groups
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property permissionGrants for groups
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ResourceSpecificPermissionGrantItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of permissions granted to apps to access the team.
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) GetMemberGroups()(*i1b79629625ba021d03ce17e2cac30d27db597a4734a4bfd14e2acb610bc708bd.GetMemberGroupsRequestBuilder) {
    return i1b79629625ba021d03ce17e2cac30d27db597a4734a4bfd14e2acb610bc708bd.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) GetMemberObjects()(*i2863f2f0c769cf58de6f6eee0e0d6882a3a1be318d45d475c3d63f3b1fce43c2.GetMemberObjectsRequestBuilder) {
    return i2863f2f0c769cf58de6f6eee0e0d6882a3a1be318d45d475c3d63f3b1fce43c2.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler a collection of permissions granted to apps to access the team.
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ResourceSpecificPermissionGrantItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable), nil
}
// Patch update the navigation property permissionGrants in groups
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property permissionGrants in groups
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ResourceSpecificPermissionGrantable, requestConfiguration *ResourceSpecificPermissionGrantItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Restore the restore property
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Restore()(*i2c781a75929cae50a139655aee3679f1e52a29600e79341f24d30fb60c23eb78.RestoreRequestBuilder) {
    return i2c781a75929cae50a139655aee3679f1e52a29600e79341f24d30fb60c23eb78.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
