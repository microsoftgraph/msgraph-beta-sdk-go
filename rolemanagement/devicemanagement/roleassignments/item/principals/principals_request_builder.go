package principals

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i46f6b4eb28a394a4151fce1c04efc0ce293bdbe502294c32263760a51431a0db "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/devicemanagement/roleassignments/item/principals/ref"
)

// PrincipalsRequestBuilder builds and executes requests for operations under \roleManagement\deviceManagement\roleAssignments\{unifiedRoleAssignmentMultiple-id}\principals
type PrincipalsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrincipalsRequestBuilderGetOptions options for Get
type PrincipalsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrincipalsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrincipalsRequestBuilderGetQueryParameters read-only collection referencing the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
type PrincipalsRequestBuilderGetQueryParameters struct {
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
// NewPrincipalsRequestBuilderInternal instantiates a new PrincipalsRequestBuilder and sets the default values.
func NewPrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrincipalsRequestBuilder) {
    m := &PrincipalsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/roleAssignments/{unifiedRoleAssignmentMultiple_id}/principals{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrincipalsRequestBuilder instantiates a new PrincipalsRequestBuilder and sets the default values.
func NewPrincipalsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation read-only collection referencing the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *PrincipalsRequestBuilder) CreateGetRequestInformation(options *PrincipalsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get read-only collection referencing the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
func (m *PrincipalsRequestBuilder) Get(options *PrincipalsRequestBuilderGetOptions)(*PrincipalsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrincipalsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*PrincipalsResponse), nil
}
func (m *PrincipalsRequestBuilder) Ref()(*i46f6b4eb28a394a4151fce1c04efc0ce293bdbe502294c32263760a51431a0db.RefRequestBuilder) {
    return i46f6b4eb28a394a4151fce1c04efc0ce293bdbe502294c32263760a51431a0db.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
