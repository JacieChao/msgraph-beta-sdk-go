package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/networkaccess"
)

// ForwardingProfilesItemPoliciesItemPolicyRequestBuilder provides operations to manage the policy property of the microsoft.graph.networkaccess.policyLink entity.
type ForwardingProfilesItemPoliciesItemPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetQueryParameters policy.
type ForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetQueryParameters
}
// NewForwardingProfilesItemPoliciesItemPolicyRequestBuilderInternal instantiates a new PolicyRequestBuilder and sets the default values.
func NewForwardingProfilesItemPoliciesItemPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ForwardingProfilesItemPoliciesItemPolicyRequestBuilder) {
    m := &ForwardingProfilesItemPoliciesItemPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/forwardingProfiles/{forwardingProfile%2Did}/policies/{policyLink%2Did}/policy{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewForwardingProfilesItemPoliciesItemPolicyRequestBuilder instantiates a new PolicyRequestBuilder and sets the default values.
func NewForwardingProfilesItemPoliciesItemPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ForwardingProfilesItemPoliciesItemPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewForwardingProfilesItemPoliciesItemPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get policy.
func (m *ForwardingProfilesItemPoliciesItemPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *ForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetRequestConfiguration)(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Policyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.CreatePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i43e723cc778f0f3f3a05d36b9df74faa56771e9360d8ed793c50bdaacec8d5d2.Policyable), nil
}
// ToGetRequestInformation policy.
func (m *ForwardingProfilesItemPoliciesItemPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ForwardingProfilesItemPoliciesItemPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ForwardingProfilesItemPoliciesItemPolicyRequestBuilder) WithUrl(rawUrl string)(*ForwardingProfilesItemPoliciesItemPolicyRequestBuilder) {
    return NewForwardingProfilesItemPoliciesItemPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
