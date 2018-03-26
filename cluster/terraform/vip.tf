resource "azurerm_public_ip" "vip" {
    name                            = "${var.team_resource_suffix}${format("%01d", count.index+100)}vip"
    count                           = "${var.team_count}"
    public_ip_address_allocation    = "static"
    location                        = "${var.location}"
    resource_group_name             = "${element(azurerm_resource_group.rg.*.name, count.index+100)}"

    tags {
        Team = "${var.team_resource_suffix}${format("%01d", count.index+100)}"
    }
}