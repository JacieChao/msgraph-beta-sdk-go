package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder provides operations to manage the collection of policyRoot entities.
type MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetQueryParameters get the list of groups that are included in a mobile app management policy. This API is available in the following national cloud deployments.
type MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetQueryParameters
}
// MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) {
    m := &MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/mobileDeviceManagementPolicies/{mobilityManagementPolicy%2Did}/includedGroups/$ref{?%24top,%24skip,%24search,%24filter,%24count,%24orderby}", pathParameters),
    }
    return m
}
// NewMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the list of groups that are included in a mobile app management policy. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mobileappmanagementpolicies-list-includedgroups?view=graph-rest-1.0
func (m *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.StringCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateStringCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.StringCollectionResponseable), nil
}
// Post add groups to be included in a mobile app management policy. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/mobileappmanagementpolicies-post-includedgroups?view=graph-rest-1.0
func (m *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceCreateable, requestConfiguration *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation get the list of groups that are included in a mobile app management policy. This API is available in the following national cloud deployments.
func (m *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation add groups to be included in a mobile app management policy. This API is available in the following national cloud deployments.
func (m *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReferenceCreateable, requestConfiguration *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) WithUrl(rawUrl string)(*MobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder) {
    return NewMobileDeviceManagementPoliciesItemIncludedGroupsRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
