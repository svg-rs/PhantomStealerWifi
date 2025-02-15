# PhantomStealerWifi

## Overview

PhantomStealerWifi is a simple Go project designed to steal Wi-Fi credentials from a Windows machine and send them to a server. It consists of a client that runs a PowerShell script to extract Wi-Fi profiles and their passwords and a server that receives and stores the stolen credentials.

## Features

*   **Client-Side Credential Harvesting:** Utilizes a PowerShell script to extract Wi-Fi SSIDs and passwords from the target machine.
*   **Server-Side Data Storage:** Receives the harvested credentials via a POST request and stores them in a JSON file.
*   **Logging:** The server logs incoming requests to a file, aiding in monitoring and debugging.
*   **Simple REST API:** Server exposes an API endpoint for data reception.

## Prerequisites

*   **Go:**  Ensure you have Go installed (version 1.16 or higher).
*   **Fiber:** You need to install the Fiber web framework.
    ```bash
    go get github.com/gofiber/fiber/v2
    go get github.com/gofiber/fiber/v2/middleware/logger
    ```

## Getting Started

1.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd PhantomStealer
    ```
2.  **Build and Run:**

    *   **Server:**
        ```bash
        go run main.go -mode=server
        ```
        This will start the server, listening on port 3000.
        Make sure that `logs/` and `server/credentials/` directories exists before running the server.  The server will create `server/credentials/creds.json` file in the correct folder and will create the log file inside the `logs` folder.

    *   **Client:**
        To run the client, you need to have a Windows environment (either a physical machine or a virtual machine)
        ```bash
        go run main.go -mode=client
        ```
        The client will execute a PowerShell script that extracts and sends Wi-Fi credentials to the server. The `creds.json` file will be updated at `server/credentials/` folder.

## Usage

1.  Start the server.
2.  Run the client on a target Windows machine. The client will gather the Wi-Fi profiles and their respective passwords and send it to the server at the specified IP and port.
3.  The stolen Wi-Fi credentials will be stored in `server/credentials/creds.json` file on the server machine.
4.  Server logs can be found at `logs/server.log`

## Security Considerations

*   **Malicious Use:** This project is provided for educational and demonstration purposes only.  Do not use it for any illegal or unethical activities.  Unauthorized access to Wi-Fi credentials is a serious offense.
*   **Antivirus Detection:**  Be aware that this kind of tool, even for legitimate use, might be flagged by antivirus software due to its credential-harvesting nature.
*   **Server Security:**  The server is currently very basic. You should implement proper security measures (e.g., authentication, input validation) if you plan to deploy it in a real-world environment.
*   **Error Handling:**  The error handling is kept to a minimum.  In a production environment, more robust error handling should be implemented.

## License

This project is licensed under the [MIT License](LICENSE).