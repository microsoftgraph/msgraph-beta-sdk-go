package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i364de94ee12d8f96dc648e7c3741312fd56499dc3a23d0298bc8da59cf990f04 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces/item/importresourceactions"
    i8eb06311e73f2c0da7a5efa55417c3517d5b7b351f6b98ff69e266e1d0c56fa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces/item/resourceactions"
    i28bbdad243bea8433bbcaf8db4e3551dcedd77a214aab3105a7ef3ad6592115a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/resourcenamespaces/item/resourceactions/item"
)

// UnifiedRbacResourceNamespaceItemRequestBuilder provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplicationMultiple entity.
type UnifiedRbacResourceNamespaceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters get resourceNamespaces from roleManagement
type UnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters
}
// UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal instantiates a new UnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRbacResourceNamespaceItemRequestBuilder) {
    m := &UnifiedRbacResourceNamespaceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRbacResourceNamespaceItemRequestBuilder instantiates a new UnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewUnifiedRbacResourceNamespaceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property resourceNamespaces for roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property resourceNamespaces for roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration get resourceNamespaces from roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get resourceNamespaces from roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property resourceNamespaces in roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property resourceNamespaces in roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property resourceNamespaces for roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property resourceNamespaces for roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler get resourceNamespaces from roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler get resourceNamespaces from roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ImportResourceActions the importResourceActions property
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) ImportResourceActions()(*i364de94ee12d8f96dc648e7c3741312fd56499dc3a23d0298bc8da59cf990f04.ImportResourceActionsRequestBuilder) {
    return i364de94ee12d8f96dc648e7c3741312fd56499dc3a23d0298bc8da59cf990f04.NewImportResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property resourceNamespaces in roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property resourceNamespaces in roleManagement
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *UnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// ResourceActions the resourceActions property
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActions()(*i8eb06311e73f2c0da7a5efa55417c3517d5b7b351f6b98ff69e266e1d0c56fa8.ResourceActionsRequestBuilder) {
    return i8eb06311e73f2c0da7a5efa55417c3517d5b7b351f6b98ff69e266e1d0c56fa8.NewResourceActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceActionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.roleManagement.deviceManagement.resourceNamespaces.item.resourceActions.item collection
func (m *UnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActionsById(id string)(*i28bbdad243bea8433bbcaf8db4e3551dcedd77a214aab3105a7ef3ad6592115a.UnifiedRbacResourceActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRbacResourceAction%2Did"] = id
    }
    return i28bbdad243bea8433bbcaf8db4e3551dcedd77a214aab3105a7ef3ad6592115a.NewUnifiedRbacResourceActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
