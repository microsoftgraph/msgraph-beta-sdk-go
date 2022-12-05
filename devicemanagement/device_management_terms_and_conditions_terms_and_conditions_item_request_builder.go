package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
type DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderGetQueryParameters the terms and conditions associated with device management of the company.
type DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderGetQueryParameters
}
// DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptanceStatuses provides operations to manage the acceptanceStatuses property of the microsoft.graph.termsAndConditions entity.
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) AcceptanceStatuses()(*DeviceManagementTermsAndConditionsItemAcceptanceStatusesRequestBuilder) {
    return NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptanceStatusesById provides operations to manage the acceptanceStatuses property of the microsoft.graph.termsAndConditions entity.
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) AcceptanceStatusesById(id string)(*DeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAcceptanceStatus%2Did"] = id
    }
    return NewDeviceManagementTermsAndConditionsItemAcceptanceStatusesTermsAndConditionsAcceptanceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.termsAndConditions entity.
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) Assignments()(*DeviceManagementTermsAndConditionsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementTermsAndConditionsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.termsAndConditions entity.
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementTermsAndConditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAssignment%2Did"] = id
    }
    return NewDeviceManagementTermsAndConditionsItemAssignmentsTermsAndConditionsAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderInternal instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewDeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) {
    m := &DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewDeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property termsAndConditions for deviceManagement
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the terms and conditions associated with device management of the company.
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property termsAndConditions in deviceManagement
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, requestConfiguration *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property termsAndConditions for deviceManagement
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the terms and conditions associated with device management of the company.
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.termsAndConditions entity.
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) GroupAssignments()(*DeviceManagementTermsAndConditionsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceManagementTermsAndConditionsItemGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.termsAndConditions entity.
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) GroupAssignmentsById(id string)(*DeviceManagementTermsAndConditionsItemGroupAssignmentsTermsAndConditionsGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsGroupAssignment%2Did"] = id
    }
    return NewDeviceManagementTermsAndConditionsItemGroupAssignmentsTermsAndConditionsGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property termsAndConditions in deviceManagement
func (m *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, requestConfiguration *DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable), nil
}
