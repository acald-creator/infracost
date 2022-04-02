provider "azurerm" {
  skip_provider_registration = true
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "example-resources"
  location = "eastus"
}

resource "azurerm_cosmosdb_account" "example" {
  name                       = "tfex-cosmosdb-account"
  resource_group_name        = azurerm_resource_group.example.name
  location                   = azurerm_resource_group.example.location
  offer_type                 = "Standard"
  analytical_storage_enabled = true

  consistency_policy {
    consistency_level = "Strong"
  }

  geo_location {
    location          = "westus"
    failover_priority = 0
  }

  backup {
    type                = "Periodic"
    interval_in_minutes = 240
    retention_in_hours  = 16
  }
}

resource "azurerm_cosmosdb_account" "continuous_backup" {
  name                = "tfex-cosmosdb-account"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  offer_type          = "Standard"

  consistency_policy {
    consistency_level = "Strong"
  }

  geo_location {
    location          = "westus"
    failover_priority = 0
  }

  geo_location {
    location          = "centralus"
    failover_priority = 1
    zone_redundant    = true
  }
  backup {
    type = "Continuous"
  }
}

resource "azurerm_cosmosdb_account" "multi-master_backup2copies" {
  name                            = "tfex-cosmosdb-account"
  resource_group_name             = azurerm_resource_group.example.name
  location                        = azurerm_resource_group.example.location
  offer_type                      = "Standard"
  enable_multiple_write_locations = true

  consistency_policy {
    consistency_level = "Strong"
  }

  geo_location {
    location          = "westus"
    failover_priority = 0
  }

  geo_location {
    location          = "centralus"
    failover_priority = 1
  }
  backup {
    type = "Periodic"
  }
}

resource "azurerm_cosmosdb_table" "non-usage_autoscale" {
  name                = "tfex-cosmos-cassandra-keyspace"
  resource_group_name = azurerm_cosmosdb_account.example.resource_group_name
  account_name        = azurerm_cosmosdb_account.example.name
  autoscale_settings {
    max_throughput = 4000
  }
}

resource "azurerm_cosmosdb_table" "autoscale" {
  name                = "tfex-cosmos-cassandra-keyspace"
  resource_group_name = azurerm_cosmosdb_account.example.resource_group_name
  account_name        = azurerm_cosmosdb_account.continuous_backup.name
  autoscale_settings {
    max_throughput = 6000
  }
}

resource "azurerm_cosmosdb_table" "provisioned" {
  name                = "tfex-cosmos-cassandra-keyspace"
  resource_group_name = azurerm_cosmosdb_account.example.resource_group_name
  account_name        = azurerm_cosmosdb_account.continuous_backup.name
  throughput          = 500
}

resource "azurerm_cosmosdb_table" "mutli-master_backup2copies" {
  name                = "tfex-cosmos-cassandra-keyspace"
  resource_group_name = azurerm_cosmosdb_account.example.resource_group_name
  account_name        = azurerm_cosmosdb_account.multi-master_backup2copies.name
  throughput          = 1000
}

resource "azurerm_cosmosdb_table" "serverless" {
  name                = "tfex-cosmos-cassandra-keyspace"
  resource_group_name = azurerm_cosmosdb_account.example.resource_group_name
  account_name        = azurerm_cosmosdb_account.example.name
}

resource "azurerm_cosmosdb_table" "account-in-another-module" {
  name                = "tfex-cosmos-cassandra-keyspace"
  resource_group_name = azurerm_cosmosdb_account.example.resource_group_name
  account_name        = "in-another-module"
}
