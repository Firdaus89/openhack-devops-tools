resource "azurerm_resource_group" "rg" {
  name     = "${var.team_resource_suffix}${format("%01d", count.index+100)}"
  count    = "${var.team_count}"
  location = "${var.location}"

  tags {
    Team = "${var.team_resource_suffix}${format("%01d", count.index+100)}"
  }
}