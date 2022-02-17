package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5024ea90c8fd45ac7dbd6e184330159a989748ef5af2324e47ab26071b412298 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers"
    ib99129148e5b54549dbf1a1827b35b5e072543abc1f750848c53c7a2bfcbbf7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/requestupgrade"
    ibbd8be26def4f1c33c115243694230b27231cfa01a930726d7cd6e16f6198b0a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelconfiguration"
    i9350d6dec8de249f17a367f9b40b072f9094b90586b5f3d60925874bf5f54981 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item"
)

// MicrosoftTunnelSiteRequestBuilder builds and executes requests for operations under \deviceManagement\microsoftTunnelSites\{microsoftTunnelSite-id}
type MicrosoftTunnelSiteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MicrosoftTunnelSiteRequestBuilderDeleteOptions options for Delete
type MicrosoftTunnelSiteRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MicrosoftTunnelSiteRequestBuilderGetOptions options for Get
type MicrosoftTunnelSiteRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MicrosoftTunnelSiteRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MicrosoftTunnelSiteRequestBuilderGetQueryParameters collection of MicrosoftTunnelSite settings associated with account.
type MicrosoftTunnelSiteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MicrosoftTunnelSiteRequestBuilderPatchOptions options for Patch
type MicrosoftTunnelSiteRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelSite;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewMicrosoftTunnelSiteRequestBuilderInternal instantiates a new MicrosoftTunnelSiteRequestBuilder and sets the default values.
func NewMicrosoftTunnelSiteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MicrosoftTunnelSiteRequestBuilder) {
    m := &MicrosoftTunnelSiteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMicrosoftTunnelSiteRequestBuilder instantiates a new MicrosoftTunnelSiteRequestBuilder and sets the default values.
func NewMicrosoftTunnelSiteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MicrosoftTunnelSiteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelSiteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSiteRequestBuilder) CreateDeleteRequestInformation(options *MicrosoftTunnelSiteRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSiteRequestBuilder) CreateGetRequestInformation(options *MicrosoftTunnelSiteRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSiteRequestBuilder) CreatePatchRequestInformation(options *MicrosoftTunnelSiteRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSiteRequestBuilder) Delete(options *MicrosoftTunnelSiteRequestBuilderDeleteOptions)(error) {
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
// Get collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSiteRequestBuilder) Get(options *MicrosoftTunnelSiteRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelSite, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMicrosoftTunnelSite() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MicrosoftTunnelSite), nil
}
func (m *MicrosoftTunnelSiteRequestBuilder) MicrosoftTunnelConfiguration()(*ibbd8be26def4f1c33c115243694230b27231cfa01a930726d7cd6e16f6198b0a.MicrosoftTunnelConfigurationRequestBuilder) {
    return ibbd8be26def4f1c33c115243694230b27231cfa01a930726d7cd6e16f6198b0a.NewMicrosoftTunnelConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MicrosoftTunnelSiteRequestBuilder) MicrosoftTunnelServers()(*i5024ea90c8fd45ac7dbd6e184330159a989748ef5af2324e47ab26071b412298.MicrosoftTunnelServersRequestBuilder) {
    return i5024ea90c8fd45ac7dbd6e184330159a989748ef5af2324e47ab26071b412298.NewMicrosoftTunnelServersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelServersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.microsoftTunnelSites.item.microsoftTunnelServers.item collection
func (m *MicrosoftTunnelSiteRequestBuilder) MicrosoftTunnelServersById(id string)(*i9350d6dec8de249f17a367f9b40b072f9094b90586b5f3d60925874bf5f54981.MicrosoftTunnelServerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelServer_id"] = id
    }
    return i9350d6dec8de249f17a367f9b40b072f9094b90586b5f3d60925874bf5f54981.NewMicrosoftTunnelServerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSiteRequestBuilder) Patch(options *MicrosoftTunnelSiteRequestBuilderPatchOptions)(error) {
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
func (m *MicrosoftTunnelSiteRequestBuilder) RequestUpgrade()(*ib99129148e5b54549dbf1a1827b35b5e072543abc1f750848c53c7a2bfcbbf7d.RequestUpgradeRequestBuilder) {
    return ib99129148e5b54549dbf1a1827b35b5e072543abc1f750848c53c7a2bfcbbf7d.NewRequestUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
