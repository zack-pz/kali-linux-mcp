# Security Tools

This project contains security tools including an MCP server for Kali Linux tools.

## Troubleshooting

### Docker Permission Denied

If you encounter a "permission denied" error when trying to use Docker, it means your user does not have permission to access the Docker daemon socket.

**Error Message:**
`permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock`

**Solution:**

You have two options:

1.  **Run with sudo:**
    ```bash
    sudo ./bin/auth-service
    ```

2.  **Add your user to the docker group (Recommended):**
    ```bash
    sudo usermod -aG docker $USER
    newgrp docker
    ```
    You may need to log out and log back in for the changes to take full effect.
