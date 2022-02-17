package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i097bfeda18c323e4c2c8ad2a58999a280c8282d0e51cffc0c46c63a73351825f "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list"
    i1201c3f43119d05817263421e5f147bf42b57283cb9f554ef5b144a491f89625 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/items"
    i2b03e5349f14f4e882066fc04d344abd4c4a92b60832759bc6992a3d82e123ca "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/root"
    i5243f1e851b438da7e97e810d9e472d42d508e31f347cb52482c0b6e22a96dc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9bb99583bee5cdfb8243ddb4415954a793021828f7cadbd6bd2f195a9d7d4ea7 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/site"
    id98d79401b99ae85b3bc19d7bb389aa17f2bfff9077e3e86b4de0856908138fd "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/driveitem"
    ife9d8be655ee27d89f6d70f3a9fe6d45a7f877b7c963a8bc1ad46f7c7a9cfee6 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/permission"
)

// SharedDriveItemRequestBuilder builds and executes requests for operations under \shares\{sharedDriveItem-id}
type SharedDriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SharedDriveItemRequestBuilderDeleteOptions options for Delete
type SharedDriveItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SharedDriveItemRequestBuilderGetOptions options for Get
type SharedDriveItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SharedDriveItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SharedDriveItemRequestBuilderGetQueryParameters get entity from shares by key
type SharedDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SharedDriveItemRequestBuilderPatchOptions options for Patch
type SharedDriveItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SharedDriveItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSharedDriveItemRequestBuilderInternal instantiates a new SharedDriveItemRequestBuilder and sets the default values.
func NewSharedDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SharedDriveItemRequestBuilder) {
    m := &SharedDriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSharedDriveItemRequestBuilder instantiates a new SharedDriveItemRequestBuilder and sets the default values.
func NewSharedDriveItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SharedDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharedDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from shares
func (m *SharedDriveItemRequestBuilder) CreateDeleteRequestInformation(options *SharedDriveItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from shares by key
func (m *SharedDriveItemRequestBuilder) CreateGetRequestInformation(options *SharedDriveItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in shares
func (m *SharedDriveItemRequestBuilder) CreatePatchRequestInformation(options *SharedDriveItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from shares
func (m *SharedDriveItemRequestBuilder) Delete(options *SharedDriveItemRequestBuilderDeleteOptions)(error) {
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
func (m *SharedDriveItemRequestBuilder) DriveItem()(*id98d79401b99ae85b3bc19d7bb389aa17f2bfff9077e3e86b4de0856908138fd.DriveItemRequestBuilder) {
    return id98d79401b99ae85b3bc19d7bb389aa17f2bfff9077e3e86b4de0856908138fd.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from shares by key
func (m *SharedDriveItemRequestBuilder) Get(options *SharedDriveItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SharedDriveItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSharedDriveItem() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SharedDriveItem), nil
}
func (m *SharedDriveItemRequestBuilder) Items()(*i1201c3f43119d05817263421e5f147bf42b57283cb9f554ef5b144a491f89625.ItemsRequestBuilder) {
    return i1201c3f43119d05817263421e5f147bf42b57283cb9f554ef5b144a491f89625.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemRequestBuilder) List()(*i097bfeda18c323e4c2c8ad2a58999a280c8282d0e51cffc0c46c63a73351825f.ListRequestBuilder) {
    return i097bfeda18c323e4c2c8ad2a58999a280c8282d0e51cffc0c46c63a73351825f.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemRequestBuilder) ListItem()(*i5243f1e851b438da7e97e810d9e472d42d508e31f347cb52482c0b6e22a96dc0.ListItemRequestBuilder) {
    return i5243f1e851b438da7e97e810d9e472d42d508e31f347cb52482c0b6e22a96dc0.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in shares
func (m *SharedDriveItemRequestBuilder) Patch(options *SharedDriveItemRequestBuilderPatchOptions)(error) {
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
func (m *SharedDriveItemRequestBuilder) Permission()(*ife9d8be655ee27d89f6d70f3a9fe6d45a7f877b7c963a8bc1ad46f7c7a9cfee6.PermissionRequestBuilder) {
    return ife9d8be655ee27d89f6d70f3a9fe6d45a7f877b7c963a8bc1ad46f7c7a9cfee6.NewPermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemRequestBuilder) Root()(*i2b03e5349f14f4e882066fc04d344abd4c4a92b60832759bc6992a3d82e123ca.RootRequestBuilder) {
    return i2b03e5349f14f4e882066fc04d344abd4c4a92b60832759bc6992a3d82e123ca.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemRequestBuilder) Site()(*i9bb99583bee5cdfb8243ddb4415954a793021828f7cadbd6bd2f195a9d7d4ea7.SiteRequestBuilder) {
    return i9bb99583bee5cdfb8243ddb4415954a793021828f7cadbd6bd2f195a9d7d4ea7.NewSiteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
