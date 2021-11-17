package managementconditions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i11ae1dda1cc3d4e637b3bde035204c5e03a394e8ce52ebcae5e9f6c3a09e32ed "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/managementconditionstatements/item/managementconditions/ref"
    ifac04764f62467483ae92e2f881527331a818498994808d123def46d76dec96c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/managementconditionstatements/item/managementconditions/getmanagementconditionsforplatformwithplatform"
)

// Builds and executes requests for operations under \deviceManagement\managementConditionStatements\{managementConditionStatement-id}\managementConditions
type ManagementConditionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type ManagementConditionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagementConditionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The management conditions associated to the management condition statement.
type ManagementConditionsRequestBuilderGetQueryParameters struct {
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
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Instantiates a new ManagementConditionsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagementConditionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementConditionsRequestBuilder) {
    m := &ManagementConditionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditionStatements/{managementConditionStatement_id}/managementConditions{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ManagementConditionsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagementConditionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementConditionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementConditionsRequestBuilderInternal(urlParams, requestAdapter)
}
// The management conditions associated to the management condition statement.
// Parameters:
//  - options : Options for the request
func (m *ManagementConditionsRequestBuilder) CreateGetRequestInformation(options *ManagementConditionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The management conditions associated to the management condition statement.
// Parameters:
//  - options : Options for the request
func (m *ManagementConditionsRequestBuilder) Get(options *ManagementConditionsRequestBuilderGetOptions)(*ManagementConditionsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementConditionsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagementConditionsResponse), nil
}
// Builds and executes requests for operations under \deviceManagement\managementConditionStatements\{managementConditionStatement-id}\managementConditions\microsoft.graph.getManagementConditionsForPlatform(platform={platform})
// Parameters:
//  - platform : Usage: platform={platform}
func (m *ManagementConditionsRequestBuilder) GetManagementConditionsForPlatformWithPlatform(platform *string)(*ifac04764f62467483ae92e2f881527331a818498994808d123def46d76dec96c.GetManagementConditionsForPlatformWithPlatformRequestBuilder) {
    return ifac04764f62467483ae92e2f881527331a818498994808d123def46d76dec96c.NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal(m.pathParameters, m.requestAdapter, platform);
}
func (m *ManagementConditionsRequestBuilder) Ref()(*i11ae1dda1cc3d4e637b3bde035204c5e03a394e8ce52ebcae5e9f6c3a09e32ed.RefRequestBuilder) {
    return i11ae1dda1cc3d4e637b3bde035204c5e03a394e8ce52ebcae5e9f6c3a09e32ed.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
