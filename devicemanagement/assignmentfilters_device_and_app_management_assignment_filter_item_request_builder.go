package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
type AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderGetQueryParameters the list of assignment filters
type AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderGetQueryParameters
}
// AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderInternal instantiates a new AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder and sets the default values.
func NewAssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) {
    m := &AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/assignmentFilters/{deviceAndAppManagementAssignmentFilter%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder instantiates a new AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder and sets the default values.
func NewAssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentFilters for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of assignment filters
// returns a DeviceAndAppManagementAssignmentFilterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable), nil
}
// GetSupportedProperties provides operations to call the getSupportedProperties method.
// returns a *AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder when successful
func (m *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) GetSupportedProperties()(*AssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilder) {
    return NewAssignmentfiltersItemGetsupportedpropertiesGetSupportedPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignmentFilters in deviceManagement
// returns a DeviceAndAppManagementAssignmentFilterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable, requestConfiguration *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable), nil
}
// ToDeleteRequestInformation delete navigation property assignmentFilters for deviceManagement
// returns a *RequestInformation when successful
func (m *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of assignment filters
// returns a *RequestInformation when successful
func (m *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentFilters in deviceManagement
// returns a *RequestInformation when successful
func (m *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable, requestConfiguration *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder when successful
func (m *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) WithUrl(rawUrl string)(*AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) {
    return NewAssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
