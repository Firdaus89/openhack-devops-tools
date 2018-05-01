# Generate a value.yaml for sentinel helm
Write-Output ""
Write-Output "**************************************************************************************************"
Write-Output "* Generating a value.yaml for the sentinel helm..."
Write-Output "**************************************************************************************************"

$template = Get-Content '.\templates\template.txt' -Raw

$services = ""

function Get-ObjectMembers {
    [CmdletBinding()]
    Param(
        [Parameter(Mandatory=$True, ValueFromPipeline=$True)]
        [PSCustomObject]$obj
    )
    $obj | Get-Member -MemberType NoteProperty | ForEach-Object {
        $key = $_.Name
        [PSCustomObject]@{Key = $key; Value = $obj."$key"}
    }
}

function Get-SubService{

    Param(
        [String] $services, 
        [int] $counter, 
        [String] $teamId, 
        [String] $endpoint, 
        [int] $servicecount) 

    $serviceId = $counter.ToString("00") + $servicecount.ToString("00")
    $subtemplate = Get-Content '.\templates\subtemplate.txt' -Raw
    $subExpand = Invoke-Expression "@`"`r`n$subtemplate`r`n`"@"
    $services = -join($services, $subExpand)
    return $services
}


$data = Get-Content '.\sample.json'` -Raw
$json = ConvertFrom-Json($data)

$counter = 1
$services = ""
$json | Get-ObjectMembers | foreach {
    
    $services = Get-SubService -services $services  -counter $counter -teamId $_.Key -endpoint  ($_.Value.endpoint01 + "/aaa/healthcheck") -servicecount 1 
    $services = Get-SubService -services $services  -counter $counter -teamId $_.Key -endpoint  ($_.Value.endpoint01 + "/bbb/healthcheck") -servicecount 2 
    $services = Get-SubService -services $services  -counter $counter -teamId $_.Key -endpoint  ($_.Value.endpoint01 + "/ccc/healthcheck") -servicecount 3                               
    $counter = $counter + 1
}


$expand = Invoke-Expression "@`"`r`n$template`r`n`"@"
Write-Host $expand
Write-Host ""

$expand | Out-File '..\..\sentinel\chart\sentinel\values.yaml' -Encoding UTF8

# Write-Host "..\..\sentinel\values.yaml has been generated"
# 
# Set-AzureKeyVaultSecret -VaultName $ExternalKeyVaultName -Name 'helmValuesYaml' -SecretValue (ConvertTo-SecureString $expand -AsPlainText -Force)
#
# Write-Host "values.yaml is published to the Key Vault: " + $ExternalKeyVaultName