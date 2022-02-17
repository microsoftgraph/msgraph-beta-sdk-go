package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i70fe0c8c492862237342af1645c6716e74e8d1fc6bd0dcea8a7c03a1aae69df5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses"
    ic490bfa15729cb25d9948b8e388970092529ca560f41391770e7660c3d7c9699 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app"
    i52b737a53d0c03e6a268168d1f82ff2bdff242cee05a2ee607462da8e246d8ba "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item"
)

// UserAppInstallStatusRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\userStatuses\{userAppInstallStatus-id}
type UserAppInstallStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserAppInstallStatusRequestBuilderDeleteOptions options for Delete
type UserAppInstallStatusRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserAppInstallStatusRequestBuilderGetOptions options for Get
type UserAppInstallStatusRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserAppInstallStatusRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserAppInstallStatusRequestBuilderGetQueryParameters the list of installation states for this mobile app.
type UserAppInstallStatusRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserAppInstallStatusRequestBuilderPatchOptions options for Patch
type UserAppInstallStatusRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserAppInstallStatus;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UserAppInstallStatusRequestBuilder) App()(*ic490bfa15729cb25d9948b8e388970092529ca560f41391770e7660c3d7c9699.AppRequestBuilder) {
    return ic490bfa15729cb25d9948b8e388970092529ca560f41391770e7660c3d7c9699.NewAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUserAppInstallStatusRequestBuilderInternal instantiates a new UserAppInstallStatusRequestBuilder and sets the default values.
func NewUserAppInstallStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserAppInstallStatusRequestBuilder) {
    m := &UserAppInstallStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/userStatuses/{userAppInstallStatus_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserAppInstallStatusRequestBuilder instantiates a new UserAppInstallStatusRequestBuilder and sets the default values.
func NewUserAppInstallStatusRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserAppInstallStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserAppInstallStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of installation states for this mobile app.
func (m *UserAppInstallStatusRequestBuilder) CreateDeleteRequestInformation(options *UserAppInstallStatusRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of installation states for this mobile app.
func (m *UserAppInstallStatusRequestBuilder) CreateGetRequestInformation(options *UserAppInstallStatusRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of installation states for this mobile app.
func (m *UserAppInstallStatusRequestBuilder) CreatePatchRequestInformation(options *UserAppInstallStatusRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of installation states for this mobile app.
func (m *UserAppInstallStatusRequestBuilder) Delete(options *UserAppInstallStatusRequestBuilderDeleteOptions)(error) {
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
func (m *UserAppInstallStatusRequestBuilder) DeviceStatuses()(*i70fe0c8c492862237342af1645c6716e74e8d1fc6bd0dcea8a7c03a1aae69df5.DeviceStatusesRequestBuilder) {
    return i70fe0c8c492862237342af1645c6716e74e8d1fc6bd0dcea8a7c03a1aae69df5.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileApps.item.userStatuses.item.deviceStatuses.item collection
func (m *UserAppInstallStatusRequestBuilder) DeviceStatusesById(id string)(*i52b737a53d0c03e6a268168d1f82ff2bdff242cee05a2ee607462da8e246d8ba.MobileAppInstallStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppInstallStatus_id"] = id
    }
    return i52b737a53d0c03e6a268168d1f82ff2bdff242cee05a2ee607462da8e246d8ba.NewMobileAppInstallStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of installation states for this mobile app.
func (m *UserAppInstallStatusRequestBuilder) Get(options *UserAppInstallStatusRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserAppInstallStatus, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserAppInstallStatus() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserAppInstallStatus), nil
}
// Patch the list of installation states for this mobile app.
func (m *UserAppInstallStatusRequestBuilder) Patch(options *UserAppInstallStatusRequestBuilderPatchOptions)(error) {
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
