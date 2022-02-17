package app

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i24efdeeca2828ffbbda31ba8d999cf141e479016e6ec383ba9ecca7d26716708 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/ref"
    i83a88b79dd560fb474962443541b74a011d3215b155004bb396ed92e9e2a67e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/updaterelationships"
    i9c64b39583d729fe92b10acc32c6506a89115949d33b387b2e4548141623f2ee "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/iosvppapp"
    id61593f4ecef7c9ca8114d08ebdda35913106925acf434b1185bc5b0c5c8f1ac "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/assign"
    if3c95e63b2ea961b05101c4b865ef29a9b66d6043796833062d636adcc37fbe8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/getrelatedappstateswithuserprincipalnamewithdeviceid"
)

// AppRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\deviceStatuses\{mobileAppInstallStatus-id}\app
type AppRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AppRequestBuilderGetOptions options for Get
type AppRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AppRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AppRequestBuilderGetQueryParameters the navigation link to the mobile app.
type AppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *AppRequestBuilder) Assign()(*id61593f4ecef7c9ca8114d08ebdda35913106925acf434b1185bc5b0c5c8f1ac.AssignRequestBuilder) {
    return id61593f4ecef7c9ca8114d08ebdda35913106925acf434b1185bc5b0c5c8f1ac.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAppRequestBuilderInternal instantiates a new AppRequestBuilder and sets the default values.
func NewAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppRequestBuilder) {
    m := &AppRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/deviceStatuses/{mobileAppInstallStatus_id}/app{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppRequestBuilder instantiates a new AppRequestBuilder and sets the default values.
func NewAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the navigation link to the mobile app.
func (m *AppRequestBuilder) CreateGetRequestInformation(options *AppRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the navigation link to the mobile app.
func (m *AppRequestBuilder) Get(options *AppRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMobileApp() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MobileApp), nil
}
// GetRelatedAppStatesWithUserPrincipalNameWithDeviceId builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\deviceStatuses\{mobileAppInstallStatus-id}\app\microsoft.graph.getRelatedAppStates(userPrincipalName='{userPrincipalName}',deviceId='{deviceId}')
func (m *AppRequestBuilder) GetRelatedAppStatesWithUserPrincipalNameWithDeviceId(userPrincipalName *string, deviceId *string)(*if3c95e63b2ea961b05101c4b865ef29a9b66d6043796833062d636adcc37fbe8.GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    return if3c95e63b2ea961b05101c4b865ef29a9b66d6043796833062d636adcc37fbe8.NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, userPrincipalName, deviceId);
}
func (m *AppRequestBuilder) IosVppApp()(*i9c64b39583d729fe92b10acc32c6506a89115949d33b387b2e4548141623f2ee.IosVppAppRequestBuilder) {
    return i9c64b39583d729fe92b10acc32c6506a89115949d33b387b2e4548141623f2ee.NewIosVppAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AppRequestBuilder) Ref()(*i24efdeeca2828ffbbda31ba8d999cf141e479016e6ec383ba9ecca7d26716708.RefRequestBuilder) {
    return i24efdeeca2828ffbbda31ba8d999cf141e479016e6ec383ba9ecca7d26716708.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AppRequestBuilder) UpdateRelationships()(*i83a88b79dd560fb474962443541b74a011d3215b155004bb396ed92e9e2a67e6.UpdateRelationshipsRequestBuilder) {
    return i83a88b79dd560fb474962443541b74a011d3215b155004bb396ed92e9e2a67e6.NewUpdateRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
