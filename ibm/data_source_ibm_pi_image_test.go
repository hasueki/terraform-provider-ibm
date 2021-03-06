// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPIImageDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_image.testacc_ds_image", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIImageDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_pi_image" "testacc_ds_image" {
		pi_image_name = "%s"
		pi_cloud_instance_id = "%s"
	}`, pi_image, pi_cloud_instance_id)

}
