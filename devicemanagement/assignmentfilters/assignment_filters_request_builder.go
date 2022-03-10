package assignmentfilters

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6320107e367b8c68fe46a4cf86ea24694ef25a7c0c2ed084170f433fab284667 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/getplatformsupportedpropertieswithplatform"
    i6df941e8c158ac5060e667569031de0e8e647a89bff7ea10194884e75f1a0bfe "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/count"
    i971b9826d0cedaebe37e1b7dbc8d3baedb6c9f0866d25b684e1da32b17da33a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/enable"
    ia4a12d8e28e11df7e1cb2aab68059eb462d6a64bdaf33642330dcd51be6b4c91 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/getstate"
    ia52abb7c9a696d072bbdd617cc112d98fa0fc464bcc7a2b2808b711dc8fd7208 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/assignmentfilters/validatefilter"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// AssignmentFiltersRequestBuilder provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
type AssignmentFiltersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AssignmentFiltersRequestBuilderGetOptions options for Get
type AssignmentFiltersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AssignmentFiltersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAssignmentFiltersRequestBuilderInternal instantiates a new AssignmentFiltersRequestBuilder and sets the default values.
func NewAssignmentFiltersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AssignmentFiltersRequestBuilder) {
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
func NewAssignmentFiltersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AssignmentFiltersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignmentFiltersRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AssignmentFiltersRequestBuilder) Count()(*i6df941e8c158ac5060e667569031de0e8e647a89bff7ea10194884e75f1a0bfe.CountRequestBuilder) {
    return i6df941e8c158ac5060e667569031de0e8e647a89bff7ea10194884e75f1a0bfe.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of assignment filters
func (m *AssignmentFiltersRequestBuilder) CreateGetRequestInformation(options *AssignmentFiltersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to assignmentFilters for deviceManagement
func (m *AssignmentFiltersRequestBuilder) CreatePostRequestInformation(options *AssignmentFiltersRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *AssignmentFiltersRequestBuilder) Enable()(*i971b9826d0cedaebe37e1b7dbc8d3baedb6c9f0866d25b684e1da32b17da33a6.EnableRequestBuilder) {
    return i971b9826d0cedaebe37e1b7dbc8d3baedb6c9f0866d25b684e1da32b17da33a6.NewEnableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of assignment filters
func (m *AssignmentFiltersRequestBuilder) Get(options *AssignmentFiltersRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceAndAppManagementAssignmentFilterCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterCollectionResponseable), nil
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
func (m *AssignmentFiltersRequestBuilder) Post(options *AssignmentFiltersRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignmentFilterable), nil
}
func (m *AssignmentFiltersRequestBuilder) ValidateFilter()(*ia52abb7c9a696d072bbdd617cc112d98fa0fc464bcc7a2b2808b711dc8fd7208.ValidateFilterRequestBuilder) {
    return ia52abb7c9a696d072bbdd617cc112d98fa0fc464bcc7a2b2808b711dc8fd7208.NewValidateFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
