package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
type DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters the available group policy uploaded definition files for this account.
type DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddLanguageFiles provides operations to call the addLanguageFiles method.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) AddLanguageFiles()(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal instantiates a new GroupPolicyUploadedDefinitionFileItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder instantiates a new GroupPolicyUploadedDefinitionFileItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the available group policy uploaded definition files for this account.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available group policy uploaded definition files for this account.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable), nil
}
// GroupPolicyOperations provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) GroupPolicyOperations()(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsRequestBuilder) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyOperationsById provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) GroupPolicyOperationsById(id string)(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyOperation%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable), nil
}
// Remove provides operations to call the remove method.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Remove()(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemRemoveRequestBuilder) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemRemoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveLanguageFiles provides operations to call the removeLanguageFiles method.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) RemoveLanguageFiles()(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemRemoveLanguageFilesRequestBuilder) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemRemoveLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateLanguageFiles provides operations to call the updateLanguageFiles method.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) UpdateLanguageFiles()(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemUpdateLanguageFilesRequestBuilder) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemUpdateLanguageFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadNewVersion provides operations to call the uploadNewVersion method.
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) UploadNewVersion()(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionRequestBuilder) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemUploadNewVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
