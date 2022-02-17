package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0639a56de300c8838fa264733c3d5834218f249aaab32ced8d97be5126cacf72 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryobjects/item/restore"
    i3862728ee9f421a8cc73557f4e6e38099fbb72d6624451de105fef3ae94552fd "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryobjects/item/getmemberobjects"
    i3e65d17a52d9eba4ea676f1e350d27aaf963bb5893e59864749abc7b8b505687 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryobjects/item/checkmemberobjects"
    i4010f9e207892cc07b7c30f271f1e014aaef4a3234180b7061ce98a1e5f1aca0 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryobjects/item/getmembergroups"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ide1a987287b4abfcc07513b59b06ed6d9c26e8ca14e50b0cb0ecb6c9c589c466 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryobjects/item/checkmembergroups"
)

// DirectoryObjectRequestBuilder builds and executes requests for operations under \directoryObjects\{directoryObject-id}
type DirectoryObjectRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryObjectRequestBuilderDeleteOptions options for Delete
type DirectoryObjectRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryObjectRequestBuilderGetOptions options for Get
type DirectoryObjectRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DirectoryObjectRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DirectoryObjectRequestBuilderGetQueryParameters get entity from directoryObjects by key
type DirectoryObjectRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DirectoryObjectRequestBuilderPatchOptions options for Patch
type DirectoryObjectRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DirectoryObjectRequestBuilder) CheckMemberGroups()(*ide1a987287b4abfcc07513b59b06ed6d9c26e8ca14e50b0cb0ecb6c9c589c466.CheckMemberGroupsRequestBuilder) {
    return ide1a987287b4abfcc07513b59b06ed6d9c26e8ca14e50b0cb0ecb6c9c589c466.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryObjectRequestBuilder) CheckMemberObjects()(*i3e65d17a52d9eba4ea676f1e350d27aaf963bb5893e59864749abc7b8b505687.CheckMemberObjectsRequestBuilder) {
    return i3e65d17a52d9eba4ea676f1e350d27aaf963bb5893e59864749abc7b8b505687.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryObjectRequestBuilderInternal instantiates a new DirectoryObjectRequestBuilder and sets the default values.
func NewDirectoryObjectRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryObjectRequestBuilder) {
    m := &DirectoryObjectRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryObjects/{directoryObject_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryObjectRequestBuilder instantiates a new DirectoryObjectRequestBuilder and sets the default values.
func NewDirectoryObjectRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DirectoryObjectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryObjectRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directoryObjects
func (m *DirectoryObjectRequestBuilder) CreateDeleteRequestInformation(options *DirectoryObjectRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from directoryObjects by key
func (m *DirectoryObjectRequestBuilder) CreateGetRequestInformation(options *DirectoryObjectRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in directoryObjects
func (m *DirectoryObjectRequestBuilder) CreatePatchRequestInformation(options *DirectoryObjectRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from directoryObjects
func (m *DirectoryObjectRequestBuilder) Delete(options *DirectoryObjectRequestBuilderDeleteOptions)(error) {
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
// Get get entity from directoryObjects by key
func (m *DirectoryObjectRequestBuilder) Get(options *DirectoryObjectRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectoryObject() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject), nil
}
func (m *DirectoryObjectRequestBuilder) GetMemberGroups()(*i4010f9e207892cc07b7c30f271f1e014aaef4a3234180b7061ce98a1e5f1aca0.GetMemberGroupsRequestBuilder) {
    return i4010f9e207892cc07b7c30f271f1e014aaef4a3234180b7061ce98a1e5f1aca0.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DirectoryObjectRequestBuilder) GetMemberObjects()(*i3862728ee9f421a8cc73557f4e6e38099fbb72d6624451de105fef3ae94552fd.GetMemberObjectsRequestBuilder) {
    return i3862728ee9f421a8cc73557f4e6e38099fbb72d6624451de105fef3ae94552fd.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directoryObjects
func (m *DirectoryObjectRequestBuilder) Patch(options *DirectoryObjectRequestBuilderPatchOptions)(error) {
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
func (m *DirectoryObjectRequestBuilder) Restore()(*i0639a56de300c8838fa264733c3d5834218f249aaab32ced8d97be5126cacf72.RestoreRequestBuilder) {
    return i0639a56de300c8838fa264733c3d5834218f249aaab32ced8d97be5126cacf72.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
