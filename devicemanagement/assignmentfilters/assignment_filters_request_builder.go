package assignmentfilters

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i6320107e367b8c68fe46a4cf86ea24694ef25a7c0c2ed084170f433fab284667 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/getplatformsupportedpropertieswithplatform"
    i6df941e8c158ac5060e667569031de0e8e647a89bff7ea10194884e75f1a0bfe "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/count"
    i971b9826d0cedaebe37e1b7dbc8d3baedb6c9f0866d25b684e1da32b17da33a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/enable"
    ia4a12d8e28e11df7e1cb2aab68059eb462d6a64bdaf33642330dcd51be6b4c91 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/getstate"
    ia52abb7c9a696d072bbdd617cc112d98fa0fc464bcc7a2b2808b711dc8fd7208 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/validatefilter"
)

// AssignmentFiltersRequestBuilder provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
type AssignmentFiltersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AssignmentFiltersRequestBuilderGetOptions options for Get
type AssignmentFiltersRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *AssignmentFiltersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AssignmentFiltersRequestBuilderGetQueryParameters the list of assignment filters
type AssignmentFiltersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// AssignmentFiltersRequestBuilderPostOptions options for Post
type AssignmentFiltersRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewAssignmentFiltersRequestBuilderInternal instantiates a new AssignmentFiltersRequestBuilder and sets the default values.
func NewAssignmentFiltersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentFiltersRequestBuilder) {
    m := &AssignmentFiltersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/assignmentFilters{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAssignmentFiltersRequestBuilder instantiates a new AssignmentFiltersRequestBuilder and sets the default values.
func NewAssignmentFiltersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignmentFiltersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentFiltersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *AssignmentFiltersRequestBuilder) Count()(*i6df941e8c158ac5060e667569031de0e8e647a89bff7ea10194884e75f1a0bfe.CountRequestBuilder) {
    return i6df941e8c158ac5060e667569031de0e8e647a89bff7ea10194884e75f1a0bfe.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of assignment filters
func (m *AssignmentFiltersRequestBuilder) CreateGetRequestInformation(options *AssignmentFiltersRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to assignmentFilters for deviceManagement
func (m *AssignmentFiltersRequestBuilder) CreatePostRequestInformation(options *AssignmentFiltersRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Enable the enable property
func (m *AssignmentFiltersRequestBuilder) Enable()(*i971b9826d0cedaebe37e1b7dbc8d3baedb6c9f0866d25b684e1da32b17da33a6.EnableRequestBuilder) {
    return i971b9826d0cedaebe37e1b7dbc8d3baedb6c9f0866d25b684e1da32b17da33a6.NewEnableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of assignment filters
func (m *AssignmentFiltersRequestBuilder) Get(options *AssignmentFiltersRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementAssignmentFilterCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterCollectionResponseable), nil
}
// GetPlatformSupportedPropertiesWithPlatform provides operations to call the getPlatformSupportedProperties method.
func (m *AssignmentFiltersRequestBuilder) GetPlatformSupportedPropertiesWithPlatform(platform *string)(*i6320107e367b8c68fe46a4cf86ea24694ef25a7c0c2ed084170f433fab284667.GetPlatformSupportedPropertiesWithPlatformRequestBuilder) {
    return i6320107e367b8c68fe46a4cf86ea24694ef25a7c0c2ed084170f433fab284667.NewGetPlatformSupportedPropertiesWithPlatformRequestBuilderInternal(m.pathParameters, m.requestAdapter, platform);
}
// GetState provides operations to call the getState method.
func (m *AssignmentFiltersRequestBuilder) GetState()(*ia4a12d8e28e11df7e1cb2aab68059eb462d6a64bdaf33642330dcd51be6b4c91.GetStateRequestBuilder) {
    return ia4a12d8e28e11df7e1cb2aab68059eb462d6a64bdaf33642330dcd51be6b4c91.NewGetStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to assignmentFilters for deviceManagement
func (m *AssignmentFiltersRequestBuilder) Post(options *AssignmentFiltersRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementAssignmentFilterable), nil
}
// ValidateFilter the validateFilter property
func (m *AssignmentFiltersRequestBuilder) ValidateFilter()(*ia52abb7c9a696d072bbdd617cc112d98fa0fc464bcc7a2b2808b711dc8fd7208.ValidateFilterRequestBuilder) {
    return ia52abb7c9a696d072bbdd617cc112d98fa0fc464bcc7a2b2808b711dc8fd7208.NewValidateFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
