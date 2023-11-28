#!/bin/bash

# Variables
ARCH=$(uname -m) # Detects architecture of the system
TARBALL_URL="https://github.com/zcubbs/mq-watch/releases/latest/download/mq-watch_Linux_$ARCH.tar.gz"
INSTALL_DIR="/opt/mq-watch"
LOG_DIR="$INSTALL_DIR/logs"
SERVICE_PATH="/etc/systemd/system/mq-watch.service"
UNINSTALL_PATH="$INSTALL_DIR/uninstall.sh"
SERVICE_USER="mq-watch" # User that will run the service

# Configuration Variables

# Ensure the script is run as root
if [ "$(id -u)" -ne 0 ]; then
    echo "Please run this script as root."
    exit 1
fi

# Download, extract and install binary
echo "Downloading and installing..."
mkdir -p $INSTALL_DIR
curl -L -o "$INSTALL_DIR/mq-watch.tar.gz" "$TARBALL_URL"
tar -xzf "$INSTALL_DIR/mq-watch.tar.gz" -C $INSTALL_DIR
rm "$INSTALL_DIR/mq-watch.tar.gz"

# Create logs directory
echo "Creating logs directory at $LOG_DIR..."
mkdir -p $LOG_DIR

# Create config directory
echo "Creating config directory at $INSTALL_DIR/config..."
mkdir -p $INSTALL_DIR/config

# Create initial configuration file
echo "Creating initial configuration file..."
cat <<EOL > $INSTALL_DIR/config/config.yaml
mqtt:
  broker: "tcp://127.0.0.1:1883" # match the broker address used in generate_data.go
  client_id: "mq-watch" # match the client ID used in generate_data.go

tenants: []

database:
  dialect: "sqlite"
  datasource: "mq-watch.db"
  auto_migrate: true

server:
  port: 8000
EOL

# Create a dedicated user and grant it necessary permissions
if ! id "$SERVICE_USER" &>/dev/null; then
    useradd -r -s /sbin/nologin $SERVICE_USER
fi
chown $SERVICE_USER: $INSTALL_DIR -R

# Configure systemd service
echo "Configuring systemd service..."
cat <<EOL > $SERVICE_PATH
[Unit]
Description=mq-watch Service
After=network.target

[Service]
ExecStart=$INSTALL_DIR/mq-watch -config $INSTALL_DIR/config
EnvironmentFile=$ENV_FILE
Restart=always
User=$SERVICE_USER
Group=nogroup
Environment=PATH=/usr/bin:/usr/local/bin:/usr/sbin
WorkingDirectory=$INSTALL_DIR
StandardOutput=append:$LOG_DIR/output.log
StandardError=append:$LOG_DIR/error.log

[Install]
WantedBy=multi-user.target
EOL

# Generate uninstall script
echo "Generating uninstall script..."
cat <<EOL > $UNINSTALL_PATH
#!/bin/bash

# Stop and disable service
systemctl stop mq-watch
systemctl disable mq-watch

# Remove systemd service
rm -f $SERVICE_PATH

# Remove sudoers permission for mq-watch user
rm -f /etc/sudoers.d/mq-watch_permissions

# Remove installation directory
rm -rf $INSTALL_DIR

echo "Uninstallation complete."

EOL
chmod +x $UNINSTALL_PATH

# Reload systemd, enable and start service
systemctl daemon-reload
systemctl enable mq-watch
systemctl start mq-watch

echo "Installation complete. Running as a systemd service."
echo "To uninstall, run $UNINSTALL_PATH"
