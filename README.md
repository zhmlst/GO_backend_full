# 🚀 Getting Started Prerequisites Go 1.22+
full backend with go

Air (for live reloading)

Environment Setup
The application uses environment variables for configuration. These are defined in the .env (or .envrc) file.

1. **Configuration File**
Create a .envrc file in the root directory:

```bash 
ADDR=":3000"
# Add other config like DB_ADDR here later
```

2. **Running the App (Windows / PowerShell)**
If you are using PowerShell, environment variables aren't automatically loaded from the file. You can set them for your session before running the server:

```powershell 
# Set the port and Run with hot-reload
  $env:ADDR=":3000" air 
```

3. **Running the App (Linux / macOS / Git Bash)**
```bash 
   export ADDR=":3000"
```

Note: If the ADDR variable is not set, the server will default to :8080.
