package getrelatedappstateswithuserprincipalnamewithdeviceid

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder provides operations to call the getRelatedAppStates method.
type GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetOptions options for Get
type GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal instantiates a new GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder and sets the default values.
func NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, deviceId *string, userPrincipalName *string)(*GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    m := &GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/microsoft.graph.getRelatedAppStates(userPrincipalName='{userPrincipalName}',deviceId='{deviceId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if deviceId != nil {
        urlTplParams["deviceId"] = *deviceId
    }
    if userPrincipalName != nil {
        urlTplParams["userPrincipalName"] = *userPrincipalName
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder instantiates a new GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder and sets the default values.
func NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function getRelatedAppStates
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) CreateGetRequestInformation(options *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getRelatedAppStates
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) Get(options *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetOptions)(GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable), nil
}
