resource "azurerm_kubernetes_cluster" "aks" {
  name                   = "${var.team_resource_suffix}${format("%01d", count.index+100)}"
  count                  = "${var.team_count}"
  location               = "${var.location}"
  resource_group_name    = "${element(azurerm_resource_group.rg.*.name, count.index+100)}"
  dns_prefix             = "${element(azurerm_resource_group.rg.*.name, count.index+100)}${substr(sha1(uuid()), 0, 5)}"
  lifecycle {
    ignore_changes = ["dns_prefix"]
  }

  linux_profile {
    admin_username = "${var.team_resource_suffix}${format("%01d", count.index+100)}"

    ssh_key {
      key_data = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCqaZoyiz1qbdOQ8xEf6uEu1cCwYowo5FHtsBhqLoDnnp7KUTEBN+L2NxRIfQ781rxV6Iq5jSav6b2Q8z5KiseOlvKA/RF2wqU0UPYqQviQhLmW6THTpmrv/YkUCuzxDpsH7DUDhZcwySLKVVe0Qm3+5N2Ta6UYH3lsDf9R9wTP2K/+vAnflKebuypNlmocIvakFWoZda18FOmsOoIVXQ8HWFNCuw9ZCunMSN62QGamCe3dL5cXlkgHYv7ekJE15IA9aOJcM7e90oeTqo+7HTcWfdu0qQqPWY5ujyMw/llas8tsXY85LFqRnr3gJ02bAscjc477+X+j/gkpFoN1QEmt terraform@demo.tld"
    }
  }

  agent_pool_profile {
    name            = "default"
    count           = "${var.node_count}"
    vm_size         = "${var.vm_size}"
    os_type         = "Linux"
  }

  service_principal {
    client_id     = "7fb4173c-3ca3-4d5b-87f8-1daac941207a"
    client_secret = "MPNSuM1auUuITefiLGBrpZZnLMDKBLw2"
  }

  tags {
    team = "${var.team_resource_suffix}${format("%01d", count.index+100)}"
  }

  provisioner "local-exec" {
    command = "az aks get-credentials -n ${azurerm_kubernetes_cluster.aks.name} -g ${azurerm_kubernetes_cluster.aks.resource_group_name}"
  }
}