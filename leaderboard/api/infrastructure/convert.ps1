# Generate a value.yaml for sentinel helm
Write-Output ""
Write-Output "**************************************************************************************************"
Write-Output "* Generating a value.yaml for the sentinel helm..."
Write-Output "**************************************************************************************************"

$template = Get-Content '.\templates\template.txt' -Raw
$subtemplate = Get-Content '.\templates\subtemplate.txt' -Raw

$services = ""

For ($teamId=1; $teamId -le $teamNum; $teamId++) {
    For ($id=1; $id -le $servicesPerTeam; $id++){
        $serviceId = $teamId.ToString("00") + $id.ToString("00")
        $subExpand = Invoke-Expression "@`"`r`n$subtemplate`r`n`"@"

        $endpoint = (Get-AzureKeyVaultSecret -VaultName $ExternalKeyVaultName -Name ("Team" + $teamId.ToString("00") + "-Endpoint" + $id.ToString("00"))).SecretValueText
        # Write-Host $subexpand
        $services = -join($services, $subExpand)
        # Write-Host $services
    }
}

$expand = Invoke-Expression "@`"`r`n$template`r`n`"@"
Write-Host $expand
Write-Host ""

$expand | Out-File '..\..\sentinel\values.yaml' -Encoding UTF8

Write-Host "..\..\sentinel\values.yaml has been generated"

Set-AzureKeyVaultSecret -VaultName $ExternalKeyVaultName -Name 'helmValuesYaml' -SecretValue (ConvertTo-SecureString $expand -AsPlainText -Force)

Write-Host "values.yaml is published to the Key Vault: " + $ExternalKeyVaultName