package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
type DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters the embedded SIM activation code pools created by this account.
type DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters
}
// DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Assign()(*DeviceManagementEmbeddedSIMActivationCodePoolsItemAssignRequestBuilder) {
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.embeddedSIMActivationCodePool entity.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Assignments()(*DeviceManagementEmbeddedSIMActivationCodePoolsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.embeddedSIMActivationCodePool entity.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementEmbeddedSIMActivationCodePoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["embeddedSIMActivationCodePoolAssignment%2Did"] = id
    }
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal instantiates a new EmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewDeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    m := &DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder instantiates a new EmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewDeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the embedded SIM activation code pools created by this account.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) DeviceStates()(*DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesRequestBuilder) {
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) DeviceStatesById(id string)(*DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["embeddedSIMDeviceState%2Did"] = id
    }
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the embedded SIM activation code pools created by this account.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable), nil
}
// Patch update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable), nil
}
