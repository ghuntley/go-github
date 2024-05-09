// Copyright 2020 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEnterpriseService_ListConsumedLicenses(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/enterprises/e/consumed-licenses", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"per_page": "2", "page": "2"})
		fmt.Fprint(w, `{
			"total_seats_consumed": 5000,
			"total_seats_purchased": 4500,
			"users": [
			  {
				"github_com_login": "monalisa",
				"github_com_name": "Mona Lisa",
				"enterprise_server_user_ids": [
				  "example_host_name.com:123",
				  "example_host_name_2:222"
				],
				"github_com_user": true,
				"enterprise_server_user": true,
				"visual_studio_subscription_user": false,
				"license_type": "enterprise",
				"github_com_profile": "https://github.com/monalisa",
				"github_com_member_roles": [
				  "org1:Owner",
				  "org2:Owner"
				],
				"github_com_enterprise_roles": [
				  "owner"
				],
				"github_com_verified_domain_emails": [
				  "monalisa@github.com"
				],
				"github_com_saml_name_id": "monalisa",
				"github_com_orgs_with_pending_invites": [
				  "org1",
				  "org2"
				],
				"github_com_two_factor_auth": true,
				"enterprise_server_emails": [
				  "monalisa@github.com"
				],
				"visual_studio_license_status": "",
				"visual_studio_subscription_email": "",
				"total_user_accounts": 3
			  },
			  {
				"github_com_login": "",
				"github_com_name": "",
				"enterprise_server_user_ids": [
				  "example_host_name:123"
				],
				"github_com_user": false,
				"enterprise_server_user": true,
				"visual_studio_subscription_user": false,
				"license_type": "enterprise",
				"github_com_profile": "",
				"github_com_member_roles": [],
				"github_com_enterprise_role": "",
				"github_com_enterprise_roles": [],
				"github_com_verified_domain_emails": [],
				"github_com_saml_name_id": "",
				"github_com_orgs_with_pending_invites": [],
				"github_com_two_factor_auth": false,
				"enterprise_server_emails": [
				  "hubot@example.com"
				],
				"visual_studio_license_status": "",
				"visual_studio_subscription_email": "",
				"total_user_accounts": 1
			  }
			]
		  }`)
	})

	opts := &ListOptions{Page: 2, PerPage: 2}
	ctx := context.Background()
	groups, _, err := client.Enterprise.ListConsumedLicenses(ctx, "e", opts)
	if err != nil {
		t.Errorf("Enterprise.ListConsumedLicenses returned error: %v", err)
	}

	want := &EnterpriseConsumedLicense{
		TotalSeatsConsumed:  Int(5000),
		TotalSeatsPurchased: Int(4500),
		Users: &[]EnterpriseConsumedLicenseUser{
			{
				GithubComLogin: String("monalisa"),
				GithubComName:  String("Mona Lisa"),
				EnterpriseServerUserIds: &[]string{
					"example_host_name.com:123",
					"example_host_name_2:222",
				},
				GithubComUser:                Bool(true),
				EnterpriseServerUser:         Bool(true),
				VisualStudioSubscriptionUser: Bool(false),
				LicenseType:                  String("enterprise"),
				GithubComProfile:             String("https://github.com/monalisa"),
				GithubComMemberRoles: &[]string{
					"org1:Owner",
					"org2:Owner",
				},
				GithubComEnterpriseRole: String("owner"),
				GithubComVerifiedDomainEmails: &[]string{
					"monalisa@github.com",
				},
				GithubComSamlNameID: String("monalisa"),
				GithubComOrgsWithPendingInvites: &[]string{
					"org1",
					"org2",
				},
				GithubComTwoFactorAuth: Bool(true),
				EnterpriseServerEmails: &[]string{
					"monalisa@github.com",
				},
				VisualStudioLicenseStatus:     String(""),
				VisualStudioSubscriptionEmail: String(""),
				TotalUserAccounts:             Int(3),
			},
			{
				GithubComLogin: String(""),
				GithubComName:  String(""),
				EnterpriseServerUserIds: &[]string{
					"example_host_name.com:123",
				},
				GithubComUser:                   Bool(false),
				EnterpriseServerUser:            Bool(true),
				VisualStudioSubscriptionUser:    Bool(false),
				LicenseType:                     String("enterprise"),
				GithubComProfile:                String(""),
				GithubComMemberRoles:            &[]string{},
				GithubComEnterpriseRole:         String(""),
				GithubComVerifiedDomainEmails:   &[]string{},
				GithubComSamlNameID:             String(""),
				GithubComOrgsWithPendingInvites: &[]string{},
				GithubComTwoFactorAuth:          Bool(true),
				EnterpriseServerEmails: &[]string{
					"hubot@github.com",
				},
				VisualStudioLicenseStatus:     String(""),
				VisualStudioSubscriptionEmail: String(""),
				TotalUserAccounts:             Int(1),
			},
		},
	}

	if !cmp.Equal(groups, want) {
		t.Errorf("Enterprise.ListConsumedLicenses returned %+v, want %+v", groups, want)
	}

	const methodName = "ListConsumedLicenses"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Enterprise.ListConsumedLicenses(ctx, "\n", opts)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Enterprise.ListConsumedLicenses(ctx, "e", opts)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})

}
