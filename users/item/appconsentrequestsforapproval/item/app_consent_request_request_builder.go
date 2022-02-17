package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i966535153f9c1713f26a89395b153ba1a62565b1a6cadf41ea99e6a5b7804c49 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/appconsentrequestsforapproval/item/userconsentrequests"
    ifb0be503d5b998b07a916333abccc7f39d2daf35db880233d26fd1608e59e943 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/appconsentrequestsforapproval/item/userconsentrequests/item"
)

// AppConsentRequestRequestBuilder builds and executes requests for operations under \users\{user-id}\appConsentRequestsForApproval\{appConsentRequest-id}
type AppConsentRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AppConsentRequestRequestBuilderDeleteOptions options for Delete
type AppConsentRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AppConsentRequestRequestBuilderGetOptions options for Get
type AppConsentRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AppConsentRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AppConsentRequestRequestBuilderGetQueryParameters get appConsentRequestsForApproval from users
type AppConsentRequestRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AppConsentRequestRequestBuilderPatchOptions options for Patch
type AppConsentRequestRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppConsentRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAppConsentRequestRequestBuilderInternal instantiates a new AppConsentRequestRequestBuilder and sets the default values.
func NewAppConsentRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppConsentRequestRequestBuilder) {
    m := &AppConsentRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/appConsentRequestsForApproval/{appConsentRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppConsentRequestRequestBuilder instantiates a new AppConsentRequestRequestBuilder and sets the default values.
func NewAppConsentRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppConsentRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppConsentRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property appConsentRequestsForApproval for users
func (m *AppConsentRequestRequestBuilder) CreateDeleteRequestInformation(options *AppConsentRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get appConsentRequestsForApproval from users
func (m *AppConsentRequestRequestBuilder) CreateGetRequestInformation(options *AppConsentRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property appConsentRequestsForApproval in users
func (m *AppConsentRequestRequestBuilder) CreatePatchRequestInformation(options *AppConsentRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property appConsentRequestsForApproval for users
func (m *AppConsentRequestRequestBuilder) Delete(options *AppConsentRequestRequestBuilderDeleteOptions)(error) {
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
// Get get appConsentRequestsForApproval from users
func (m *AppConsentRequestRequestBuilder) Get(options *AppConsentRequestRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppConsentRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAppConsentRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AppConsentRequest), nil
}
// Patch update the navigation property appConsentRequestsForApproval in users
func (m *AppConsentRequestRequestBuilder) Patch(options *AppConsentRequestRequestBuilderPatchOptions)(error) {
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
func (m *AppConsentRequestRequestBuilder) UserConsentRequests()(*i966535153f9c1713f26a89395b153ba1a62565b1a6cadf41ea99e6a5b7804c49.UserConsentRequestsRequestBuilder) {
    return i966535153f9c1713f26a89395b153ba1a62565b1a6cadf41ea99e6a5b7804c49.NewUserConsentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserConsentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.appConsentRequestsForApproval.item.userConsentRequests.item collection
func (m *AppConsentRequestRequestBuilder) UserConsentRequestsById(id string)(*ifb0be503d5b998b07a916333abccc7f39d2daf35db880233d26fd1608e59e943.UserConsentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userConsentRequest_id"] = id
    }
    return ifb0be503d5b998b07a916333abccc7f39d2daf35db880233d26fd1608e59e943.NewUserConsentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
