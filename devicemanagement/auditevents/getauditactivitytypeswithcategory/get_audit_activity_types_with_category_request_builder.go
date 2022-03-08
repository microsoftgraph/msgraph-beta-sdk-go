package getauditactivitytypeswithcategory

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetAuditActivityTypesWithCategoryRequestBuilder provides operations to call the getAuditActivityTypes method.
type GetAuditActivityTypesWithCategoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetAuditActivityTypesWithCategoryRequestBuilderGetOptions options for Get
type GetAuditActivityTypesWithCategoryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetAuditActivityTypesWithCategoryRequestBuilderInternal instantiates a new GetAuditActivityTypesWithCategoryRequestBuilder and sets the default values.
func NewGetAuditActivityTypesWithCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, category *string)(*GetAuditActivityTypesWithCategoryRequestBuilder) {
    m := &GetAuditActivityTypesWithCategoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/auditEvents/microsoft.graph.getAuditActivityTypes(category='{category}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if category != nil {
        urlTplParams["category"] = *category
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAuditActivityTypesWithCategoryRequestBuilder instantiates a new GetAuditActivityTypesWithCategoryRequestBuilder and sets the default values.
func NewGetAuditActivityTypesWithCategoryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetAuditActivityTypesWithCategoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAuditActivityTypesWithCategoryRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getAuditActivityTypes
func (m *GetAuditActivityTypesWithCategoryRequestBuilder) CreateGetRequestInformation(options *GetAuditActivityTypesWithCategoryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function getAuditActivityTypes
func (m *GetAuditActivityTypesWithCategoryRequestBuilder) Get(options *GetAuditActivityTypesWithCategoryRequestBuilderGetOptions)(GetAuditActivityTypesWithCategoryResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetAuditActivityTypesWithCategoryResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetAuditActivityTypesWithCategoryResponseable), nil
}
