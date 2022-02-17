package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1cd82da38eaeb9697597d203dbe05121215aa83c60f4bb16b4c4bb0ae1e829f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item/gethealthmetrictimeseries"
    i9c30fe64f3cadb8062e1e80b68b8d12a9b9727c1f84ad283d585cd25e965ef20 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item/createserverlogcollectionrequest"
    ia9fb0847e9b7d23ecba981329a874216f6257bc66e680123b5b5de57e3c5785f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item/gethealthmetrics"
)

// MicrosoftTunnelServerRequestBuilder builds and executes requests for operations under \deviceManagement\microsoftTunnelSites\{microsoftTunnelSite-id}\microsoftTunnelServers\{microsoftTunnelServer-id}
type MicrosoftTunnelServerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MicrosoftTunnelServerRequestBuilderDeleteOptions options for Delete
type MicrosoftTunnelServerRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MicrosoftTunnelServerRequestBuilderGetOptions options for Get
type MicrosoftTunnelServerRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MicrosoftTunnelServerRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MicrosoftTunnelServerRequestBuilderGetQueryParameters a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
type MicrosoftTunnelServerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MicrosoftTunnelServerRequestBuilderPatchOptions options for Patch
type MicrosoftTunnelServerRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelServer;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewMicrosoftTunnelServerRequestBuilderInternal instantiates a new MicrosoftTunnelServerRequestBuilder and sets the default values.
func NewMicrosoftTunnelServerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MicrosoftTunnelServerRequestBuilder) {
    m := &MicrosoftTunnelServerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite_id}/microsoftTunnelServers/{microsoftTunnelServer_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMicrosoftTunnelServerRequestBuilder instantiates a new MicrosoftTunnelServerRequestBuilder and sets the default values.
func NewMicrosoftTunnelServerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MicrosoftTunnelServerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelServerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelServerRequestBuilder) CreateDeleteRequestInformation(options *MicrosoftTunnelServerRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelServerRequestBuilder) CreateGetRequestInformation(options *MicrosoftTunnelServerRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelServerRequestBuilder) CreatePatchRequestInformation(options *MicrosoftTunnelServerRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MicrosoftTunnelServerRequestBuilder) CreateServerLogCollectionRequest()(*i9c30fe64f3cadb8062e1e80b68b8d12a9b9727c1f84ad283d585cd25e965ef20.CreateServerLogCollectionRequestRequestBuilder) {
    return i9c30fe64f3cadb8062e1e80b68b8d12a9b9727c1f84ad283d585cd25e965ef20.NewCreateServerLogCollectionRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelServerRequestBuilder) Delete(options *MicrosoftTunnelServerRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelServerRequestBuilder) Get(options *MicrosoftTunnelServerRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelServer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMicrosoftTunnelServer() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelServer), nil
}
func (m *MicrosoftTunnelServerRequestBuilder) GetHealthMetrics()(*ia9fb0847e9b7d23ecba981329a874216f6257bc66e680123b5b5de57e3c5785f.GetHealthMetricsRequestBuilder) {
    return ia9fb0847e9b7d23ecba981329a874216f6257bc66e680123b5b5de57e3c5785f.NewGetHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MicrosoftTunnelServerRequestBuilder) GetHealthMetricTimeSeries()(*i1cd82da38eaeb9697597d203dbe05121215aa83c60f4bb16b4c4bb0ae1e829f2.GetHealthMetricTimeSeriesRequestBuilder) {
    return i1cd82da38eaeb9697597d203dbe05121215aa83c60f4bb16b4c4bb0ae1e829f2.NewGetHealthMetricTimeSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelServerRequestBuilder) Patch(options *MicrosoftTunnelServerRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
