package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder provides operations to manage the definitions property of the microsoft.graph.groupPolicyDefinitionFile entity.
type DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters the group policy definitions associated with the file.
type DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetQueryParameters
}
// NewDeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal instantiates a new GroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyDefinitionFiles/{groupPolicyDefinitionFile%2Did}/definitions/{groupPolicyDefinition%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder instantiates a new GroupPolicyDefinitionItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the group policy definitions associated with the file.
func (m *DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the group policy definitions associated with the file.
func (m *DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionable), nil
}
