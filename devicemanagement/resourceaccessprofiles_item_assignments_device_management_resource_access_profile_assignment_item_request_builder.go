package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.deviceManagementResourceAccessProfileBase entity.
type ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderGetQueryParameters the list of assignments for the device configuration profile.
type ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderGetQueryParameters
}
// ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderInternal instantiates a new ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder and sets the default values.
func NewResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) {
    m := &ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/resourceAccessProfiles/{deviceManagementResourceAccessProfileBase%2Did}/assignments/{deviceManagementResourceAccessProfileAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder instantiates a new ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder and sets the default values.
func NewResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of assignments for the device configuration profile.
// returns a DeviceManagementResourceAccessProfileAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileAssignmentable), nil
}
// Patch update the navigation property assignments in deviceManagement
// returns a DeviceManagementResourceAccessProfileAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileAssignmentable, requestConfiguration *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of assignments for the device configuration profile.
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in deviceManagement
// returns a *RequestInformation when successful
func (m *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileAssignmentable, requestConfiguration *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder when successful
func (m *ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*ResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) {
    return NewResourceaccessprofilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
