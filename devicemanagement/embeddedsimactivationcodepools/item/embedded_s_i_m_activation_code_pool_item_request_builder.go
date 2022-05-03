package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assignments"
    i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assign"
    ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/devicestates"
    i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assignments/item"
    if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/devicestates/item"
)

// EmbeddedSIMActivationCodePoolItemRequestBuilder provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
type EmbeddedSIMActivationCodePoolItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters the embedded SIM activation code pools created by this account.
type EmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters
}
// EmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) Assign()(*i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81.AssignRequestBuilder) {
    return i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) Assignments()(*i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e.AssignmentsRequestBuilder) {
    return i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.embeddedSIMActivationCodePools.item.assignments.item collection
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) AssignmentsById(id string)(*i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45.EmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["embeddedSIMActivationCodePoolAssignment%2Did"] = id
    }
    return i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45.NewEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEmbeddedSIMActivationCodePoolItemRequestBuilderInternal instantiates a new EmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedSIMActivationCodePoolItemRequestBuilder) {
    m := &EmbeddedSIMActivationCodePoolItemRequestBuilder{
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
// NewEmbeddedSIMActivationCodePoolItemRequestBuilder instantiates a new EmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the embedded SIM activation code pools created by this account.
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the embedded SIM activation code pools created by this account.
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// DeviceStates the deviceStates property
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) DeviceStates()(*ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3.DeviceStatesRequestBuilder) {
    return ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.embeddedSIMActivationCodePools.item.deviceStates.item collection
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) DeviceStatesById(id string)(*if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01.EmbeddedSIMDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["embeddedSIMDeviceState%2Did"] = id
    }
    return if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01.NewEmbeddedSIMDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GetWithResponseHandler the embedded SIM activation code pools created by this account.
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) GetWithResponseHandler(requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the embedded SIM activation code pools created by this account.
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) GetWithResponseHandler(requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable), nil
}
// PatchWithResponseHandler update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
