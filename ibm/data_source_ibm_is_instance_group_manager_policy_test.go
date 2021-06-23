// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISInstanceGroupManagerPolicy_dataBasic(t *testing.T) {
	randInt := acctest.RandIntRange(1200, 1300)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDmdgitxJde3s6PFDmWbyoF8S5YsC8+l3Qo3vRzBzXf05n5b3JL0t1sswZ7XNNyO2y8jTX7sCpGMzv4Q3WksCzkU12OPbr89Zmf+mC+11o3Lp/NpejiNtYf8hVtHWUAUrKLNywjFnm28pn64pf9KFgdkkp9quBZQgis8osfeygknYaSBBzkZKaZPszGuixTqaRAaomfwDP7QJJvS3Bo8bAe2kK+4EsW2DfP7h1G6BhHoxjoinVshbfE1nsJ2zlQigidjyjFL5YbCUYygjz5kq2khoxWmaNNKPVxAZ8fqvIHNi8F8sLCKW6VTruxPQIlW2A/D1YIJ4ME/Y6Goje9l40dA1W/mnygD0mZVYiLtYtlUM6ylKoQKNGeV+ugA554UK0lA++FVg5xOm8SNSvWWf6hyN/mK6atbpSBzRLoUQc95XsG1u7eQtz/zA1+pKaVsASpCMMbZFcTaeOiPLIcIVlzYDBcKap0MglFlsJsoKSRJ4uxIGm+CHoCWdc7VgaVvz8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	instanceGroupManager := fmt.Sprintf("testinstancegroupmanager%d", randInt)
	instanceGroupManagerPolicy := fmt.Sprintf("testinstancegroupmanagerpolicy%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupManagerPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupManagerPolicyConfigd(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group_manager"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "metric_type"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "metric_value"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "policy_type"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerPolicyConfigd(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy string) string {
	return testAccCheckIBMISInstanceGroupManagerPolicyConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy) + fmt.Sprintf(`
	
	data "ibm_is_instance_group_manager_policy" "instance_group_manager_policy" {
		instance_group = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group
		instance_group_manager = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group_manager
		name = ibm_is_instance_group_manager_policy.cpuPolicy.name
	}
	`)

}
