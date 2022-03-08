package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i258a858c85da90b801e765428af11e65bc154f922525178ab29e97c7a14c345e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/certificateconnectordetails/item/gethealthmetrictimeseries"
    i7e1773f4927e9225f8c986534d86959f50bad1e76f4272d85d379daeb2d90bee "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/certificateconnectordetails/item/gethealthmetrics"
)

// CertificateConnectorDetailsItemRequestBuilder provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
type CertificateConnectorDetailsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CertificateConnectorDetailsItemRequestBuilderDeleteOptions options for Delete
type CertificateConnectorDetailsItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CertificateConnectorDetailsItemRequestBuilderGetOptions options for Get
type CertificateConnectorDetailsItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CertificateConnectorDetailsItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CertificateConnectorDetailsItemRequestBuilderGetQueryParameters collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
type CertificateConnectorDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CertificateConnectorDetailsItemRequestBuilderPatchOptions options for Patch
type CertificateConnectorDetailsItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CertificateConnectorDetailsable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCertificateConnectorDetailsItemRequestBuilderInternal instantiates a new CertificateConnectorDetailsItemRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CertificateConnectorDetailsItemRequestBuilder) {
    m := &CertificateConnectorDetailsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCertificateConnectorDetailsItemRequestBuilder instantiates a new CertificateConnectorDetailsItemRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CertificateConnectorDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateConnectorDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property certificateConnectorDetails for deviceManagement
func (m *CertificateConnectorDetailsItemRequestBuilder) CreateDeleteRequestInformation(options *CertificateConnectorDetailsItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *CertificateConnectorDetailsItemRequestBuilder) CreateGetRequestInformation(options *CertificateConnectorDetailsItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property certificateConnectorDetails in deviceManagement
func (m *CertificateConnectorDetailsItemRequestBuilder) CreatePatchRequestInformation(options *CertificateConnectorDetailsItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete delete navigation property certificateConnectorDetails for deviceManagement
func (m *CertificateConnectorDetailsItemRequestBuilder) Delete(options *CertificateConnectorDetailsItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *CertificateConnectorDetailsItemRequestBuilder) Get(options *CertificateConnectorDetailsItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CertificateConnectorDetailsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCertificateConnectorDetailsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CertificateConnectorDetailsable), nil
}
func (m *CertificateConnectorDetailsItemRequestBuilder) GetHealthMetrics()(*i7e1773f4927e9225f8c986534d86959f50bad1e76f4272d85d379daeb2d90bee.GetHealthMetricsRequestBuilder) {
    return i7e1773f4927e9225f8c986534d86959f50bad1e76f4272d85d379daeb2d90bee.NewGetHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CertificateConnectorDetailsItemRequestBuilder) GetHealthMetricTimeSeries()(*i258a858c85da90b801e765428af11e65bc154f922525178ab29e97c7a14c345e.GetHealthMetricTimeSeriesRequestBuilder) {
    return i258a858c85da90b801e765428af11e65bc154f922525178ab29e97c7a14c345e.NewGetHealthMetricTimeSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property certificateConnectorDetails in deviceManagement
func (m *CertificateConnectorDetailsItemRequestBuilder) Patch(options *CertificateConnectorDetailsItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
