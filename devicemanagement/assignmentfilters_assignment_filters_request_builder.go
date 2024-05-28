package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignmentfiltersAssignmentFiltersRequestBuilder provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
type AssignmentfiltersAssignmentFiltersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AssignmentfiltersAssignmentFiltersRequestBuilderGetQueryParameters the list of assignment filters
type AssignmentfiltersAssignmentFiltersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AssignmentfiltersAssignmentFiltersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersAssignmentFiltersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AssignmentfiltersAssignmentFiltersRequestBuilderGetQueryParameters
}
// AssignmentfiltersAssignmentFiltersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignmentfiltersAssignmentFiltersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceAndAppManagementAssignmentFilterId provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
// returns a *AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) ByDeviceAndAppManagementAssignmentFilterId(deviceAndAppManagementAssignmentFilterId string)(*AssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceAndAppManagementAssignmentFilterId != "" {
        urlTplParams["deviceAndAppManagementAssignmentFilter%2Did"] = deviceAndAppManagementAssignmentFilterId
    }
    return NewAssignmentfiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAssignmentfiltersAssignmentFiltersRequestBuilderInternal instantiates a new AssignmentfiltersAssignmentFiltersRequestBuilder and sets the default values.
func NewAssignmentfiltersAssignmentFiltersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersAssignmentFiltersRequestBuilder) {
    m := &AssignmentfiltersAssignmentFiltersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/assignmentFilters{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAssignmentfiltersAssignmentFiltersRequestBuilder instantiates a new AssignmentfiltersAssignmentFiltersRequestBuilder and sets the default values.
func NewAssignmentfiltersAssignmentFiltersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentfiltersAssignmentFiltersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentfiltersAssignmentFiltersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AssignmentfiltersCountRequestBuilder when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) Count()(*AssignmentfiltersCountRequestBuilder) {
    return NewAssignmentfiltersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enable provides operations to call the enable method.
// returns a *AssignmentfiltersEnableRequestBuilder when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) Enable()(*AssignmentfiltersEnableRequestBuilder) {
    return NewAssignmentfiltersEnableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of assignment filters
// returns a DeviceAndAppManagementAssignmentFilterCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) Get(ctx context.Context, requestConfiguration *AssignmentfiltersAssignmentFiltersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementAssignmentFilterCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterCollectionResponseable), nil
}
// GetPlatformSupportedPropertiesWithPlatform provides operations to call the getPlatformSupportedProperties method.
// returns a *AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) GetPlatformSupportedPropertiesWithPlatform(platform *string)(*AssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    return NewAssignmentfiltersGetplatformsupportedpropertieswithplatformGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, platform)
}
// GetState provides operations to call the getState method.
// returns a *AssignmentfiltersGetstateGetStateRequestBuilder when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) GetState()(*AssignmentfiltersGetstateGetStateRequestBuilder) {
    return NewAssignmentfiltersGetstateGetStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to assignmentFilters for deviceManagement
// returns a DeviceAndAppManagementAssignmentFilterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable, requestConfiguration *AssignmentfiltersAssignmentFiltersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the list of assignment filters
// returns a *RequestInformation when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AssignmentfiltersAssignmentFiltersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to assignmentFilters for deviceManagement
// returns a *RequestInformation when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable, requestConfiguration *AssignmentfiltersAssignmentFiltersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// ValidateFilter provides operations to call the validateFilter method.
// returns a *AssignmentfiltersValidatefilterValidateFilterRequestBuilder when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) ValidateFilter()(*AssignmentfiltersValidatefilterValidateFilterRequestBuilder) {
    return NewAssignmentfiltersValidatefilterValidateFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AssignmentfiltersAssignmentFiltersRequestBuilder when successful
func (m *AssignmentfiltersAssignmentFiltersRequestBuilder) WithUrl(rawUrl string)(*AssignmentfiltersAssignmentFiltersRequestBuilder) {
    return NewAssignmentfiltersAssignmentFiltersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
