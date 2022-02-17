package directoryscopes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i88e0e7e2f1e3e4368ab5f710c4e4cb200619cdaec5e65074d4d0aa48693e0581 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item/directoryscopes/ref"
)

// DirectoryScopesRequestBuilder builds and executes requests for operations under \roleManagement\cloudPC\roleAssignments\{unifiedRoleAssignmentMultiple-id}\directoryScopes
type DirectoryScopesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryScopesRequestBuilderGetOptions options for Get
type DirectoryScopesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryScopesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryScopesRequestBuilderGetQueryParameters read-only collection referencing the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
type DirectoryScopesRequestBuilderGetQueryParameters struct {
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
// NewDirectoryScopesRequestBuilderInternal instantiates a new DirectoryScopesRequestBuilder and sets the default values.
func NewDirectoryScopesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryScopesRequestBuilder) {
    m := &DirectoryScopesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/cloudPC/roleAssignments/{unifiedRoleAssignmentMultiple_id}/directoryScopes{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryScopesRequestBuilder instantiates a new DirectoryScopesRequestBuilder and sets the default values.
func NewDirectoryScopesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryScopesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryScopesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation read-only collection referencing the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *DirectoryScopesRequestBuilder) CreateGetRequestInformation(options *DirectoryScopesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get read-only collection referencing the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *DirectoryScopesRequestBuilder) Get(options *DirectoryScopesRequestBuilderGetOptions)(*DirectoryScopesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryScopesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DirectoryScopesResponse), nil
}
func (m *DirectoryScopesRequestBuilder) Ref()(*i88e0e7e2f1e3e4368ab5f710c4e4cb200619cdaec5e65074d4d0aa48693e0581.RefRequestBuilder) {
    return i88e0e7e2f1e3e4368ab5f710c4e4cb200619cdaec5e65074d4d0aa48693e0581.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
