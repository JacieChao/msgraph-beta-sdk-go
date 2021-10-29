package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i458957d039676a87049f925f9a0ac58100c522d4c5494921661924e3d7de9536 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item/excludes"
    ibb33c8122164f9a21b54efbd3817c55a46e13f7c672ae3565862893e0225a426 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item/includes"
    i8b8ea18691791149f6486e79113d56f56a6921846e0823954a25159321c50bdc "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item/includes/item"
    if3d69c085cc31553bffa6238747fd2f166e883710ce3cb7738b69051017fde73 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/serviceprincipalcreationpolicies/item/excludes/item"
)

// Builds and executes requests for operations under \policies\servicePrincipalCreationPolicies\{servicePrincipalCreationPolicy-id}
type ServicePrincipalCreationPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ServicePrincipalCreationPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ServicePrincipalCreationPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ServicePrincipalCreationPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get servicePrincipalCreationPolicies from policies
type ServicePrincipalCreationPolicyRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ServicePrincipalCreationPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalCreationPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new ServicePrincipalCreationPolicyRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewServicePrincipalCreationPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalCreationPolicyRequestBuilder) {
    m := &ServicePrincipalCreationPolicyRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/policies/servicePrincipalCreationPolicies/{servicePrincipalCreationPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ServicePrincipalCreationPolicyRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewServicePrincipalCreationPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalCreationPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalCreationPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property servicePrincipalCreationPolicies for policies
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalCreationPolicyRequestBuilder) CreateDeleteRequestInformation(options *ServicePrincipalCreationPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get servicePrincipalCreationPolicies from policies
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalCreationPolicyRequestBuilder) CreateGetRequestInformation(options *ServicePrincipalCreationPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update the navigation property servicePrincipalCreationPolicies in policies
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalCreationPolicyRequestBuilder) CreatePatchRequestInformation(options *ServicePrincipalCreationPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete navigation property servicePrincipalCreationPolicies for policies
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalCreationPolicyRequestBuilder) Delete(options *ServicePrincipalCreationPolicyRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ServicePrincipalCreationPolicyRequestBuilder) Excludes()(*i458957d039676a87049f925f9a0ac58100c522d4c5494921661924e3d7de9536.ExcludesRequestBuilder) {
    return i458957d039676a87049f925f9a0ac58100c522d4c5494921661924e3d7de9536.NewExcludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.policies.servicePrincipalCreationPolicies.item.excludes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServicePrincipalCreationPolicyRequestBuilder) ExcludesById(id string)(*if3d69c085cc31553bffa6238747fd2f166e883710ce3cb7738b69051017fde73.ServicePrincipalCreationConditionSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationConditionSet_id"] = id
    }
    return if3d69c085cc31553bffa6238747fd2f166e883710ce3cb7738b69051017fde73.NewServicePrincipalCreationConditionSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get servicePrincipalCreationPolicies from policies
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalCreationPolicyRequestBuilder) Get(options *ServicePrincipalCreationPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalCreationPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewServicePrincipalCreationPolicy() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ServicePrincipalCreationPolicy), nil
}
func (m *ServicePrincipalCreationPolicyRequestBuilder) Includes()(*ibb33c8122164f9a21b54efbd3817c55a46e13f7c672ae3565862893e0225a426.IncludesRequestBuilder) {
    return ibb33c8122164f9a21b54efbd3817c55a46e13f7c672ae3565862893e0225a426.NewIncludesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.policies.servicePrincipalCreationPolicies.item.includes.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ServicePrincipalCreationPolicyRequestBuilder) IncludesById(id string)(*i8b8ea18691791149f6486e79113d56f56a6921846e0823954a25159321c50bdc.ServicePrincipalCreationConditionSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationConditionSet_id"] = id
    }
    return i8b8ea18691791149f6486e79113d56f56a6921846e0823954a25159321c50bdc.NewServicePrincipalCreationConditionSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update the navigation property servicePrincipalCreationPolicies in policies
// Parameters:
//  - options : Options for the request
func (m *ServicePrincipalCreationPolicyRequestBuilder) Patch(options *ServicePrincipalCreationPolicyRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}