# 👻 Phantom Stealer 👻

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://go.dev/)

## 🚀 Overview

Phantom Stealer is a simple Go-based project designed to **steal Wi-Fi passwords** 😈 from a Windows machine and send them to a server. This project demonstrates basic client-server communication and is intended for educational purposes only.  **Use responsibly!**

## 🛠️ Project Structure
├───client          # 📡 Client-side code (PowerShell script)
├───logs            # 🪵 Server logs
├───server          # 🌐 Server-side code (Fiber web server)
└───static          # 📁 Static files (Not used in this example)
main.go           # 🚦 Entry point - handles client/server mode
## ⚙️ Installation and Usage

1.  **Prerequisites:**

    *   Go installed on your machine.
    *   A Windows environment to run the client.

2.  **Clone the repository:**

    ```bash
    git clone <repository_url>
    cd <repository_directory>
    ```

3.  **Run the Server:**

    ```bash
    go run main.go -mode=server
    ```

    This will start the server, listening on port `3000`.  Logs will be saved to `logs/server.log`.

4.  **Run the Client:**

    *   Open a new terminal or command prompt.
    *   Navigate to the project directory.
    *   Run the client:

        ```bash
        go run main.go -mode=client
        ```

    This will execute the PowerShell script within the `client` directory, collect Wi-Fi profiles, and send them to the server.

## 💡 How it Works

1.  **Client:**

    *   The client-side code, located in `client/client.go`, uses `exec.Command` to run a PowerShell script.
    *   The PowerShell script does the following:
        *   Uses `netsh wlan show profiles` to get a list of saved Wi-Fi profiles.
        *   For each profile, uses `netsh wlan show profile ... key=clear` to retrieve the password (if available).
        *   Formats the collected data into JSON.
        *   Uses `Invoke-RestMethod` to send a POST request to the server endpoint `/api/input` with the JSON data.

2.  **Server:**

    *   The server-side code, located in `server/server.go`, uses the [Fiber](https://gofiber.io/) web framework.
    *   The server defines a POST endpoint `/api/input`.
    *   When a request is received at this endpoint:
        *   It parses the JSON data in the request body.
        *   It extracts the SSID and password from the received Wi-Fi profiles.
        *   It prints the received SSID and password to the console.
        *   Logs all requests and responses to `logs/server.log`.

## ⚠️ Disclaimer

This project is for educational purposes only.  It is **crucial** to use this project ethically and responsibly.  Unauthorized access to Wi-Fi passwords is illegal and can have serious consequences.  The author is not responsible for any misuse of this code. Always obtain proper authorization before attempting to access or collect any sensitive information.

## ✨ Features

*   Client-server architecture.
*   Wi-Fi password extraction from Windows using PowerShell.
*   Data transmission using JSON.
*   Simple server with logging.
*   Uses Fiber, a Go web framework.

## 🤝 Contributing

Contributions are welcome!  Feel free to submit pull requests or open issues.

## 📜 License

[MIT License](LICENSE)