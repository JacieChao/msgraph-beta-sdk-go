package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/definitions"
    i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/parent"
    ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/definitionfile"
    id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicycategories/item/children"
)

// GroupPolicyCategoryItemRequestBuilder builds and executes requests for operations under \deviceManagement\groupPolicyCategories\{groupPolicyCategory-id}
type GroupPolicyCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupPolicyCategoryItemRequestBuilderDeleteOptions options for Delete
type GroupPolicyCategoryItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyCategoryItemRequestBuilderGetOptions options for Get
type GroupPolicyCategoryItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyCategoryItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyCategoryItemRequestBuilderGetQueryParameters the available group policy categories for this account.
type GroupPolicyCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupPolicyCategoryItemRequestBuilderPatchOptions options for Patch
type GroupPolicyCategoryItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyCategory;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *GroupPolicyCategoryItemRequestBuilder) Children()(*id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d.ChildrenRequestBuilder) {
    return id2be9225a42cb4e206ea5df101d0b19a8153d5b9b239d1b508a6a626f022511d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupPolicyCategoryItemRequestBuilderInternal instantiates a new GroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewGroupPolicyCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyCategoryItemRequestBuilder) {
    m := &GroupPolicyCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyCategories/{groupPolicyCategory_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyCategoryItemRequestBuilder instantiates a new GroupPolicyCategoryItemRequestBuilder and sets the default values.
func NewGroupPolicyCategoryItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyCategoryItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) CreateGetRequestInformation(options *GroupPolicyCategoryItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyCategoryItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyCategoryItemRequestBuilder) DefinitionFile()(*ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f.DefinitionFileRequestBuilder) {
    return ib08bb2c1c958799079296de4709174f5e45ba516b4aba985d1cb9bc3f3277f8f.NewDefinitionFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyCategoryItemRequestBuilder) Definitions()(*i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c.DefinitionsRequestBuilder) {
    return i077f9bb18a3b8a45bbe869035f3d5b79f48a4e8c97eeec6eee3ae05847933d3c.NewDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) Delete(options *GroupPolicyCategoryItemRequestBuilderDeleteOptions)(error) {
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
// Get the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) Get(options *GroupPolicyCategoryItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyCategory, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyCategory() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyCategory), nil
}
func (m *GroupPolicyCategoryItemRequestBuilder) Parent()(*i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252.ParentRequestBuilder) {
    return i821765f6697e8edae076c3f190855ac0dea6daad75cfdcb0ceb7c24fc78d5252.NewParentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the available group policy categories for this account.
func (m *GroupPolicyCategoryItemRequestBuilder) Patch(options *GroupPolicyCategoryItemRequestBuilderPatchOptions)(error) {
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