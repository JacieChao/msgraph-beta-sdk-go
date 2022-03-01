package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ic9fabc845b02722e73b3fb9e2f8552139c9b07bf659e6677a52a631de914b0b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item/definition"
    ie0ffa233916a11feb6aca33fa5315a7307c871c11d9e88f6dca82ade1eaa693b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item/presentationvalues"
    i133b5aa14c1c7989ad2fec6712baeefe6f3323249842963eb73e49f9e71226da "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item/presentationvalues/item"
)

// GroupPolicyDefinitionValueItemRequestBuilder builds and executes requests for operations under \deviceManagement\groupPolicyConfigurations\{groupPolicyConfiguration-id}\definitionValues\{groupPolicyDefinitionValue-id}
type GroupPolicyDefinitionValueItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupPolicyDefinitionValueItemRequestBuilderDeleteOptions options for Delete
type GroupPolicyDefinitionValueItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyDefinitionValueItemRequestBuilderGetOptions options for Get
type GroupPolicyDefinitionValueItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters the list of enabled or disabled group policy definition values for the configuration.
type GroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupPolicyDefinitionValueItemRequestBuilderPatchOptions options for Patch
type GroupPolicyDefinitionValueItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionValue;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGroupPolicyDefinitionValueItemRequestBuilderInternal instantiates a new GroupPolicyDefinitionValueItemRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionValueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyDefinitionValueItemRequestBuilder) {
    m := &GroupPolicyDefinitionValueItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration_id}/definitionValues/{groupPolicyDefinitionValue_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyDefinitionValueItemRequestBuilder instantiates a new GroupPolicyDefinitionValueItemRequestBuilder and sets the default values.
func NewGroupPolicyDefinitionValueItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyDefinitionValueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyDefinitionValueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyDefinitionValueItemRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyDefinitionValueItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyDefinitionValueItemRequestBuilder) CreateGetRequestInformation(options *GroupPolicyDefinitionValueItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyDefinitionValueItemRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyDefinitionValueItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyDefinitionValueItemRequestBuilder) Definition()(*ic9fabc845b02722e73b3fb9e2f8552139c9b07bf659e6677a52a631de914b0b9.DefinitionRequestBuilder) {
    return ic9fabc845b02722e73b3fb9e2f8552139c9b07bf659e6677a52a631de914b0b9.NewDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyDefinitionValueItemRequestBuilder) Delete(options *GroupPolicyDefinitionValueItemRequestBuilderDeleteOptions)(error) {
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
// Get the list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyDefinitionValueItemRequestBuilder) Get(options *GroupPolicyDefinitionValueItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionValue, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyDefinitionValue() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyDefinitionValue), nil
}
// Patch the list of enabled or disabled group policy definition values for the configuration.
func (m *GroupPolicyDefinitionValueItemRequestBuilder) Patch(options *GroupPolicyDefinitionValueItemRequestBuilderPatchOptions)(error) {
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
func (m *GroupPolicyDefinitionValueItemRequestBuilder) PresentationValues()(*ie0ffa233916a11feb6aca33fa5315a7307c871c11d9e88f6dca82ade1eaa693b.PresentationValuesRequestBuilder) {
    return ie0ffa233916a11feb6aca33fa5315a7307c871c11d9e88f6dca82ade1eaa693b.NewPresentationValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresentationValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyConfigurations.item.definitionValues.item.presentationValues.item collection
func (m *GroupPolicyDefinitionValueItemRequestBuilder) PresentationValuesById(id string)(*i133b5aa14c1c7989ad2fec6712baeefe6f3323249842963eb73e49f9e71226da.GroupPolicyPresentationValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentationValue_id"] = id
    }
    return i133b5aa14c1c7989ad2fec6712baeefe6f3323249842963eb73e49f9e71226da.NewGroupPolicyPresentationValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}