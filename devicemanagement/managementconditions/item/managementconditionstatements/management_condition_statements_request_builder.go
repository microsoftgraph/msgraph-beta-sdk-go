package managementconditionstatements

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ie653e6d497ebe0b3127cb9b98857f9799213e0b1bd7e4906a4d8cd8a5280330c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/managementconditions/item/managementconditionstatements/ref"
    if8ca9966930d7ad505e9fc8a4fd9fc00488cdc68bc25732da4bff69daa542a27 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/managementconditions/item/managementconditionstatements/getmanagementconditionstatementsforplatformwithplatform"
)

// ManagementConditionStatementsRequestBuilder builds and executes requests for operations under \deviceManagement\managementConditions\{managementCondition-id}\managementConditionStatements
type ManagementConditionStatementsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagementConditionStatementsRequestBuilderGetOptions options for Get
type ManagementConditionStatementsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagementConditionStatementsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagementConditionStatementsRequestBuilderGetQueryParameters the management condition statements associated to the management condition.
type ManagementConditionStatementsRequestBuilderGetQueryParameters struct {
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
// NewManagementConditionStatementsRequestBuilderInternal instantiates a new ManagementConditionStatementsRequestBuilder and sets the default values.
func NewManagementConditionStatementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementConditionStatementsRequestBuilder) {
    m := &ManagementConditionStatementsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditions/{managementCondition_id}/managementConditionStatements{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagementConditionStatementsRequestBuilder instantiates a new ManagementConditionStatementsRequestBuilder and sets the default values.
func NewManagementConditionStatementsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementConditionStatementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementConditionStatementsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the management condition statements associated to the management condition.
func (m *ManagementConditionStatementsRequestBuilder) CreateGetRequestInformation(options *ManagementConditionStatementsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the management condition statements associated to the management condition.
func (m *ManagementConditionStatementsRequestBuilder) Get(options *ManagementConditionStatementsRequestBuilderGetOptions)(*ManagementConditionStatementsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementConditionStatementsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagementConditionStatementsResponse), nil
}
// GetManagementConditionStatementsForPlatformWithPlatform builds and executes requests for operations under \deviceManagement\managementConditions\{managementCondition-id}\managementConditionStatements\microsoft.graph.getManagementConditionStatementsForPlatform(platform={platform})
func (m *ManagementConditionStatementsRequestBuilder) GetManagementConditionStatementsForPlatformWithPlatform(platform *string)(*if8ca9966930d7ad505e9fc8a4fd9fc00488cdc68bc25732da4bff69daa542a27.GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder) {
    return if8ca9966930d7ad505e9fc8a4fd9fc00488cdc68bc25732da4bff69daa542a27.NewGetManagementConditionStatementsForPlatformWithPlatformRequestBuilderInternal(m.pathParameters, m.requestAdapter, platform);
}
func (m *ManagementConditionStatementsRequestBuilder) Ref()(*ie653e6d497ebe0b3127cb9b98857f9799213e0b1bd7e4906a4d8cd8a5280330c.RefRequestBuilder) {
    return ie653e6d497ebe0b3127cb9b98857f9799213e0b1bd7e4906a4d8cd8a5280330c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
