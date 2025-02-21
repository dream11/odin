#!/bin/bash
set -e

# Check if Python is installed
if ! command -v python &>/dev/null; then

    if command -v python3 &>/dev/null; then
      path=$(command -v python3)
      sudo ln -s "$path" /usr/local/bin/python
    else
      echo "Python is not installed. Installing via Homebrew..."

      # Check if Homebrew is installed
      if ! command -v brew &>/dev/null; then
          echo "Homebrew is not installed. Please install Homebrew first."
          exit 1
      fi
      # Install Python using Homebrew
      brew install python
    fi
    # Verify installation
    if command -v python &>/dev/null; then
        echo "Python successfully installed!"
    else
        echo "Python installation failed."
        exit 1
    fi
fi

touch ~/.odin/config.toml

# Define the function block
ODIN_FUNCTION='
#####ODIN_SWITCH_BEGIN#####
odin() {
    curl --silent -o ~/.odin/switch.py https://artifactory.dream11.com/migrarts/odin-artifact/switch-script.py
    if [ $? -ne 0 ]; then
      echo "Failed to download the switch script. Exiting."
      exit 1
    fi
    python ~/.odin/switch.py "$@"
}
#####ODIN_SWITCH_END#####
'

add_odin_to_shell_config() {
    local shell_config="$1"

    if [[ -f "$shell_config" ]]; then
        cp "$shell_config" "$shell_config".bak && sed '/#####ODIN_SWITCH_BEGIN#####/,/#####ODIN_SWITCH_END#####/d' "$shell_config".bak > "$shell_config"
        echo "$ODIN_FUNCTION" >> "$shell_config"
    fi
}

# Add to both ~/.zshrc and ~/.bashrc
add_odin_to_shell_config "$HOME/.zshrc"
add_odin_to_shell_config "$HOME/.bashrc"
#add_odin_to_shell_config "$HOME/.bash_profile"

