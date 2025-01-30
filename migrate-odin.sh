#!/bin/bash
set -e

# Revert function to restore old odin binary
revert_odin() {
  echo "Reverting to old odin binary..."

  # Check if there is a backup (old-odin) somewhere
  if [ -f /usr/local/bin/old-odin ]; then
    # If backup exists, remove any current odin binary
    if [ -f /usr/local/bin/odin ]; then
      sudo rm -f /usr/local/bin/odin
      echo "Removed existing odin binary from /usr/local/bin"
    fi

    if [ -f /opt/homebrew/bin/odin ]; then
      sudo rm -f /opt/homebrew/bin/odin
      echo "Removed existing odin binary from /opt/homebrew/bin"
    fi

    # Restore the old-odin backup
    sudo mv /usr/local/bin/old-odin /usr/local/bin/odin
    chmod 755 /usr/local/bin/odin
    echo "Reverted to the old odin binary successfully from /usr/local/bin, configure again to start using odin"
  
  elif [ -f /opt/homebrew/bin/old-odin ]; then
    # If backup exists, remove any current odin binary
    if [ -f /opt/homebrew/bin/odin ]; then
      sudo rm -f /opt/homebrew/bin/odin
      echo "Removed existing odin binary from /opt/homebrew/bin"
    fi

    if [ -f /usr/local/bin/odin ]; then
      sudo rm -f /usr/local/bin/odin
      echo "Removed existing odin binary from /usr/local/bin"
    fi

    # Restore the old-odin backup
    sudo mv /opt/homebrew/bin/old-odin /opt/homebrew/bin/odin
    chmod 755 /opt/homebrew/bin/odin
    echo "Reverted to the old odin binary successfully from /opt/homebrew/bin, configure again to start using odin"
  
  else
    echo "No backup odin binary found to revert to."
    exit 1
  fi
}



# Check for revert flag
if [ "$1" == "revert" ]; then
  revert_odin
  exit 0
fi


# Fetching the odin binary from GitHub
curl -L -o ./odin https://raw.githubusercontent.com/dream11/odin/chore/add-binary/odin
if [ $? -ne 0 ]; then
  echo "Failed to download the odin binary. Exiting."
  exit 1
fi
chmod +x ./odin

echo "Downloaded the odin binary successfully."


if [ -f /usr/local/bin/odin ] && [ ! -f /usr/local/bin/old-odin ] && [ ! -f /opt/homebrew/bin/old-odin ]; then
  echo "Existing odin binary found, moving it to backup."
  sudo mv /usr/local/bin/odin /usr/local/bin/old-odin
fi

# Handle /opt/homebrew/bin case if it exists
if [ -f /opt/homebrew/bin/odin ] && [ ! -f /opt/homebrew/bin/old-odin ] && [ ! -f /usr/local/bin/old-odin ]; then
  echo "Homebrew odin binary found, moving it to backup."
  sudo mv /opt/homebrew/bin/odin /opt/homebrew/bin/old-odin
fi

sudo mv ./odin /usr/local/bin/
chmod 755 /usr/local/bin/odin


# Enable app verification and remove quarantine attributes
sudo spctl --master-enable
xattr -dr com.apple.quarantine /usr/local/bin/odin

# Verify the new odin installation
echo "New odin installed, running odin version to verify [Should output 2.0.0]"
odin version || { echo "Odin version verification failed."; exit 1; }


# Read access and secret keys from ~/.odin/config if they exist
CONFIG_FILE=~/.odin/config
if [ -f "$CONFIG_FILE" ]; then
  echo "Reading access and secret keys from config."
  ODIN_ACCESS_KEY=$(awk '/access_key:/ && !/secret_access_key:/ {print $2}' $CONFIG_FILE)
  ODIN_SECRET_ACCESS_KEY=$(awk '/secret_access_key:/ {print $2}' $CONFIG_FILE)
else
  echo "Config file not found. Prompting user for keys."
  read -p "Enter ODIN_ACCESS_KEY: " ODIN_ACCESS_KEY
  read -p "Enter ODIN_SECRET_ACCESS_KEY: " ODIN_SECRET_ACCESS_KEY
fi

# If keys are empty or file doesn't exist, prompt user for input
if [ -z "${ODIN_ACCESS_KEY:-}" ] || [ -z "${ODIN_SECRET_ACCESS_KEY:-}" ]; then
  echo "Access keys are missing"
  read -p "Enter ODIN_ACCESS_KEY: " ODIN_ACCESS_KEY
  read -p "Enter ODIN_SECRET_ACCESS_KEY: " ODIN_SECRET_ACCESS_KEY
fi

# Backup existing configuration if it exists
if [ -f ~/.odin/config ]; then
  echo "Existing odin config found, creating a backup."
  mv ~/.odin/config ~/.odin/old-config
fi

# Set environment variables
export ODIN_ACCESS_KEY=${ODIN_ACCESS_KEY}
export ODIN_SECRET_ACCESS_KEY=${ODIN_SECRET_ACCESS_KEY}
export ODIN_BACKEND_ADDRESS="odin-deployer.dss-platform.private:80"

# Configure odin
odin configure
if [ $? -eq 0 ]; then
  echo "Odin configured successfully."
else
  echo "Failed to configure odin."
  exit 1
fi
