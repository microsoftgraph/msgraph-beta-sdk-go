package app

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6c8873a2917406b0fb3f914a42484e42b2c5136b48520bce34114e5e16d9ab33 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app/iosvppapp"
    i8f7a908a43d617c31758f6f6191470629d663c61ebec8d28652145d6ca5de70c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app/assign"
    i9bdb0a85763d0a60eda1572d933abe726db35f7a1d827e18c23f13dc5ebcc221 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app/getrelatedappstateswithuserprincipalnamewithdeviceid"
    if5bde5c338c63f67ab86e75f7ab2df462c2280f068e4be67e5823ecd7d96939f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app/updaterelationships"
    ifa01d92abdc8ebd05b3c57289ef897db762e4230f82ec635ead0b141a0dd3516 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app/ref"
)

// AppRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\userStatuses\{userAppInstallStatus-id}\app
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
func (m *AppRequestBuilder) Assign()(*i8f7a908a43d617c31758f6f6191470629d663c61ebec8d28652145d6ca5de70c.AssignRequestBuilder) {
    return i8f7a908a43d617c31758f6f6191470629d663c61ebec8d28652145d6ca5de70c.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAppRequestBuilderInternal instantiates a new AppRequestBuilder and sets the default values.
func NewAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppRequestBuilder) {
    m := &AppRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/userStatuses/{userAppInstallStatus_id}/app{?select,expand}";
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
// GetRelatedAppStatesWithUserPrincipalNameWithDeviceId builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\userStatuses\{userAppInstallStatus-id}\app\microsoft.graph.getRelatedAppStates(userPrincipalName='{userPrincipalName}',deviceId='{deviceId}')
func (m *AppRequestBuilder) GetRelatedAppStatesWithUserPrincipalNameWithDeviceId(userPrincipalName *string, deviceId *string)(*i9bdb0a85763d0a60eda1572d933abe726db35f7a1d827e18c23f13dc5ebcc221.GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    return i9bdb0a85763d0a60eda1572d933abe726db35f7a1d827e18c23f13dc5ebcc221.NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, userPrincipalName, deviceId);
}
func (m *AppRequestBuilder) IosVppApp()(*i6c8873a2917406b0fb3f914a42484e42b2c5136b48520bce34114e5e16d9ab33.IosVppAppRequestBuilder) {
    return i6c8873a2917406b0fb3f914a42484e42b2c5136b48520bce34114e5e16d9ab33.NewIosVppAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AppRequestBuilder) Ref()(*ifa01d92abdc8ebd05b3c57289ef897db762e4230f82ec635ead0b141a0dd3516.RefRequestBuilder) {
    return ifa01d92abdc8ebd05b3c57289ef897db762e4230f82ec635ead0b141a0dd3516.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AppRequestBuilder) UpdateRelationships()(*if5bde5c338c63f67ab86e75f7ab2df462c2280f068e4be67e5823ecd7d96939f.UpdateRelationshipsRequestBuilder) {
    return if5bde5c338c63f67ab86e75f7ab2df462c2280f068e4be67e5823ecd7d96939f.NewUpdateRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
