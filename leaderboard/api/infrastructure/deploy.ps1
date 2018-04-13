Param(
    [string] [Parameter(Mandatory=$true)] $ResourceGroupName,
    [string] [Parameter(Mandatory=$true)] $Location,
    [string] [Parameter(Mandatory=$true)] $CosmosdbAccountName,
    [string] [Parameter(Mandatory=$true)] $FunctionAppBaseName
)

# NOTE: Since this script works on the VSTS, I skip the login script.
# Login-AzureRmAccount

# Create a ResourceGroup
Get-AzureRmResourceGroup -Name $ResourceGroupName -ErrorVariable notPresent -ErrorAction SilentlyContinue

if(!$notPresent)
{
    Remove-AzureRmResourceGroup -Name $ResourceGroupName -Force
}

New-AzureRmResourceGroup -Name $ResourceGroupName -Location $Location -Force

# Create a CosmosDB

Function Get-PrimaryKey
{
    [CmdletBinding()]
  Param
  (
        [Parameter(Mandatory=$true)][String]$DocumentDBApi,
        [Parameter(Mandatory=$true)][String]$ResourceGroupName,
        [Parameter(Mandatory=$true)][String]$CosmosdbAccountName
    )
    try
    {
  
        $keys=Invoke-AzureRmResourceAction -Action listKeys -ResourceType "Microsoft.DocumentDb/databaseAccounts" -ApiVersion $DocumentDBApi -ResourceGroupName $ResourceGroupName -Name $CosmosdbAccountName -Force
        $connectionKey=$keys[0].primaryMasterKey
        return $connectionKey
    }
    catch 
    {
        Write-Host "ErrorStatusDescription:" $_
    }
}


$locations = @(@{"locationName"="japaneast"; "failoverPriority"=0})

$consistencyPolicy = @{"defaultConsistencyLevel"="Session";
                        "maxIntervalInSeconds"="10";
                        "maxStalenessPrefix"="200"}

$DBProperties = @{"databaseAccountOfferType"="Standard";
                    "locations"=$Locations;
                    "consistencyPolicy"=$consistencyPolicy
                    }

$ResourceName = $CosmosdbAccountName
$DBProperties
New-AzureRmResource -ResourceType "Microsoft.DocumentDb/databaseAccounts"`
                    -ApiVersion "2015-04-08"`
                    -ResourceGroupName $ResourceGroupName `
                    -Location $Location `
                    -Name $CosmosdbAccountName `
                    -Properties $DBProperties

## Retrive CosmosDB ConnectionString
$cosmosPrimaryKey = Get-PrimaryKey -DocumentDBApi "2015-04-08" -ResourceGroupName $ResourceGroupName -CosmosdbAccountName $CosmosdbAccountName
$cosmosDBConnectionString = "AccountEndpoint=https://" + $CosmosdbAccountName + ".documents.azure.com:443/;AccountKey=" + $cosmosPrimaryKey + ";"

# Create a Function App with Function App V2
# Set AppSettings to the Function App

$currentSubscriptionId = (Get-AzureRmContext).Subscription.Id
$hostingPlanName = $FunctionAppBaseName + "Plan"

$random = Get-Random -minimum 1000000 -maximum 9999999;([String]$random).SubString(1,6)
$storageName = $FunctionAppBaseName + $random
New-AzureRmResourceGroupDeployment -Name LeaderBoardBackendDeployment -ResourceGroup $ResourceGroupName -Templatefile scripts/template.json -functionName $FunctionAppBaseName -storageName $StorageName -hostingPlanName $hostingPlanName -location $Location -sku Standard -workerSize 0 -serverFarmResourceGroup $ResourceGroupName -skuCode "S1" -subscriptionId $currentSubscriptionId -cosmosConnectionString $cosmosDBConnectionString




